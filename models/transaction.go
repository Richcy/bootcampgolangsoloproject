package models

import (
	"gorm.io/gorm"
)

type Transaction struct {
	gorm.Model
	Id        int     `form:"id" json:"id" validate:"required"`
	Name      string  `form:"name" json:"name" validate:"required"`
	Image     string  `form:"image" json:"image" validate:"required"`
	Deskripsi string  `form:"desc" json:"desc" validate:"required"`
	Quantity  int     `form:"quantity" json:"quantity" validate:"required"`
	Price     float32 `form:"price" json:"price" validate:"required"`
	ProductID int     `form:"productid" json:"productid" validate:"required"`
}

// CRUD
func CreateTransaction(db *gorm.DB, newTransaction *Transaction) (err error) {
	err = db.Create(newTransaction).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadTransactions(db *gorm.DB, transactions *[]Transaction) (err error) {
	err = db.Find(transactions).Error
	if err != nil {
		return err
	}
	return nil
}
func ReadTransactionById(db *gorm.DB, transactions *Transaction, id int) (err error) {
	err = db.Where("id=?", id).First(transactions).Error
	if err != nil {
		return err
	}
	return nil
}
func UpdateTransaction(db *gorm.DB, transaction *Transaction) (err error) {
	db.Save(transaction)

	return nil
}
func DeleteTransactionById(db *gorm.DB, transaction *Transaction, id int) (err error) {
	db.Where("id=?", id).Delete(transaction)

	return nil
}
