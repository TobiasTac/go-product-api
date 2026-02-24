package handlers

import "github.com/TobiasTac/go-product-api/internal/infra/database"

type UserHandler struct {
	UserDB database.UserInterface
}

func NewUserHandler(userDB database.UserInterface) *UserHandler {
	return &UserHandler{UserDB: userDB}
}
