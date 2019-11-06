package controller

/*
	created by Ali Sadikin
	this is phone controller that invoked by phone route in main app

	## function list
	GetData()
	GetOne()
	PostData()
	PutOne()
	DeleteOne()
*/

import (
	general "phonebook/general"
	phoneRepo "phonebook/repository"
	"strconv"

	"math"

	"github.com/gin-gonic/gin"
)

/*
	-- func PostData() will insert new data with checking phone number first
*/
func PostData(contect *gin.Context) {
	meta := general.GetMetaData(0, "")
	contactData := phoneRepo.GetContactStruct()
	contect.Bind(&contactData)
	result, err := phoneRepo.PostData(contactData)
	if err != nil {
		contect.JSON(200, general.DoResponse("011", err.Error(), result, meta))
		return
	}

	contect.JSON(200, general.DoResponse("030", general.CustomResponseCode["030"], result, meta))
}

/*
	-- func GetData() will retriev all data with query string given
*/
func GetData(contect *gin.Context) {
	queryParams := map[string]interface{}{
		"page":  contect.Query("page"),
		"limit": contect.Query("limit"),
	}

	contactData, err := phoneRepo.GetData(queryParams)
	countRecords, _ := phoneRepo.CountRecords(queryParams)
	limit, _ := strconv.ParseFloat(queryParams["limit"].(string), 64)
	pageCount := math.Ceil(float64(countRecords) / limit)
	baseUrl := contect.Request.URL.Path
	queryParams["page"] = strconv.Itoa(1)
	pageUrl := general.SetQueryParams(queryParams).(string)
	meta := general.GetMetaData(uint8(pageCount), baseUrl+pageUrl)
	if err != nil {
		contect.JSON(200, general.DoResponse("033", err.Error(), contactData, meta))
		return
	}

	contect.JSON(200, general.DoResponse("034", general.CustomResponseCode["034"], contactData, meta))
}

/*
	-- func GetOne() will retrive only one data by its ID
*/
func GetOne(contect *gin.Context) {
	meta := general.GetMetaData(0, "")
	param := "id"
	contactData, err := phoneRepo.GetOne(param, contect.Param(param))
	if err != nil {
		contect.JSON(200, general.DoResponse("033", err.Error(), contactData, meta))
		return
	}

	contect.JSON(200, general.DoResponse("034", general.CustomResponseCode["034"], contactData, meta))
}

/*
	-- func PutData() will update one data by its ID
*/
func PutOne(contect *gin.Context) {
	meta := general.GetMetaData(0, "")
	contactStruct := phoneRepo.GetContactStruct()
	contect.Bind(&contactStruct)
	id, _ := strconv.Atoi(contect.Param("id"))
	result, err := phoneRepo.PutOne(id, contactStruct)
	if err != nil {
		contect.JSON(200, general.DoResponse("033", err.Error(), result, meta))
		return
	}

	contect.JSON(200, general.DoResponse("031", general.CustomResponseCode["031"], result, meta))
}

/*
	-- func DeleteData() will soft delete one data by its ID
*/
func DeleteOne(contect *gin.Context) {
	//check wether id is exitst
	meta := general.GetMetaData(0, "")
	id, _ := strconv.Atoi(contect.Param("id"))
	result, err := phoneRepo.DeleteOne(id)
	if err != nil {
		contect.JSON(200, general.DoResponse("013", err.Error(), result, meta))
		return
	}

	contect.JSON(200, general.DoResponse("032", general.CustomResponseCode["032"], result, meta))
}
