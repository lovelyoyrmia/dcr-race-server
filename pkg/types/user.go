package types

type UserLocation struct {
	Uid        string   `json:"uid"`
	Email     string  `json:"email"`
	Fullname  string  `json:"fullname"`
	Latitude  float64 `json:"latitude"`
	Longitude float64 `json:"longitude"`
	Altitude  float64 `json:"altitude"`
	Category  string  `json:"category"`
	Timestamp string  `json:"timestamp"`
}

type UserLocationParams struct {
	Uid     string  `json:"uid" form:"uid"`
	Email     string  `json:"email" form:"email"`
	Fullname  string  `json:"fullname" form:"fullname"`
	Latitude  float64 `json:"latitude" form:"latitude"`
	Longitude float64 `json:"longitude" form:"longitude"`
	Altitude  float64 `json:"altitude" form:"altitude"`
	Category  string  `json:"category" form:"category"`
}

type ListUserLocationParams struct {
	Category *string `query:"category"`
	Count    *int64  `query:"count"`
}
