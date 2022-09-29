package contracts

import (
	"context"
	"crypto/ecdsa"
	"errors"
	"fmt"
	"golang_service/online_test"
	"math/big"
	"os"

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
