package dto

type CreateReviewRequest struct {
	Description string `json:"description" validate:"required"`
	Rating      uint32 `json:"rating" validate:"required,gte=1,lte=5"`
}

type CreateReviewRequestService struct {
	UserId       uint32 `json:"userId"`
	RestaurantId uint32 `json:"restaurantId"`
	Description  string `json:"description"`
	Rating       uint32 `json:"rating"`
}
