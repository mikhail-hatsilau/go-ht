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

func main() {
	err := Perform(parseArgs(), os.Stdout)
	if err != nil {
		panic(err)
	}
}

type User struct {
	Id    string `json:"id"`
	Email string `json:"email"`
	Age   int    `json:"age"`
}

type Arguments map[string]string

func parseArgs() Arguments {
	operation := flag.String("operation", "", "operation")
	filename := flag.String("fileName", "", "filename")
	item := flag.String("item", "", "item")
	id := flag.String("id", "", "id")
	flag.Parse()
	return Arguments{
		"operation": *operation,
		"fileName":  *filename,
		"item":      *item,
		"id":        *id,
	}
}

func Perform(args Arguments, writer io.Writer) (err error) {
	operation := args["operation"]
	if operation == "" {
		return fmt.Errorf("-operation flag has to be specified")
	}
	filename := args["fileName"]
	if filename == "" {
		return fmt.Errorf("-fileName flag has to be specified")
	}
	switch operation {
	case "add":
		item := args["item"]
		if item == "" {
			return fmt.Errorf("-item flag has to be specified")
		}
		var newUser User
		err = json.Unmarshal([]byte(item), &newUser)
		if err != nil {
			return
		}
		var users []User
		var bytes []byte
		bytes, err = ioutil.ReadFile(filename)
		if err != nil {
			return
		}
		err = json.Unmarshal(bytes, &users)
		if err != nil {
			return
		}
		for _, user := range users {
			if user.Id == newUser.Id {
				errStr := "Item with id " + newUser.Id + " already exists"
				_, err = writer.Write([]byte(errStr))
				return
			}
		}
		users = append(users, newUser)
		bytes, err = json.Marshal(users)
		if err != nil {
			return
		}
		return ioutil.WriteFile(filename, bytes, 512)
	case "findById":
		id := args["id"]
		if len(id) == 0 {
			return errors.New("-id flag has to be specified")
		}
		var users []User
		bytes, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		err = json.Unmarshal(bytes, &users)
		var user User
		for _, value := range users {
			if value.Id == id {
				user = value
				break
			}
		}
		userBytes := []byte("")
		if len(user.Id) != 0 {
			bytes, err := json.Marshal(user)
			if err != nil {
				return err
			}
			userBytes = bytes
		}
		_, err = writer.Write(userBytes)
		return err
	case "remove":
		id := args["id"]
		if len(id) == 0 {
			return errors.New("-id flag has to be specified")
		}
		bytesUsers, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		var users []User
		err = json.Unmarshal(bytesUsers, &users)
		var resultingSlice []User
		isRemoved := false
		for i, user := range users {
			if user.Id == id {
				resultingSlice = append(users[:i], users[i+1:]...)
				isRemoved = true
				break
			}
		}
		if !isRemoved {
			_, err = writer.Write([]byte("Item with id " + id + " not found"))
			return err
		}
		bytes, err := json.Marshal(resultingSlice)
		if err != nil {
			return err
		}
		return ioutil.WriteFile(filename, bytes, 512)
	case "list":
		users, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		_, err = writer.Write(users)
		return err
	default:
		return fmt.Errorf("Operation %s not allowed!", operation)
	}
}
