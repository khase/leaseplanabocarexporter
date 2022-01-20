{{- $last := dict -}}
{{- $current := dict -}}
{{- range $car := index . 0 -}}
{{- $_ := set $current $car.RentalObject.Ident $car  -}}
{{- end -}}
{{- range $car := index . 1 -}}
{{- $_ := set $last $car.RentalObject.Ident $car  -}}
{{- end -}}
{{- $added := list -}}
{{- range $index, $car := $current -}}
{{- if not (hasKey $last $index) -}}
{{- $added = append $added $car -}}
{{- end -}}
{{- end -}}
{{- $removed := list -}}
{{- range $index, $car := $last -}}
{{- if not (hasKey $current $index) -}}
{{- $removed = append $removed $car -}}
{{- end -}}
{{- end -}}
{{- if gt len($added) 0 }}
*Added ({{ len($added) }})*:
{{- range $car := $added}}
    [{{ $car.OfferTypeName }}](https://www.leaseplan-abocar.de/offer-details/{{ $car.Ident }}/{{ $car.RentalObject.Ident }}) {{ $car.RentalObject.PowerHp }}PS ({{ $car.RentalObject.PriceProducer1 }}€)
{{- end}}
{{- end}}
{{- if gt len($removed) 0 }}
{{- if gt len($added) 0 }}

{{ end -}}
*Removed ({{ len($removed) }})*:
{{- range $car := $removed}}
    [{{ $car.OfferTypeName }}](https://www.leaseplan-abocar.de/offer-details/{{ $car.Ident }}/{{ $car.RentalObject.Ident }}) {{ $car.RentalObject.PowerHp }}PS ({{ $car.RentalObject.PriceProducer1 }}€)
{{- end}}
{{- end}}