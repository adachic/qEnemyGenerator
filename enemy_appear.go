package main

type AIType string
const (
	EQPTypeAttacker AIType = "attacker"
	EQPTypeDefender AIType = "defender"
	EQPTypeHealer AIType = "healer"
)

type CharacterId int
const (
	CharacterIdSword CharacterId = iota
	CharacterIdArcher
	CharacterIdMage
	CharacterIdHealer
	CharacterIdThief
	CharacterIdWarlock
	CharacterIdNinja
	CharacterIdSlimeB
	CharacterIdSlimeR
	CharacterIdSlimeY
	CharacterIdSlimeG
	CharacterIdSlimeD
	CharacterIdSkeleton
	CharacterIdPenguin
	CharacterIdGoblin
	CharacterIdGoblinC
	CharacterIdLizard
	CharacterIdLizardP
	CharacterIdLizardC
	CharacterIdFrog
	CharacterIdFrogP
	CharacterIdBat
	CharacterIdBatI
	CharacterIdBatP
	CharacterIdGhost
	CharacterIdGhostI
	CharacterIdSpore
	CharacterIdSporeP
	CharacterIdSporeC
	CharacterIdNecR
	CharacterIdNecP
	CharacterIdNecD
	CharacterIdWitch
	CharacterIdWitchV
	CharacterIdWitchS
	CharacterIdGigantI
	CharacterIdGigantP
	CharacterIdGigantC
	CharacterIdSkeletonW
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
quest JsonGameQuestIn, eqps map[string]JsonGameEqp, questEnvironment QuestEnvironment) (enemyAppears []EnemyAppear, enemySamples []EnemySample, zones []JsonZone, questsOut []JsonGameQuestOut) {
	//敵の数を決定


	//組み合わせを決定


	//敵の強さを決定


	//ペースに従って配置


	return enemyAppears, enemySamples, zones, questsOut
}
