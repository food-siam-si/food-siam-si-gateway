package dto

type AveragePrice string

const (
	LowerThanHundreds       AveragePrice = "LowerThanHundreds"
	HundredToTwoHundred     AveragePrice = "HundredToTwoHundred"
	TwoHundredToFiveHundred AveragePrice = "TwoHundredToFiveHundred"
	MoreThanFiveHundred     AveragePrice = "MoreThanFiveHundred"
	MoreThanOneThousand     AveragePrice = "MoreThanOneThousand"
)

type CreateRestaurantRequest struct {
	Name              string       `json:"name" validate:"required"`
	Description       string       `json:"description" validate:"required"`
	PhoneNumber       string       `json:"phoneNumber" validate:"required,phoneNumber"`
	LocationLat       float32      `json:"locationLat" validate:"required,latitude"`
	LocationLong      float32      `json:"locationLong" validate:"required,longitude"`
	AveragePrice      AveragePrice `json:"averagePrice" validate:"required"`
	ImageUrl          string       `json:"imageUrl" validate:"omitempty,url"`
	RestaurantTypeIds []uint32     `json:"restaurantTypeIds" validate:"required"`
}

type UpdateRestaurantRequest struct {
	Name              string       `json:"name"`
	Description       string       `json:"description"`
	PhoneNumber       string       `json:"phoneNumber" validate:"omitempty,phoneNumber"`
	LocationLat       float32      `json:"locationLat" validate:"omitempty,latitude"`
	LocationLong      float32      `json:"locationLong" validate:"omitempty,longitude"`
	AveragePrice      AveragePrice `json:"averagePrice"`
	ImageUrl          string       `json:"imageUrl" validate:"omitempty,url"`
	RestaurantTypeIds []uint32     `json:"restaurantTypeIds"`
	IsInService       bool         `json:"isInService"`
}

type RestaurantType struct {
	Id   uint32 `json:"id"`
	Name string `json:"name"`
}

type Restaurant struct {
	Id             uint32
	Name           string           `json:"name"`
	Description    string           `json:"description"`
	PhoneNumber    string           `json:"phoneNumber"`
	LocationLat    float32          `json:"locationLat"`
	LocationLong   float32          `json:"locationLong"`
	AveragePrice   AveragePrice     `json:"averagePrice"`
	ImageUrl       string           `json:"imageUrl"`
	RestaurantType []RestaurantType `json:"restaurantType"`
}
