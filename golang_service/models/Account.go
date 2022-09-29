package models

type Account struct {
	Address    string
	PrivateKey string
	Role       string
}

func (account *Account) SetAccount(address string, privateKey string, role string) {
	account.Address = address
	account.PrivateKey = privateKey
	account.Role = role
}

// func (account Account) getAccount() string {
//     return account.name
// }
