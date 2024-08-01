package main

type PredictorConfig struct {
	Cert    string `json:"cert"`
	Privkey string `json:"privkey"`
}

type PredictOptions struct {
	Home  string `json:"home"`
	Guest string `json:"guest"`
}

type Predict struct {
	Result string `json:"result"`
	Total  string `json:"total"`
}
