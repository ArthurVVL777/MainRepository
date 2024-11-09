package main

type Message struct {
	ID   int    `json:"id"`
	Text string `json:"text"`
}

type Task struct {
	ID      int    `json:"id" gorm:"primaryKey"`
	Task    string `json:"task"`
	IsDone  bool   `json:"is_done"`
	Message string `json:"message"`
}

type Response struct {
	Status  string `json:"status"`
	Message string `json:"message"`
}
