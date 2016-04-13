package main

import (
	"io/ioutil"
	"encoding/json"
//	"fmt"
)

type JsonPanel struct {
	X  int `json:"x"`
	Y  int `json:"y"`
	Z  int `json:"z"`
	Id string `json:"id"`
}

//座標
type GameMapPosition struct {
	X int `json:"x"`
	Y int `json:"y"`
	Z int `json:"z"`
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
	GameParts []GameParts `json:"gameParts"`

	AllyStartPoint   GameMapPosition `json:"allyStartPoint"`
	EnemyStartPoints []GameMapPosition `json:"enemyStartPoints"`
	Category         Category `json:"category"`

	JungleGym3 [][][] *GameParts
	MapId int
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
		Dlogln("Read Error: ", err)
	}

	var jsonGameMap JsonGameMap

	json_err := json.Unmarshal(file, &jsonGameMap)
	if json_err != nil {
		Dlogln("Format Error1: ", json_err)
	}

//	fmt.Printf("%+v\n", jsonGameMap)

	return jsonGameMap
}

//１次元配列を３次元配列に転換
func (game_map *JsonGameMap) allocJungle3(gamePartsDict map[string]GameParts){
	game_map.JungleGym3 = make([][][]*GameParts, game_map.MaxZ + 1)
	for z := 0; z < game_map.MaxZ + 1; z++ {
		game_map.JungleGym3[z] = make([][]*GameParts, game_map.MaxY)
		for y := 0; y < game_map.MaxY; y++ {
			game_map.JungleGym3[z][y] = make([]*GameParts, game_map.MaxX)
		}
	}
	/*
	for z := 0; z < game_map.MaxZ; z++ {
		for y := 0; y < game_map.MaxY; y++ {
			for x := 0; x < game_map.MaxX; x++ {
				game_map.JungleGym[z][y][x] = false;
			}
		}
	}
	*/
	for _, value := range game_map.JungleGym{
		parts := gamePartsDict[value.Id]
		game_map.JungleGym3[value.Z][value.Y][value.X] = &GameParts{
			Id:parts.Id,
			Walkable:parts.Walkable,
			MacroTypes:parts.MacroTypes,
			WaterType:parts.WaterType,
			Category:parts.Category,
			StructureType:parts.StructureType,
			PavementType:parts.PavementType,
		}
	}
}
