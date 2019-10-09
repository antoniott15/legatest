package main

type Investor struct {
	ID                string `json:"id"`
	Name              string `json:"name"`
	LastName          string `json:"last_name"`
	Gender            string `json:"gender"`
	Creationdate      string `json:"creation_date"`
	Email             string `json:"email"`
	Phone             string `json:"phone"`
	TotalInvest       string `json:"total_invest"`
	Country           string `json:"country"`
	NumberTransaction string `json:"number_of_transaction"`
	Birthday          string `json:"birthday"`
}

type Project struct {
	ID           string `json:"id"`
	Location     string `json:"location"`
	Name         string `json:"name"`
	Status       string `json:"status"`
	TotalActions string `json:"totalactions"`
	CreationDate string `json:"creationdate"`
}
