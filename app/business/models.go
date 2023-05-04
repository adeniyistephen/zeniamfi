package business

import (
	"time"

	"gorm.io/gorm"
)

type BankAccount struct {
	gorm.Model
	ID                string        `json:"bankAccountID"`
	AccountNumber     string        `json:"accountNumber"`
	AccountHolderName string        `json:"accountHolderName"`
	AccountType       string        `json:"accountType"`
	Balance           float64       `json:"balance"`
	Ledger            []Transaction `json:"ledger"`
}

type Transaction struct {
	gorm.Model
	ID            string      `json:"transactionID"`
	Amount        float64     `json:"amount"`
	Type          string      `json:"type"`
	Date          time.Time   `json:"date"`
	BankAccountID string      `json:"_"`
	Account       BankAccount `json:"account,omitempty" gorm:"foreignKey:BankAccountID;references:ID"`
}
