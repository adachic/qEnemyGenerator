package main

type AIType int
const (
	EQPTypeAttacker AIType = "attacker"
)

type CharacterId int
const (
	CharacterIdSlime CharacterId = iota
)

type EnemyAppear struct {
	Id           int
	Quest        JsonGameQuestIn
	Sample       EnemySample
	Zone         JsonZone
	AIType       AIType

	AppearTime   int
	IntervalTime int
	Num          int
}

type EnemySample struct {
	Id           int
	CharacterId  CharacterId
	UnitLevel    int
	mainEqp      JsonGameEqp
	mainEqpLevel int
	subEqp1      JsonGameEqp
	subEqp2      JsonGameEqp
	subEqp3      JsonGameEqp
}

type JsonZone struct {
	Id    int `json:"zoneId"`
	MapId int `json:"mapid"`
	pos1  int `json:"pos1"`
	pos2  int `json:"pos2"`
	pos3  int `json:"pos3"`
	pos4  int `json:"pos4"`
	pos5  int `json:"pos5"`
	pos6  int `json:"pos6"`
}

//敵出現情報,クエスト情報を返す
func CreateEnemyAppears(gameMap JsonGameMap,
quests JsonGameQuestIn, eqps JsonGameEqp, questId int, questEnvironment QuestEnvironment) (enemyAppears []EnemyAppear, enemySamples []EnemySample, zones []JsonZone, quests []JsonGameQuestOut) {

	//敵の強さを決定


	//敵の数を決定






	return enemyAppears, enemySamples, zones, quests
}
