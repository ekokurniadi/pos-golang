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

type jenisHandler struct {
	service service.JenisService
}

func NewJenisHandler(service service.JenisService) *jenisHandler {
	return &jenisHandler{service}
}
func (h *jenisHandler) GetJenis(c *gin.Context) {
	var input input.InputIDJenis
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Jenis", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	jenisDetail, err := h.service.JenisServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Jenis", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail Jenis", http.StatusOK, "success", formatter.FormatJenis(jenisDetail))
	c.JSON(http.StatusOK, response)
}

func (h *jenisHandler) GetJeniss(c *gin.Context) {
	jeniss, err := h.service.JenisServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get Jeniss", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Jeniss", http.StatusOK, "success", formatter.FormatJeniss(jeniss))
	c.JSON(http.StatusOK, response)
}

func (h *jenisHandler) CreateJenis(c *gin.Context) {
	var input input.JenisInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create Jenis failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newJenis, err := h.service.JenisServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create Jenis failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create Jenis", http.StatusOK, "success", formatter.FormatJenis(newJenis))
	c.JSON(http.StatusOK, response)
}
func (h *jenisHandler) UpdateJenis(c *gin.Context) {
	var inputID input.InputIDJenis
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Jeniss", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.JenisInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update Jenis failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedJenis, err := h.service.JenisServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get Jeniss", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update Jenis", http.StatusOK, "success", formatter.FormatJenis(updatedJenis))
	c.JSON(http.StatusOK, response)
}
func (h *jenisHandler) DeleteJenis(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDJenis
	inputID.ID = id
	_, err := h.service.JenisServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Jeniss", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.JenisServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Jeniss", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete Jenis", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
