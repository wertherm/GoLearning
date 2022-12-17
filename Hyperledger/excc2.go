package main

import (
	"fmt"
	"strconv"
	"github.com/hyperledger/fabric/core/chaincode/shim"
	pb "github.com/hyperledger/fabric/protos/peer"
)

// SimpleChaincode é a implementação do chaincode
type SimpleChaincode struct{}

// Init é a função de inicialização do chaincode
func (t *SimpleChaincode) Init(stub shim.ChaincodeStubInterface) pb.Response {
	// Inicializar as variáveis aqui
	var err error

	// Atualizar o valor das variáveis aqui
	err = stub.PutState("counter", []byte("0"))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// Invoke é a função principal do chaincode
func (t *SimpleChaincode) Invoke(stub shim.ChaincodeStubInterface) pb.Response {
	// Obter os parâmetros do chaincode aqui
	function, args := stub.GetFunctionAndParameters()

	// Executar a lógica do chaincode aqui
	if function == "increment" {
		return t.increment(stub, args)
	} else if function == "get" {
		return t.get(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"increment
