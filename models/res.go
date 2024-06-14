package models

type MergeResInfo struct {
	ID      int      `json:"id"`
	Title   string   `json:"title"`
	Sha     string   `json:"sha"`
	Message []string `json:"message"` // 错误的时候返回
}
