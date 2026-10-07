package main

import (
	"fmt"
	"math"
	"reflect"
	"rpc/scheme"
)

func handleReq(method scheme.Method) scheme.Response {
	var response scheme.Response
	failed := scheme.Response{
		Status: 500,
		Value:  nil,
		Error:  "",
	}
	switch method.Name {
	case "Add":
		result, err := add(method.Args)
		if err != nil {
			failed.Error = err.Error()
			return failed
		}
		return scheme.Response{
			Status: 200,
			Value:  result,
			Error:  "",
		}
	case "Divide":
		result, err := divide(method.Args)
		if err != nil {
			failed.Error = err.Error()
			return failed
		}
		return scheme.Response{
			Status: 200,
			Value:  result,
			Error:  "",
		}
	case "Sqrt":
		result, err := sqrt(method.Args)
		if err != nil {
			failed.Error = err.Error()
			return failed
		}
		return scheme.Response{
			Status: 200,
			Value:  result,
			Error:  "",
		}
	case "IsPrime":
		result, err := isPrime(method.Args)
		if err != nil {
			failed.Error = err.Error()
			return failed
		}
		return scheme.Response{
			Status: 200,
			Value:  result,
			Error:  "",
		}
	case "Average":
		result, err := avg(method.Args)
		if err != nil {
			failed.Error = err.Error()
			return failed
		}
		return scheme.Response{
			Status: 200,
			Value:  result,
			Error:  "",
		}
	case "Reverse":
		result, err := reverse(method.Args)
		if err != nil {
			failed.Error = err.Error()
			return failed
		}
		return scheme.Response{
			Status: 200,
			Value:  result,
			Error:  "",
		}
	case "IsPalindrome":
		result, err := isPalindrome(method.Args)
		if err != nil {
			failed.Error = err.Error()
			return failed
		}
		return scheme.Response{
			Status: 200,
			Value:  result,
			Error:  "",
		}
	case "Sort":
		result, err := sort(method.Args)
		if err != nil {
			failed.Error = err.Error()
			return failed
		}
		return scheme.Response{
			Status: 200,
			Value:  result,
			Error:  "",
		}
	case "Hash":
		result, err := hash(method.Args)
		if err != nil {
			failed.Error = err.Error()
			return failed
		}
		return scheme.Response{
			Status: 200,
			Value:  result,
			Error:  "",
		}
	case "Ping":
		result, err := ping(method.Args)
		if err != nil {
			failed.Error = err.Error()
			return failed
		}
		return scheme.Response{
			Status: 200,
			Value:  result,
			Error:  "",
		}

	}

	return response
}

func add(args []reflect.Value) (float64, error) {
	var result float64 = 0

	for _, arg := range args {
		result += float64(arg.Float())
	}
	return result, nil
}

func divide(args []reflect.Value) (float64, error) {
	var result float64 = 0

	result = float64(args[0].Float() / args[1].Float())
	return result, nil
}

func sqrt(args []reflect.Value) (float64, error) {
	var result float64 = 0
	result = math.Sqrt(float64(args[0].Float()))
	fmt.Println("result : ", result)
	return result, nil
}

func isPrime(args []reflect.Value) (bool, error) {
	num := int(args[0].Int())
	if num < 2 {
		return false, nil
	}
	if num == 2 {
		return true, nil
	}
	if num%2 == 0 {
		return false, nil
	}
	for i := 3; i*i < num; i += 2 {
		if num%i == 0 {
			return false, nil
		}
	}
	return true, nil
}

func avg(args []reflect.Value) (float64, error) {
	var result float64 = 0
	for _, arg := range args {
		result += float64(arg.Float())
	}
	return result / float64(len(args)), nil
}

func reverse(args []reflect.Value) (string, error) {
	str := string(args[0].String())
	b := []byte(str)
	j := len(b) - 1
	for i := range b {
		if i == j || i+1 == j {
			break
		}
		temp := b[i]
		b[i] = b[j]
		b[j] = temp
		j--
	}
	return string(b), nil
}

func isPalindrome(args []reflect.Value) (bool, error) {
	str := string(args[0].String())
	b := []byte(str)
	j := len(b) - 1
	for i := range b {
		if i == j || i+1 == j {
			break
		}
		if b[i] != b[j] {
			return false, nil
		}
		j--
	}
	return true, nil
}

func sort(args []reflect.Value) ([]float64, error) {
	arr := make([]float64, len(args))
	for i := range args {
		arr[i] = float64(args[i].Float())
	}
	scheme.QuickSort(arr, 0, len(arr)-1)
	return arr, nil
}

func hash(args []reflect.Value) (string, error) {
	str := string(args[0].String())
	hashed_string := scheme.Hash(str)
	return hashed_string, nil
}

func ping(args []reflect.Value) (string, error) {
	ip := string(args[0].String())
	body, err := scheme.Ping(ip)
	if err != nil {
		return "", err
	}
	return body, nil
}
