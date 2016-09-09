package facebook

// AnswerError is a fb opengraph error description
type AnswerError struct {
	Message   string `json:"message"`
	Type      string `json:"type"`
	Code      int    `json:"code"`
	FbTraceID string `json:"fbtrace_id"`
}
