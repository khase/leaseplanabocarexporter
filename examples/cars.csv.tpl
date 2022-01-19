Name;Marke;Model;PS;BLP;Fahrzeugtyp;Farbe;Farbfamilie;Schaltung;Treibstoff;Antrieb;TürAnz;SitzAnz
{{- range $car := .}}
{{ $car.OfferTypeName }};{{ $car.RentalObject.CarLabel }};{{ $car.RentalObject.CarModell }};{{ $car.RentalObject.PowerHp }};{{ $car.RentalObject.PriceProducer1 }};{{ $car.RentalObject.CarType }};{{ $car.RentalObject.Color }};{{ $car.RentalObject.ColorFamily }};{{ $car.RentalObject.KindOfGear }};{{ $car.RentalObject.KindOfFuel }};{{ $car.RentalObject.KindOfDrive }};{{ $car.RentalObject.NoOfDoors }};{{ $car.RentalObject.NoOfSeats }}
{{- end}}