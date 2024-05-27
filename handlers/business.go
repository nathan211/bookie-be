package handlers

import (
    "bookie-be/db"
    "bookie-be/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

// CreateBusiness creates a new business
func CreateBusiness(c *gin.Context) {
    var business models.Business
    if err := c.ShouldBindJSON(&business); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.MustGet("userID").(uint)
    business.UserID = userID

    if err := db.DB.Create(&business).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create business"})
        return
    }

    c.JSON(http.StatusOK, business)
}

// GetBusinesses fetches all businesses
func GetBusinesses(c *gin.Context) {
    var businesses []models.Business
    userID := c.MustGet("userID").(uint)

    if err := db.DB.Where("user_id = ?", userID).Find(&businesses).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch businesses"})
        return
    }

    c.JSON(http.StatusOK, businesses)
}

// UpdateBusiness updates an existing business
func UpdateBusiness(c *gin.Context) {
    var business models.Business
    if err := c.ShouldBindJSON(&business); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.MustGet("userID").(uint)
    if err := db.DB.Where("id = ? AND user_id = ?", business.ID, userID).First(&business).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
        return
    }

    if err := db.DB.Save(&business).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update business"})
        return
    }

    c.JSON(http.StatusOK, business)
}

// DeleteBusiness deletes a business
func DeleteBusiness(c *gin.Context) {
    var business models.Business
    if err := c.ShouldBindJSON(&business); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    userID := c.MustGet("userID").(uint)
    if err := db.DB.Where("id = ? AND user_id = ?", business.ID, userID).First(&business).Error; err != nil {
        c.JSON(http.StatusNotFound, gin.H{"error": "Business not found"})
        return
    }

    if err := db.DB.Delete(&business).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete business"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Business deleted successfully"})
}