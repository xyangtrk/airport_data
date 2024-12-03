package models

type Airline struct {
	ID       string `json:"id"`
	Name     string `json:"name"`
	Alias    string `json:"alias"`
	IATA     string `json:"iata"`
	ICAO     string `json:"icao"`
	Callsign string `json:"callsign"`
	Country  string `json:"country"`
	Active   string `json:"active"`
}
