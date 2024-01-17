package model

import "time"

type Event struct {
	Id     int       `json:"id"`
	UserId int       `json:"user_id"`
	Name   string    `json:"name"`
	Date   time.Time `json:"date"`
}
