package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Client struct {
	file        string
	name        string
	id          int
	phoneNumber string
	home        string
}

func (c *Client) StringFileFormat() string {
	return fmt.Sprintf("%s,%s,%d,%s,%s\n", c.file, c.name, c.id, c.phoneNumber, c.home)
}

func (c *Client) CheckDuplicate() error {
	for _, client := range Clients {
		if client == *c {
			return NewInvalidClientError("client already exists")
		}
	}
	return nil
}

func (c *Client) CheckValidFields() (int, error) {
	var errorCounter int = 0

	if c.file == "" {
		errorCounter += 1
	}
	if c.name == "" {
		errorCounter += 1
	}
	if c.id == 0 {
		errorCounter += 1
	}
	if c.phoneNumber == "" {
		errorCounter += 1
	}
	if c.home == "" {
		errorCounter += 1
	}

	if errorCounter != 0 {
		return errorCounter, NewInvalidClientError("not all fields are filled")
	}
	return errorCounter, nil
}

type InvalidClientError struct {
	msg string
}

func (e InvalidClientError) Error() string {
	return fmt.Sprintf("InvalidClientError: %s", e.msg)
}

func NewInvalidClientError(msg string) InvalidClientError {
	return InvalidClientError{msg: msg}
}

var Clients []Client = []Client{}

func main() {
	defer func() {
		err := recover()
		if err != nil {
			fmt.Println("Several errors were detected at runtime")
		}
		fmt.Println("End of execution")
	}()

	file, err := os.Open("./client-data/customers.txt")
	defer file.Close()

	defer func() {
		fmt.Println("ejecucion finalizada")
	}()
	if err != nil {
		fmt.Println(err)
		panic("the indicated file was not found or is damaged")
	}

	fileScanner := bufio.NewScanner(file)
	fileScanner.Split(bufio.ScanLines)

	var id int

	for fileScanner.Scan() {
		text := strings.Split(fileScanner.Text(), ",")
		id, _ = strconv.Atoi(text[2])
		Clients = append(Clients, Client{text[0], text[1], id, text[3], text[4]})
	}

	// newClient := Client{
	// 	file: "file2", name: "dog", id: 2, phoneNumber: "12312", home: "the moon",
	// }

	newClient := Client{
		file: "file1", name: "baby driver", id: 1, phoneNumber: "123", home: "the sun",
	}

	defer func() {
		err := recover()
		if err != nil {
			fmt.Println(err)
		}
	}()

	if err = newClient.CheckDuplicate(); err != nil {
		panic("Error: client already exists")
	}

	if _, err = newClient.CheckValidFields(); err != nil {
		panic("Error: invalid client")
	}

	fmt.Println("all right!")
}
