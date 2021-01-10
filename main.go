package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

//constants are used as commands and filenames and return values
const (
	fileName = "text.json"
	SET      = "set"
	GET      = "get"
	INC      = "inc"
	DEC      = "dec"
	DEL      = "del"
	NotExist = "Variable not exist"
)

//Data is struct type of json
type Data struct {
	Name  string `json:"Name"`
	Value string `json:"value"`
}

//CheckError checks the whether is there any error
func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

//WriteData writes the key value pairs into the json file
func WriteData(m1 map[string]string) {
	var readData []Data
	for i, v := range m1 {
		data1 := Data{Name: i, Value: v}
		readData = append(readData, data1)
	}
	bytedata, err := json.MarshalIndent(readData, "", " ")
	CheckError(err)
	ioutil.WriteFile(fileName, bytedata, os.ModePerm)

}

//IsDataExists is a fucntion that helps to identify whether the data exists or not
func IsDataExists(key string) (map[string]string, bool) {
	if _, err := os.Stat(fileName); os.IsNotExist(err) {
		f, err := os.Create(fileName)
		CheckError(err)
		_, err = io.WriteString(f, "[]")
		defer f.Close()
	}

	databyte, err := ioutil.ReadFile(fileName)
	CheckError(err)

	var readData []Data
	err = json.Unmarshal(databyte, &readData)
	CheckError(err)
	m1 := make(map[string]string)
	for _, val := range readData {
		m1[val.Name] = val.Value
	}
	if _, ok := m1[key]; ok {
		return m1, true
	}
	return m1, false

}

//SetData sets the key value pairs into the json file
func SetData(key, value string) string {
	readData, err := IsDataExists(key)
	if err != true {
		//m1 := make(map[string]string)
		readData[key] = value
		WriteData(readData)
		return "variable succesfully added"

	}
	return "Already Exist"

}

//GetData gets the key value if key exists
func GetData(key string) string {
	m1, err := IsDataExists(key)
	if err != false {
		val := m1[key]
		return val
	}
	return NotExist

}

//DeleteData deletes the key value pairs if key exists
func DeleteData(key string) string {
	m1, err := IsDataExists(key)
	if err != false {
		if _, ok := m1[key]; ok {
			delete(m1, key)
		}
		WriteData(m1)
		return "Data Deleted"

	}
	return NotExist

}

//UpdateData updates the value of a key by increment or decrement based on command
func UpdateData(key, cmd string) string {
	m1, err := IsDataExists(key)
	if err != false {
		val, _ := strconv.Atoi(m1[key])
		if cmd == INC {
			val++
		} else {
			val--
		}
		m1[key] = strconv.Itoa(val)
		WriteData(m1)
		return "Updated"
	}
	return NotExist
}

func main() {
	cmd := flag.String("cmd", "GET", "Command Prompt")
	key := flag.String("k", " ", "Key Name")
	value := flag.String("v", " ", "Value")
	flag.Parse()
	switch *cmd {
	case SET:
		fmt.Println(SetData(*key, *value))

	case GET:
		fmt.Println(GetData(*key))

	case INC:
		fmt.Println(UpdateData(*key, *cmd))

	case DEC:
		fmt.Println(UpdateData(*key, *cmd))

	case DEL:
		fmt.Println(DeleteData(*key))
	default:
		fmt.Println("command not recognised")
		return
	}

}
