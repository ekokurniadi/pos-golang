package handler

import (
	"net/http"

	"github.com/ekokurniadi/pos-golang/service"
	"github.com/gin-gonic/gin"
)

type sessionHandler struct {
	service service.UserService
}

func NewSessionHandler(service service.UserService) *sessionHandler {
	return &sessionHandler{service}
}

func (h *sessionHandler) Dashboard(c *gin.Context) {
	c.HTML(http.StatusOK, "header", gin.H{"nama": "Eko Kurniadi", "title": "Dashboard"})
	c.HTML(http.StatusOK, "dashboard.html", nil)
	c.HTML(http.StatusOK, "footer", nil)
}
