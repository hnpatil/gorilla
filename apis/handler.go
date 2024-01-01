package apis

import (
	"fmt"
	"hanmantpatil/gorilla/config"
	"hanmantpatil/gorilla/usecase"

	"github.com/gin-gonic/gin"
	"github.com/loopfz/gadgeto/tonic"
	"github.com/sirupsen/logrus"
	ginlogrus "github.com/toorop/gin-logrus"
)

type Handler struct {
	api    *gin.RouterGroup
	auth   usecase.Auth
	books  usecase.Books
	users  usecase.Users
	engine *gin.Engine
}

func NewHander(cfg *config.Config, books usecase.Books, users usecase.Users, auth usecase.Auth, logger logrus.FieldLogger) *Handler {
	tonic.SetErrorHook(tonicErrorHook)

	router := gin.New()
	router.Use(loggerMiddleware(logger))
	router.Use(ginlogrus.Logger(logger), gin.Recovery())

	api := router.Group(fmt.Sprintf("/api/%s", cfg.GetValue(config.VERSION)))

	return &Handler{
		api:    api,
		auth:   auth,
		books:  books,
		users:  users,
		engine: router,
	}
}

func (hanlder *Handler) RegisterRoutes() *Handler {
	hanlder.registerPingRoutes()
	hanlder.registerBooksRoutes()
	hanlder.registerUsersRoutes()
	hanlder.registerAuthRoutes()

	return hanlder
}

func (handler *Handler) Start(address string) {
	if err := handler.engine.Run(address); err != nil {
		panic(err.Error())
	}
}

type APIError struct {
	Error     string `json:"error"`
	ErrorCode int    `json:"error_code"`
}

func (handler *Handler) jwtAuthMiddleware() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		bearerToken := extractToken(ctx)
		if bearerToken == "" {
			abortUnauthenticated(ctx)

			return
		}

		usrID, err := handler.auth.Authenticate(usecase.NewContext(ctx), bearerToken)
		if err != nil {
			abortUnauthenticated(ctx)

			return
		}

		ctx.Set("identifier", usrID)
		ctx.Next()
	}
}

func loggerMiddleware(logger logrus.FieldLogger) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		lg := logger.WithFields(logrus.Fields{
			"path":      ctx.Request.URL.Path,
			"method":    ctx.Request.Method,
			"hostname":  ctx.Request.Host,
			"client_ip": ctx.ClientIP(),
		})

		ctx.Set("logger", lg)
		ctx.Next()
	}
}
