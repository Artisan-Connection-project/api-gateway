package models

type AddRatingRequest struct {
	UserId  string  `json:"user_id,omitempty"`
	Rating  float64 `json:"rating,omitempty"`
	Comment string  `json:"comment,omitempty"`
}
