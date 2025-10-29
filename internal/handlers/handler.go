package httphandlers

import (
	"github.com/wb-go/wbf/ginext"

	"github.com/sunr3d/optimization-practice/internal/interfaces/services"
)

type handler struct {
	svc services.StatsService
}

func New(svc services.StatsService) *handler {
	return &handler{
		svc: svc,
	}
}

func (h *handler) RegisterHandlers() *ginext.Engine {
	router := ginext.New("")
	router.Use(ginext.Logger(), ginext.Recovery())

	router.POST("/stats", h.calculateStats)

	return router
}
