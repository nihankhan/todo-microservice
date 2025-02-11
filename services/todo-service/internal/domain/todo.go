package domain

type Todo struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Description string `json:"description"`
	Completed   bool   `json:"completed"`
	UserID      string `json:"user_id"`
}

// type Todo struct {
// 	ID          string
// 	Title       string
// 	Description string
// 	Completed   bool
// }
