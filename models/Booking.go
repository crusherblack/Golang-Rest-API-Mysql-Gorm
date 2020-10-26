package models

type Booking struct {
	Id      int    `json:"id"`
	User    string `json:"user"`
	Members string `json:"members"`
}
