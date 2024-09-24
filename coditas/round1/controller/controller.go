package controller

import (
	"encoding/json"
	"fmt"
	"net/http"
	"regexp"

	"github.com/VallabhLakadeTech/coditas/round1/model"
	"github.com/go-playground/validator"
)

func Middleware(w http.ResponseWriter, r *http.Request) {

	// startTime = time.Now()
	// r.

}

func SavePANDetails(w http.ResponseWriter, r *http.Request) {
	panDetails := model.PANDetails{}

	decode := json.NewDecoder(r.Body)
	err := decode.Decode(&panDetails)
	if err != nil {
		statusCode := http.StatusInternalServerError
		customError := model.CustomResponse{
			Msg:        "error decoding input",
			StatusCode: statusCode,
			Err:        err,
		}
		sendResponse(w, customError)
	}

	validator := validator.New()
	validator.RegisterValidation("panValidator", panValidator)
	validator.RegisterValidation("mobileValidator", mobileValidator)

	err = validator.Struct(panDetails)

	if err != nil {
		// var errMsg []string
		// for _, err := range err.(validator.ValidationErrors) {
		// 	errMsg = append(errMsg, fmt.Sprintf("Field '%s' failed validation: %s\n", err.Field(), err.Tag()))
		// }
		statusCode := http.StatusBadRequest
		customError := model.CustomResponse{
			Msg:        fmt.Sprintf("invalid input: %v", err.Error()),
			StatusCode: statusCode,
			Err:        err,
		}
		sendResponse(w, customError)
	}

	w.WriteHeader(http.StatusOK)
}

func sendResponse(w http.ResponseWriter, customError model.CustomResponse) {
	encoder := json.NewEncoder(w)
	err := encoder.Encode(customError)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	w.WriteHeader(customError.StatusCode)
}

func mobileValidator(fl validator.FieldLevel) bool {
	mobile := fl.Field().Int()
	regex := regexp.MustCompile("[0-9]{10}$")
	return regex.MatchString(mobile)

}

func panValidator(fl validator.FieldLevel) bool {
	regex := regexp.MustCompile("[A-Z]{5}[0-9]{4}[A-Z]{1}")
	panNumber := fl.Field().String()
	return regex.MatchString(panNumber)
}
