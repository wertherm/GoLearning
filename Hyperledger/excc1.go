package main

import (
	"fmt"
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
	err = stub.PutState("foo", []byte("bar"))
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
	if function == "set" {
		return t.set(stub, args)
	} else if function == "get" {
		return t.get(stub, args)
	}

	return shim.Error("Invalid invoke function name. Expecting \"set\" or \"get\"")
}

// set é uma função auxiliar que atualiza o valor de uma variável
func (t *SimpleChaincode) set(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, value string
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	key = args[0]
	value = args[1]

	// Atualizar o valor da variável aqui
	err = stub.PutState(key, []byte(value))
	if err != nil {
		return shim.Error(err.Error())
	}

	return shim.Success(nil)
}

// get é uma função auxiliar que obtém o valor de uma variável
func (t *SimpleChaincode) get(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var key, value string
	var err error

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	key = args[0]

	// Obter o valor da variável aqui
	valueBytes, err := stub.GetState(key)
	if err != nil {
		return shim.Error(err.Error())
	}

	value = string
