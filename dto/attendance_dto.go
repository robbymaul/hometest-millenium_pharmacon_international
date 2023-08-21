package dto

type ClockInRequstDTO struct {
	UserID       int    `json:"userId" form:"userId"`
	ClockInImage string `json:"clockInImage" form:"clockInImage"`
}

type ClockOutRequestDTO struct {
	ClockOutImage string `json:"clockOutImage" from:"clockOutImage"`
}
