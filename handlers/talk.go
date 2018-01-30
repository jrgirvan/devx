package handlers

import (
	"github.com/jinzhu/gorm"
	"github.com/jrgirvan/devx/models"
	"github.com/labstack/echo"
	"github.com/labstack/gommon/log"
	"net/http"
	"net/url"
)

type H map[string]interface{}

// GetTalks endpoint
func GetTalks(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var talks []models.Talk
		// var p = page(c)
		// if err := db.Limit(PaginationLimit).Order("rank").Offset(p * PaginationLimit).Find(&talks).Error; err != nil {
		if err := db.Order("rank").Find(&talks).Error; err != nil {
			log.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)

		}
		return c.JSON(http.StatusOK, talks)
	}
}

// CreateTalk endpoint
func CreateTalk(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Instantiate a new talk
		var talk models.Talk
		var totalCount int
		// Map imcoming JSON body to the new Talk
		db.Find(&talk).Count(&totalCount)
		c.Bind(&talk)
		talk.ID = models.NewID()
		talk.Rank = totalCount
		talk.Completed = false
		talk.Pinned = false
		// Add a talk using our new model
		if err := db.Create(&talk).Error; err != nil {
			log.Debug("Error creating talk", err)
			return echo.NewHTTPError(http.StatusBadRequest)
		}
		log.Debug("Created talk", talk.ID)
		return c.JSON(http.StatusCreated, talk)
	}
}

// PutTalk endpoint
func PutTalk(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var talk models.Talk
		id, _ := url.QueryUnescape(c.Param("id"))
		if err := db.Where("id = ?", id).First(&talk).Error; err != nil {
			log.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}
		c.Bind(&talk)
		db.Save(&talk)
		log.Debug("Updated talk", talk.ID)
		return c.JSON(http.StatusOK, talk)
	}
}

// DeleteTalk endpoint
func DeleteTalk(db *gorm.DB) echo.HandlerFunc {
	return func(c echo.Context) error {
		var talk models.Talk
		id, _ := url.QueryUnescape(c.Param("id"))
		if err := db.Where("id = ?", id).First(&talk).Error; err != nil {
			log.Debug(err)
			return echo.NewHTTPError(http.StatusNotFound)
		}

		db.Delete(&talk)
		log.Debug("Deleted talk")
		return c.JSON(http.StatusOK, H{
			"deleted": c.Param("id"),
		})
	}
}
