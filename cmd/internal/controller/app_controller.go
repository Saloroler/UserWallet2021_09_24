package controller

import (
	"UserWallet2021_09_24/cmd/internal/application/registration"
	"UserWallet2021_09_24/cmd/internal/application/transaction"
	"UserWallet2021_09_24/cmd/internal/models"
	"encoding/json"
	"net/http"
	"strings"
)

type Controller struct {
	transactionProcess  transaction.Process
	registrationProcess registration.Process
}

func NewAppController(transactionProcess transaction.Process, registrationProcess registration.Process) Controller {
	return Controller{
		transactionProcess:  transactionProcess,
		registrationProcess: registrationProcess,
	}
}

func (c *Controller) UserRegistration(w http.ResponseWriter, r *http.Request) {
	auth := strings.SplitN(r.Header.Get("Authorization"), " ", 2)
	if len(auth) != 2 || auth[0] != "Basic" {
		printHTTPResult(w, http.StatusUnauthorized, "Need Basic authorization")
		return
	}
	credentials := auth[1]

	var registrationRequest models.EmailRegistrationRequest
	if err := json.NewDecoder(r.Body).Decode(&registrationRequest); err != nil {
		printHTTPResult(w, http.StatusBadRequest, "Invalid json")
		return
	}

	if err := registrationRequest.Validate(); err != nil {
		printHTTPResult(w, http.StatusBadRequest, err.Error())
	}

	user, err := c.registrationProcess.NewUser(registrationRequest.Email, credentials)
	if err != nil {
		printHTTPResult(w, http.StatusInternalServerError, err.Error())
		return
	}

	token, err := c.registrationProcess.CreateAuthTokenForUser(user)
	if err != nil {
		printHTTPResult(w, http.StatusInternalServerError, err.Error())
		return
	}

	printHTTPResult(w, http.StatusCreated, models.UserRegistrationResponse{
		User:  user,
		Token: token,
	})
}
