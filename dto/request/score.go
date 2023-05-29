package request

type ScoreRequest struct {
	Subject    string  `json:"subject"`
	Score      float64 `json:"score"`
	CreatedBy  string  `json:"createdBy"`
	ModifiedBy string  `json:"modifiedBy"`
}
