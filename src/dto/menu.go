package dto

type MenuAddons struct {
	MenuId uint   `json:"menu_id"`
	Addons string `json:"addon"`
}

type Menu struct {
	Id          uint         `json:"id"`
	Title       string       `json:"title"`
	Description string       `json:"description"`
	Price       uint         `json:"price"`
	IsRecom     bool         `json:"is_recom"`
	ImageUrl    string       `json:"image_url"`
	RestId      uint         `json:"rest_id"`
	Addons      []MenuAddons `json:"addons"`
}

type GetRecommendMenuResponseService = []Menu

type GetRecommendMenuResponse = []Menu

type GetMenuResponseService struct {
	Menu
}

type GetMenuResponse struct {
	Menu
}

type GetMenusResponseService = []Menu

type GetMenusResponse = []Menu
