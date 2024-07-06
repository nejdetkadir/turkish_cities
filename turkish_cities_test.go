package turkish_cities

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestLoad(t *testing.T) {
	t.Run("should load the data", func(t *testing.T) {
		tc := Load()

		assert.NotEmpty(t, tc.GetCities())
		assert.Equal(t, 81, len(tc.GetCities()))
	})
}

func TestTurkishCities_GetCities(t *testing.T) {
	t.Run("should return cities", func(t *testing.T) {
		tc := TurkishCities{
			Cities: []City{
				{
					BaseEntity: BaseEntity{
						ID:   1,
						Name: "Adana",
					},
				},
				{
					BaseEntity: BaseEntity{
						ID:   2,
						Name: "Adıyaman",
					},
				},
			},
		}

		assert.Equal(t, 2, len(tc.GetCities()))
	})

	t.Run("should return empty cities", func(t *testing.T) {
		tc := TurkishCities{}

		assert.Empty(t, tc.GetCities())
	})

	t.Run("should return loaded cities", func(t *testing.T) {
		tc := Load()

		assert.NotEmpty(t, tc.GetCities())
	})
}

func TestCity_GetTowns(t *testing.T) {
	t.Run("should return towns", func(t *testing.T) {
		c := City{
			Towns: []Town{
				{
					BaseEntity: BaseEntity{
						ID:   1,
						Name: "Seyhan",
					},
				},
				{
					BaseEntity: BaseEntity{
						ID:   2,
						Name: "Yüreğir",
					},
				},
			},
		}

		assert.Equal(t, 2, len(c.GetTowns()))
	})

	t.Run("should return empty towns", func(t *testing.T) {
		c := City{}

		assert.Empty(t, c.GetTowns())
	})

	t.Run("should return loaded towns", func(t *testing.T) {
		tc := Load()
		c := tc.GetCities()[0]

		assert.NotEmpty(t, c.GetTowns())
	})
}

func TestTown_GetDistricts(t *testing.T) {
	t.Run("should return districts", func(t *testing.T) {
		town := Town{
			Districts: []District{
				{
					BaseEntity: BaseEntity{
						ID:   1,
						Name: "Seyhan",
					},
				},
				{
					BaseEntity: BaseEntity{
						ID:   2,
						Name: "Yüreğir",
					},
				},
			},
		}

		assert.Equal(t, 2, len(town.GetDistricts()))
	})

	t.Run("should return empty districts", func(t *testing.T) {
		town := Town{}

		assert.Empty(t, town.GetDistricts())
	})

	t.Run("should return loaded districts", func(t *testing.T) {
		tc := Load()
		c := tc.GetCities()[0]
		town := c.GetTowns()[0]

		assert.NotEmpty(t, town.GetDistricts())
	})
}

func TestDistrict_GetQuarters(t *testing.T) {
	t.Run("should return quarters", func(t *testing.T) {
		district := District{
			Quarters: []Quarter{
				{
					BaseEntity: BaseEntity{
						ID:   1,
						Name: "Seyhan",
					},
				},
				{
					BaseEntity: BaseEntity{
						ID:   2,
						Name: "Yüreğir",
					},
				},
			},
		}

		assert.Equal(t, 2, len(district.GetQuarters()))
	})

	t.Run("should return empty quarters", func(t *testing.T) {
		district := District{}

		assert.Empty(t, district.GetQuarters())
	})

	t.Run("should return loaded quarters", func(t *testing.T) {
		tc := Load()
		c := tc.GetCities()[0]
		town := c.GetTowns()[0]
		district := town.GetDistricts()[0]

		assert.NotEmpty(t, district.GetQuarters())
	})
}

func TestTurkishCities_GetCountry(t *testing.T) {
	t.Run("should return country", func(t *testing.T) {
		tc := Load()

		assert.NotEmpty(t, tc.GetCountry())
		assert.Equal(t, "Türkiye", tc.GetCountry().Name)
		assert.Equal(t, "+90", tc.GetCountry().PhoneCode)
		assert.Equal(t, "TR", tc.GetCountry().Alpha2Code)
		assert.Equal(t, "TUR", tc.GetCountry().Alpha3Code)
		assert.Equal(t, "TR", tc.GetCountry().Abbreviation)
	})
}

func TestTurkishCities_GetCityByID(t *testing.T) {
	t.Run("should return city by id", func(t *testing.T) {
		tc := Load()

		assert.NotNil(t, tc.GetCityByID(1))
		assert.Equal(t, "Adana", tc.GetCityByID(1).Name)
	})

	t.Run("should return nil for unknown city id", func(t *testing.T) {
		tc := Load()

		assert.Nil(t, tc.GetCityByID(100))
	})
}

func TestTurkishCities_GetTownByID(t *testing.T) {
	t.Run("should return town by id", func(t *testing.T) {
		tc := Load()

		randomCity := tc.GetCities()[0]
		randomTown := randomCity.GetTowns()[0]

		result := tc.GetTownByID(randomCity.ID, randomTown.ID)

		assert.NotNil(t, result)
		assert.Equal(t, randomTown.Name, result.Name)
	})

	t.Run("should return nil for unknown town id", func(t *testing.T) {
		tc := Load()

		result := tc.GetTownByID(1, 999)

		assert.Nil(t, result)
	})
}

func TestTurkishCities_GetDistrictByID(t *testing.T) {
	t.Run("should return district by id", func(t *testing.T) {
		tc := Load()

		randomCity := tc.GetCities()[0]
		randomTown := randomCity.GetTowns()[0]
		randomDistrict := randomTown.GetDistricts()[0]

		result := tc.GetDistrictByID(randomCity.ID, randomTown.ID, randomDistrict.ID)

		assert.NotNil(t, result)
		assert.Equal(t, randomDistrict.Name, result.Name)
	})

	t.Run("should return nil for unknown district id", func(t *testing.T) {
		tc := Load()

		result := tc.GetDistrictByID(1, 1, 999)

		assert.Nil(t, result)
	})
}

func TestTurkishCities_GetQuarterByID(t *testing.T) {
	t.Run("should return quarter by id", func(t *testing.T) {
		tc := Load()

		randomCity := tc.GetCities()[0]
		randomTown := randomCity.GetTowns()[0]
		randomDistrict := randomTown.GetDistricts()[0]
		randomQuarter := randomDistrict.GetQuarters()[0]

		result := tc.GetQuarterByID(randomCity.ID, randomTown.ID, randomDistrict.ID, randomQuarter.ID)

		assert.NotNil(t, result)
		assert.Equal(t, randomQuarter.Name, result.Name)
	})

	t.Run("should return nil for unknown quarter id", func(t *testing.T) {
		tc := Load()

		result := tc.GetQuarterByID(1, 1, 1, 999)

		assert.Nil(t, result)
	})
}
