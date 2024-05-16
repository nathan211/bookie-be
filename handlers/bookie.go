package handlers

import (
	"bookie-be/db"
	"bookie-be/models"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func CreateBookie(c *gin.Context) {
	var bookie models.Bookie
	if err := c.ShouldBindJSON(&bookie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	bookie.UserID = userID
	bookie.Status = "new"
	bookie.BookedTime = time.Now()

	if err := db.DB.Create(&bookie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create bookie"})
		return
	}

	c.JSON(http.StatusOK, bookie)
}

func GetBookies(c *gin.Context) {
	var bookies []models.Bookie
	userID := c.MustGet("userID").(uint)

	if err := db.DB.Where("user_id = ?", userID).Find(&bookies).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch bookies"})
		return
	}

	c.JSON(http.StatusOK, bookies)
}

func UpdateBookie(c *gin.Context) {
	var bookie models.Bookie
	if err := c.ShouldBindJSON(&bookie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	if err := db.DB.Where("id = ? AND user_id = ?", bookie.ID, userID).First(&bookie).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bookie not found"})
		return
	}

	if err := db.DB.Save(&bookie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update bookie"})
		return
	}

	c.JSON(http.StatusOK, bookie)
}

func DeleteBookie(c *gin.Context) {
	var bookie models.Bookie
	if err := c.ShouldBindJSON(&bookie); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID := c.MustGet("userID").(uint)
	if err := db.DB.Where("id = ? AND user_id = ?", bookie.ID, userID).First(&bookie).Error; err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Bookie not found"})
		return
	}

	if err := db.DB.Delete(&bookie).Error; err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete bookie"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Bookie deleted successfully"})
}
