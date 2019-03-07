package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

type Person struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Arguments map[string]string

func Perform(args Arguments, writer io.Writer) error {
	var err error
	operation := args["operation"]
	if operation == "" {
		return errors.New("-operation flag has to be specified")
	}
	fileName := args["fileName"]
	if fileName == "" {
		return errors.New("-fileName flag has to be specified")
	}
	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	data, err := ioutil.ReadAll(f)
	err = f.Close()
	if err != nil {
		return err
	}
	persons := make([]Person, 0)
	if len(data) != 0 {
		err = json.Unmarshal(data, &persons)
		if err != nil {
			return err
		}
	}
	item := args["item"]
	id := args["id"]

	switch operation {
	case "add":
		{
			if item == "" {
				return errors.New("-item flag has to be specified")
			}
			var person Person
			err = json.Unmarshal([]byte(item), &person)
			if err != nil {
				return err
			}
			isSame := false
			for _, value := range persons {
				if value.Id == person.Id {
					errString := `Item with id ` + person.Id + ` already exists`
					writer.Write([]byte(errString))
					isSame = true
				}
			}
			if !isSame {
				persons = append(persons, person)
				if err = StoreToFile(fileName, persons); err != nil {
					return err
				}
			}
		}
	case "list":
		{
			_, err = writer.Write(data)
			if err != nil {
				return err
			}
		}
	case "findById":
		{
			if id == "" {
				return errors.New("-id flag has to be specified")
			}
			for _, value := range persons {
				if value.Id == id {
					bytesW, err := json.Marshal(value)
					if err != nil {
						return err
					}
					_, err = writer.Write(bytesW)
					if err != nil {
						return err
					}
				}
			}
		}
	case "remove":
		{
			if id == "" {
				return errors.New("-id flag has to be specified")
			}
			newList := make([]Person, 0)
			passed := false
			for _, value := range persons {
				if value.Id != id {
					newList = append(newList, value)
				}
				if value.Id == id {
					passed = true
				}
			}
			if !passed {
				errString := "Item with id " + item + " not found"
				_, err = writer.Write([]byte(errString))
				if err != nil {
					return err
				}
			}
			if passed {
				if err = StoreToFile(fileName, newList); err != nil {
					return err
				}
			}
		}
	default:
		return fmt.Errorf("Operation %s not allowed!", operation)
	}
	return nil
}

func StoreToFile(filename string, persons []Person) error {
	data, err := json.Marshal(persons)
	if err != nil {
		return err
	}
	os.Remove(filename)
	file, err := os.OpenFile(filename, os.O_CREATE, 0644)
	if err != nil {
		return err
	}
	defer file.Close()
	_, err = file.WriteString(string(data))
	if err != nil {
		return err
	}
	return nil
}

func parseArgs() Arguments {
	var id = flag.String("id", "", "id")
	var item = flag.String("item", "", "item")
	var operation = flag.String("operation", "", "operation")
	var fileName = flag.String("fileName", "", "fileName")
	args := Arguments{
		"id":        *id,
		"item":      *item,
		"operation": *operation,
		"fileName":  *fileName,
	}
	return args
}

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}
