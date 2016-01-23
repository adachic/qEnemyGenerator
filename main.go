package main

import (
	"flag"
	"fmt"
	"strconv"
	"encoding/json"
	"os"
)


func main() {
	fmt.Printf("Hello, world3.\n")

	//マップ読み込み
	//クエスト一覧読み込み
	//EQP一覧読み込み
	var mapFilePath string
	var questFilePath string
	var eqpFilePath string
	var enemyFilePath string
	var questId int
	flag.StringVar(&mapFilePath, "map", "map.json", "APP_PARTS_FILE_PATH")
	flag.StringVar(&questFilePath, "quest", "quest.json", "APP_PARTS_FILE_PATH")
	flag.StringVar(&eqpFilePath, "eqp", "eqp.json", "APP_PARTS_FILE_PATH")
	flag.StringVar(&enemyFilePath, "character", "character.json", "APP_PARTS_FILE_PATH")
	flag.IntVar(&questId, "questId", 0, "APP_PARTS_FILE_PATH")
	flag.Parse()

	//敵出現位置取得
	//questId取得
	fmt.Println("==map=")
	gameMap := CreateGameMap(mapFilePath)
	gamePartsDict := CreateGamePartsDict("./assets/IntegratedPartsAll2.json") //harfId対応済み
	fmt.Println("==quest=")
	quests := CreateGameQuests(questFilePath)
	fmt.Println("==eqp=")
	eqps := CreateGameEqps(eqpFilePath)
	fmt.Println("==characters=")
	CreateEnemySamplesJ(enemyFilePath)

	currentQuest := quests[strconv.Itoa(questId)]
	//出現難度配分
	//出現タイミング
	//組み合わせ比率
	questEnvironment := CreateQuestEnvironment(currentQuest)

	//種類・強さ・数・タイミング・位置
	//enemyAppears, enemySamples, zones, quests := CreateEnemyAppears(gameMap, quests, eqps, questId, questEnvironment)
	enemyAppears, zones:= CreateEnemyAppears(gamePartsDict, gameMap, currentQuest, eqps, questEnvironment)

	//敵サンプル出力 enemy_sample.json/csv
	CreateJsonAndCsv(enemyAppears, zones)

	fmt.Printf("Hello, world4.\n")
}

//Json/CSV出力
func CreateJsonAndCsv(enemyAppears []*EnemyAppear, zones []JsonZone) {
	fmt.Printf("==output json==\n")

	/*
	jsonStub := JsonGameMap{
		MaxX:game_map.Size.MaxX,
		MaxY:game_map.Size.MaxY,
		MaxZ:game_map.Size.MaxZ,
		AspectX:32,
		AspectY:16,
		AspectT:16,
		AllyStartPoint:game_map.AllyStartPoint,
		EnemyStartPoints:game_map.EnemyStartPoints,
		Category:game_map.Category,
	}
	*/
	//	fmt.Printf("%+v\n", jsonStub)

	bytes, json_err := json.Marshal(enemyAppears)
	if json_err != nil {
		fmt.Println("Json Encode Error: ", json_err)
	}

	//	fmt.Printf("bytes:%+v\n", string(bytes))

	file, err := os.Create("./output/" + "" + "1.json")
	_, err = file.Write(bytes)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
}
