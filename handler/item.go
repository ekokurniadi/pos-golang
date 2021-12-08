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

type itemHandler struct {
	service service.ItemService
}

func NewItemHandler(service service.ItemService) *itemHandler {
	return &itemHandler{service}
}

func (h *itemHandler) SearchItem(c *gin.Context) {
	var input input.InputFilter
	err := c.ShouldBindJSON(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	filtered, err := h.service.ItemServiceSearch(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail Item", http.StatusOK, "success", formatter.FormatItems(filtered))
	c.JSON(http.StatusOK, response)

}

func (h *itemHandler) GetItem(c *gin.Context) {
	var input input.InputIDItem
	err := c.ShouldBindUri(&input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	itemDetail, err := h.service.ItemServiceGetByID(input)
	if err != nil {
		response := helper.ApiResponse("Failed to get Item", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Detail Item", http.StatusOK, "success", formatter.FormatItem(itemDetail))
	c.JSON(http.StatusOK, response)
}

func (h *itemHandler) GetItems(c *gin.Context) {
	items, err := h.service.ItemServiceGetAll()
	if err != nil {
		response := helper.ApiResponse("Failed to get Items", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("List of Items", http.StatusOK, "success", formatter.FormatItems(items))
	c.JSON(http.StatusOK, response)
}

func (h *itemHandler) CreateItem(c *gin.Context) {
	var input input.ItemInput
	err := c.ShouldBindJSON(&input)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Create Item failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	newItem, err := h.service.ItemServiceCreate(input)
	if err != nil {
		response := helper.ApiResponse("Create Item failed", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Create Item", http.StatusOK, "success", formatter.FormatItem(newItem))
	c.JSON(http.StatusOK, response)
}
func (h *itemHandler) UpdateItem(c *gin.Context) {
	var inputID input.InputIDItem
	err := c.ShouldBindUri(&inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Items", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var inputData input.ItemInput
	err = c.ShouldBindJSON(&inputData)
	if err != nil {
		errors := helper.FormatValidationError(err)
		errorMessage := gin.H{"errors": errors}
		response := helper.ApiResponse("Update Item failed", http.StatusBadRequest, "error", errorMessage)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	updatedItem, err := h.service.ItemServiceUpdate(inputID, inputData)
	if err != nil {
		response := helper.ApiResponse("Failed to get Items", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Update Item", http.StatusOK, "success", formatter.FormatItem(updatedItem))
	c.JSON(http.StatusOK, response)
}
func (h *itemHandler) DeleteItem(c *gin.Context) {
	param := c.Param("id")
	id, _ := strconv.Atoi(param)
	var inputID input.InputIDItem
	inputID.ID = id
	_, err := h.service.ItemServiceGetByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Items", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	_, err = h.service.ItemServiceDeleteByID(inputID)
	if err != nil {
		response := helper.ApiResponse("Failed to get Items", http.StatusBadRequest, "error", nil)
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := helper.ApiResponse("Successfully Delete Item", http.StatusOK, "success", nil)
	c.JSON(http.StatusOK, response)
}
