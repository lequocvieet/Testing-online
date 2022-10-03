package contracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"golang_service/online_test"
	"math/big"
	"os"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

func GetContract(privateKeyAddress string, value uint64) (*online_test.OnlineTest, *bind.TransactOpts) {
	// address of etherum testnet node
	local_ganache := "http://127.0.0.1:7545"
	client, err := ethclient.Dial(local_ganache)
	if err != nil {
		panic(err)
	}
	deployed_contract_address := os.Getenv("DEPLOYED_CONTRACT_ADDRESS")
	conn, err := online_test.NewOnlineTest(common.HexToAddress(deployed_contract_address), client)
	if err != nil {
		panic(err)
	}

	var auth *bind.TransactOpts
	if privateKeyAddress != "" {
		auth, err = GetAccountAuth(client, privateKeyAddress, value)
		if err != nil {
			panic(err)
		}
	}

	return conn, auth
}

func GetAccountAuth(client *ethclient.Client, privateKeyAddress string, value uint64) (*bind.TransactOpts, error) {

	privateKey, err := crypto.HexToECDSA(privateKeyAddress)
	if err != nil {
		return nil, err
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, err
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	fmt.Println(fromAddress)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		return nil, errors.New("PendingNonce error!")
	}
	fmt.Println("nounce=", nonce)
	chainID, err := client.ChainID(context.Background())
	if err != nil {
		return nil, errors.New("ChainID error!")
	}
	fmt.Println("ChainID", chainID)
	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		return nil, errors.New("NewKeyedTransactorWithChainID error!")
	}
	auth.Nonce = big.NewInt(int64(nonce))
	if value != 0 {
		auth.Value = big.NewInt(int64(value)) // in wei
	} else {
		auth.Value = big.NewInt(0) // in wei
	}
	auth.GasLimit = uint64(3000000) // in units
	auth.GasPrice = big.NewInt(1000000)

	return auth, nil
}

func ListenChooseAdminEvent() *online_test.OnlineTestChooseAdmin {
	// address of etherum testnet node
	local_ganache := "http://127.0.0.1:7545"
	client, err := ethclient.Dial(local_ganache)
	if err != nil {
		panic(err)
	}
	deployed_contract_address := os.Getenv("DEPLOYED_CONTRACT_ADDRESS")
	conn, err := online_test.NewOnlineTest(common.HexToAddress(deployed_contract_address), client)
	if err != nil {
		panic(err)
	}
	// Watch for a  event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}

	// Setup a channel for results
	channel := make(chan *online_test.OnlineTestChooseAdmin)
	// Start a goroutine which watches new events

	go func() {
		sub, err := conn.WatchChooseAdmin(watchOpts, channel, nil, nil)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()
	}()
	// Receive events from the channel
	for {
		select {
		case event := <-channel:
			fmt.Println("event", event)
			return event
		case <-time.After(8 * time.Second):
			fmt.Println("Timeout")
			return nil
		}

	}
}

func ListenCreateExamEvent(adminAddress common.Address) *online_test.OnlineTestCreateExam {

	// address of etherum testnet node
	local_ganache := "ws://127.0.0.1:7545"
	client, err := ethclient.Dial(local_ganache)
	if err != nil {
		fmt.Println("error", err)
		panic(err)
	}
	deployed_contract_address := os.Getenv("DEPLOYED_CONTRACT_ADDRESS")
	conn, err := online_test.NewOnlineTest(common.HexToAddress(deployed_contract_address), client)
	if err != nil {
		panic(err)
	}
	var adminAddresses []common.Address
	adminAddresses = append(adminAddresses, adminAddress)

	// //test filter event
	// // Filter for a Deposited event
	// fmt.Println("pass here")
	// fmt.Println()
	// filterOpts := &bind.FilterOpts{Context: context.Background(), Start: 0, End: nil}
	// itr, err := conn.FilterCreateExam(filterOpts, adminAddresses)
	// if err != nil {
	// 	fmt.Println(err)
	// 	panic(err)
	// }
	// return itr.Event
	// Loop over all found events
	// for itr.Next() {
	// 	event := itr.Event
	// 	// Print out all caller addresses
	// 	//fmt.Printf(event.Addr.Hex())
	// 	return event
	// }

	// Watch for a  event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}

	// Setup a channel for results
	channel := make(chan *online_test.OnlineTestCreateExam)
	// Start a goroutine which watches new events

	go func() {
		sub, err := conn.WatchCreateExam(watchOpts, channel, nil)
		if err != nil {
			fmt.Println(err)
			panic(err)
		}
		defer sub.Unsubscribe()
	}()
	// Receive events from the channel
	for {
		select {
		case event := <-channel:
			fmt.Println("event", event)
			return event
		case <-time.After(8 * time.Second):
			fmt.Println("Timeout")
			return nil
		}

	}

}

func ListenSubmitTestEvent() *online_test.OnlineTestSubmitTest {
	// address of etherum testnet node
	local_ganache := "http://127.0.0.1:7545"
	client, err := ethclient.Dial(local_ganache)
	if err != nil {
		panic(err)
	}
	deployed_contract_address := os.Getenv("DEPLOYED_CONTRACT_ADDRESS")
	conn, err := online_test.NewOnlineTest(common.HexToAddress(deployed_contract_address), client)
	if err != nil {
		panic(err)
	}
	// Watch for a  event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}

	// Setup a channel for results
	channel := make(chan *online_test.OnlineTestSubmitTest)
	// Start a goroutine which watches new events

	go func() {
		sub, err := conn.WatchSubmitTest(watchOpts, channel, nil)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()
	}()
	// Receive events from the channel
	for {
		select {
		case event := <-channel:
			fmt.Println("event", event)
			return event
		case <-time.After(8 * time.Second):
			fmt.Println("Timeout")
			return nil
		}

	}
}

func ListenEndSubmitEvent() *online_test.OnlineTestEndSubmit {
	// address of etherum testnet node
	local_ganache := "http://127.0.0.1:7545"
	client, err := ethclient.Dial(local_ganache)
	if err != nil {
		panic(err)
	}
	deployed_contract_address := os.Getenv("DEPLOYED_CONTRACT_ADDRESS")
	conn, err := online_test.NewOnlineTest(common.HexToAddress(deployed_contract_address), client)
	if err != nil {
		panic(err)
	}
	// Watch for a  event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}

	// Setup a channel for results
	channel := make(chan *online_test.OnlineTestEndSubmit)
	// Start a goroutine which watches new events

	go func() {
		sub, err := conn.WatchEndSubmit(watchOpts, channel, nil)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()
	}()
	// Receive events from the channel
	for {
		select {
		case event := <-channel:
			fmt.Println("event", event)
			return event
		case <-time.After(8 * time.Second):
			fmt.Println("Timeout")
			return nil
		}

	}
}

func ListenSubmitAnswerEvent() *online_test.OnlineTestSubmitAnswer {
	// address of etherum testnet node
	local_ganache := "http://127.0.0.1:7545"
	client, err := ethclient.Dial(local_ganache)
	if err != nil {
		panic(err)
	}
	deployed_contract_address := os.Getenv("DEPLOYED_CONTRACT_ADDRESS")
	conn, err := online_test.NewOnlineTest(common.HexToAddress(deployed_contract_address), client)
	if err != nil {
		panic(err)
	}
	// Watch for a  event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}

	// Setup a channel for results
	channel := make(chan *online_test.OnlineTestSubmitAnswer)
	// Start a goroutine which watches new events

	go func() {
		sub, err := conn.WatchSubmitAnswer(watchOpts, channel, nil)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()
	}()
	// Receive events from the channel
	for {
		select {
		case event := <-channel:
			fmt.Println("event", event)
			return event
		case <-time.After(8 * time.Second):
			fmt.Println("Timeout")
			return nil
		}

	}
}

func ListenCaculateWinnerEvent() *online_test.OnlineTestCaculateWinner {
	// address of etherum testnet node
	local_ganache := "http://127.0.0.1:7545"
	client, err := ethclient.Dial(local_ganache)
	if err != nil {
		panic(err)
	}
	deployed_contract_address := os.Getenv("DEPLOYED_CONTRACT_ADDRESS")
	conn, err := online_test.NewOnlineTest(common.HexToAddress(deployed_contract_address), client)
	if err != nil {
		panic(err)
	}
	// Watch for a  event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}

	// Setup a channel for results
	channel := make(chan *online_test.OnlineTestCaculateWinner)
	// Start a goroutine which watches new events

	go func() {
		sub, err := conn.WatchCaculateWinner(watchOpts, channel, nil)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()
	}()
	// Receive events from the channel
	for {
		select {
		case event := <-channel:
			fmt.Println("event", event)
			return event
		case <-time.After(8 * time.Second):
			fmt.Println("Timeout")
			return nil
		}

	}
}

func ListenWithdrawEvent() *online_test.OnlineTestWithDraw {
	// address of etherum testnet node
	local_ganache := "ws://127.0.0.1:7545"
	client, err := ethclient.Dial(local_ganache)
	if err != nil {
		panic(err)
	}
	deployed_contract_address := os.Getenv("DEPLOYED_CONTRACT_ADDRESS")
	conn, err := online_test.NewOnlineTest(common.HexToAddress(deployed_contract_address), client)
	if err != nil {
		panic(err)
	}
	// Watch for a  event
	watchOpts := &bind.WatchOpts{Context: context.Background(), Start: nil}

	// Setup a channel for results
	channel := make(chan *online_test.OnlineTestWithDraw)
	// Start a goroutine which watches new events

	go func() {
		sub, err := conn.WatchWithDraw(watchOpts, channel)
		if err != nil {
			panic(err)
		}
		defer sub.Unsubscribe()
	}()
	// Receive events from the channel
	for {
		select {
		case event := <-channel:
			fmt.Println("event", event)
			return event
		case <-time.After(8 * time.Second):
			fmt.Println("Timeout")
			return nil
		}

	}
}
