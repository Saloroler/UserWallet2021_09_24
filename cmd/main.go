package main

import (
	"UserWallet2021_09_24/cmd/internal/application/registration"
	"UserWallet2021_09_24/cmd/internal/application/transaction"
	"UserWallet2021_09_24/cmd/internal/controller"
	"UserWallet2021_09_24/cmd/internal/db/mysql/repo"
	"fmt"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"net"
	"net/http"
	"os"
)

func main() {
	dbConnection := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	db, err := gorm.Open(os.Getenv("DB_DRIVER"), dbConnection)
	if err != nil {
		log.Fatal("Error happened on IgniteGorm err:", err) // here will be retry on again
	}
	defer db.Close()

	//Repo
	userRepo := repo.NewUserRepo(db)
	transactionRepo := repo.NewTransactionRepo(db)
	tokenRepo := repo.NewTokenRepo(db)

	registrationProcess := registration.NewProcess(userRepo, tokenRepo)
	transactionProcess := transaction.NewProcess(transactionRepo)

	// controller
	appController := controller.NewAppController(transactionProcess, registrationProcess)

	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	authRouter := apiRouter.PathPrefix("/auth").Subrouter()
	//transactionRouter := router.PathPrefix("/transaction").Subrouter()

	//router.Use(middlewares) // for middleware

	authRouter.HandleFunc("/email/registration", appController.UserRegistration).Methods(http.MethodPost)
	//transactionRouter.HandleFunc("/transfer", appController.Transfer).Methods(http.MethodPost)
	tcpAddr := net.TCPAddr{Port: 8080}
	log.Println("[INFO] Server is starting on port", 8080)
	if err := http.ListenAndServe(tcpAddr.String(), router); err != nil {
		log.Fatal("Fail to listen: ", err.Error())
	}
}
