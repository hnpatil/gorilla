package apis

import (
	"hanmantpatil/gorilla/entity"
	"hanmantpatil/gorilla/usecase"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
)

func (handler *Handler) registerBooksRoutes() {
	booksCreate := handler.api.Group("/books")
	booksCreate.Use(handler.jwtAuthMiddleware())
	booksCreate.POST("", tonic.Handler(handler.handleCreateBook, http.StatusOK))

	books := handler.api.Group("/books")

	books.GET("", tonic.Handler(handler.handleListBooks, http.StatusOK))
	books.GET("/:book_id", tonic.Handler(handler.handleGetBook, http.StatusOK))
	books.DELETE("/:book_id", tonic.Handler(handler.handleDeleteBook, http.StatusNoContent))
}

func (handler *Handler) handleListBooks(c *gin.Context) ([]*entity.Book, error) {
	return handler.books.ListBooks(usecase.NewContext(c))
}

type BookIDInput struct {
	ID int `path:"book_id" validate:"required"`
}

func (handler *Handler) handleGetBook(c *gin.Context, input *BookIDInput) (*entity.Book, error) {
	return handler.books.GetBook(usecase.NewContext(c), input.ID)
}

func (handler *Handler) handleDeleteBook(c *gin.Context, input *BookIDInput) error {
	return handler.books.DeleteBook(usecase.NewContext(c), input.ID)
}

type CreateBookInput struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description" validate:"required"`
}

func (handler *Handler) handleCreateBook(c *gin.Context, input *CreateBookInput) (*entity.Book, error) {
	return handler.books.AddBook(usecase.NewContext(c), input.Title, input.Description)
}
