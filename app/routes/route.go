package routes

import (
	"github.com/adeniyistephen/zeinamfi/app/business"
	"github.com/gofiber/fiber/v2"
)

type Routes struct {
	Database business.DB
}

func (r *Routes) CreateAccount(c *fiber.Ctx) error {
	newAccount := new(business.BankAccount)

	err := c.BodyParser(newAccount)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	result, err := r.Database.CreateBankAccount(*newAccount)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"data":    result,
		"success": true,
		"message": "Task added!",
	})
	return nil
}

func (r *Routes) GetBalance(c *fiber.Ctx) error {
	accountNumber := c.Params("accntnum")
	if accountNumber == "" {

		return c.Status(500).JSON(&fiber.Map{

			"message": "accntnum cannot be empty",
		})
	}

	result, err := r.Database.Balance(accountNumber)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data":    result,
		"success": true,
		"message": "",
	})
}

func (r *Routes) CreateDeposit(c *fiber.Ctx) error {
	accntnum := c.Params("accntnum")

	if accntnum == "" {

		return c.Status(500).JSON(&fiber.Map{

			"message": "accntnum cannot be empty",
		})
	}

	newDeposit := new(business.Transaction)

	err := c.BodyParser(newDeposit)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	result, err := r.Database.Deposit(accntnum, newDeposit.Amount)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"data":    result,
		"success": true,
		"message": "Account Updated!",
	})
	return nil
}

func (r *Routes) CreateWithdraw(c *fiber.Ctx) error {
	accntnum := c.Params("accntnum")

	if accntnum == "" {

		return c.Status(500).JSON(&fiber.Map{

			"message": "accntnum cannot be empty",
		})
	}

	newDeposit := new(business.Transaction)

	err := c.BodyParser(newDeposit)
	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	result, err := r.Database.Withdraw(accntnum, newDeposit.Amount)

	if err != nil {
		c.Status(400).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
		return err
	}

	c.Status(200).JSON(&fiber.Map{
		"data":    result,
		"success": true,
		"message": "Account Updated!",
	})
	return nil
}

func (r *Routes) GetLedger(c *fiber.Ctx) error {
	accountNumber := c.Params("accntnum")
	if accountNumber == "" {

		return c.Status(500).JSON(&fiber.Map{

			"message": "accntnum cannot be empty",
		})
	}

	result, err := r.Database.PrintLedger(accountNumber)
	if err != nil {
		return c.Status(500).JSON(&fiber.Map{
			"data":    nil,
			"success": false,
			"message": err,
		})
	}

	return c.Status(200).JSON(&fiber.Map{
		"data":    result,
		"success": true,
		"message": "",
	})
}
