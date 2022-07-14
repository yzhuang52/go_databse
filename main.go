package main

import (
	"fmt"
	"os"
	"encoding/json"
	"sync"
	"github.com/jcelliott/lumber"
)

const Version = "1.0.1"

type (
	Logger interface{
		Fatal(string, ...interface{})
		Error(string, ...interface{})
		Warn(string, ...interface{})
		Info(string, ...interface{})
		Debug(string, ...interface{})
		Trace(string, ...interface{})
	}
	Driver struct{
		mutex sync.Mutex
		mutexes map[string]*sync.Mutex
		dir string 
		log Logger 
	}
)

type Address struct{
	City string 
	State string 
	Country string 
	Pincode json.Number
}

type User struct{
	Name string 
	Age json.Number
	Contact string 
	Company string 
	Address Address 
}



func main(){
	
	dir := "./"
	db, err := New(dir, nil)
	if err != nil{
		fmt.Println("Error", err)
	}

	employees := []User{
		{"John1", "23", "18332344213", "Yan's tech", Address{"Qingdao", "Shandong", "China", "266000"}},
		{"John2", "23", "18332344213", "Yan's tech", Address{"Qingdao", "Shandong", "China", "266000"}},
		{"John3", "23", "18332344213", "Yan's tech", Address{"Qingdao", "Shandong", "China", "266000"}},
		{"John4", "23", "18332344213", "Yan's tech", Address{"Qingdao", "Shandong", "China", "266000"}},
		{"John5", "23", "18332344213", "Yan's tech", Address{"Qingdao", "Shandong", "China", "266000"}},
	}

	for _, value := range employees{
		db.write("users", value.Name, User{
			Name: value.Name,
			Age: value.Age,
			Contact: value.Contact,
			Company: value.Company,
			Address: value.Address,
		})
	}

	records, err := db.ReadAll("users")
	if err != nil{
		fmt.Println("Error", err)
	}
	fmt.Println(records)

	allUsers := []User{}

	for _, value := range records{
		employeeFound := User{}
		if err := json.Unmarshal([]byte(value), &employeeFound); err != nil{
			fmt.Println("Error", err)
		}
		allUsers = append(allUsers, employeeFound)
	}
	fmt.Println(allUsers)
	// if err := db.Delete("user", "john"); err != nil{
	// 	fmt.Println("Error", err)
	// }

	// if err := db.Delete("user", ""); err != nil{
	// 	fmt.Println("Error", err)
	// }
}