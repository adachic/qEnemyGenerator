package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type EQPType string
const (
	EQPTypeWeapon EQPType = "weapon"
)

type EQPSubType string
const (
	EQPSubTypeMelee EQPSubType = "melee"
)

type JsonGameEqp struct {
	Id      int `json:"eqpId"`
	Type    EQPType `json:"type"`
	SubType EQPSubType `json:"subType"`
}

// Jsonからパースする
func CreateGameEqps(filePath string) map[string]JsonGameEqp {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read Error: ", err)
	}

	var jsonGameEqps map[string]JsonGameEqp

	json_err := json.Unmarshal(file, &jsonGameEqps)
	if json_err != nil {
		fmt.Println("Format Error: ", json_err)
	}

	fmt.Printf("%+v\n", jsonGameEqps)

	return jsonGameEqps
}
