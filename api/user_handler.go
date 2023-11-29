package api

import (
	"context"

	"github.com/gofiber/fiber/v2"
	"github.com/nahaktarun/hotel-reservation/db"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUsers(c *fiber.Ctx) error {
	return c.JSON("James")
}
func (h *UserHandler) HandleGetUser(c *fiber.Ctx) error {

	id := c.Params("id")
	ctx := context.Background()
	user, err := h.userStore.GetUserByID(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)

}
