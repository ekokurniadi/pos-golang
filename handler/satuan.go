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

type satuanHandler struct {
	service service.SatuanService
}

func NewSatuanHandler(service service.SatuanService) *satuanHandler {
	return &satuanHandler{service}
}
func (h *satuanHandler) GetSatuan(c *gin.Context) {
	var input input.InputIDSatuan
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Satuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	satuanDetail, err := h.service.SatuanServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Satuan", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail Satuan", http.StatusOK, "success", formatter.FormatSatuan(satuanDetail))
	c.JSON(http.StatusOK, response)
}

func (h *satuanHandler) GetSatuans(c *gin.Context) {
	satuans, err := h.service.SatuanServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get Satuans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Satuans", http.StatusOK, "success", formatter.FormatSatuans(satuans))
	c.JSON(http.StatusOK, response)
}

func (h *satuanHandler) CreateSatuan(c *gin.Context) {
	var input input.SatuanInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create Satuan failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newSatuan, err := h.service.SatuanServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create Satuan failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create Satuan", http.StatusOK, "success", formatter.FormatSatuan(newSatuan))
	c.JSON(http.StatusOK, response)
}
func (h *satuanHandler) UpdateSatuan(c *gin.Context) {
	var inputID input.InputIDSatuan
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Satuans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.SatuanInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update Satuan failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedSatuan, err := h.service.SatuanServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get Satuans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update Satuan", http.StatusOK, "success", formatter.FormatSatuan(updatedSatuan))
	c.JSON(http.StatusOK, response)
}
func (h *satuanHandler) DeleteSatuan(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDSatuan
	inputID.ID = id
	_, err := h.service.SatuanServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Satuans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.SatuanServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Satuans", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete Satuan", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
