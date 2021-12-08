package handler

import (
	"net/http"
	"strconv"

	"github.com/ekokurniadi/pos-golang/formatter"
	"github.com/ekokurniadi/pos-golang/helper"
	"github.com/ekokurniadi/pos-golang/input"
	"github.com/ekokurniadi/pos-golang/service"
	"github.com/gin-gonic/gin"
)

type lokasiHandler struct {
	service service.LokasiService
}

func NewLokasiHandler(service service.LokasiService) *lokasiHandler {
	return &lokasiHandler{service}
}
func (h *lokasiHandler) GetLokasi(c *gin.Context) {
	var input input.InputIDLokasi
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Lokasi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	lokasiDetail, err := h.service.LokasiServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Lokasi", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail Lokasi", http.StatusOK, "success", formatter.FormatLokasi(lokasiDetail))
	c.JSON(http.StatusOK, response)
}

func (h *lokasiHandler) GetLokasis(c *gin.Context) {
	lokasis, err := h.service.LokasiServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get Lokasis", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Lokasis", http.StatusOK, "success", formatter.FormatLokasis(lokasis))
	c.JSON(http.StatusOK, response)
}

func (h *lokasiHandler) CreateLokasi(c *gin.Context) {
	var input input.LokasiInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create Lokasi failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newLokasi, err := h.service.LokasiServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create Lokasi failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create Lokasi", http.StatusOK, "success", formatter.FormatLokasi(newLokasi))
	c.JSON(http.StatusOK, response)
}
func (h *lokasiHandler) UpdateLokasi(c *gin.Context) {
	var inputID input.InputIDLokasi
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Lokasis", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.LokasiInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update Lokasi failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedLokasi, err := h.service.LokasiServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get Lokasis", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update Lokasi", http.StatusOK, "success", formatter.FormatLokasi(updatedLokasi))
	c.JSON(http.StatusOK, response)
}
func (h *lokasiHandler) DeleteLokasi(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDLokasi
	inputID.ID = id
	_, err := h.service.LokasiServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Lokasis", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.LokasiServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Lokasis", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete Lokasi", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
