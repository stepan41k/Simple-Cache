package models

type User struct {
	ID   int    `json:"id"`
	Name string `json:"name" redis:"name"`
	Data string `json:"data" redis:"data"`
}