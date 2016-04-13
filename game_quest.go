package main

import (
	"io/ioutil"
	"encoding/json"
//	"fmt"
)

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

//クエスト
type JsonGameQuestIn struct {
	Id        int `json:"questId"`
	SpendAP   int `json:"spendAP"`
	DropRank  int `json:"dropRank"`

	Difficult int `json:"difficult"`
	MapId     int `json:"mapId"`
}

// Jsonからパースする
func CreateGameQuests(filePath string) map[string]JsonGameQuestIn {

	var jsonGameQuests map[string]JsonGameQuestIn

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		Dlogln("Read Error: ", err)
	}

	json_err := json.Unmarshal(file, &jsonGameQuests)
	if json_err != nil {
		Dlogln("Format Error: ", json_err)
	}

//	fmt.Printf("%+v\n", jsonGameQuests)

	return jsonGameQuests
}

//difficultを落としこむ
//出現難度配分
//出現タイミング
//組み合わせ比率
type QuestEnvironment struct {
	DifficurtQuest        int

	PointPerOne           int //1体あたりの目安評価値
	SecondsPerQuest       int //クエスト秒数
	SecondsPerSlice       int //スライスあたり秒数
	BasePointPerSlice     int //スライスあたりの目安評価値のベース体数
	IncreasePointPerSlice int //スライスあたりの増加評価値のベース体数
}

//スライスの数
func (questEnvironment QuestEnvironment)timeSliceCount() int {
	return int(questEnvironment.SecondsPerQuest / questEnvironment.SecondsPerSlice);
}

//ステージの期待値を返す
func (questEnvironment QuestEnvironment)criteriaStateEvaluation() int {
	criteriaStateEvaluation := questEnvironment.timeSliceCount() *
		questEnvironment.BasePointPerSlice *
		questEnvironment.PointPerOne

	for i := 0; i < questEnvironment.timeSliceCount(); i++ {
		criteriaStateEvaluation +=
			questEnvironment.IncreasePointPerSlice *
			questEnvironment.PointPerOne *
			i
	}

	return criteriaStateEvaluation
}

//任意のスライスの期待値を返す
//@param sliceIndex 何番目のスライスか(0が最初)
func (questEnvironment QuestEnvironment)criteriaEvaluationPerSliceAtIndex(sliceIndex int) int {
	return questEnvironment.BasePointPerSlice * questEnvironment.PointPerOne +
	questEnvironment.IncreasePointPerSlice * questEnvironment.PointPerOne * sliceIndex
}

func CreateQuestEnvironment(jsonGameQuestIn JsonGameQuestIn) QuestEnvironment {
	return QuestEnvironment{
		DifficurtQuest 		   : jsonGameQuestIn.Difficult,
		PointPerOne            : 70,
		SecondsPerQuest        : 60,
		SecondsPerSlice        : 5,
		BasePointPerSlice      : 5,
		IncreasePointPerSlice  : 1,
	}
}


//敵の総量を決定する
func CreateEnemyNum(questEnvironment QuestEnvironment) int {
	return questEnvironment.DifficurtQuest * 10;
}
