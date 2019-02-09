package main

import (
	"encoding/json"
	"fmt"
	"strconv"
	"time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

type LogisticsChaincode struct {
}

type Logistics struct {
	LogisticsID          string   `json:"LID"`
	LogisticsCompanyName string   `json:"LCN"`
	ShippingAddress      string   `json:"SA"`
	Shipper              string   `json:"shipper"`
	ReceivingAddress     string   `json:"RA"`
	Receiver             string   `json:"receiver"`
	Actions              []Action `json:"actions"`
	IsCollected          bool     `json:"isCollected"`
	IsArrived            bool     `json:"isArrived"`
	IsSigned             bool     `json:"isSigned"`
}

type Action struct {
	Info string `json:"info"`
	Time int64  `json:"time"`
}

const (
	LOGISTICS_NO_KEY = "LOGISTICS_NO_KEY"
)

func (lcc *LogisticsChaincode) initLogisticsNO(stub shim.ChaincodeStubInterface) error {
	return stub.PutState(LOGISTICS_NO_KEY, []byte(strconv.Itoa(0)))
}

func (lcc *LogisticsChaincode) fetchLogisticsID(stub shim.ChaincodeStubInterface) (string, error) {
	noAsBytes, err := stub.GetState(LOGISTICS_NO_KEY)
	if err != nil {
		return "", err
	}
	noAsStr := string(noAsBytes)
	no, err := strconv.Atoi(noAsStr)
	if err != nil {
		return "", err
	}
	no++
	err = stub.PutState(LOGISTICS_NO_KEY, []byte(strconv.Itoa(no)))
	if err != nil {
		return "", err
	}
	logisticsID := "LOGISTICS_ID_" + noAsStr
	return logisticsID, nil
}

func (lcc *LogisticsChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "init" {
		return lcc.init(stub, args)
	}
	return shim.Error("Invalid Logistics Chaincode function name.")
}

func (lcc *LogisticsChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	function, args := stub.GetFunctionAndParameters()
	if function == "queryLogisticsByID" {
		return lcc.queryLogisticsByID(stub, args)
	} else if function == "queryLogisticsesByShipper" {
		return lcc.queryLogisticsesByShipper(stub, args)
	} else if function == "queryLogisticsesByReceiver" {
		return lcc.queryLogisticsesByReceiver(stub, args)
	} else if function == "queryLogisticsesByLCN" {
		return lcc.queryLogisticsesByLCN(stub, args)
	} else if function == "createLogistics" {
		return lcc.createLogistics(stub, args)
	} else if function == "collectLogistics" {
		return lcc.collectLogistics(stub, args)
	} else if function == "transportLogistics" {
		return lcc.transportLogistics(stub, args)
	} else if function == "signLogistics" {
		return lcc.signLogistics(stub, args)
	}
	return shim.Error("Invalid Logistics Chaincode function name.")
}

func (lcc *LogisticsChaincode) init(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 0 {
		return shim.Error("Incorrect number of arguments. Expecting 0")
	}
	err := lcc.initLogisticsNO(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (lcc *LogisticsChaincode) queryLogisticsByID(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	logisticsAsBytes, err := stub.GetState(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(logisticsAsBytes)
}

func (lcc *LogisticsChaincode) queryLogisticsesByShipper(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	shipper := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"shipper\":\"%s\"}}", shipper)
	iterator, err := stub.GetQueryResult(queryString) //must CouchDB
	if err != nil {
		return shim.Error(err.Error())
	}
	defer iterator.Close()
	logisticses := make(map[string]string)
	for iterator.HasNext() {
		result, err := iterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		logisticses[result.Key] = string(result.Value)
	}
	logisticsesAsBytes, err := json.Marshal(logisticses)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(logisticsesAsBytes)
}

func (lcc *LogisticsChaincode) queryLogisticsesByReceiver(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	receiver := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"receiver\":\"%s\"}}", receiver)
	iterator, err := stub.GetQueryResult(queryString) //must CouchDB
	if err != nil {
		return shim.Error(err.Error())
	}
	defer iterator.Close()
	logisticses := make(map[string]string)
	for iterator.HasNext() {
		result, err := iterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		logisticses[result.Key] = string(result.Value)
	}
	logisticsesAsBytes, err := json.Marshal(logisticses)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(logisticsesAsBytes)
}

func (lcc *LogisticsChaincode) queryLogisticsesByLCN(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	logisticsCompanyName := args[0]
	queryString := fmt.Sprintf("{\"selector\":{\"LCN\":\"%s\"}}", logisticsCompanyName)
	iterator, err := stub.GetQueryResult(queryString) //must CouchDB
	if err != nil {
		return shim.Error(err.Error())
	}
	defer iterator.Close()
	logisticses := make(map[string]string)
	for iterator.HasNext() {
		result, err := iterator.Next()
		if err != nil {
			return shim.Error(err.Error())
		}
		logisticses[result.Key] = string(result.Value)
	}
	logisticsesAsBytes, err := json.Marshal(logisticses)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(logisticsesAsBytes)
}

func (lcc *LogisticsChaincode) createLogistics(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	logisticsID, err := lcc.fetchLogisticsID(stub)
	if err != nil {
		return shim.Error(err.Error())
	}
	createAction := Action{
		Info: "createAction",
		Time: time.Now().Unix(),
	}
	logistics := Logistics{
		LogisticsID:          logisticsID,
		LogisticsCompanyName: args[0],
		ShippingAddress:      args[1],
		Shipper:              args[2],
		ReceivingAddress:     args[3],
		Receiver:             args[4],
		Actions:              []Action{createAction},
		IsCollected:          false,
		IsArrived:            false,
		IsSigned:             false,
	}
	logisticsAsBytes, err := json.Marshal(logistics)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(logisticsID, logisticsAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (lcc *LogisticsChaincode) collectLogistics(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	logisticsID := args[0]
	logisticsAsBytes, err := stub.GetState(logisticsID)
	if err != nil {
		return shim.Error(err.Error())
	}
	logistics := Logistics{}
	err = json.Unmarshal(logisticsAsBytes, &logistics)
	if err != nil {
		return shim.Error(err.Error())
	}
	if logistics.IsCollected {
		return shim.Error("Logistics is collected")
	}
	collectAction := Action{
		Info: fmt.Sprintf("collector:%s", args[1]),
		Time: time.Now().Unix(),
	}
	logistics.Actions = append(logistics.Actions, collectAction)
	logistics.IsCollected = true
	logisticsAsBytes, err = json.Marshal(logistics)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(logisticsID, logisticsAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (lcc *LogisticsChaincode) transportLogistics(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}
	logisticsID := args[0]
	logisticsAsBytes, err := stub.GetState(logisticsID)
	if err != nil {
		return shim.Error(err.Error())
	}
	logistics := Logistics{}
	err = json.Unmarshal(logisticsAsBytes, &logistics)
	if err != nil {
		return shim.Error(err.Error())
	}
	if !logistics.IsCollected || logistics.IsArrived {
		return shim.Error("Logistics cant transport")
	}
	transportAction := Action{
		Info: fmt.Sprintf("transporter:%s,from:%s,to:%s", args[1], args[2], args[3]),
		Time: time.Now().Unix(),
	}
	logistics.Actions = append(logistics.Actions, transportAction)
	if args[4] == "true" {
		logistics.IsArrived = true
	}
	logisticsAsBytes, err = json.Marshal(logistics)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(logisticsID, logisticsAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func (lcc *LogisticsChaincode) signLogistics(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	logisticsID := args[0]
	logisticsAsBytes, err := stub.GetState(logisticsID)
	if err != nil {
		return shim.Error(err.Error())
	}
	logistics := Logistics{}
	err = json.Unmarshal(logisticsAsBytes, &logistics)
	if err != nil {
		return shim.Error(err.Error())
	}
	if !logistics.IsArrived || logistics.IsSigned {
		return shim.Error("Logistics cant sign")
	}
	signAction := Action{
		Info: fmt.Sprintf("sign"),
		Time: time.Now().Unix(),
	}
	logistics.Actions = append(logistics.Actions, signAction)
	logistics.IsSigned = true
	logisticsAsBytes, err = json.Marshal(logistics)
	if err != nil {
		return shim.Error(err.Error())
	}
	err = stub.PutState(logisticsID, logisticsAsBytes)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}

func main() {
	err := shim.Start(new(LogisticsChaincode))
	if err != nil {
		fmt.Printf("Error creating new Logistics Chaincode: %s", err)
	}
}
