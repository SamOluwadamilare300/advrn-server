package routes

import (
	// "advrn-server/models"
	"advrn-server/models/storage"
	"advrn-server/models/utils"
	"net/http"
	"strconv"

	"github.com/kataras/iris/v12"
)

func Get360Images(ctx iris.Context) {
	propertyID, err := strconv.ParseUint(ctx.Params().Get("id"), 10, 32)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid property ID"})
		return
	}

	var virtualTour models.VirtualTour
	if err := storage.DB.Where("property_id = ? AND type = ?", propertyID, models.TourType360Photos).
		Preload("Images360").
		First(&virtualTour).Error; err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(iris.Map{"error": "360° virtual tour not found"})
		return
	}

	ctx.JSON(iris.Map{
		"status": "success",
		"data":   virtualTour.Images360,
	})
}

func ActivateVirtualTour(ctx iris.Context) {
	tourID, err := strconv.ParseUint(ctx.Params().Get("id"), 10, 32)
	if err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid tour ID"})
		return
	}

	var req struct {
		IsActive bool `json:"is_active"`
	}

	if err := ctx.ReadJSON(&req); err != nil {
		ctx.StatusCode(http.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid request body"})
		return
	}

	var virtualTour models.VirtualTour
	if err := storage.DB.Preload("Property").First(&virtualTour, tourID).Error; err != nil {
		ctx.StatusCode(http.StatusNotFound)
		ctx.JSON(iris.Map{"error": "Virtual tour not found"})
		return
	}

	// Verify user owns the property
	claims := utils.GetJWTClaims(ctx)
	if virtualTour.Property.UserID != claims.UserID {
		ctx.StatusCode(http.StatusForbidden)
		ctx.JSON(iris.Map{"error": "Access denied"})
		return
	}

	virtualTour.IsActive = req.IsActive

	if err := storage.DB.Save(&virtualTour).Error; err != nil {
		ctx.StatusCode(http.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to update virtual tour status"})
		return
	}

	ctx.JSON(iris.Map{
		"status": "success",
		"data":   virtualTour,
	})
}
