package routes

import (
	"fmt"
	"golang_service/contracts"
	"math/big"
	"net/http"
	_ "reflect"
	"strconv"

	"golang_service/models"
	res "golang_service/utils"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"

	"github.com/asaskevich/govalidator"
)

var adminEntranceFee string
var userSubmitFee string
var contractOwnerAddr string = "0x09Fa09f6170Aa006fF0582949f0A5428107E82C5"
var contractOwnerPrivateKey string = "da92007b3455ce0af600748143d211107ded4222185ee499354b8762f41128da"
var adminAddr string = "0x374B54B1D0e2019E0D969D38598282E365Aa991A"
var adminPrivateKey string = "c62deaf79f1bb77a3f1877b04bbcf286f959cf2b82cb09aa0a2d0abb6145ec6a"
var user1Addr string = "0x8459846257e0774d650E881d1306DCff6430c625"
var user1PrivateKey string = "fdcfa2cf0ac9e829aec4a0b5e02a82d85b928e949b30a75c06012e5f66a576dc"
var user2Addr string = ""
var user2PrivateKey string = ""

func ChooseAdmin(w http.ResponseWriter, r *http.Request) {
	contractOwner := models.Account{}
	contractOwner.SetAccount(contractOwnerAddr, contractOwnerPrivateKey, "contract owner") //set contract owner account
	adminAddr := r.PostFormValue("adminAddress")
	adminEntranceFee = r.PostFormValue("adminEntranceFee")
	//ToDo: fill variables to model and save to db
	if govalidator.IsNull(adminAddr) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	if govalidator.IsNull(adminEntranceFee) {
		res.JSON(w, 400, "Data can not empty")
		return
	}

	//create transaction on blockchain
	conn, auth := contracts.GetContract(contractOwner.PrivateKey, 0)
	n := new(big.Int)
	adminEntranceFees, _ := n.SetString(adminEntranceFee, 10)
	hexAddress := common.HexToAddress(adminAddr)
	result, err := conn.ChooseAdmin(auth, hexAddress, adminEntranceFees)

	if err != nil {
		res.JSON(w, 500, "chooseAdmin has failed")
		return
	}

	res.JSON(w, 201, result)

}

func CreateExams(w http.ResponseWriter, r *http.Request) {
	admin := models.Account{}
	admin.SetAccount(adminAddr, adminPrivateKey, "admin") //set admin account
	reward := r.PostFormValue("reward")
	userSubmitFee = r.PostFormValue("userSubmitFee")
	if govalidator.IsNull(reward) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	if govalidator.IsNull(userSubmitFee) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	fmt.Println(adminEntranceFee)
	adminEntranceFees, _ := strconv.ParseUint(adminEntranceFee, 10, 64)

	//create transaction on blockchain
	conn, auth := contracts.GetContract(admin.PrivateKey, adminEntranceFees)
	n := new(big.Int)
	m := new(big.Int)
	rewards, _ := n.SetString(reward, 10)
	userSubmitFees, _ := m.SetString(userSubmitFee, 10)
	result, err := conn.CreateExams(auth, rewards, userSubmitFees)
	fmt.Println("adminEntranceFee", adminEntranceFees)
	fmt.Println("error", err)

	if err != nil {
		res.JSON(w, 500, "createExams has failed")
		return
	}

	res.JSON(w, 201, result)
}

func SubmitTest(w http.ResponseWriter, r *http.Request) {
	user := models.Account{}
	user.SetAccount(user1Addr, user1PrivateKey, "user") //set user account
	examCode := r.PostFormValue("examCode")
	answer := r.PostFormValue("answer")
	if govalidator.IsNull(examCode) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	if govalidator.IsNull(answer) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	userSubmitFees, _ := strconv.ParseUint(userSubmitFee, 10, 64)

	//create transaction on blockchain
	conn, auth := contracts.GetContract(user.PrivateKey, userSubmitFees)
	n := new(big.Int)
	examCodes, _ := n.SetString(examCode, 10)
	result, err := conn.SubmitTest(auth, examCodes, answer)

	if err != nil {
		res.JSON(w, 500, "submitTest has failed")
		return
	}

	res.JSON(w, 201, result)
}

func EndSubmit(w http.ResponseWriter, r *http.Request) {
	admin := models.Account{}
	admin.SetAccount(adminAddr, adminPrivateKey, "admin") //set admin account
	examCode := r.PostFormValue("examCode")
	if govalidator.IsNull(examCode) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	//create transaction on blockchain
	conn, auth := contracts.GetContract(admin.PrivateKey, 0)
	n := new(big.Int)
	examCodes, _ := n.SetString(examCode, 10)
	result, err := conn.EndSubmit(auth, examCodes)

	if err != nil {
		res.JSON(w, 500, "endSubmit has failed")
		return
	}

	res.JSON(w, 201, result)

}

func SubmitAnswer(w http.ResponseWriter, r *http.Request) {
	admin := models.Account{}
	admin.SetAccount(adminAddr, adminPrivateKey, "admin") //set admin account
	examCode := r.PostFormValue("examCode")
	answer := r.PostFormValue("answer")
	if govalidator.IsNull(examCode) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	if govalidator.IsNull(answer) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	//create transaction on blockchain
	conn, auth := contracts.GetContract(admin.PrivateKey, 0)
	n := new(big.Int)
	examCodes, _ := n.SetString(examCode, 10)
	result, err := conn.SubmitAnswer(auth, examCodes, answer)

	if err != nil {
		res.JSON(w, 500, "submitAnswer has failed")
		return
	}

	res.JSON(w, 201, result)

}

func CaculateWinner(w http.ResponseWriter, r *http.Request) {
	admin := models.Account{}
	admin.SetAccount(adminAddr, adminPrivateKey, "admin") //set admin account
	examCode := r.PostFormValue("examCode")
	if govalidator.IsNull(examCode) {
		res.JSON(w, 400, "Data can not empty")
		return
	}
	//create transaction on blockchain
	conn, auth := contracts.GetContract(admin.PrivateKey, 0)
	n := new(big.Int)
	examCodes, _ := n.SetString(examCode, 10)
	result, err := conn.CaculateWinner(auth, examCodes)

	if err != nil {
		res.JSON(w, 500, "caculateWinner has failed")
		return
	}

	res.JSON(w, 201, result)

}

func Withdraw(w http.ResponseWriter, r *http.Request) {
	contractOwner := models.Account{}
	contractOwner.SetAccount(contractOwnerAddr, contractOwnerPrivateKey, "contract owner") //set contract owner account
	//create transaction on blockchain
	conn, auth := contracts.GetContract(contractOwner.PrivateKey, 0)
	result, err := conn.Withdraw(auth)

	if err != nil {
		res.JSON(w, 500, "Withdraw has failed")
		return
	}

	res.JSON(w, 201, result)
}

func CheckAccountBalance(w http.ResponseWriter, r *http.Request) {
	accountAddr := r.PostFormValue("accountAddr")
	conn, _ := contracts.GetContract("", 0)
	result, err := conn.GetAccountBalance(&bind.CallOpts{From: common.HexToAddress(accountAddr)})

	if err != nil {
		res.JSON(w, 500, "CheckBalance has failed")
		return
	}

	res.JSON(w, 201, result)
}
func CheckContractBalance(w http.ResponseWriter, r *http.Request) {
	conn, _ := contracts.GetContract("", 0)
	result, err := conn.GetAccountBalance(&bind.CallOpts{})

	if err != nil {
		res.JSON(w, 500, "CheckContractBalance has failed")
		return
	}

	res.JSON(w, 201, result)
}
