package general

/*
	created by Ali Sadikin
	this file will provide all general function that all service (controller,model,repo) can use it

	## Function List
	DoResponse(params)
	WriteResponse(params)
	GetMetaData(params)
*/
import (
	"log"
	"os"
	"strconv"
)

/*
	-- this variable display all custom response code and its message
*/
var CustomResponseCode = map[string]string{
	"011": "Phone number Already Exists",
	"030": "Contact Succesfully Inserted",
	"031": "Contact Succesfully Updated",
	"032": "Contact Succesfully Deleted",
	"033": "Data Not Found",
	"034": "Data Found",
	"035": "Field Number can not be blank",
	"036": "Field Name can not be blank",
}

type Response struct {
	CC       string      `json:"ResponseCode"`
	Message  interface{} `json:"Message"`
	Data     interface{} `json:"Data"`
	MetaData Meta        `json:"Meta"`
}

type QueryParams struct {
	Page  uint8 `json:"Page"`
	Limit uint8 `json:"Limit"`
}

type Meta struct {
	Page Pagination `json:"Pagination"`
}

type Pagination struct {
	Count uint8  `json:"Count"`
	URL   string `json:"URL"`
}

/*
	-- func DoResponse(params) will display json data when invoked (generally) by controller
*/
func DoResponse(customeResposeCode string, message interface{}, data interface{}, meta interface{}) Response {
	// assertion
	assertedMeta := meta.(Meta)
	res := Response{CC: customeResposeCode, Message: message, Data: data, MetaData: assertedMeta}
	return res
}

/*
	-- func WriteErrorLog(param) will write error log and save it in to error.log file
*/
func WriteErrorLog(errorText string) {
	// If the file doesn't exist, create it, or append to the file
	f, err := os.OpenFile("error.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if _, err := f.Write([]byte(errorText + "\n")); err != nil {
		log.Fatal(err)
	}
	if err := f.Close(); err != nil {
		log.Fatal(err)
	}

}

/*
	-- func GetMetaData(params) will display meta data / pagination on json response
*/
func GetMetaData(PageCount uint8, PageURL string) interface{} {
	if PageCount < 1 {
		PageURL = "Not Available"
	}
	meta := Meta{
		Page: Pagination{
			Count: PageCount,
			URL:   PageURL,
		},
	}

	return meta
}

/*
	-- func SetQueryParams(param) will set query string params
*/
func SetQueryParams(allKeys map[string]interface{}) interface{} {
	keys := ""
	mark := "?"
	index := 0
	for key, val := range allKeys {
		if index > 0 {
			mark = "&"
		}

		index++
		keys = keys + mark + key + "=" + val.(string)
	}

	return keys
}

func isNumeric(s string) bool {
	_, err := strconv.ParseFloat(s, 64)
	return err == nil
}
