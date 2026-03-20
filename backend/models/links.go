package models

type Link struct {
	ID    int    `json:"id"`
	Owner int    `json:"owner"`
	Path  string `json:"path"`
	URL   string `json:"url"`
}
