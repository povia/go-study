package controllers

import (
	"github.com/gin-gonic/gin"
)

type OrderController struct {
}

type orderReq struct {
	symbol    string `json:"symbol"`
	price     string `json:"price"`
	qty       string `json:"qty"`
	orderType string `json:"type"`
	side      string `json:"side"`
}

func (OrderController) Place(c *gin.Context) {
	var o orderReq
	if err := c.BindJSON(&o); err != nil {
		return
	}
	// TODO order 처리 추가
}
