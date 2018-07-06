package main

import (
	"crypto/sha1"
	"encoding/hex"
	"encoding/json"
	"fmt"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	"github.com/hyperledger/fabric/protos/peer"
)

type HouseProfile struct {
	//HouseID        string
	ProjectID      string
	DoorNo         string
	Building       string
	Street         string
	Suburb         string
	City           string
	ZipCode        string
	Country        string
	Owner          string
	Builder        string
	ApprovalStatus string
}

// Simple structure for a Project
type Project struct {
	ProjectName string
}

// implement a simple chaincode to manage an asset
type SimpleAsset struct {
}

var bcFunctions = map[string]func(shim.ChaincodeStubInterface, []string) peer.Response{
	"add_record":  addRecord,
	"add_project": addProjectRecord,
	//"update_record":  updateRecord,
	"list_record": queryRecord,
	//"list_project": queryProjectRecord,
	//"approve_record": approveRecord,
}

// Init is called during chaincode instatiation to initialise any data
func (t *SimpleAsset) Init(stub shim.ChaincodeStubInterface) peer.Response {
	/*
		// Get the args from the transaction proposal
		args := stub.GetStringArgs()
		if len(args) != 2 {
			return shim.Error("Incorrect arguments. Expecting a key and value pair")
		}
		bcFunc := bcFunctions[args[0]]
		return bcFunc(stub, []string(args[1]))
	*/

	/*
		//We store the key and value in the ledger
		err := stub.PutState(args[0], []byte(args[1]))
		if err != nil {
			return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
		}
	*/
	fmt.Println("Chaincode has been initialized")
	return shim.Success(nil)
}

// Invoke is called per transaction on the chaincode. Each transaction is
// either a 'get' or a 'set' on the asset created by Init function. The 'set'
// method may create a new asset by specifying a new key-value pair.
func (t *SimpleAsset) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	// Extract the function and args from the transaction proposal
	fn, args := stub.GetFunctionAndParameters()

	bcFunc := bcFunctions[fn]
	if bcFunc == nil {
		return shim.Error("Invalid invoke of chaincode")
	}

	return bcFunc(stub, args)
}

func addRecord(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Invalid argument count.")
	}
	var house HouseProfile
	if err := json.Unmarshal([]byte(args[0]), &house); err != nil {
		return shim.Error(err.Error())
	}

	HouseID := sha1Hash(house.ProjectID + house.DoorNo)
	fmt.Println("The composite key for House is: ", HouseID)

	/*err = json.Unmarshal([]byte(args[0]), &ct)
	if err != nil {
		return shim.Error(err.Error())
	}*/

	/*key, err := stub.CreateCompositeKey("profileType", []string{partial.HouseID})
	if err != nil {
		return shim.Error(err.Error())
	} */

	/* value, err := json.Marshal(ct)
	   if err != nil {
	       return shim.Error(err.Error())
	   } */

	err := stub.PutState(HouseID, []byte(args[0]) /*value*/)
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

func queryRecord(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	partial := struct {
		HouseID   string
		DoorNo    string
		ProjectID string
	}{}

	if len(args) == 1 {
		err := json.Unmarshal([]byte(args[0]), &partial)
		if err != nil {
			return shim.Error(err.Error())
		}
		// fmt.Println("The HouseID received is: ", partial.HouseID)
	}

	// direct search if the house id has been passed
	if partial.HouseID != "" {
		house, err := stub.GetState(partial.HouseID)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(house)
	}

	// door no is passed and the project id is also present, then form the
	// house id to search
	if partial.DoorNo != "" && partial.ProjectID != "" {
		partial.HouseID = sha1Hash(partial.ProjectID + partial.DoorNo)
		house, err := stub.GetState(partial.HouseID)
		if err != nil {
			return shim.Error(err.Error())
		}
		return shim.Success(house)
	}

	// if only project id is passed, then we shall search for all houses
	// to that project
	if partial.ProjectID != "" {
		resutlIterator, err := stub.GetQueryResult(partial.ProjectID)
		if err != nil {
			return shim.Error(err.Error())
		}
		defer resutlIterator.Close()

		var houses [][]byte

		for resutlIterator.HasNext() {
			kvResult, err := resutlIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			houses = append(houses, kvResult.Value)
		}

	}

	// at this stage, there is no focus on searching for states by range
	// this code will be commented for now
	/*
		resutlIterator, err := stub.GetStateByRange("000000", "999999")
		if err != nil {
			return shim.Error(err.Error())
		}
		defer resutlIterator.Close()

		for resutlIterator.HasNext() {
			kvResult, err := resutlIterator.Next()
			if err != nil {
				return shim.Error(err.Error())
			}
			var hp HouseProfile
			if err = json.Unmarshal(kvResult.Value, &hp); err != nil {
				return shim.Error(err.Error())
			}
			//fmt.Println("The HouseID found in the ledger is: ", hp.HouseID)
			if hp.HouseID == partial.HouseID {
				// So we have now got the Housing profile that we are querying for.
				// We shall return the value and close the loop now
				return shim.Success(kvResult.Value)
			}
		} */

	// if we have reached so far, it only means that the house that is being
	// searched for couldn't be found.
	return shim.Error("Queried housing profile couldn't be found! Search with different value")
}

/*
// Set stores asset (both key and the value) on the ledger. If the key exists,
// it overrides the value with the new one
func set(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 2 {
        return "", fmt.Errorf("Incorrect arguments. Expecting a key and a value")
    }

    err := stub.PutState(args[0], []byte(args[1]))
    if err!= nil {
        return "", Errorf("Failed to get asset: %s with error : %s", args[0], err)
    }
    return args[1], nil
}

// Get returns the value of the specified asset key
func get(stub shim.ChaincodeStubInterface, args []string) (string, error) {
    if len(args) != 1 {
        return "", fmt.Error("Incorrect arguments. Expecting a key")
    }
    value, err := stub.GetState(args[0]) && err != nil {
        return "", Errorf("Failed to get asset for the assetid: %s with error: %s", args[0], err)
    }
    if value == nil {
        return "", Errorf("Asset not found %s", args[0])
    }
    return args[1], nil
}
*/

func sha1Hash(input string) string {
	h := sha1.New()
	h.Write([]byte(input))
	hash := h.Sum(nil)
	return hex.EncodeToString(hash)
}

// main function starts up the chaincode in the container during instatiate
func main() {
	if err := shim.Start(new(SimpleAsset)); err != nil {
		fmt.Printf("SimpleAsset chaincode fialed to start with error: %s", err)
	}
}
