package request

type StudentRequest struct {
	Name       string `json:"name"`
	Age        int    `json:"age"`
	CreatedBy  string `json:"createdBy"`
	ModifiedBy string `json:"modifiedBy"`
}
