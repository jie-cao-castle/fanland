package eth

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/ethclient"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
)

type EthContractClient struct {
}

const key = `paste the contents of your *testnet* key json here`

func init() {
	// Create an IPC based RPC connection to a remote node and instantiate a contract binding
	conn, err := ethclient.Dial("/home/karalabe/.ethereum/testnet/geth.ipc")
	if err != nil {
		log.Fatalf("Failed to connect to the Ethereum client: %v", err)
	}

	fromKeystore, err := ioutil.ReadFile(fromKeyStoreFile)
	fromKey, err := keystore.DecryptKey(fromKeystore, password)
	fromPrivkey := fromKey.PrivateKey

	// Create an authorized transactor and spend 1 unicorn
	auth := bind.NewKeyedTransactor(fromPrivkey)
	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0)     // in wei
	auth.GasLimit = uint64(300000) // in units
	auth.GasPrice = gasPrice

	address, types, erc721Token, err := DeployErc721(auth, conn, "", "")
	if err != nil {
		return
	}
	erc721Token.
}
