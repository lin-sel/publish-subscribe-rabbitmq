package model

// Offer have hotel, room, rateplan
type Offer struct {
	Hotel    Hotel    `json:"hotel"`
	Room     Room     `json:"room"`
	RatePlan RatePlan `json:"rate_plan"`
}
