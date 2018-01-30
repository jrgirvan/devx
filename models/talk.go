package models

// Talk is a struct containing Talk data
type Talk struct {
	ID          string `json:"id,omitempty",gorm:"primary_key"`
	Topic       string `json:"topic"`
	Name        string `json:"name"`
	Duration    int    `json:"duration"`
	Rank        int    `json:"rank"`
	Completed   bool   `json:"completed"`
	Pinned      bool   `json:"pinned"`
	Description string `json:"description"`
}
