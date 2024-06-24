package http

import (
	"github.com/gin-gonic/gin"
	"smart-contracts-research/internal/usecase"
)

func NewRouter(handler *gin.Engine) {
	_ = handler.Group("/api/v1")
	{
	}
}
