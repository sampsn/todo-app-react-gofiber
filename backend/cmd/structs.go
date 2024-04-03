package main

type Todo struct {
	Id   int    `json:"id"`
	Item string `json:"item"`
}

type ListTodosResponse struct {
	Data []Todo `json:"data"`
}
