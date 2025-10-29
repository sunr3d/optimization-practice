package httphandlers

type calcStatsReq struct {
	Values []float64 `json:"values" binding:"required"`
}

type calcStatsResp struct {
	Count     int     `json:"count"`
	Sum       float64 `json:"sum"`
	Mean      float64 `json:"mean"`
	Median    float64 `json:"median"`
	Variance  float64 `json:"variance"`
	Deviation float64 `json:"deviation"`
	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
}
