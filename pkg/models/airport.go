package models

// input
type Airport struct {
	FS                 string  `json:"fs"`
	IATA               string  `json:"iata"`
	ICAO               string  `json:"icao"`
	FAA                string  `json:"faa"`
	Name               string  `json:"name"`
	City               string  `json:"city"`
	StateCode          string  `json:"stateCode"`
	CountryCode        string  `json:"countryCode"`
	CountryName        string  `json:"countryName"`
	RegionName         string  `json:"regionName"`
	TimeZoneRegionName string  `json:"timeZoneRegionName"`
	WeatherZone        string  `json:"weatherZone"`
	LocalTime          string  `json:"localTime"`
	UtcOffsetHours     float64 `json:"utcOffsetHours"`
	Latitude           float64 `json:"latitude"`
	Longitude          float64 `json:"longitude"`
	ElevationFeet      int     `json:"elevationFeet"`
	Classification     int     `json:"classification"`
	Active             bool    `json:"active"`
	WeatherUrl         string  `json:"weatherUrl"`
	DelayIndexUrl      string  `json:"delayIndexUrl"`
}

type AirportsData struct {
	Airports []Airport `json:"airports"`
}

// output
type OutputAirport struct {
	UTCOffset          float64 `json:"utcOffset"`
	Latitude           float64 `json:"lat"`
	Longitude          float64 `json:"lon"`
	TimeZoneRegionName string  `json:"timeZoneRegionName"`
}
