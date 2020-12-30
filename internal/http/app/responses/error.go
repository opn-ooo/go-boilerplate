package responses

// Error ... Common Error Structure
type Error struct {
	Code    string `json:"code"`
	Message string `json:"message"`
}
