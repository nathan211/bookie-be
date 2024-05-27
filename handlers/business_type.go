package handlers

import (
    "bookie-be/db"
    "bookie-be/models"
    "net/http"

    "github.com/gin-gonic/gin"
)

// CreateBusinessType creates a new business type
func CreateBusinessType(c *gin.Context) {
    var businessType models.BusinessType
    if err := c.ShouldBindJSON(&businessType); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Create(&businessType).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create business type"})
        return
    }

    c.JSON(http.StatusOK, businessType)
}

// GetBusinessTypes fetches all business types
func GetBusinessTypes(c *gin.Context) {
    var businessTypes []models.BusinessType
    if err := db.DB.Find(&businessTypes).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch business types"})
        return
    }

    c.JSON(http.StatusOK, businessTypes)
}

// UpdateBusinessType updates an existing business type
func UpdateBusinessType(c *gin.Context) {
    var businessType models.BusinessType
    if err := c.ShouldBindJSON(&businessType); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Save(&businessType).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update business type"})
        return
    }

    c.JSON(http.StatusOK, businessType)
}

// DeleteBusinessType deletes a business type
func DeleteBusinessType(c *gin.Context) {
    var businessType models.BusinessType
    if err := c.ShouldBindJSON(&businessType); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
        return
    }

    if err := db.DB.Delete(&businessType).Error; err != nil {
        c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete business type"})
        return
    }

    c.JSON(http.StatusOK, gin.H{"message": "Business type deleted successfully"})
}