package response

type ScoreResponse struct {
	Subject string  `json:"subject"`
	Score   float64 `json:"score"`
}
