package Routes

import (
	"URLShortener/API"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func RegisterRoutes(r *gin.Engine, db *gorm.DB) {
	r.POST("/api/shorten", API.CreateShortURL(db))
	r.GET("/:short_code", API.Redirect(db))
}
