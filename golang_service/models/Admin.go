package models

type Admin struct {
	Username         string `json:"username"`
	Email            string `json:"email"`
	Password         string `json:"password"`
	adminEntranceFee uint64
	adminAddr        string
}
