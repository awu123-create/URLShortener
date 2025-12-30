package API

import (
	"URLShortener/Models"
	"URLShortener/Utils"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

// 此API用于创建短链接
func CreateShortURL(db *gorm.DB) gin.HandlerFunc {
	return func(c *gin.Context) {
		var requestBody struct {
			URL string `json:"url"`
		}

		if err := c.ShouldBindJSON(&requestBody); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid request",
			})
			return
		}

		// 校验 URL 格式
		if !strings.HasPrefix(requestBody.URL, "http://") && !strings.HasPrefix(requestBody.URL, "https://") {
			c.JSON(http.StatusBadRequest, gin.H{
				"error": "Invalid URL format",
			})
			return
		}

		url := Models.URL{
			OriginalURL: requestBody.URL,
		}
		if err := db.Create(&url).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error",
			})
			return
		}

		shortCode := Utils.EncodeBase62(url.ID)
		if err := db.Model(&url).Update("short_code", shortCode).Error; err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": "Database error",
			})
			return
		}
		c.JSON(http.StatusOK, gin.H{
			"short_url": os.Getenv("BASE_URL") + "/" + shortCode,
		})
	}
}
