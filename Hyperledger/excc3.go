// sum é uma função auxiliar que calcula a soma de dois números
func (t *SimpleChaincode) sum(stub shim.ChaincodeStubInterface, args []string) pb.Response {
	var a, b int
	var err error

	if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}

	a, err = strconv.Atoi(args[0])
	if err != nil {
		return shim.Error(err.Error())
	}

	b, err = strconv.Atoi(args[1])
	if err != nil {
		return shim.Error(err.Error())
	}
