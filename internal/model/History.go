package model

type History struct {
    ID        string `json:"id"`
    UserID string `json:"user_id"`
    Activity  string `json:"activity"` 
    Timestamp string `json:"timestamp"`
}