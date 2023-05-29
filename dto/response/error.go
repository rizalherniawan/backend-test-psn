package response

type Error struct {
	Success bool   `json:"Success"`
	Message string `json:"Message"`
}
