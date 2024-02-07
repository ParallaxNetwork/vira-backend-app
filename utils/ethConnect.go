package utils

import (
	"context"
	"crypto/ecdsa"
	"log"
	"math/big"
	"os"
	"vira-backend-app/contracts"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/joho/godotenv"
)

func ContractConnect() (*contracts.Vira, *ethclient.Client) {
	err := godotenv.Load(".env")
	rpcUrl := os.Getenv("POLYGON_RPC_URL")

	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	client, err := ethclient.Dial(rpcUrl)

	if err != nil {
		log.Fatal(err)
	}

	vira, err := contracts.NewVira(common.HexToAddress("0x13772B90a3174f05Fd6b62d300623F8F19C41d47"), client)
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	return vira, client
}

func WalletConnect(client *ethclient.Client) *bind.TransactOpts {
	err := godotenv.Load(".env")
	walletPrivateKey := os.Getenv("PRIVATEKEY_ADMIN")
	if err != nil {
		log.Fatalf("Error loading .env file: %s", err.Error())
	}

	privateKey, err := crypto.HexToECDSA(walletPrivateKey)
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("error casting public key to ECDSA")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	gasPrice, err := client.SuggestGasPrice(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, big.NewInt(80001))
	if err != nil {
		log.Fatal(err)
	}
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)
	auth.GasLimit = uint64(300000)
	auth.GasPrice = gasPrice

	return auth
}
