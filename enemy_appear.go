package main

//import "fmt"

type AIType string

const (
	EQPTypeAttacker AIType = "attacker"
	EQPTypeDefender AIType = "defender"
	EQPTypeHealer AIType = "healer"
)

type EnemyAppear struct {
	Id           int
	Sample       *EnemySample
	Zone         JsonZone
	AIType       AIType
	AppearTime   int
	QuestId      int

	Quest        JsonGameQuestIn
	IntervalTime int
	Num          int
}

type EnemySample struct {
	Id             int
	CharacterId    CharacterId
	CharacterIdStr string
	UnitLevel      int
	MainEqp        JsonGameEqp
	MainEqpLevel   int
	SubEqp1        JsonGameEqp
	SubEqp2        JsonGameEqp
	SubEqp3        JsonGameEqp
}

type JsonZone struct {
	Id    int `json:"zoneId"`
	MapId int `json:"mapid"`
	Pos1  int `json:"pos1"`
	Pos2  int `json:"pos2"`
	Pos3  int `json:"pos3"`
	Pos4  int `json:"pos4"`
	Pos5  int `json:"pos5"`
	Pos6  int `json:"pos6"`
}

func (zone JsonZone)getMapPosition() GameMapPosition {
	pos := zone.Pos1
	return GameMapPosition{
		X:pos % 100,
		Y:(pos / 100) % 100,
		Z:(pos / 10000) % 100,
	}
}

//position1,2の単純な引き算した距離を返す
func (position1 GameMapPosition)distanceAbs(position2 GameMapPosition) int {
	distanceAbs := 0
	{
		xdiff := position2.X - position1.X
		if (xdiff < 0) {
			distanceAbs += -xdiff
		}else {
			distanceAbs += xdiff
		}
	}
	{
		ydiff := position2.Y - position1.Y
		if (ydiff < 0) {
			distanceAbs += -ydiff
		}else {
			distanceAbs += ydiff
		}
	}
	{
		zdiff := position2.Z - position1.Z
		if (zdiff < 0) {
			distanceAbs += -zdiff
		}else {
			distanceAbs += zdiff
		}
	}
	return distanceAbs
}

func (zone JsonZone) getFit(role Role, geneEnvironment GeneEnvironment) int {
	fit := 0
	switch role {
	case RoleTank:
		fallthrough
	case RoleDpsMelee:
		fit = zone.getFitForMelee(geneEnvironment)
	case RoleHealer:
		fallthrough
	case RoleDpsRanged:
		fallthrough
	case RoleDpsAoe:
		fallthrough
	case RoleBuff:
		fallthrough
	case RoleDeBuff:
		fit = zone.getFitForRanged(geneEnvironment)
	default:
	}
	return fit
}

func (zone JsonZone) getFitForMelee(geneEnvironment GeneEnvironment) int {
	fit := 0
	distanceAbs := geneEnvironment.JsonGameMap.AllyStartPoint.distanceAbs(zone.getMapPosition())
	switch {
	case distanceAbs > 20:
		fit = 1
	case distanceAbs > 15:
		fit = 5
	case distanceAbs > 10:
		fit = 10
	case distanceAbs > 5:
		fit = 15
	default:
		fit = 20
	}
	return fit
}

func (zone JsonZone) getFitForRanged(geneEnvironment GeneEnvironment) int {
	provisionalFitXYZ := 0
	{
		distanceAbs := geneEnvironment.JsonGameMap.AllyStartPoint.distanceAbs(zone.getMapPosition())
		switch {
		case distanceAbs > 20:
			provisionalFitXYZ = 1
		case distanceAbs > 15:
			provisionalFitXYZ = 2
		case distanceAbs > 10:
			provisionalFitXYZ = 3
		case distanceAbs > 5:
			provisionalFitXYZ = 5
		default:
			provisionalFitXYZ = 10
		}
	}
	provisionalFitZ := 0
	{
		zdiff := zone.getMapPosition().Z - geneEnvironment.JsonGameMap.AllyStartPoint.Z
		switch  {
		case zdiff > 10:
			provisionalFitZ = 20
		case zdiff > 4:
			provisionalFitZ = 10
		case zdiff > 1:
			provisionalFitZ = 5
		default:
			provisionalFitZ = 0
		}
	}
	provisionalFit := provisionalFitXYZ + provisionalFitZ
	fit := 0
	{
		switch  {
		case provisionalFit > 25:
			fit = 65
		case provisionalFit > 20:
			fit = 40
		case provisionalFit > 15:
			fit = 25
		case provisionalFit > 10:
			fit = 12
		case provisionalFit > 5:
			fit = 5
		default:
			fit = 1
		}
	}
	return fit
}

type GameZone struct {
	GameMapPositions []GameMapPosition
}

//敵出現情報,クエスト情報を返す
func CreateEnemyAppears(gamePartsDict map[string]GameParts, gameMap JsonGameMap,
quest JsonGameQuestIn, eqps map[string]JsonGameEqp, questEnvironment QuestEnvironment) (enemyAppears []*EnemyAppear,
zones []JsonZone) {

	gameMap.allocJungle3(gamePartsDict)

	//地点・ゾーンの確保
	zones = CreateZones(questEnvironment, gameMap, gamePartsDict)

	Dlogln("===zones====")
	Dlog("zones count:%d\n", len(zones))

	geneEnvironment := CreateGeneEnvironment(zones, questEnvironment, gameMap);

	//求めたい評価値
	creteriaEvaluation := questEnvironment.criteriaStateEvaluation()
	Dlog("creteriaEvaluation %+v\n", creteriaEvaluation)

	//スライス単位でナップザックする
	for i := 0; i < questEnvironment.timeSliceCount(); i++ {
		//このスライスの理想評価値
		creteriaEvaluationPerSlice := questEnvironment.criteriaEvaluationPerSliceAtIndex(i)
		Dlog("[%d]creteriaEvaluationPerSlice:%+v\n", i, creteriaEvaluationPerSlice)
		enemyAppearsPerSlice := EnemiesWithZone(creteriaEvaluationPerSlice, zones, questEnvironment, geneEnvironment, i)
		enemyAppears = append(enemyAppears, enemyAppearsPerSlice...)
	}

	//AppearのIdとか足りてないやつをセットする
	id := 1
	for _, enemyAppear := range enemyAppears {
		enemyAppear.Id = id
		enemyAppear.QuestId = quest.Id
		enemyAppear.Sample.UnitLevel = quest.Difficult
		enemyAppear.Sample.MainEqpLevel = quest.Difficult
		id++
	}
	return enemyAppears, zones
}


//出現ゾーン生成
func CreateZones(questEnvironment QuestEnvironment, gameMap JsonGameMap, gamePartsDict map[string]GameParts) (jsonZones []JsonZone) {
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

	Dlog("aho1:gamemapENEM:%+v\n", gameMap.EnemyStartPoints)
	Dlog("aho1:gamemapALLY:%+v\n", gameMap.AllyStartPoint)

	//敵地点をゾーンに変換
	var gameZones []GameZone
	for _, value := range gameMap.EnemyStartPoints {
		DDlog("zone2]%+v\n", value)
		positions := CreateNearlyGamePositions(value, gameMap, xy)
		gameZone := NewGameZone(positions)
		gameZones = append(gameZones, *gameZone)
		DDlog("zone0]%+v\n", gameZones)
		DDlog("zone1]%+v\n", positions)
	}
	DDlogln("aho2")

	//JSON形式に変換
	jsonZones = ConvertToJsonZone(gameZones, gameMap.MapId)
	DDlogln("aho3")
	return jsonZones
}

//GameZoneをJsonZoneに変換
func ConvertToJsonZone(gameZones []GameZone, mapId int) (jsonZones []JsonZone) {
	id := 1
	for _, positions := range gameZones {
		var jsonZone JsonZone
		try := 1
		for _, position := range positions.GameMapPositions {
			jsonZone.Id = id
			jsonZone.MapId = mapId
			id++
			pos := position.Z * 10000 + position.Y * 100 + position.X
			switch try {
			case 1:
				jsonZone.Pos1 = pos
			case 2:
				jsonZone.Pos2 = pos
			case 3:
				jsonZone.Pos3 = pos
			case 4:
				jsonZone.Pos4 = pos
			case 5:
				jsonZone.Pos5 = pos
			case 6:
				jsonZone.Pos6 = pos
			}
			try++
			jsonZones = append(jsonZones, jsonZone)
		}
	}
	return jsonZones
}

//positionの周辺のマスから歩行可能、高低差１以内の場所を配列として返す(ゾーンの内容)
func CreateNearlyGamePositions(position GameMapPosition, gameMap JsonGameMap, xy [][]bool) (gameMapPositions []GameMapPosition) {
	createMaxNum := 5
	createdNum := 0

	xOffs := [...]int{-1, 0, 1, 0, 0, -1, 1, -1, 1}
	yOffs := [...]int{0, -1, 0, 1, 0, -1, 1, 1, -1}

	for i := 0; i < 9; i++ {
		x := position.X + xOffs[i]
		y := position.Y + yOffs[i]
		z := position.Z
		DDlog("x:%d, y:%d, z%d\n", x,y,z)

		if (x >= gameMap.MaxX || x < 0 || y >= gameMap.MaxY || y < 0 || z == 0) {
			DDlogln("unko-1:", z)
			continue
		}
		//他のゾーンで取られている
		existXY := xy[y][x]
		if (existXY) {
			DDlogln("unko0:", z)
			continue
		}
		//harfなら1マス下が肉抜きされている
		cube1 := gameMap.JungleGym3[z][y][x]
		if (cube1 != nil){
			//歩行不能
			if (!cube1.Walkable) {
				DDlogln("unko22:", z)
				continue
			}
			cube2 := gameMap.JungleGym3[z+1][y][x]
			//ブロックで埋まっている
			if (cube2 != nil) {
				DDlogln("unko33:", z)
				continue
			}
		}else{
			cube := gameMap.JungleGym3[z - 1][y][x]
			//足元がない
			if (cube == nil) {
				DDlogln("unko1:", z)
				continue
			}
			//歩行不能
			if (!cube.Walkable) {
				DDlogln("unko2:", z)
				continue
			}
			cube2 := gameMap.JungleGym3[z][y][x]
			//ブロックで埋まっている
			if (cube2 != nil) {
				DDlogln("unko3:", z)
				continue
			}
		}
		DDlogln("add:")
		xy[y][x] = true
		gameMapPositions = append(gameMapPositions, GameMapPosition{X:x, Y:y, Z:z})
		createdNum++
		if createdNum >= createMaxNum {
			break
		}
	}
Dlog("gameMapPositions:%+v",gameMapPositions)
	return gameMapPositions
}

//ゾーン生成
func NewGameZone(gameMapPositions []GameMapPosition) *GameZone {
	game_zone := &GameZone{
		GameMapPositions:gameMapPositions,
	}
	return game_zone
}

