package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type notRouteHandler struct {
}

func NewNotRouteHandler() *notRouteHandler {
	return &notRouteHandler{}
}
func (h *notRouteHandler) Index(c *gin.Context) {
	c.HTML(http.StatusNotFound, "404", nil)
}
