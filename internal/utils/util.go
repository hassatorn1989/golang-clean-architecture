package utils

import "github.com/gofiber/fiber/v3"

func GetPaginationParams(c fiber.Ctx) (int, int) {
	page := fiber.Query[int](c, "page", 1) // default = 1
	if page < 1 {
		page = 1
	}
	limit := fiber.Query[int](c, "limit", 10)
	if limit <= 0 {
		limit = 10
	}
	offset := (page - 1) * limit

	return limit, offset
}
