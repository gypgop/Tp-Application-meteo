package models

type Datas struct {
	Centre Center `json:"centre"`
}

type Center struct {
	Coordinates []float64 `json:"coordinates"`
}

type Code struct {
	Codes int `json:"codesPostaux"`
}
