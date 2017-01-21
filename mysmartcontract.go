/*
  This is smart contract to be deployed in BlockChain
*/

package main

import (
	"fmt"
	"github.com/hyperledger/fabric/core/chaincode/shim"
)

type SimpleChaincode struct {
}

var userIndexStr = "myindex"

func main() {
	fmt.Println("ELITE HACKERS... Era begins...");
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {

	fmt.Println("EHLLO");


	return nil, nil
}
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	
	
	if function == "addKYC" {
        return t.addKYC(stub, args)
    } 
    fmt.Println("invoke did not find func: " + function)

	return nil, nil
}


func (t *SimpleChaincode) addKYC(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


	var err error
	var str string


    fmt.Println("running addKYC@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@()")
    fmt.Println("running addKYC@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@()")
    fmt.Println("running addKYC@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@()")
    fmt.Println("running addKYC@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@()")
    fmt.Println("running addKYC@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@@()")

    fmt.Println( args[0] )
    fmt.Println( args[1] )
    fmt.Println( args[2] )
    fmt.Println( args[3] )
    fmt.Println( args[4] )


	str = `{"id": "` + args[0] + `", "name": "` + args[1] + `", "bname": "` + args[2] + `", "kycdate": "` + args[3] + `", "validTill": "` + args[4] + `", "document": "` + args[5] + `"}`

	err = stub.PutState(args[0], []byte(str))

	if err != nil {
		return nil, err
	}

	return []byte("Congratulations!, We guarantee your records"), nil
}

func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
    fmt.Println("query is running " + function)

    // Handle different functions
    if function == "getKYC" {                       //read a variable
        return t.getKYC(stub, args)
    }
    fmt.Println("query did not find func: " + function)

	return []byte("Received unknown function query: " + function), nil
}

func (t *SimpleChaincode) getKYC(stub shim.ChaincodeStubInterface, args []string) ([]byte, error) {


	var err error
	var accountId string
	var jsonResp string

	accountId = args[0]
	valAsbytes, err := stub.GetState(accountId) //get the var from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + accountId + "\"}"
		return []byte(jsonResp), nil
	}

	return valAsbytes, nil
}
