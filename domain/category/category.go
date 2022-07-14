package category

type Category string

// Category 是用來在輸出log 時，區別 log 的來源是出自系統的那一個功能模組。
const (
	// User 功能模組
	User Category = "user"
	// Config 功能模組
	Config   Category = "config"
)
