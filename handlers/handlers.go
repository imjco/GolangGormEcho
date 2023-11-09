package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	validation "github.com/go-ozzo/ozzo-validation/v4"
	"github.com/go-ozzo/ozzo-validation/v4/is"
	"github.com/labstack/echo/v4"
	"github.com/shopspring/decimal"
	"gorm.io/gorm"
	"io"
	"profile/db"
	"profile/model"
)

func GetCustomerList(c echo.Context) error {

	var customerList []model.Customer
	db := db.Db()

	res := db.Preload("GenderType").Find(&customerList)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return c.JSON(404, "No records found")
	}

	return c.JSON(200, customerList)

}

func GetCustomerbyID(e echo.Context) error {

	var customer model.Customer
	db := db.Db()
	res := db.Where("id = ?", e.Param("customerid")).First(&customer)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return e.JSON(404, "No records found")
	}

	return e.JSON(200, customer)

}

func AddCustomer(e echo.Context) error {

	defer e.Request().Body.Close()

	jsonbody, err := ParseRequestBody(e.Request().Body)
	if err != nil {
		return echo.ErrBadRequest
	}

	validationresult := validation.Validate(jsonbody, validation.Map(
		validation.Key("firstname", validation.Required, validation.Length(1, 20)),
		validation.Key("lastname", validation.Required, validation.Length(1, 20)),
		validation.Key("middlename", validation.Skip),
		validation.Key("email", validation.Required, is.Email),
		validation.Key("phone", validation.Required, is.Digit),
		validation.Key("address", validation.Required, validation.Length(1, 250)),
		validation.Key("dateofbirth", validation.Required),
		validation.Key(`gendertypeid`, validation.Required),
	))
	if validationresult != nil {
		return e.JSON(404, validationresult.Error())
	}

	gendertypeidint, err := ToUintAny(jsonbody["gendertypeid"])
	if err != nil {
		return e.JSON(404, "gendertypeid is not a number")
	}
	customer := model.Customer{
		FirstName:    jsonbody["firstname"].(string),
		MiddleName:   jsonbody["middlename"].(string),
		LastName:     jsonbody["lastname"].(string),
		Email:        jsonbody["email"].(string),
		Phone:        jsonbody["phone"].(string),
		Address:      jsonbody["address"].(string),
		DateofBirth:  jsonbody["dateofbirth"].(string),
		GenderTypeID: gendertypeidint,
	}
	db := db.Db()
	res := db.Create(&customer)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return e.JSON(404, "No records found")
	}

	return e.JSON(200, customer.ID)

}

func UpdateCustomer(e echo.Context) error {

	defer e.Request().Body.Close()

	jsonbody, err := ParseRequestBody(e.Request().Body)
	if err != nil {
		return echo.ErrBadRequest
	}

	validationresult := validation.Validate(jsonbody, validation.Map(
		validation.Key("customerid", validation.Required),
		validation.Key("firstname", validation.Required, validation.Length(1, 20)),
		validation.Key("lastname", validation.Required, validation.Length(1, 20)),
		validation.Key("middlename", validation.Skip),
		validation.Key("email", validation.Required, is.Email),
		validation.Key("phone", validation.Required, is.Digit),
		validation.Key("address", validation.Required, validation.Length(1, 250)),
		validation.Key("dateofbirth", validation.Required),
		validation.Key(`gendertypeid`, validation.Required),
	))
	if validationresult != nil {
		return e.JSON(404, validationresult.Error())
	}

	var customer model.Customer
	db := db.Db()
	res := db.Where("id = ?", jsonbody["customerid"].(string)).First(&customer)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return e.JSON(404, "No records found")
	}

	gendertypeidint, err := ToUintAny(jsonbody["gendertypeid"])
	if err != nil {
		return e.JSON(404, "gendertypeid is not a number")
	}

	customer.FirstName = jsonbody["firstname"].(string)
	customer.MiddleName = jsonbody["middlename"].(string)
	customer.LastName = jsonbody["lastname"].(string)
	customer.Email = jsonbody["email"].(string)
	customer.Phone = jsonbody["phone"].(string)
	customer.Address = jsonbody["address"].(string)
	customer.DateofBirth = jsonbody["dateofbirth"].(string)
	customer.GenderTypeID = gendertypeidint

	db.Save(&customer)
	if errors.Is(res.Error, gorm.ErrRecordNotFound) {
		return e.JSON(404, "No records found")
	}

	return e.JSON(200, "Customer Updated")

}

func ParseRequestBody(body io.Reader) (map[string]interface{}, error) {
	var jsonBody map[string]interface{}
	err := json.NewDecoder(body).Decode(&jsonBody)

	if err != nil {
		return nil, err
	} else {
		return jsonBody, nil
	}
}
func ToString(value interface{}) string {
	str := fmt.Sprintf("%v", value)
	return str
}
func ToUintAny(value interface{}) (uint, error) {
	number, err := decimal.NewFromString(ToString(value))

	if err != nil {
		return 0, err
	}

	return uint(number.IntPart()), nil
}
