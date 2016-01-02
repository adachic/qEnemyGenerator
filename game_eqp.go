package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type EQPType int
const (
	EQPTypeWeapon EQPType = "weapon"
)

type EQPSubType int
const (
	EQPSubTypeMelee EQPSubType = "melee"
)

type JsonGameEqp struct {
	Id      int `json:"eqpId"`
	Type    EQPType `json:"type"`
	SubType EQPSubType `json:"subType"`
}

// Jsonからパースする
func CreateGameEqps(filePath string) []JsonGameEqp {
	// Loading jsonfile
	file, err := ioutil.ReadFile(filePath)
	// 指定したDataset構造体が中身になるSliceで宣言する

	var jsonGameEqps []JsonGameEqp

	json_err := json.Unmarshal(file, &jsonGameEqps)
	if err != nil {
		fmt.Println("Format Error: ", json_err)
	}

	fmt.Printf("%+v\n", jsonGameEqps)

	return jsonGameEqps
}
