package main

import (
	"fmt"
	"log"
	"net/http"

	"golang_service/middlewares"
	"golang_service/routes"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error getting env, %v", err)
	}

	router := mux.NewRouter()

	router.Handle("/auth/login", http.HandlerFunc(routes.Login)).Methods("POST")
	router.Handle("/auth/register", http.HandlerFunc(routes.Register)).Methods("POST")
	router.Handle("/choose-admin", http.HandlerFunc(middlewares.CheckJwt(routes.ChooseAdmin))).Methods("POST")
	router.Handle("/create-exam", http.HandlerFunc(middlewares.CheckJwt(routes.CreateExams))).Methods("POST")
	router.Handle("/submit-test", http.HandlerFunc(middlewares.CheckJwt(routes.SubmitTest))).Methods("POST")
	router.Handle("/end-submit", http.HandlerFunc(middlewares.CheckJwt(routes.EndSubmit))).Methods("POST")
	router.Handle("/submit-answer", http.HandlerFunc(middlewares.CheckJwt(routes.SubmitAnswer))).Methods("POST")
	router.Handle("/caculate-winner", http.HandlerFunc(middlewares.CheckJwt(routes.CaculateWinner))).Methods("POST")
	router.Handle("/withdraw", http.HandlerFunc(middlewares.CheckJwt(routes.Withdraw))).Methods("POST")
	router.Handle("/check-account-balance", http.HandlerFunc(middlewares.CheckJwt(routes.CheckAccountBalance))).Methods("POST")
	router.Handle("/check-contract-balance", http.HandlerFunc(middlewares.CheckJwt(routes.CheckContractBalance))).Methods("GET")
	fmt.Println("Listening to port 8000")
	log.Fatal(http.ListenAndServe(":8000", router))
}
