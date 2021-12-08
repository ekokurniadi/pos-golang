package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type transaksiHandler struct {
}

func NewTransaksiHandler() *transaksiHandler {
	return &transaksiHandler{}
}

func (h *transaksiHandler) Index(c *gin.Context) {
	c.HTML(http.StatusOK, "header", gin.H{"nama": "Eko Kurniadi", "title": "Dashboard"})
	c.HTML(http.StatusOK, "form-transaksi", nil)
	c.HTML(http.StatusOK, "footer", nil)
}
