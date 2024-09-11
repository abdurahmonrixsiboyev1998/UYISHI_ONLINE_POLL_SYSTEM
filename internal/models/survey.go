package models

type Survey struct {
	ID      string            `json:"id"`
	Title   string            `json:"title"`
	Options map[string]int    `json:"options"` 
	Votes   map[string]string `json:"votes"`   
}
