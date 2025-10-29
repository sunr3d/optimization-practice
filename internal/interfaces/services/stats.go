package services

import (
	"context"

	"github.com/sunr3d/optimization-practice/internal/models"
)

type StatsService interface {
	Calculate(ctx context.Context, data []float64) (*models.Stats, error)
}
