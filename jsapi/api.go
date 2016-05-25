package jsapi

type ErrorResponse struct {
	Errors []Error `json:"errors"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}

type Error struct {
	Status string `json:"status,omitempty"`
	Title string `json:"title,omitempty"`
	Meta map[string]interface{} `json:"meta,omitempty"`
}
