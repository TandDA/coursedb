package model

type Building struct {
	Id             string
	Class          int    `json:"class"`
	NumberOfFloors int    `json:"number_of_floors"`
	Address        string `json:"address"`
}

type Room struct {
	Id             string  `json:"id"`
	NumberOfRooms          int    `json:"number_of_rooms"`
	RegularPrice int    `json:"regular_price"`
	FloorId        string `json:"floor_id"`
}