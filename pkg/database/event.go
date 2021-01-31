package database

type RequestCreate struct {
	Title       string `json:"title"`
	Description string `json:"description"`
}

type Event struct {
	ID        string  `json:"id"`
	EventName string  `json:"event_name"`
	CreatedAt string  `json:"created_at"`
	Payload   Article `json:"payload"`
}
