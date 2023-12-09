package main

type Event struct {
	Id          int    `json:"id"`
	Date        string `json:"date"`
	Name        string `json:"name"`
	Description string `json:"description"`
}
