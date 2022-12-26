package models

type Debt struct {
	Id     int `json:"id"`
	UserId int `json:"user_id"`
	Value  int `json:"value"`
}
