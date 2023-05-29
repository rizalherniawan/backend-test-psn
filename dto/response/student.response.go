package response

type StudentResponse struct {
	Id      int             `json:"id"`
	Name    string          `json:"name"`
	Age     int             `json:"age"`
	Subject []ScoreResponse `json:"subject"`
}
