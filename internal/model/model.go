package model

type Building struct {
	Id             string
	Class          int    `json:"class"`
	NumberOfFloors int    `json:"number_of_floors"`
	Address        string `json:"address"`
}
