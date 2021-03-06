package pkg

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"
	"strings"

	"github.com/khase/leaseplanabocarexporter/dto"
)

func GetAllCars(token string, page int, count int) ([]dto.Item, error) {
	fullCarList := []dto.Item{}

	pageIndex := 1

	for {
		page, err := getCarPage(token, pageIndex, count)

		if err != nil {
			return nil, err
		}

		if len(page.Items) == 0 {
			break
		}

		fullCarList = append(fullCarList, page.Items...)

		pageIndex++
	}

	return fullCarList, nil
}

func GetPageItems(token string, page int, count int) ([]dto.Item, error) {
	pageResponse, err := getCarPage(token, page, count)

	if err != nil {
		return nil, err
	}

	return pageResponse.Items, nil
}

func getCarPage(token string, page int, count int) (dto.CarResponse, error) {
	data := dto.CarRequest{
		Bookmark:     false,
		ItemsPerPage: count,
		OrderAsc:     false,
		OrderBy:      "DateRegistration",
		Page:         page}

	resp, err := doPostJson(
		"https://rowebapiservice-autoabo.azurewebsites.net/api/carcharter/v1/Shop/SearchOfferTypes?ROType=carch--prd&shopSubdomain=leaseplan-abocar",
		data,
		token)

	if err != nil {
		return dto.CarResponse{}, err
	}

	var res dto.CarResponse
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}

func GetToken(mail string, pass string) (string, error) {
	var res dto.LoginResponse

	r, err := login(mail, pass)

	if err != nil {
		return "", err
	}

	res = r

	token := fmt.Sprintf("%v", res.Content.AddressToken)
	return token, nil
}

func login(mail string, pass string) (dto.LoginResponse, error) {
	data := map[string]string{"Email": mail, "Password": pass}
	resp, err := doPostJson(
		"https://rowebapiservice-autoabo.azurewebsites.net/api/carcharter/v1/Shop/Login?ROType=carch--prd&shopSubdomain=leaseplan-abocar",
		data,
		"")

	if err != nil {
		return dto.LoginResponse{}, err
	}

	var res dto.LoginResponse
	json.NewDecoder(resp.Body).Decode(&res)

	return res, nil
}

func doPostJson(url string, data interface{}, token string) (*http.Response, error) {
	json_data, err := json.Marshal(data)

	if err != nil {
		return nil, err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(json_data))

	if err != nil {
		return nil, err
	}

	req.Header.Set("Content-Type", "application/json")

	if len(token) > 0 {
		req.Header.Set("Address-Token", token)
	}

	resp, err := http.DefaultClient.Do(req)

	if err != nil {
		return nil, err
	}

	if resp.StatusCode != 200 {
		buf := new(strings.Builder)
		_, err := io.Copy(buf, resp.Body)
		if err != nil {
			return nil, err
		}
		return nil, errors.New(buf.String())
	}

	return resp, nil
}
