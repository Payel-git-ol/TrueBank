package models

import "time"

type HistoryTransaction struct {
	Username        string    `json:"username"`
	NameTransaction string    `json:"name_transaction"`
	Sum             string    `json:"sum"`
	NumberCard      string    `json:"number_card"`
	DateCreated     time.Time `json:"date_created"`
}
