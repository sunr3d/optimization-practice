package statssvc

import (
	"context"
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Тесты.
func TestService_Calculate(t *testing.T) {
	svc := New()
	ctx := context.Background()

	t.Run("пустой массив", func(t *testing.T) {
		result, err := svc.Calculate(ctx, []float64{})

		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, 0, result.Count)
		assert.Equal(t, 0.0, result.Sum)
		assert.Equal(t, 0.0, result.Mean)
		assert.Equal(t, 0.0, result.Median)
		assert.Equal(t, 0.0, result.Variance)
		assert.Equal(t, 0.0, result.Deviation)
		assert.Equal(t, 0.0, result.Min)
		assert.Equal(t, 0.0, result.Max)
	})

	t.Run("один элемент", func(t *testing.T) {
		result, err := svc.Calculate(ctx, []float64{5})

		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, 1, result.Count)
		assert.Equal(t, 5.0, result.Sum)
		assert.Equal(t, 5.0, result.Mean)
		assert.Equal(t, 5.0, result.Median)
		assert.Equal(t, 0.0, result.Variance)
		assert.Equal(t, 0.0, result.Deviation)
		assert.Equal(t, 5.0, result.Min)
		assert.Equal(t, 5.0, result.Max)
	})

	t.Run("два элемента", func(t *testing.T) {
		result, err := svc.Calculate(ctx, []float64{1, 3})

		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, 2, result.Count)
		assert.Equal(t, 4.0, result.Sum)
		assert.Equal(t, 2.0, result.Mean)
		assert.Equal(t, 2.0, result.Median)
		assert.Equal(t, 1.0, result.Variance)
		assert.Equal(t, 1.0, result.Deviation)
		assert.Equal(t, 1.0, result.Min)
		assert.Equal(t, 3.0, result.Max)
	})

	t.Run("нечетное количество элементов", func(t *testing.T) {
		result, err := svc.Calculate(ctx, []float64{1, 2, 3})

		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, 3, result.Count)
		assert.Equal(t, 6.0, result.Sum)
		assert.Equal(t, 2.0, result.Mean)
		assert.Equal(t, 2.0, result.Median)
		assert.Equal(t, 2.0/3.0, result.Variance)
		assert.InDelta(t, math.Sqrt(2.0/3.0), result.Deviation, 1e-10)
		assert.Equal(t, 1.0, result.Min)
		assert.Equal(t, 3.0, result.Max)
	})

	t.Run("четное количество элементов", func(t *testing.T) {
		result, err := svc.Calculate(ctx, []float64{1, 2, 3, 4})

		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, 4, result.Count)
		assert.Equal(t, 10.0, result.Sum)
		assert.Equal(t, 2.5, result.Mean)
		assert.Equal(t, 2.5, result.Median)
		assert.Equal(t, 1.25, result.Variance)
		assert.InDelta(t, math.Sqrt(1.25), result.Deviation, 1e-10)
		assert.Equal(t, 1.0, result.Min)
		assert.Equal(t, 4.0, result.Max)
	})

	t.Run("отрицательные числа", func(t *testing.T) {
		result, err := svc.Calculate(ctx, []float64{-1, 0, 1})

		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, 3, result.Count)
		assert.Equal(t, 0.0, result.Sum)
		assert.Equal(t, 0.0, result.Mean)
		assert.Equal(t, 0.0, result.Median)
		assert.Equal(t, 2.0/3.0, result.Variance)
		assert.InDelta(t, math.Sqrt(2.0/3.0), result.Deviation, 1e-10)
		assert.Equal(t, -1.0, result.Min)
		assert.Equal(t, 1.0, result.Max)
	})

	t.Run("одинаковые числа", func(t *testing.T) {
		result, err := svc.Calculate(ctx, []float64{5, 5, 5})

		require.NoError(t, err)
		require.NotNil(t, result)

		assert.Equal(t, 3, result.Count)
		assert.Equal(t, 15.0, result.Sum)
		assert.Equal(t, 5.0, result.Mean)
		assert.Equal(t, 5.0, result.Median)
		assert.Equal(t, 0.0, result.Variance)
		assert.Equal(t, 0.0, result.Deviation)
		assert.Equal(t, 5.0, result.Min)
		assert.Equal(t, 5.0, result.Max)
	})
}

// Бенчмарк.
func BenchmarkService_Calculate(b *testing.B) {
	svc := New()
	ctx := context.Background()

	testCases := []struct {
		name string
		data []float64
	}{
		{"10 элементов", generateData(10)},
		{"100 элементов", generateData(100)},
		{"1000 элементов", generateData(1000)},
		{"10000 элементов", generateData(10000)},
		{"100000 элементов", generateData(100000)},
		{"1000000 элементов", generateData(1000000)},
	}

	for _, tc := range testCases {
		b.Run(tc.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_, err := svc.Calculate(ctx, tc.data)
				if err != nil {
					b.Fatal(err)
				}
			}
		})
	}
}

// Вспомогательная функция для генерации тестовых данных.
func generateData(size int) []float64 {
	data := make([]float64, size)

	for i := 0; i < size; i++ {
		data[i] = float64(i)
	}

	return data
}
