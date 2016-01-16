package main
import "fmt"

type AIType string
const (
	EQPTypeAttacker AIType = "attacker"
	EQPTypeDefender AIType = "defender"
	EQPTypeHealer AIType = "healer"
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

func (zone JsonZone) getFit() int{
	//TODO 正しくzoneの評価をする
	return 10
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

	fmt.Println("===zones====")
	fmt.Printf("%+v\n", zones)

	//求めたい評価値
	creteriaEvaluation := questEnvironment.criteriaStateEvaluation()
	fmt.Printf("creteriaEvaluation %+v\n", creteriaEvaluation)

	//スライス単位でナップザックする
	for i := 0; i < questEnvironment.timeSliceCount() ; i++ {
		//このスライスの理想評価値
		creteriaEvaluationPerSlice := questEnvironment.criteriaEvaluationPerSliceAtIndex(i)
		fmt.Printf("[%d]%+v\n",i, creteriaEvaluationPerSlice)

		_ = EnemiesWithZone(creteriaEvaluationPerSlice, zones, questEnvironment)
	}

	//ナップザック結果の結合




	return enemyAppears, enemySamples, zones, questsOut
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

//	fmt.Println("aho1")
	//敵地点をゾーンに変換
	var gameZones []GameZone
	for _, value := range gameMap.EnemyStartPoints {
		positions := CreateNearlyGamePositions(value, gameMap, xy)
		gameZone := NewGameZone(positions)
		gameZones = append(gameZones, *gameZone)
//		fmt.Printf("%+v\n", gameZones)
	}
//	fmt.Println("aho2")

	//JSON形式に変換
	jsonZones = ConvertToJsonZone(gameZones, gameMap.MapId)
//	fmt.Println("aho3")
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
				jsonZone.pos1 = pos
			case 2:
				jsonZone.pos2 = pos
			case 3:
				jsonZone.pos3 = pos
			case 4:
				jsonZone.pos4 = pos
			case 5:
				jsonZone.pos5 = pos
			case 6:
				jsonZone.pos6 = pos
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

	xOffs := [...]int{-1, 0, 1, 0, -1, 1, -1, 1}
	yOffs := [...]int{0, -1, 0, 1, -1, 1, 1, -1}

	for i := 0; i < 8; i++ {
		x := position.X + xOffs[i]
		y := position.Y + yOffs[i]
		z := position.Z
		if(x >= gameMap.MaxX || x < 0 || y >= gameMap.MaxY || y < 0 || z == 0){
			continue
		}
		//他のゾーンで取られている
		existXY := xy[y][x]
		if (existXY) {
			continue
		}
//		fmt.Print("unko0:",z)
		cube := gameMap.JungleGym3[z - 1][y][x]
		//足元がない
		if (cube == nil) {
			continue
		}
		//歩行不能
		if (!cube.Walkable) {
			continue
		}
//		fmt.Print("unko1:",z)
		cube2 := gameMap.JungleGym3[z][y][x]
		//ブロックで埋まっている
		if (cube2 != nil) {
			continue
		}
		xy[y][x] = true
		gameMapPositions = append(gameMapPositions, GameMapPosition{X:x, Y:y, Z:z})
		createdNum++
		if createdNum >= createMaxNum {
			break
		}
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
