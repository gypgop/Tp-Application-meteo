package models

type Datas struct {
	Centre Center `json:"centre"`
}

type Center struct {
	Coordinates []float64 `json:"coordinates"`
}

type Code struct {
	Codes []string `json:"codesPostaux"`
}

type Reponse struct {
	Hourly Hourly `json:"hourly"`
}

type Hourly struct {
	Time                     Hours `json:"time"`
	Temperature              Temps `json:"temperature_2m"`
	PrecipitationProbability []int `json:"precipitation_probability"`
}

type Hours []string

type Temps []float64
