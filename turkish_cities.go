package turkish_cities

import (
	"encoding/json"
	"io"
	"os"
)

type (
	TurkishCities struct {
		Cities  []City
		Country Country
	}
	Context interface {
		GetCities() []City
		GetCityByID(id int) *City
		GetTownByID(cityID, townID int) *Town
		GetDistrictByID(cityID, townID, districtID int) *District
		GetQuarterByID(cityID, townID, districtID, quarterID int) *Quarter
		GetCountry() Country
	}
	BaseEntity struct {
		ID       int       `json:"id"`
		Name     string    `json:"name"`
		Location *Location `json:"location"`
	}
	Location struct {
		Latitude  *float64 `json:"lat"`
		Longitude *float64 `json:"lon"`
	}
	City struct {
		BaseEntity
		Towns []Town `json:"townsData"`
	}
	Town struct {
		BaseEntity
		Districts []District `json:"districtsData"`
	}
	District struct {
		BaseEntity
		Quarters []Quarter `json:"quarters"`
	}
	Quarter struct {
		BaseEntity
	}
	Country struct {
		Name         string
		PhoneCode    string
		Alpha2Code   string
		Alpha3Code   string
		Abbreviation string
	}
)

func Load() Context {
	rawJsonFile, err := os.Open("./static/data.json")

	if err != nil {
		panic(err)
	}

	defer func(rawJsonFile *os.File) {
		err := rawJsonFile.Close()
		if err != nil {
			panic(err)
		}
	}(rawJsonFile)

	byteValue, err := io.ReadAll(rawJsonFile)

	if err != nil {
		panic(err)
	}

	var cities []City

	err = json.Unmarshal(byteValue, &cities)

	if err != nil {
		panic(err)
	}

	return &TurkishCities{
		Cities: cities,
		Country: Country{
			Name:         "Türkiye",
			PhoneCode:    "+90",
			Alpha2Code:   "TR",
			Alpha3Code:   "TUR",
			Abbreviation: "TR",
		},
	}
}

func (t *TurkishCities) GetCities() []City {
	return t.Cities
}

func (t *TurkishCities) GetCityByID(id int) *City {
	for _, city := range t.Cities {
		if city.ID == id {
			return &city
		}
	}

	return nil
}

func (t *TurkishCities) GetTownByID(cityID, townID int) *Town {
	city := t.GetCityByID(cityID)

	if city == nil {
		return nil
	}

	for _, town := range city.Towns {
		if town.ID == townID {
			return &town
		}
	}

	return nil
}

func (t *TurkishCities) GetDistrictByID(cityID, townID, districtID int) *District {
	town := t.GetTownByID(cityID, townID)

	if town == nil {
		return nil
	}

	for _, district := range town.Districts {
		if district.ID == districtID {
			return &district
		}
	}

	return nil
}

func (t *TurkishCities) GetQuarterByID(cityID, townID, districtID, quarterID int) *Quarter {
	district := t.GetDistrictByID(cityID, townID, districtID)

	if district == nil {
		return nil
	}

	for _, quarter := range district.Quarters {
		if quarter.ID == quarterID {
			return &quarter
		}
	}

	return nil
}

func (t *TurkishCities) GetCountry() Country {
	return t.Country
}

func (c *City) GetTowns() []Town {
	return c.Towns
}

func (t *Town) GetDistricts() []District {
	return t.Districts
}

func (d *District) GetQuarters() []Quarter {
	return d.Quarters
}
