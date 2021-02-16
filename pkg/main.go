package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"

	keystore "github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/crypto"
	filekeystore "github.com/ethersphere/bee/pkg/keystore/file"
	"github.com/pborman/uuid"
)

func main() {

	if len(os.Args) != 3 {
		fmt.Println("exportSwarmKey <sourceDir> <password>")
		return
	}

	sourceDir := os.Args[1]
	password := os.Args[2]

	files, err := ioutil.ReadDir(sourceDir)
	if err != nil {
		fmt.Println("error reading source dir :, ", err.Error())
		return
	}
	for _, f := range files {
		fks := filekeystore.New(sourceDir)
		sourceFile := filepath.Join(sourceDir, f.Name())
		keyName := strings.Split(f.Name(), ".")
		privateKeyECDSA, _, err := fks.Key(keyName[0], password)
		if err != nil {
			fmt.Println("error reading key : ", err.Error())
			return
		}

		id := uuid.NewRandom()
		key := &keystore.Key{
			Id:         id,
			Address:    crypto.PubkeyToAddress(privateKeyECDSA.PublicKey),
			PrivateKey: privateKeyECDSA,
		}

		content, err := json.Marshal(key)
		if err != nil {
			fmt.Println("error marshalling key :", err.Error())
			return
		}

		fmt.Println(sourceFile, " : ", string(content))

	}






}



