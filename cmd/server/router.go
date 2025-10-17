package server

import (
	"orbit/cmd/server/handlers"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

/* -------------------------------------------------------------------------- */
/*                                   ROUTER                                   */
/* -------------------------------------------------------------------------- */
func SetupRouter(db *gorm.DB) *gin.Engine {
	r := gin.Default()
	api := r.Group("/api")
	{
		api.GET("/runs", handlers.ListRunsHandler(db))
		api.GET("/runs/:id", handlers.GetRunDetailHandler(db))
	}
	r.POST("/webhook", handlers.WebhookHandler(db))
	return r
}
