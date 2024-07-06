![Build and test](https://github.com/nejdetkadir/turkish_cities/actions/workflows/main.yml/badge.svg?branch=main)
![Go Version](https://img.shields.io/badge/go_version-_1.22.2-007d9c.svg)


![cover](docs/cover.png)

# TurkishCities

TurkishCities is a comprehensive library for managing Turkish cities and their administrative divisions (towns, districts, quarters) in Go. It provides an easy way to load, query, and interact with this hierarchical data structure.

## Features
- Latest data for Turkish cities, towns, districts, and quarters
- Query cities, towns, districts, and quarters by IDs
- Access country-level information

## Installation
To install TurkishCities, use go get:

```bash
go get github.com/nejdetkadir/turkish_cities
```

## Usage

### Loading Data
Load city data from a JSON file:

```go
package main

import (
    "fmt"
    "github.com/nejdetkadir/turkish_cities"
)

func main() {
    tc := turkish_cities.Load()

    firstCity := tc.GetCities()[0]
    firstTown := firstCity.GetTowns()[0]
    firstDistrict := firstTown.GetDistricts()[0]
    firstQuarter := firstDistrict.GetQuarters()[0]

    fmt.Println("total cities count: ", len(tc.GetCities()))
    // total cities count:  81

    fmt.Println("city id: ", firstCity.ID)
    // city id:  69

    fmt.Println("city name: ", firstCity.Name)
    // city name:  Bayburt

    fmt.Println("city latitude/longitude: ", firstCity.Location.Latitude, firstCity.Location.Longitude)
    // city latitude/longitude:  0x140000a7828 0x140000a7830

    fmt.Println("towns of city count: ", len(firstCity.GetTowns()))
    // towns of city count:  3

    fmt.Println("town id: ", firstTown.ID)
    // town id:  884

    fmt.Println("town name: ", firstTown.Name)
    // town name:  Aydıntepe

    fmt.Println("town latitude/longitude: ", firstTown.Location.Latitude, firstTown.Location.Longitude)
    // town latitude/longitude:  0x140000a6928 0x140000a6930

    fmt.Println("districts of town count: ", len(firstTown.GetDistricts()))
    // districts of town count:  2

    fmt.Println("district id: ", firstDistrict.ID)
    // district id:  4166

    fmt.Println("district name: ", firstDistrict.Name)
    // district name:  Aydıntepe

    fmt.Println("district latitude/longitude: ", firstDistrict.Location.Latitude, firstDistrict.Location.Longitude)
    // district latitude/longitude:  <nil> <nil>

    fmt.Println("quarters of district count: ", len(firstDistrict.GetQuarters()))
    // quarters of district count:  5

    fmt.Println("quarter id: ", firstQuarter.ID)
    // quarter id:  47917

    fmt.Println("quarter name: ", firstQuarter.Name)
    // quarter name:  Ahmet Yesevi Mah.

    fmt.Println("quarter latitude/longitude: ", firstQuarter.Location.Latitude, firstQuarter.Location.Longitude)
    // quarter latitude/longitude:  0x1400000e720 0x1400000e728
}
```

### Country Information
Access country-level information:

```go
package main

import (
    "fmt"
    "github.com/nejdetkadir/turkish_cities"
)

func main() {
    tc := turkish_cities.Load()

    country := tc.GetCountry()

    fmt.Println("country name:", country.Name)
    // country name: Türkiye
    
    fmt.Println("country a2c:", country.Alpha2Code)
    // country a2c: TR
    
    fmt.Println("country a3c:", country.Alpha3Code)
    // country a3c: TUR
    
    fmt.Println("country abbreviation:", country.Abbreviation)
    // country abbreviation: TR
    
    fmt.Println("country phone code:", country.PhoneCode)
    // country phone code: +90
}
```

### Querying Cities
Query cities by ID:

```go
package main

import (
    "fmt"
    "github.com/nejdetkadir/turkish_cities"
)

func main() {
    tc := turkish_cities.Load()
    
    foundCity := tc.GetCityByID(55)
    
    if foundCity != nil {
        fmt.Println("city found: ", foundCity.Name)
        // city found:  Samsun
    } else {
        fmt.Println("city not found")
    }
}
```

### Querying Towns
Query towns by ID:

```go
package main

import (
    "fmt"
    "github.com/nejdetkadir/turkish_cities"
)

func main() {
    tc := turkish_cities.Load()

    cityId := 55
    townId := 737

    foundTown := tc.GetTownByID(cityId, townId)
    
    if foundTown != nil {
        fmt.Println("town found: ", foundTown.Name)
        // town found:  Bafra
    } else {
        fmt.Println("town not found")
    }
}
```

### Querying Districts
Query districts by ID:

```go
package main

import (
    "fmt"
    "github.com/nejdetkadir/turkish_cities"
)

func main() {
    tc := turkish_cities.Load()

    cityId := 55
    townId := 737
    districtId := 5387

    foundDistrict := tc.GetDistrictByID(cityId, townId, districtId)

    if foundDistrict != nil {
        fmt.Println("district found: ", foundDistrict.Name)
        // district found:  Merkez
    } else {
        fmt.Println("district not found")
    }
}
```

### Querying Quarters
Query quarters by ID:

```go
package main

import (
    "fmt"
    "github.com/nejdetkadir/turkish_cities"
)

func main() {
    tc := turkish_cities.Load()

    cityId := 55
    townId := 737
    districtId := 5387
    quarterId := 37805

    foundQuarter := tc.GetQuarterByID(cityId, townId, districtId, quarterId)

    if foundQuarter != nil {
        fmt.Println("quarter found: ", foundQuarter.Name)
        // quarter found:  İshaklı Mh.
    } else {
        fmt.Println("quarter not found")
    }
}
```

## Testing
To run the tests, use go test:

```bash
$ go test ./...
```

## Contributing
Bug reports and pull requests are welcome on GitHub at https://github.com/nejdetkadir/turkish_cities. This project is intended to be a safe, welcoming space for collaboration, and contributors are expected to adhere to the [code of conduct](https://github.com/nejdetkadir/turkish_cities/blob/main/CODE_OF_CONDUCT.md).

## License
This project is licensed under the MIT License. See the [LICENSE](LICENSE) file for details.