package main

import (
	"bytes"
	"io/ioutil"
	"os"
	"testing"
)

const fileName = "test.json"
const filePermission = 0644

func checkError(t *testing.T, err error) {
	if err != nil {
		t.Fatal(err)
	}
}

// Common validation tests
func TestOperationMissingError(t *testing.T) {
	var buffer bytes.Buffer

	expectedError := "-operation flag has to be specified"
	args := Arguments{
		id:        -1,
		operation: "",
		item:      "",
		fileName:  fileName,
	}
	err := Perform(args, &buffer)

	if err == nil {
		t.Error("Expect error when -operation flag is missing")
	}

	if err.Error() != expectedError {
		t.Errorf("Expect error to be '%s', but got '%s'", expectedError, err.Error())
	}
}

func TestWrongOperationError(t *testing.T) {
	var buffer bytes.Buffer
	args := Arguments{
		id:        -1,
		operation: "abcd",
		item:      "",
		fileName:  fileName,
	}
	expectedError := "Operation abcd not allowed!"

	err := Perform(args, &buffer)

	if err == nil {
		t.Error("Expect error when wrong -operation passed")
	}

	if err.Error() != expectedError {
		t.Errorf("Expect error to be '%s', but got '%s'", expectedError, err.Error())
	}
}

func TestFileNameMissingError(t *testing.T) {
	var buffer bytes.Buffer
	args := Arguments{
		id:        -1,
		operation: "list",
		item:      "",
		fileName:  "",
	}
	expectedError := "-fileName flag has to be specified"

	err := Perform(args, &buffer)

	if err == nil {
		t.Error("Expect error when -fileName flag is missing")
	}

	if err.Error() != expectedError {
		t.Errorf("Expect error to be '%s', but got '%s'", expectedError, err.Error())
	}
}

// List operation tests
func TestListOperation(t *testing.T) {
	args := Arguments{
		id:        -1,
		operation: "list",
		item:      "",
		fileName:  fileName,
	}
	var buffer bytes.Buffer

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, filePermission)
	defer os.Remove(fileName)
	checkError(t, err)

	existingItems := "[{\"id\": 1, \"email\": \"test@test.com\", \"age\": 34}, {\"id\": 2, \"email\": \"tes2@test.com\", \"age\": 32}]"

	file.Write([]byte(existingItems))
	file.Close()

	err = Perform(args, &buffer)
	checkError(t, err)

	file, err = os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, filePermission)
	checkError(t, err)

	bytes, err := ioutil.ReadAll(file)
	checkError(t, err)

	result := buffer.String()
	if result != existingItems {
		t.Errorf("Expect output to equal %s, but got %s", existingItems, result)
	}
	if string(bytes) != existingItems {
		t.Errorf("Expect file content to equal %s, but got %s", existingItems, string(bytes))
	}
}

// Adding operation tests
func TestAddingOperationMissingItem(t *testing.T) {
	var buffer bytes.Buffer
	args := Arguments{
		id:        -1,
		operation: "add",
		item:      "",
		fileName:  fileName,
	}
	expectedError := "-item flag has to be specified"

	err := Perform(args, &buffer)

	if err == nil {
		t.Error("Expect error when -item flag is missing")
	}

	if err.Error() != expectedError {
		t.Errorf("Expect error to be '%s', but got '%s'", expectedError, err.Error())
	}
}

func TestAddingOperationSameID(t *testing.T) {
	var buffer bytes.Buffer

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, filePermission)
	defer os.Remove(fileName)

	checkError(t, err)

	existingItem := "[{\"id\": 1, \"email\": \"test@test.com\", \"age\": 34}]"

	file.Write([]byte(existingItem))
	file.Close()

	item := "{\"id\": 1, \"email\": \"test@test.com\", \"age\": 34}"
	args := Arguments{
		id:        -1,
		operation: "add",
		item:      item,
		fileName:  fileName,
	}
	expectedOutput := "Item with id 1 already exists"

	err = Perform(args, &buffer)
	checkError(t, err)

	resultOutput := buffer.String()

	if resultOutput != expectedOutput {
		t.Errorf("Expect error to be '%s', but got '%s'", expectedOutput, resultOutput)
	}
}

func TestAddingOperation(t *testing.T) {
	var buffer bytes.Buffer

	expectedFileContent := "[{\"id\":1,\"email\":\"test@test.com\",\"age\":34}]"
	itemToAdd := "{\"id\": 1, \"email\": \"test@test.com\", \"age\": 34}"
	args := Arguments{
		id:        -1,
		operation: "add",
		item:      itemToAdd,
		fileName:  fileName,
	}
	defer os.Remove(fileName)
	err := Perform(args, &buffer)
	checkError(t, err)

	file, err := os.OpenFile(fileName, os.O_RDONLY, filePermission)
	defer file.Close()

	checkError(t, err)

	bytes, err := ioutil.ReadAll(file)
	checkError(t, err)

	if string(bytes) != expectedFileContent {
		t.Errorf("Expect file content to be %s, but got %s", expectedFileContent, bytes)
	}
}

// FindByID operation tests
func TestFindByIdOperationMissingID(t *testing.T) {
	var buffer bytes.Buffer
	args := Arguments{
		id:        -1,
		operation: "findById",
		item:      "",
		fileName:  fileName,
	}
	expectedError := "-id flag has to be specified"

	err := Perform(args, &buffer)

	if err == nil {
		t.Error("Expect error when -id flag is missing")
	}

	if err.Error() != expectedError {
		t.Errorf("Expect error to be '%s', but got '%s'", expectedError, err.Error())
	}
}

func TestFindByIdOperation(t *testing.T) {
	var buffer bytes.Buffer

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, filePermission)
	defer os.Remove(fileName)

	checkError(t, err)

	existingItems := "[{\"id\":1,\"email\":\"test@test.com\",\"age\":34},{\"id\":2,\"email\":\"test2@test.com\",\"age\":31}]"

	file.Write([]byte(existingItems))
	file.Close()

	expectedOutput := "{\"id\":2,\"email\":\"test2@test.com\",\"age\":31}"
	args := Arguments{
		id:        2,
		operation: "findById",
		item:      "",
		fileName:  fileName,
	}
	err = Perform(args, &buffer)

	checkError(t, err)

	resultString := buffer.String()

	if resultString != expectedOutput {
		t.Errorf("Expect output to be '%s', but got '%s'", expectedOutput, resultString)
	}
}

func TestFindByIdOperationWrongID(t *testing.T) {
	var buffer bytes.Buffer

	file, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, filePermission)
	defer os.Remove(fileName)

	checkError(t, err)

	existingItems := "[{\"id\":1,\"email\":\"test@test.com\",\"age\":34},{\"id\":2,\"email\":\"test2@test.com\",\"age\":31}]"

	file.Write([]byte(existingItems))
	file.Close()

	expectedOutput := ""
	args := Arguments{
		id:        3,
		operation: "findById",
		item:      "",
		fileName:  fileName,
	}
	err = Perform(args, &buffer)

	checkError(t, err)

	resultString := buffer.String()

	if resultString != expectedOutput {
		t.Errorf("Expect output to be '%s', but got '%s'", expectedOutput, resultString)
	}
}

// Removing operations tests

func TestRemovingOperationMissingID(t *testing.T) {
	var buffer bytes.Buffer
	args := Arguments{
		id:        -1,
		operation: "remove",
		item:      "",
		fileName:  fileName,
	}

	expectedError := "-id flag has to be specified"

	err := Perform(args, &buffer)

	if err == nil {
		t.Error("Error has to be shown when -id flag is missing")
	}

	if err.Error() != expectedError {
		t.Errorf("Expect error to be '%s', but got '%s'", expectedError, err.Error())
	}
}

func TestRemovingOperationWrongID(t *testing.T) {
	var buffer bytes.Buffer

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, filePermission)
	defer os.Remove(fileName)

	checkError(t, err)

	existingItems := "[{\"id\":1,\"email\":\"test@test.com\",\"age\":34}]"

	file.Write([]byte(existingItems))
	file.Close()

	expectedOutput := "Item with id 2 not found"
	args := Arguments{
		id:        2,
		operation: "remove",
		item:      "",
		fileName:  fileName,
	}
	err = Perform(args, &buffer)

	checkError(t, err)

	resultOutput := buffer.String()

	if resultOutput != expectedOutput {
		t.Errorf("Expect output to be '%s', but got '%s'", expectedOutput, resultOutput)
	}
}

func TestRemovingOperation(t *testing.T) {
	var buffer bytes.Buffer

	file, err := os.OpenFile(fileName, os.O_WRONLY|os.O_CREATE, filePermission)
	defer os.Remove(fileName)

	checkError(t, err)

	existingItems := "[{\"id\":1,\"email\":\"test@test.com\",\"age\":34},{\"id\":2,\"email\":\"test2@test.com\",\"age\":31}]"

	file.Write([]byte(existingItems))
	file.Close()
	expectedFileContent := "[{\"id\":2,\"email\":\"test2@test.com\",\"age\":31}]"
	args := Arguments{
		id:        1,
		operation: "remove",
		item:      "",
		fileName:  fileName,
	}

	err = Perform(args, &buffer)

	checkError(t, err)

	file, err = os.OpenFile(fileName, os.O_RDONLY, filePermission)
	defer file.Close()

	checkError(t, err)

	bytes, err := ioutil.ReadAll(file)

	if string(bytes) != expectedFileContent {
		t.Errorf("Expect file content to be '%s', but got '%s'", expectedFileContent, bytes)
	}
}
