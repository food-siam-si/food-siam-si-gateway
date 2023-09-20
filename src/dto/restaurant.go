package dto

type CreateRestaurantRequest struct {
	Name              string   `json:"name" validate:"required"`
	Description       string   `json:"description" validate:"required"`
	PhoneNumber       string   `json:"phoneNumber" validate:"required,phoneNumber"`
	LocationLat       float32  `json:"locationLat" validate:"required,latitude"`
	LocationLong      float32  `json:"locationLong" validate:"required,longitude"`
	AveragePrice      string   `json:"averagePrice" validate:"required"`
	ImageUrl          string   `json:"imageUrl" validate:"omitempty,url"`
	RestaurantTypeIds []uint32 `json:"restaurantTypeIds" validate:"required"`
}

type UpdateRestaurantRequest struct {
	Name              string   `json:"name"`
	Description       string   `json:"description"`
	PhoneNumber       string   `json:"phoneNumber" validate:"omitempty,phoneNumber"`
	LocationLat       float32  `json:"locationLat" validate:"omitempty,latitude"`
	LocationLong      float32  `json:"locationLong" validate:"omitempty,longitude"`
	AveragePrice      string   `json:"averagePrice"`
	ImageUrl          string   `json:"imageUrl" validate:"omitempty,url"`
	RestaurantTypeIds []uint32 `json:"restaurantTypeIds"`
}
