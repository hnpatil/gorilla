package apis

import (
	"errors"
	"hanmantpatil/gorilla/entity"
	"hanmantpatil/gorilla/usecase"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
)

func tonicErrorHook(ctx *gin.Context, err error) (int, interface{}) {
	bindErr := &tonic.BindError{}
	if errors.As(err, bindErr) {
		return http.StatusBadRequest, &APIError{Error: bindErr.Error()}
	}

	if entity.IsNotFound(err) {
		return http.StatusNotFound, &APIError{Error: err.Error()}
	}

	if entity.IsConstraintError(err) {
		return http.StatusBadRequest, &APIError{Error: err.Error()}
	}

	if ue := usecase.GetUsecaseError(err); ue != nil {
		return http.StatusBadRequest, &APIError{Error: err.Error(), ErrorCode: ue.ErrorCode}
	}

	return http.StatusInternalServerError, &APIError{Error: err.Error()}
}

func getUserID(ctx *gin.Context) string {
	usr, ok := ctx.Get("identifier")
	if !ok {
		return ""
	}

	return usr.(string)
}

func abortUnauthenticated(ctx *gin.Context) {
	ctx.JSON(http.StatusUnauthorized, &APIError{Error: "Unauthorized"})
	ctx.Abort()
}

func extractToken(ctx *gin.Context) string {
	bearerToken := ctx.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}

	return ""
}
