package models

type User struct {
	Id   uint64 `json:"id"`
	Name string `json:"name"`
	Rank string `json:"rank"`
}
