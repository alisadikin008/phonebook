package repository

/*
	Created by Ali Sadikin
	this file is another layer between logic and model database

	# Function List
	GetContactStruct()
	OfNumber(param)
	IsNumberExist(param)
	GetData(params)
	GetOne(params)
	PostData(params)
	CountRecords

*/
import (
	"errors"
	general "phonebook/general"
	phoneMod "phonebook/model"
)

/*
	-- func GetContactStruct() is intended to return contact object/struct
*/
func GetContactStruct() phoneMod.Contact {
	var contactStruct phoneMod.Contact
	return contactStruct
}

/*
	-- func OfNumber(number) is layer to invoke model data by phone number
*/
func OfNumber(stringNumber string) (phoneMod.Contact, error) {
	data := phoneMod.OfNumber(stringNumber)
	return data, nil
}

/*
	-- func IsEmailExist(param) to return whether email exists in database or no
*/
func IsNumberExist(stringNumber string) bool {
	data := phoneMod.OfNumber(stringNumber)
	if data.Number != "" {
		return true
	}

	return false
}

/*
	-- func GetData(param) is a layer to invoke all contact data by giving query string / params on request
*/
func GetData(params map[string]interface{}) (interface{}, error) {
	data := phoneMod.GetData(params)
	if len(data) < 1 {
		return nil, errors.New(general.CustomResponseCode["033"])
	}
	return data, nil
}

/*
	-- func GetOne(params) is a layer to invoke contact data/model by its attribute
*/
func GetOne(attribute interface{}, value interface{}) (interface{}, error) {
	data := phoneMod.GetOne(attribute, value)
	if len(data) < 1 {
		return nil, errors.New(general.CustomResponseCode["033"])
	}
	return data, nil
}

/*
	-- func PostData(params) will invoke contact data /model to insert new one
*/
func PostData(data phoneMod.Contact) (interface{}, error) {
	if data.Number == "" {
		return nil, errors.New(general.CustomResponseCode["035"])
	}

	if data.Name == "" {
		return nil, errors.New(general.CustomResponseCode["036"])
	}

	if IsNumberExist(data.Number) {
		return nil, errors.New(general.CustomResponseCode["011"])
	}

	trx := phoneMod.PostData(data)
	return trx, nil
}

/*
	-- func CountRecords(param) will invoke a sum of records in contact data/model
*/
func CountRecords(params map[string]interface{}) (int, error) {
	number := phoneMod.CountRecords(params)
	if number < 1 {
		return 0, errors.New(general.CustomResponseCode["033"])
	}
	return number, nil
}

/*
	-- func PutOne(id,data) will invoke contact model to update one data by ID
*/
func PutOne(id int, data phoneMod.Contact) (interface{}, error) {
	if data.Number == "" {
		return nil, errors.New(general.CustomResponseCode["035"])
	}

	if data.Name == "" {
		return nil, errors.New(general.CustomResponseCode["036"])
	}

	result, err := phoneMod.PutOne(id, data)
	return result, err
}

/*
	-- func DeleteOne(id) will soft delete contact by its ID
*/

func DeleteOne(id int) (interface{}, error) {
	//check by ID
	isExist := phoneMod.GetOne("id", id)
	if len(isExist) == 0 {
		return nil, errors.New(general.CustomResponseCode["034"])
	}

	result, err := phoneMod.DeleteOne(id)
	return result, err
}
