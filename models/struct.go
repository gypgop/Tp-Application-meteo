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
	Hourly Hourly `json:"daily"`
}

type Hourly struct { //not hourly but i don't want to break smth changing it 👀
	Time                     Hours     `json:"time"`
	Temperature              Temps     `json:"temperature_2m_min"`
	PrecipitationProbability []float64 `json:"precipitation_sum"`
}

type Hours []string

type Temps []float64
