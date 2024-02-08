package utils

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"net/http"
	"regexp"
	"strings"
	"unicode"

	"github.com/wepull/bugTracker/models"
)

var (
	httpClient *http.Client
)

func IsValidEmail(email string) bool {
	emailPattern := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`

	// Compile the pattern
	regex := regexp.MustCompile(emailPattern)

	// Check if the email matches the pattern
	return regex.MatchString(email)
}

func IsValidPassword(password string) (bool, error) {
	var err error
	if len(password) < 6 || len(password) > 20 {
		err = fmt.Errorf("Password must be longer than 6 Charecters and shorter than 20 charecters")
		return false, err
	}
	var (
		hasLowercase bool
		hasUppercase bool
		hasSymbol    bool
	)

	for _, char := range password {
		switch {
		case unicode.IsLower(char):
			hasLowercase = true
		case unicode.IsUpper(char):
			hasUppercase = true
		case strings.IndexAny(string(char), "!@#$%^&*()-_=+,.?") != -1:
			hasSymbol = true
		}
	}

	isValid := hasLowercase && hasUppercase && hasSymbol
	if !isValid {
		err = fmt.Errorf("Password should contain at least one lowecase charecter, one uppercase charecter and one symbol")
		return isValid, err
	}

	return isValid, nil
}

func getUnmarshallErrorString(unMarshalErr error) error {
	if ute, ok := unMarshalErr.(*json.UnmarshalTypeError); ok {
		return errors.New("Input " + ute.Value + " for field " + ute.Field + " is incorrect.")
	} else {
		return unMarshalErr
	}
}

func ReadAndParseInput(w http.ResponseWriter, r *http.Request, input interface{}) error {

	body, err := io.ReadAll(r.Body)
	if err != nil {
		log.Printf("Error in Reading request Body %+v", err)
		return err
	}
	if err1 := r.Body.Close(); err1 != nil {
		log.Printf("Error in Closing body %s\n", err1.Error())
		return err1
	}

	if err2 := json.Unmarshal(body, input); err2 != nil {

		err2 = getUnmarshallErrorString(err2)
		log.Printf("Unmarshalling Error. %+v ", err2)

		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		output := models.ApiResp{
			ResponseCode:        1,
			ResponseDescription: err2.Error(),
		}
		if err3 := json.NewEncoder(w).Encode(output); err3 != nil {
			log.Printf("Json Encoding Error. %+v", err3)
		}
		return err2
	}
	return nil
}
