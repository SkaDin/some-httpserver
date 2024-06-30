package models

type Car struct {
	Id           uint64 `json:"id"`
	Owner        User   `json:"owner"`
	Colour       string `json:"colour"`
	Brand        string `json:"brand"`
	LicencePlate string `json:"licence_plate"`
}
