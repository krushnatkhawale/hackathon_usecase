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

var userIndexStr = "_userindex"

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
	fmt.Printf("Running invoke")
	
	fmt.Println("EHL123LO");
 

	return nil, nil
}
func (t *SimpleChaincode) Query(stub shim.ChaincodeStubInterface, function string, args []string) ([]byte, error) {
	fmt.Printf("Running invoke")
	
	fmt.Println("QQQ");
 

	return nil, nil
}
