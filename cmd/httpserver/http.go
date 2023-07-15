package httpserver

import (
	"fmt"
	"net/http"
	"pi/config"
	"pi/internal/dto"
	"pi/internal/handlers/userhdl"
	"pi/pkg/meta"
	"strings"

	_ "pi/docs"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	echoSwagger "github.com/swaggo/echo-swagger"
)

func NewHTTPServer(
	uh *userhdl.Handler,
) *echo.Echo {
	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	// allow cors
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"*"},
		AllowMethods: []string{"*"},
	}))
	e.HTTPErrorHandler = customHTTPErrorHandler

	// auth middleware if needed
	// authMiddleware := echojwt.WithConfig(echojwt.Config{
	// 	SigningKey: []byte(config.Get().JWT.Secret),
	// 	NewClaimsFunc: func(c echo.Context) jwt.Claims {
	// 		return new(auth.JWTCustomClaims)
	// 	},
	// })

	// swagger
	if config.Get().App.EnableSwagger {
		e.GET("/swagger/*", echoSwagger.WrapHandler)
		fmt.Printf("Enable swager local: http://localhost:%s/swagger/index.html \n", config.Get().Endpoint.Port)
	}

	e.GET("/health", func(c echo.Context) error {
		return c.String(http.StatusOK, "server is running...")
	})

	// add prefix for all routes below
	v1 := e.Group("/api/v1")
	// user
	user := v1.Group("/user")
	user.POST("", uh.CreateUser)
	user.GET("/:id", uh.GetUserByID)
	user.PUT("/:id", uh.UpdateUser)
	user.DELETE("/:id", uh.DeleteUser)

	return e
}

func customHTTPErrorHandler(err error, c echo.Context) {
	var m *meta.MetaError

	if metaErr, ok := meta.IsError(err); ok {
		m = metaErr
	} else if he, ok := err.(*echo.HTTPError); ok {
		m = meta.NewError(he.Code).AppendMessage(1000, strings.ToLower(he.Message.(string)))
	} else {
		m = meta.MetaErrorInternalServer.AppendError(1000, err)
	}

	c.JSON(m.HttpStatus, dto.BaseErrorResponse{
		BaseResponse: dto.BaseResponse{
			Code: m.Code,
		},
		Message: m.Message,
	})
}
