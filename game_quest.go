package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

//クエスト
type JsonGameQuestIn struct {
	Id        int
	Type      string
	Rule      string
	SpendAP   int
	DropRank  int

	Difficult int
	MapId     int
}

type JsonGameQuestOut struct {
	Id          int
	Difficult   int
	MapId       int

	//ここからがOut
	Type        string
	Rule        string
	SpendAP     int
	DropRank    int
	RewordLimes int
	RewordExp   int

	UnitExp     int
	EQPExp      int

	BtFlagX     int
	BtFlagY     int
	BtFlagZ     int

	TimeLimit   int
	PurpTime    int
	Cond1       string
	Cond2       string

	Title       string
	Description string
}

/*
	AllyStartPoint   GameMapPosition
	EnemyStartPoints []GameMapPosition
	Category         Category
*/

// Jsonからパースする
func CreateGameQuests(filePath string) []JsonGameQuestIn {
	// Loading jsonfile
	file, err := ioutil.ReadFile(filePath)
	// 指定したDataset構造体が中身になるSliceで宣言する

	var jsonGameQuests []JsonGameQuestIn

	json_err := json.Unmarshal(file, &jsonGameQuests)
	if err != nil {
		fmt.Println("Format Error: ", json_err)
	}

	fmt.Printf("%+v\n", jsonGameQuests)

	return jsonGameQuests
}

//difficultを落としこむ
//出現難度配分
//出現タイミング
//組み合わせ比率
type QuestEnvironment struct {
	DifficurtQuest int
	MonsPerHum float32
	IncreaseAppearPerSec float32
	InpulseVolume float32
	InpulsePerQuest float32
}

func CreateQuestEnvironment(jsonGameQuestIn JsonGameQuestIn) QuestEnvironment {
	return QuestEnvironment{
		DifficurtQuest: jsonGameQuestIn.Difficult,
		MonsPerHum: 5.0,
		IncreaseAppearPerSec: 0.2,
		InpulseVolume: 5.0,
		InpulsePerQuest: 3,
	}
}
