package model

type WebResponse[T any] struct {
	Data    T      `json:"data"`
	Message string `json:"message,omitempty"`
	Success bool   `json:"success"`
	// Paging is optional
	Paging *PageMetadata `json:"paging,omitempty"`
}

type PageMetadata struct {
	Page      int   `json:"page"`
	Size      int   `json:"size"`
	TotalItem int64 `json:"total_item"`
	TotalPage int64 `json:"total_page"`
}
