package model

type Building struct {
	Id             string
	Class          int    `json:"class"`
	NumberOfFloors int    `json:"number_of_floors"`
	Address        string `json:"address"`
}

type Room struct {
	Id            string `json:"id"`
	NumberOfRooms int    `json:"number_of_rooms"`
	RegularPrice  int    `json:"regular_price"`
	FloorId       string `json:"floor_id"`
}

type Complain struct {
	Id           string `json:"id"`
	ComplainText string `json:"complain_text"`
	GuestId      string `json:"guest_id,omitempty"`
}

type GuestAndComplain struct {
	Id          string `json:"id,omitempty"`
	FirstName   string `json:"first_name,omitempty"`
	LastName    string `json:"last_name,omitempty"`
	DateOfEntry string `json:"date_of_entry" json:"date_of_entry,omitempty"`
	Complain    `json:"complain"`
}
