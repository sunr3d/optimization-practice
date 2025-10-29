package httphandlers

import (
	"net/http"

	"github.com/wb-go/wbf/ginext"
)

func (h *handler) calculateStats(c *ginext.Context) {
	var req calcStatsReq
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(
			http.StatusBadRequest,
			ginext.H{"error": "некорректный запрос"},
		)
		return
	}

	stats, err := h.svc.Calculate(c.Request.Context(), req.Values)
	if err != nil {
		c.JSON(
			http.StatusInternalServerError,
			ginext.H{"error": "ошибка при расчете статистики"},
		)
		return
	}

	resp := calcStatsResp{
		Count:     stats.Count,
		Sum:       roundToTwo(stats.Sum),
		Mean:      roundToTwo(stats.Mean),
		Median:    roundToTwo(stats.Median),
		Variance:  roundToTwo(stats.Variance),
		Deviation: roundToTwo(stats.Deviation),
		Min:       roundToTwo(stats.Min),
		Max:       roundToTwo(stats.Max),
	}

	c.JSON(http.StatusOK, resp)
}
