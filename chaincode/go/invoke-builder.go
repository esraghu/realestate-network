package main

import (
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

func addProjectRecord(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument count")
	}

	// obtain the project details from arguments
	var project Project
	if err := json.Unmarshal([]byte(args[0]), &project); err != nil {
		return shim.Error(err.Error())
	}

	ProjectID := sha1Hash(project.ProjectName)
	fmt.Println("The key for the project is: ", ProjectID)

	err := stub.PutState(ProjectID, []byte(args[0]))
	if err != nil {
		return shim.Error(err.Error())
	} else {
		return shim.Success(nil)
	}

}
