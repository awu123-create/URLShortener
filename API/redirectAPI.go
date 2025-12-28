package API

import (
	"URLShortener/Service"
	"net/http"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Redirect(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		shortCode := c.Param("short_code")

		url, err := Service.GetAndIncreasesVisitCount(db, shortCode)

		if err != nil {
			if err.Error() == "Short URL not found" {
				c.JSON(http.StatusNotFound, gin.H{
					"error": "Short URL not found",
				})
			} else {
				c.JSON(http.StatusInternalServerError, gin.H{
					"error": "Database error",
				})
			}
			return
		}

		c.Redirect(http.StatusFound, url.OriginalURL)
	}
}
