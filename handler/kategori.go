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

type kategoriHandler struct {
	service service.KategoriService
}

func NewKategoriHandler(service service.KategoriService) *kategoriHandler {
	return &kategoriHandler{service}
}
func (h *kategoriHandler) GetKategori(c *gin.Context) {
	var input input.InputIDKategori
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Kategori", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	kategoriDetail, err := h.service.KategoriServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Kategori", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail Kategori", http.StatusOK, "success", formatter.FormatKategori(kategoriDetail))
	c.JSON(http.StatusOK, response)
}

func (h *kategoriHandler) GetKategoris(c *gin.Context) {
	kategoris, err := h.service.KategoriServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get Kategoris", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Kategoris", http.StatusOK, "success", formatter.FormatKategoris(kategoris))
	c.JSON(http.StatusOK, response)
}

func (h *kategoriHandler) CreateKategori(c *gin.Context) {
	var input input.KategoriInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create Kategori failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newKategori, err := h.service.KategoriServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create Kategori failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create Kategori", http.StatusOK, "success", formatter.FormatKategori(newKategori))
	c.JSON(http.StatusOK, response)
}
func (h *kategoriHandler) UpdateKategori(c *gin.Context) {
	var inputID input.InputIDKategori
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Kategoris", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.KategoriInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update Kategori failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedKategori, err := h.service.KategoriServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get Kategoris", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update Kategori", http.StatusOK, "success", formatter.FormatKategori(updatedKategori))
	c.JSON(http.StatusOK, response)
}
func (h *kategoriHandler) DeleteKategori(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDKategori
	inputID.ID = id
	_, err := h.service.KategoriServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Kategoris", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.KategoriServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Kategoris", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete Kategori", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
