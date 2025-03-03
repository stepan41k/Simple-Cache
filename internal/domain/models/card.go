package models

type Card struct {
	Name string `json:"name" redis:"name"`
	Data string `json:"data" redis:"data"`
}
