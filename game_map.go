package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type JsonPanel struct {
	X  int `json:"x"`
	Y  int `json:"y"`
	Z  int `json:"z"`
	Id string `json:"id"`
}

//マップ
type JsonGameMap struct {
	MaxX      int `json:"maxX"`
	MaxY      int `json:"maxY"`
	MaxZ      int `json:"maxZ"`
	AspectX   int `json:"aspectX"`
	AspectY   int `json:"aspectY"`
	AspectT   int `json:"aspectT"`
	JungleGym []JsonPanel `json:"jungleGym"`
}

/*
	AllyStartPoint   GameMapPosition
	EnemyStartPoints []GameMapPosition
	Category         Category
*/

// Jsonからパースする
func CreateGameMap(filePath string) JsonGameMap{
	// Loading jsonfile
	file, err := ioutil.ReadFile(filePath)
	// 指定したDataset構造体が中身になるSliceで宣言する
	if err != nil {
		fmt.Println("Read Error: ", err)
	}

	var jsonGameMap JsonGameMap

	json_err := json.Unmarshal(file, &jsonGameMap)
	if json_err != nil {
		fmt.Println("Format Error1: ", json_err)
	}

	fmt.Printf("%+v\n", jsonGameMap)

	return jsonGameMap
}
