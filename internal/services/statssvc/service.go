package statssvc

import (
	"context"
	"math"

	"github.com/sunr3d/optimization-practice/internal/interfaces/services"
	"github.com/sunr3d/optimization-practice/internal/models"
)

var _ services.StatsService = (*service)(nil)

type service struct{}

func New() *service {
	return &service{}
}

func (s *service) Calculate(ctx context.Context, data []float64) (*models.Stats, error) {
	if len(data) == 0 {
		return &models.Stats{}, nil
	}

	sum := getSum(data)
	mean := sum / float64(len(data))
	min, max := getMinMax(data)
	median := getMedian(data)
	variance := getVariance(data, mean)
	deviation := math.Sqrt(variance)

	return &models.Stats{
		Count:     len(data),
		Sum:       sum,
		Mean:      mean,
		Median:    median,
		Variance:  variance,
		Deviation: deviation,
		Min:       min,
		Max:       max,
	}, nil
}
