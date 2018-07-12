package parse

// MeasuredLocation is a measurment location with associated measurements
type MeasuredLocation struct {
	Lat      float64   `json:"lat"`
	Lon      float64   `json:"lon"`
	ID       string    `json:"id"`
	Name     string    `json:"name"`
	Data     []float64 `json:"data"`
	Type     string    `json:"type"`
	Chemical string    `json:"chemical"`
}
