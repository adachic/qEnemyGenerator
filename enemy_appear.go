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

type GameZone struct {
	GameMapPositions []GameMapPosition
}

//敵出現情報,クエスト情報を返す
func CreateEnemyAppears(gamePartsDict map[string]GameParts, gameMap JsonGameMap,
quest JsonGameQuestIn, eqps map[string]JsonGameEqp, questEnvironment QuestEnvironment) (enemyAppears []EnemyAppear,
enemySamples []EnemySample, zones []JsonZone, questsOut []JsonGameQuestOut) {

	gameMap.allocJungle3(gamePartsDict)

	//地点・ゾーンの確保
	zones = CreateZones(questEnvironment, gameMap, gamePartsDict)

	//敵の数を決定


	//組み合わせを決定


	//敵の強さを決定


	//ペースに従って配置


	return enemyAppears, enemySamples, zones, questsOut
}

func CreateZones(questEnvironment QuestEnvironment, gameMap JsonGameMap, gamePartsDict map[string]GameParts) {
	var zones []GameZone
	//出現地点をゾーンに変換
	//	gameMap.AllyStartPoint

	//ゾーンかぶらないための判別マップ
	var xy [][]bool
	xy = make([][]bool, gameMap.MaxY)
	for y := 0; y < gameMap.MaxY; y++ {
		xy[y] = make([]bool, gameMap.MaxX)
	}
	for y := 0; y < gameMap.MaxY; y++ {
		for x := 0; x < gameMap.MaxX; x++ {
			xy[y][x] = false
		}
	}

	//敵地点をゾーンに変換
	for _, value := range gameMap.EnemyStartPoints {
		positions := CreateNearlyGamePositions(value, gameMap, xy)
		zone := NewGameZone(positions)
		zones = append(zones, zone)
	}

	//難度に応じてゾーン増やす
	CreateZonesForEdge(zones, gameMap, questEnvironment)

	return zones
}

//難度に応じてゾーン増やす
func CreateZonesForEdge(zones []GameZone, gameMap JsonGameMap, questEnvironment QuestEnvironment){
	//必要なゾーン数
	alreadyNum := len(zones)





}

//positionの周辺のマスから歩行可能、高低差１以内の場所を配列として返す(ゾーンの内容)
func CreateNearlyGamePositions(position GameMapPosition, gameMap JsonGameMap, xy [][]bool) (gameMapPositions []GameMapPosition) {
	createNum := 5
	xOffs := [...]int{-1, 0, 1, 0, -1, 1, -1, 1}
	yOffs := [...]int{0, -1, 0, 1, -1, 1, 1, -1}

	for i := 0; i <= createNum; i++ {
		x := position.X + xOffs[i]
		y := position.Y + yOffs[i]
		z := position.Z
		cube, exist := gameMap.JungleGym3[z - 1][y][x]
		//他のゾーンで取られている
		_, existXY := xy[y][x]
		if (existXY) {
			continue
		}
		//足元がない
		if (!exist) {
			continue
		}
		//歩行不能
		if (!cube.Walkable) {
			continue
		}
		_, exist = gameMap.JungleGym3[z][y][x]
		//ブロックで埋まっている
		if (exist) {
			continue
		}
		xy[y][x] = true
		gameMapPositions = append(gameMapPositions, GameMapPosition{X:x, Y:y, Z:z})
	}

	return gameMapPositions
}


//ゾーン生成
func NewGameZone(gameMapPositions []GameMapPosition) *GameZone {
	game_zone := &GameZone{
		GameMapPositions:gameMapPositions,
	}
	return game_zone
}

