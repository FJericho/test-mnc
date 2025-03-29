package model

type Payment struct {
    ID          string  `json:"id"`
    UserID  string  `json:"user_id"`
    Amount      float64 `json:"amount"`
    Timestamp   string  `json:"timestamp"`
}