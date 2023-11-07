package dto

type CreateReviewRequest struct {
	Description string  `json:"description" validate:"required"`
	Rating      float32 `json:"rating" validate:"required,gte=1,lte=5"`
}

type CreateReviewRequestService struct {
	UserId       uint32  `json:"userId"`
	RestaurantId uint32  `json:"restaurantId"`
	Description  string  `json:"description"`
	Rating       float32 `json:"rating"`
}

type RestaurantReview struct {
	Id          string  `json:"id"`
	Description string  `json:"description"`
	Rating      float32 `json:"rate"`
	WrittenDate string  `json:"writtenDate"`
}

type RestaurantReviewResponse = []RestaurantReview
