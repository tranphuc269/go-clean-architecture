package http

import (
	"github.com/labstack/echo/v4"

	"go-clean-architecture/internal/middleware"
	"go-clean-architecture/internal/news"
)

// MapNewsRoutes routes
func MapNewsRoutes(newsGroup *echo.Group, h news.Handlers, mw *middleware.MiddlewareManager) {
	newsGroup.POST("/create", h.Create(), mw.AuthSessionMiddleware, mw.CSRF)
	newsGroup.PUT("/:news_id", h.Update(), mw.AuthSessionMiddleware, mw.CSRF)
	newsGroup.DELETE("/:news_id", h.Delete(), mw.AuthSessionMiddleware, mw.CSRF)
	newsGroup.GET("/:news_id", h.GetByID())
	newsGroup.GET("/search", h.SearchByTitle())
	newsGroup.GET("", h.GetNews())
}
