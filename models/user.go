package models

type User struct {
	Id    int    `json:"id"`
	Age   int    `json:"age"`
	Name  string `json:"name"`
	Debts []Debt `json:"debts"`
}
