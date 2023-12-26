package apis

import (
	"hanmantpatil/gorilla/entity"
	"hanmantpatil/gorilla/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
)

func (handler *Handler) registerUsersRoutes() {
	users := handler.api.Group("/users")

	users.Use(handler.jwtAuthMiddleware())
	users.POST("", tonic.Handler(handler.handleCreateUser, http.StatusCreated))
}

type CreateUserInput struct {
	FirstName string `json:"first_name" validate:"required"`
	LastName  string `json:"last_name" validate:"required"`
}

func (handler *Handler) handleCreateUser(c *gin.Context, input *CreateUserInput) (*entity.User, error) {
	return handler.users.CreateUser(usecase.NewContext(c), input.FirstName, input.LastName)
}
