package model

/*
	created by Ali Sadikin
	this file is intended to return data collection of phone from database

	## Function List
	GetData(param)
	GetOne(params)
	PostData(params)
	CountRecords(param)
	PutOne(params)
	DeleteOne(id)
*/

import (
	config "phonebook/config"
	general "phonebook/general"
	"strconv"

	"github.com/jinzhu/gorm"
)

type Contact struct {
	gorm.Model
	Number string `gorm:"type:varchar(20)" form:"number"`
	Name   string `gorm:"type:varchar(100)" form:"name"`
}

func init() {
	db, _ := config.ConnectDB()
	db.AutoMigrate(&Contact{})
}

/*
	-- func OfNumber() will return one contact by contact number
*/

func OfNumber(stringNumber string) Contact {
	var contactData Contact
	db, _ := config.ConnectDB()
	if err := db.Where("number=?", stringNumber).First(&contactData).Error; err != nil {
		general.WriteErrorLog(err.Error())
	}

	return contactData
}

/*
	-- func GetData() will return all contacts data by query string from database
*/
func GetData(params map[string]interface{}) []*Contact {
	var contactData []*Contact
	db, _ := config.ConnectDB()
	limit, _ := strconv.Atoi(params["limit"].(string))
	if limit == 0 {
		limit = 20
	}

	page, _ := strconv.Atoi(params["page"].(string))
	if page == 0 || page == 1 {
		page = 1
	}

	offset := limit * (page - 1)
	query := db
	query = query.Limit(limit).Offset(offset).Find(&contactData)
	return contactData

}

/*
	-- func GetOne(params) will return one data (contact) by attribute / URI from database
*/
func GetOne(attribute interface{}, value interface{}) []*Contact {
	var contactData []*Contact
	db, _ := config.ConnectDB()
	if err := db.First(&contactData, value).Error; err != nil {
		general.WriteErrorLog(err.Error())
	}

	return contactData
}

/*
	-- func PostData(params) will insert data to cotact table
*/
func PostData(contactData Contact) interface{} {
	db, _ := config.ConnectDB()
	tx := db.Begin()
	if err := tx.Create(&contactData).Error; err != nil {
		tx.Rollback()
		general.WriteErrorLog(err.Error())
		return err
	}

	tx.Commit()
	return contactData
}

/*
	-- func PutOne(param) will update the contact data on database
*/
func PutOne(id int, data Contact) (interface{}, error) {
	db, _ := config.ConnectDB()
	var contactData Contact
	tx := db.Begin()
	if err := tx.First(&contactData, id).Error; err != nil {
		tx.Rollback()
		general.WriteErrorLog(err.Error())
		return nil, err
	}

	tx.Model(&contactData).Update(data)
	tx.Commit()
	return contactData, nil
}

/*
	-- func CountRecords(param) will return sum of all contact by query params
*/
func CountRecords(params map[string]interface{}) int {
	var contactData []*Contact
	var count int
	db, _ := config.ConnectDB()
	query := db
	query = query.Find(&contactData).Count(&count)
	return count
}

/*
	-- func DeleteOne(id) will delete contact by its ID
*/

func DeleteOne(id int) (interface{}, error) {
	db, _ := config.ConnectDB()
	var contactData Contact
	tx := db.Begin()
	if err := tx.First(&contactData, id).Error; err != nil {
		tx.Rollback()
		return nil, err
	}

	tx.Delete(&contactData, id)
	tx.Commit()
	return contactData, nil
}
