package business

import (
	"log"
	"math/rand"
	"os"
	"time"

	_ "github.com/lib/pq"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type DB struct {
	Db *gorm.DB
}

func (d *DB) ConnectDb() {
	// To start DB run:
	// docker run --name postgres_db  -p 5432:5432 -e POSTGRES_USER=postgres -e POSTGRES_PASSWORD=postgres -e POSTGRES_DB=postgres -d postgres:13
	dsn_test := "host=172.17.0.2 user=postgres password=pass dbname=crud port=5432 sslmode=disable TimeZone=Asia/Shanghai"

	db, err := gorm.Open(postgres.Open(dsn_test), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		log.Fatal("Failed to connect to database. \n", err)
		os.Exit(2)
	}

	log.Println("connected")
	db.Logger = logger.Default.LogMode(logger.Info)

	log.Println("running migrations")
	db.AutoMigrate(&BankAccount{}, &Transaction{})

	d.Db = db
}

func (d *DB) CreateBankAccount(bankAccount BankAccount) (BankAccount, error) {
	id := generateID()

	accountNum := ""
	for i := 0; i < 10; i++ {
		accountNum += (string)(rand.Intn(10) + 48)
	}

	bankAccount = BankAccount{
		ID:                id,
		AccountNumber:     accountNum,
		AccountHolderName: bankAccount.AccountHolderName,
		AccountType:       bankAccount.AccountType,
		Balance:           bankAccount.Balance,
		Ledger:            bankAccount.Ledger,
	}

	d.Db.Create(&bankAccount)

	return bankAccount, nil
}

func (d *DB) Balance(accntNumber string) (float64, error) {
	var bankAccount BankAccount

	d.Db.Where("account_Number = ?", accntNumber).First(&bankAccount)

	return bankAccount.Balance, nil
}

func (d *DB) Deposit(accntNumber string, amount float64) (float64, error) {
	var bankAccount BankAccount
	// get account with account number
	d.Db.Where("account_Number = ?", accntNumber).First(&bankAccount)

	depositID := generateID() // generate a random ID
	transactionDeposit := Transaction{ID: depositID, Amount: amount, Type: "Deposit", Date: time.Now(), BankAccountID: bankAccount.ID} // create a transaction
	bankAccount.Ledger = append(bankAccount.Ledger, transactionDeposit) // add the transaction to the ledger
	bankAccount.Balance += amount // update the balance

	d.Db.Where("account_Number = ?", accntNumber).Updates(&bankAccount) // update the account

	return bankAccount.Balance, nil
}

func (d *DB) Withdraw(accntNumber string, amount float64) (float64, error) {
	var bankAccount BankAccount
	d.Db.Where("account_Number = ?", accntNumber).First(&bankAccount) // get the bank account

	withdrawID := generateID() // generate a random ID
	transactionWithdraw := Transaction{ID: withdrawID, Amount: amount, Type: "Withdraw", Date: time.Now(), BankAccountID: bankAccount.ID} // create a transaction
	bankAccount.Ledger = append(bankAccount.Ledger, transactionWithdraw) // add the transaction to the ledger
	bankAccount.Balance -= amount // update the balance

	d.Db.Where("account_Number = ?", accntNumber).Updates(&bankAccount) // update the account

	return bankAccount.Balance, nil
}

func (d *DB) PrintLedger(accntNumber string) ([]Transaction, error) {
	var bankAccount BankAccount

	d.Db.Where("account_Number = ?", accntNumber).First(&bankAccount) // get the bank account
	log.Printf("Found Bank Account: %s", bankAccount.ID)

	var transactions []Transaction // get the transactions
	d.Db.Where("bank_account_id = ?", bankAccount.ID).Find(&transactions)

	return transactions, nil
}
