/*
Copyright IBM Corp. 2016 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

		 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package main

//WARNING - this chaincode's ID is hard-coded in chaincode_example04 to illustrate one way of
//calling chaincode from a chaincode. If this example is modified, chaincode_example04.go has
//to be modified as well with the new ID of chaincode_example02.
//chaincode_example05 show's how chaincode ID can be passed in as a parameter instead of
//hard-coding.

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode example simple Chaincode implementation
type Chaincode struct {
}
type Researcher struct {
	FirstName string `json:"first_name"`
	LastName  string `json:"last_name"`
	RegistrationNumber      int32  `json:"registration_number"`
	Company string `json:"company"`
	Countary string `json:"country"`
	CompanyUrl string `json:"company_url"`
}

func (t *Chaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {

	fmt.Println("hello test")
	return shim.Success([]byte("Chaincode initialize successfully"))
}

// this is the invoke method thid will get execute at the time ogf invocation

func (t *Chaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "addResearcher" {
		return addResearcher(stub, args)
	} else if function == "getResearcher" {
		return getResearcher(stub, args)
	}
	fmt.Println("hello invoke")
	return shim.Error("Pleas eneter a valid function name!!!!!!!!!!!")
}

func addResearcher(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	if len(args) != 1 {
		return shim.Error("Argument for insert should be equal to 1")
	}

	user := Researcher{}
	userparseerr := json.Unmarshal([]byte(args[0]), &user)
	if userparseerr != nil {
		return shim.Error(userparseerr.Error())
	}
	key := stub.GetTxID()
	userbytes, usermarserr := json.Marshal(user)
	if usermarserr != nil {
		return shim.Error(usermarserr.Error())
	}
	fmt.Println(user)
	err0 := stub.PutState(key, userbytes)
	if err0 != nil {
		return shim.Error(err0.Error())
	}

	fmt.Println("Printed all the args as given")

	return shim.Success([]byte(key))

}

func getResearcher(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Please provide single parameter as key !!!!!!!!!!!!1")
	}
	returndebytes, err := stub.GetState(args[0])

	if err != nil {
		return shim.Error("Unable to fetch the given key something went wrong ")
	}
	return shim.Success(returndebytes)
}

// Transaction makes payment of X units from A to B

func main() {
	err := shim.Start(new(Chaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
