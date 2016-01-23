package main

import (
	"flag"
	"fmt"
	"strconv"
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
	//	enemyAppears, enemySamples, zones, quests := CreateEnemyAppears(gameMap, quests, eqps, questId, questEnvironment)
	enemyAppears, enemySamples, zones, questsOut := CreateEnemyAppears(gamePartsDict, gameMap, currentQuest, eqps, questEnvironment)

	//敵サンプル出力 enemy_sample.json/csv
	//クエストへのひも付け quest_enemy.json/csv
	//出現位置 enemy_appear_point.json/csv
	CreateJsonAndCsv(enemyAppears, enemySamples, zones, questsOut)

	fmt.Printf("Hello, world4.\n")
}

//CSV出力
func CreateJsonAndCsv(enemyAppears []EnemyAppear, enemySmaples []EnemySample, zones []JsonZone, quests []JsonGameQuestOut) {

}
