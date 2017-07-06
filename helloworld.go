

package main

import (
	"errors"
	"fmt"
	"strconv"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

var myLogger = logging.MustGetLogger("asset_mgm")

// HelloWorld example simple Chaincode implementation
type HelloWorld struct {
}

// Init is called during Deploy transaction after the container has been
func (t *HelloWorld) Init(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	myLogger.Debug("HelloWorld Init is called!")

	if len(args) != 0 {
		return nil, errors.New("Incorrect number of arguments. Expecting 0")
	}

	// Set the admin
	// The metadata will contain the certificate of the administrator

	err := stub.PutState("noHelloWorlds", []byte("0"))
	if err != nil {
		return nil, err
	}

	myLogger.Debug("Init HelloWorld...done")

	return nil, nil
}

// Invoke is called for every Invoke transactions. The chaincode may change
// its state variables
func (t *HelloWorld) Invoke(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	fmt.Printf("HelloWorld in Invoke with function %s!\n", function)
	if function != "invoke" {
		return nil, errors.New("Invalid invoke function name. Expecting \"invoke\"")
	}

	b, err := stub.GetState("noHelloWorlds")
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	noevts, _ := strconv.Atoi(string(b))

	var theArgs string

	for _, s := range args {
		theArgs = theArgs + "," + s
	}
	fmt.Printf("theArgs: %s ", theArgs)

	err = stub.PutState("noHelloWorlds", []byte(strconv.Itoa(noevts+1)))
	if err != nil {
		return nil, err
	}

	return nil, nil
}

// Query is called for Query transactions. The chaincode may only read
// (but not modify) its state variables and return the result
func (t *HelloWorld) Query(stub *shim.ChaincodeStub, function string, args []string) ([]byte, error) {
	myLogger.Debug("HelloWorld Query is called!\n")
	if function != "query" {
		return nil, errors.New("Invalid query function name. Expecting \"query\"")
	}

	b, err := stub.GetState("noHelloWorlds")
	if err != nil {
		return nil, errors.New("Failed to get state")
	}
	jsonResp := "{\"NoHelloWorlds\":\"" + string(b) + "\"}"
	return []byte(jsonResp), nil
}

func main() {
	myLogger.Debug("HelloWorld Main\n")
	err := shim.Start(new(HelloWorld))
	if err != nil {
		fmt.Printf("Error starting HelloWorld chaincode: %s\n", err)
	}
}
