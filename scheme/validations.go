package scheme

import (
	"encoding/json"
	"errors"
	"fmt"
	"reflect"
	"strings"
)

// define allowed methods with there types using reflect and number of arguments
// we use reflect to know the type recieved at run time rather than at compile time
var Methods = map[string]Method{
	"Add": {
		Name:    "Add",
		ArgType: reflect.TypeFor[float64](),
		Count:   -1,
	},
	"Divide": {
		Name:    "Divide",
		ArgType: reflect.TypeFor[float64](),
		Count:   2,
	},
	"Sqrt": {
		Name:    "Sqrt",
		ArgType: reflect.TypeFor[float64](),
		Count:   1,
	},
	"IsPrime": {
		Name:    "IsPrime",
		ArgType: reflect.TypeFor[int](),
		Count:   1,
	},
	"Average": {
		Name:    "Average",
		ArgType: reflect.TypeFor[float64](),
		Count:   -1,
	},
	"Reverse": {
		Name:    "Reverse",
		ArgType: reflect.TypeFor[string](),
		Count:   1,
	},
	"IsPalindrome": {
		Name:    "IsPalindrome",
		ArgType: reflect.TypeFor[string](),
		Count:   1,
	},
	"Sort": {
		Name:    "Sort",
		ArgType: reflect.TypeFor[float64](),
		Count:   -1,
	},
	"Hash": {
		Name:    "Hash",
		ArgType: reflect.TypeFor[string](),
		Count:   1,
	},
	"Ping": {
		Name:    "Ping",
		ArgType: reflect.TypeFor[string](),
		Count:   1,
	},
}

func ValidateRequest(req Request) (Method, error) {
	//if method name is not provided return appropriate error
	if strings.TrimSpace(req.Method) == "" {
		return Method{}, errors.New("Method name is required")
	}
	//check if args array entries are valid json
	for i, arg := range req.Args {
		if !json.Valid(arg) {
			return Method{}, fmt.Errorf("Argument %d contains invalid json", i)
		}
	}
	value, exists := Methods[req.Method]
	//check if provided method name exists
	if !exists {
		return Method{}, errors.New("Method not found")
	}
	//insure a positive number of arguments is provided
	//if this method args count is not variable and doesnt follow the count rules throw error
	//if the count is -1 thats okay because it can be any number of arguments
	args_length := len(req.Args)
	if args_length == 0 {
		return Method{}, errors.New("No arguments provided")
	}
	if value.Count != -1 && args_length != value.Count {
		return Method{}, errors.New("invalid number of arguments")
	}

	//loop through args which is an array of array of bytes
	//create with reflect.New a zero valued pointer of the Method type
	//serialize the rawArg into the arg using Unmarshal and arg.Interface
	//write the outputs to and array and return itH
	args := make([]reflect.Value, 0, args_length)
	for i, rawArg := range req.Args {
		arg := reflect.New(value.ArgType)
		err := json.Unmarshal(rawArg, arg.Interface())
		if err != nil {
			return Method{}, fmt.Errorf("Invalid argument %d: %w", i, err)
		}
		args = append(args, arg.Elem())
	}
	value.Args = args
	return value, nil
}
