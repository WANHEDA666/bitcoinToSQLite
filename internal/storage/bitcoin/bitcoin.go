package bitcoin

type Response struct {
	ChartName string `json:"chartName"`
	Time      Time   `json:"time"`
	BPI       BPI    `json:"bpi"`
}

type Time struct {
	Updated string `json:"updated"`
}

type BPI struct {
	USD USD `json:"USD"`
	GBP GBP `json:"GBP"`
	EUR EUR `json:"EUR"`
}

type USD struct {
	Rate string `json:"rate"`
}

type GBP struct {
	Rate string `json:"rate"`
}

type EUR struct {
	Rate string `json:"rate"`
}
