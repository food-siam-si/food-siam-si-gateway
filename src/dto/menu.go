package dto

type MenuAddonsServiceData struct {
	MenuId uint   `json:"menu_id"`
	Addons string `json:"addons"`
}

type MenuServiceData struct {
	Id          uint                    `json:"id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Price       uint                    `json:"price"`
	IsRecom     bool                    `json:"is_recom"`
	ImageUrl    string                  `json:"image_url"`
	RestId      uint                    `json:"rest_id"`
	Addons      []MenuAddonsServiceData `json:"addons"`
}

type MenuAddons struct {
	MenuId uint   `json:"menuId"`
	Addons string `json:"label"`
}

type Menu struct {
	Id          uint                    `json:"id"`
	Title       string                  `json:"title"`
	Description string                  `json:"description"`
	Price       uint                    `json:"price"`
	IsRecom     bool                    `json:"isRecom"`
	ImageUrl    string                  `json:"imageUrl"`
	Addons      []MenuAddonsServiceData `json:"addons"`
}

type GetRecommendMenuResponseService struct {
	Menu []MenuServiceData `json:"menus"`
}

type GetRecommendMenuResponse = []Menu

type GetMenuResponseService struct {
	Menu MenuServiceData `json:"menus"`
}

type GetMenuResponse = Menu

type UpdateRecommendMenuRequestBody struct {
	IsRecom bool `json:"isRecom"`
}

type UpdateRecommendMenuRequestBodyService struct {
	IsRecom bool   `json:"is_recom"`
	UserId  uint32 `json:"user_id"`
	MenuId  uint32 `json:"menu_id"`
}

type GetMenusResponseService struct {
	Menu []MenuServiceData `json:"menus"`
}

type GetMenusResponse = []Menu

type CreateMenuRequestBody struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Price       uint32   `json:"price" validate:"required"`
	IsRecom     bool     `json:"isRecom"`
	ImageUrl    string   `json:"imageUrl" validate:"omitempty,url"`
	Addons      []string `json:"addons"`
}

type CreateMenuRequestBodyService struct {
	UserId      uint32   `json:"user_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       uint32   `json:"price"`
	IsRecom     bool     `json:"is_recom"`
	ImageUrl    string   `json:"image_url"`
	Addons      []string `json:"addons"`
}

type UpdateMenuRequestBody struct {
	Title       string   `json:"title" validate:"required"`
	Description string   `json:"description" validate:"required"`
	Price       uint32   `json:"price" validate:"required"`
	IsRecom     bool     `json:"isRecom"`
	ImageUrl    string   `json:"imageUrl" validate:"omitempty,url"`
	Addons      []string `json:"addons"`
}

type UpdateMenuRequestBodyService struct {
	UserId      uint32   `json:"user_id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Price       uint32   `json:"price"`
	IsRecom     bool     `json:"is_recom"`
	ImageUrl    string   `json:"image_url"`
	Addons      []string `json:"addons"`
	MenuId      uint32   `json:"menu_id"`
}
