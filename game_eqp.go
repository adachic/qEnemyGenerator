package main

import (
	"io/ioutil"
	"encoding/json"
	"fmt"
)

type EQPType string
const (
	EQPTypeWeapon EQPType = "weapon"
)

type EQPSubType string
const (
	EQPSubTypeMelee EQPSubType = "melee"
)

type JsonGameEqp struct {
	Id      int `json:"eqpId"`
	Type    EQPType `json:"type"`
	SubType EQPSubType `json:"subType"`
}

//Fitを得る
func (eqp JsonGameEqp) getFit() int{
	//TODO: 適切にfitを算出
	return 10
}

// Jsonからパースする
func CreateGameEqps(filePath string) map[string]JsonGameEqp {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read Error: ", err)
	}

	var jsonGameEqps map[string]JsonGameEqp

	json_err := json.Unmarshal(file, &jsonGameEqps)
	if json_err != nil {
		fmt.Println("Format Error: ", json_err)
	}

	fmt.Printf("%+v\n", jsonGameEqps)

	return jsonGameEqps
}

//敵種別に対応するEQPを返す
func PickUpEQPWithSampleEnemy(enemySample Enemy) JsonGameEqp {
	//TODO,EQPの選定
	switch enemySample.characterId {
	case CharacterIdSword :
	case CharacterIdArcher:
	case CharacterIdMage:
	case CharacterIdHealer:
	case CharacterIdThief:
	case CharacterIdWarlock:
	case CharacterIdNinja:
	case CharacterIdSlimeB:
	case CharacterIdSlimeR:
	case CharacterIdSlimeY:
	case CharacterIdSlimeG:
	case CharacterIdSlimeD:
	case CharacterIdSkeleton:
	case CharacterIdPenguin:
	case CharacterIdGoblin:
	case CharacterIdGoblinC:
	case CharacterIdLizard:
	case CharacterIdLizardP:
	case CharacterIdLizardC:
	case CharacterIdFrog:
	case CharacterIdFrogP:
	case CharacterIdBat:
	case CharacterIdBatI:
	case CharacterIdBatP:
	case CharacterIdGhost:
	case CharacterIdGhostI:
	case CharacterIdSpore:
	case CharacterIdSporeP:
	case CharacterIdSporeC:
	case CharacterIdNecR:
	case CharacterIdNecP:
	case CharacterIdNecD:
	case CharacterIdWitch:
	case CharacterIdWitchV:
	case CharacterIdWitchS:
	case CharacterIdGigantI:
	case CharacterIdGigantP:
	case CharacterIdGigantC:
	case CharacterIdSkeletonW:
	}
	return JsonGameEqp{Id:1}
}
