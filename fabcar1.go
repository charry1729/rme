package main

import (
	"bytes"
	"encoding/json"
	"fmt"
        "strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------
//   Sl.No           Name                   Date                    Description
//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------
//    01          Rohith Seetha            23-07-2019               Initial Version
//-----------------------------------------------------------------------------------------------------------------------------------------------------------------------------      


// SimpleChaincode example simple Chaincode implementation
type SimpleChaincode struct {
}

type user struct {
	ObjectType            string  `json:"docType"` //docType is used to distinguish the various types of objects in state database
        Uuid           string  `json:"Uuid"`            //SN-04
        FirstName         string  `json:"FirstName"`
        LastName           string  `json:"LastName"`
        UserName         string  `json:"UserName"`
        AddressOne           string  `json:"AddressOne"`
        AddressTwo          string  `json:"AddressTwo"`
        City           string  `json:"City"`
        State               string  `json:"State"`
        Country          string  `json:"Country"`
        PinCode              string  `json:"Pincode"`
        DOB           string  `json:"DOB"`
        Contact     string  `json:"Contact"`
        CompanyName        string  `json:"CompanyName"`

        UserStatus   string  `json:"UserStatus"`
        
        Occupations    []string  `json:"Occupations"`
        Specializations         []string  `json:"Specializations"`
        MarketPlaces            []string  `json:"MarketPlaces"`
        ParticipationTypes            []string  `json:"ParticipationTypes"`
        MembershipPools           []string  `json:"MembershipPools"`
        UserTypes  string  `json:"UserTypes"`
        IsEmailVerified      string  `json:"IsEmailVerified"`
        KYCstatus         string `json:"KYCstatus"`
        KYTstatus           string  `json:"KYTstatus"`
        RefferalCode               string  `json:"RefferalCode"`
        RefferedByCode  string  `json:"RefferedByCode"`
//KYT
        Pancard string `json:"pancard"`
        KytdoctypeStatus          string     `json:"kytdoctypeStatus"`
        Pancardnotpresentflag string `json:"pancardnotpresentflag"`
        OtherOccupation            string  `json:"OtherOccupation"`
        SupplierOf            string  `json:"SupplierOf"`
        UserSecretKey                RmeUserSecretKey  `json:"UserSecretKey"`
     }

type RmeUserSecretKey struct {
        ObjectType string  `json:"docType"`
        Id string  `json:"Id"`           
        Password string  `json:"Password"`
        PrivateKey string  `json:"PrivateKey"`
        PublicKey         string  `json:"PublicKey"`
        RmeUser           string  `json:"RmeUser"`
        UserBackupPhrase          string  `json:"UserBackupPhrase"`
        UserWalletAddress           string  `json:"UserWalletAddress"`
 } 




// ===================================================================================
// Main
// ===================================================================================
func main() {
	err := shim.Start(new(SimpleChaincode))
	if err != nil {
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}

// Init initializes chaincode
// ===========================
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
        fmt.Println("entering invoice init successfully")
	return shim.Success(nil)
}

// Invoke - Our entry point for Invocations
// ========================================
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	//fmt.Println("invoke is running " + function)

	// Handle different functions
	if function == "signup" { //create a new User
		return t.signup(stub, args)
	// } else if function == "initUser" { //read a User
	// 	return t.siginitUsernup(stub, args)	
        }else if function == "readUser" { //read a User
		return t.readUser(stub, args)	
        } else if function == "deleteUser" {
                return t.deleteUser(stub, args)
        }

	fmt.Println("invoke did not find func: " + function) //error
	return shim.Error("Received unknown function invocation")
}

// ============================================================
// initUser - create a new User, store into chaincode state
// ============================================================
func (t *SimpleChaincode) signup(stub shim.ChaincodeStubInterface, args []string) pb.Response {

	//var poVal string

	// if len(args) != 28 {
	// 	return shim.Error("Incorrect number of arguments. Expecting 28")
	// }

        fmt.Println("- start Signup user")
        
      //  Uuid            :=  args[0]
        FirstName          :=  args[0]
        LastName           :=  args[1]
        UserName         :=  args[2]
        AddressOne           :=  args[3]
        AddressTwo          :=  args[4]
        City           :=  args[5]
        State                :=  args[6]
        Country          :=  args[7]
        PinCode              :=  args[8]
        DOB           :=  args[9]
        Contact     :=  args[10]
        CompanyName := args[11]
        //UserStatus   :=  args[12]
        Status1,err := strconv.Atoi(args[12])
        if err != nil {
                return shim.Error("userstatus argument must be a numeric string")
        }


        UserStatus := strconv.Itoa(Status1)


		var oc []string 
		oc = append(oc,args[13])
		Occupations    :=  oc
		
		var sp []string 
		sp = append(sp,args[14])
		Specializations    :=  sp

                var mp []string 
		mp = append(mp,args[15])
		MarketPlaces    :=  mp
		
		var pt []string 
		pt = append(pt,args[16])
		ParticipationTypes    :=  pt
		
		var mpl []string 
		mpl = append(mpl,args[17])
		MembershipPools    :=  mpl
		
		// MarketPlaces            :=  args[15]
        // ParticipationTypes            :=  args[16]
        // MembershipPools           :=  args[17]
        UserTypes :=  args[18]
        IsEmailVerified   :=  args[19]
        KYCstatus          :=  args[20]
	KYTstatus :=  args[21]
        RefferalCode      :=  args[22]
        RefferedByCode   :=  args[23]
             
       // var UserSecretKey RmeUserSecretKey
	Id  := args[24]
	Password    := args[25]
	PrivateKey  := args[26]
	PublicKey   := args[27]
	RmeUser := args[28]
	UserBackupPharse  := args[29]
        UserWalletAddress  := args[30]
          
        Pancard := args[31]
        DocStatus1,err := strconv.Atoi(args[32])
        if err != nil {
                return shim.Error("doc type argument must be a numeric string")
        }
        
//	Status:= stats(Status1)

        KytdoctypeStatus := strconv.Itoa(DocStatus1)
        Pancardnotpresentflag := args[33]
        OtherOccupation:= args[34]
        SupplierOf:= args[35]


	uObjectType := "remusersecretkey"
        UserSecretKey := RmeUserSecretKey{uObjectType,	  
                  Id,
		  Password,
		  PrivateKey,
		  PublicKey,
		  RmeUser,
		  UserBackupPharse,
                  UserWalletAddress}

          
//KycdoctypeStatus 

        //   kytStatus3,err := strconv.Atoi(args[30])
        //           if err != nil {
        //                   return shim.Error("10th argument must be a numeric string")
        //           }
                  
        //   //	Status:= stats(Status1)
                    
        //   kytStatus4:= strconv.Itoa(kytStatus3)


				objectType := "user"
                                user := user{objectType, UserName, FirstName, LastName, UserName, AddressOne, AddressTwo, City , State,
                                         Country , PinCode, DOB, Contact, CompanyName, UserStatus, Occupations , Specializations, MarketPlaces, 
					ParticipationTypes, MembershipPools, UserTypes, IsEmailVerified,KYCstatus,KYTstatus, RefferalCode, 
					RefferedByCode, Pancard,KytdoctypeStatus,Pancardnotpresentflag,OtherOccupation,SupplierOf,UserSecretKey}  //SN-04	
				userJSONasBytes, err3 := json.Marshal(user)
				if err3 != nil {
					return shim.Error(err3.Error())
				}

			// === Save invoice to state ===
				err4 := stub.PutState(UserName, userJSONasBytes)
				if err4 != nil {
					return shim.Error(err4.Error())
                                } 
                            
	fmt.Println("- end init invoice")
	return shim.Success([]byte("Success"))
}



// ========================================================
// function to fetch invoice request id for auto-increment
// ========================================================

func fetch(stub shim.ChaincodeStubInterface, args []string) (string, error) {
        if len(args) != 1 {
                return "", fmt.Errorf("Incorrect arguments. Expecting a key")
        }
    
        value, err10 := stub.GetState(args[0])
        if err10 != nil {
                return "", fmt.Errorf("Failed to get asset: %s with error: %s", args[0], err10)
        }
        if value == nil {
                return "", fmt.Errorf("Asset not found: %s", args[0])
        }
        return string(value), nil
    }
    
func insert(stub shim.ChaincodeStubInterface, args []string) (string, error) {

        // Set up any variables or assets here by calling stub.PutState()
    
        // We store the key and the value on the ledger
        err9 := stub.PutState(args[0], []byte(args[1]))
        if err9 != nil {
                return "", fmt.Errorf("Failed to set asset: %s", args[0])
    //            return shim.Error(fmt.Sprintf("Failed to create asset: %s", args[0]))
        }
        return args[1], nil
    }
    



// ===============================================
// readInvoice - read an invoice from chaincode state
// ===============================================
func (t *SimpleChaincode) readUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var userName, jsonResp string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting PO number to query")
	}

	userName = args[0]
	valAsbytes, err := stub.GetState(userName) //get the invoice from chaincode state
	if err != nil {
		jsonResp = "{\"Error\":\"Failed to get state for " + userName + "\"}"
		return shim.Error(jsonResp)
	} else if valAsbytes == nil {
		jsonResp = "{\"Error\":\"Invoice does not exist: " + userName + "\"}"
		return shim.Error(jsonResp)
	}

	return shim.Success([]byte(valAsbytes))
}


// =========================================================================================
// getQueryResultForQueryString executes the passed in query string.
// Result set is built and returned as a byte array containing the JSON results.
// =========================================================================================
func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) (string,error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)
        var count int

	resultsIterator, err1 := stub.GetQueryResult(queryString)
	if err1 != nil {
		return "", err1
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
         count = count + 1
		queryResponse, err2 := resultsIterator.Next()
		if err2 != nil {
			return  "", err2
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")
        totalCount := strconv.Itoa(count)
       // countAsBytes, err3 := json.Marshal(totalCount)

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", totalCount)
        return (totalCount), nil
}

// ===============================================
// deleteInvoice - delete an invoice request from chaincode state
// ===============================================
func (t *SimpleChaincode) deleteUser(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var userName string

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting PO number to query")
	}

	userName = args[0]
	err := stub.DelState(userName) //get the invoice from chaincode state
	if err != nil {
             return shim.Error(fmt.Sprintf("Failed to delete invoice req: %s", args[0]))
        }
        return shim.Success(nil);

}

