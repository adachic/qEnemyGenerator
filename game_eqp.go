package main

import (
	"io/ioutil"
	"encoding/json"
//	"fmt"
	"github.com/adachic/lottery"
)

type EQPType string
const (
	EQPTypeWeapon EQPType = "weapon"
	EQPTypeRepair EQPType = "repair"
	EQPTypeHeal EQPType = "heal"

	EQPTypeAccesory EQPType = "accessory"
)

type EQPSubType string
const (
	EQPSubTypeMelee EQPSubType = "melee"
	EQPSubTypeFire EQPSubType = "fire"
	EQPSubTypeArea EQPSubType = "area"
	EQPSubTypeArrow EQPSubType = "arrow"
	EQPSubTypeInstant EQPSubType = "instant"

	EQPSubTypeShield EQPSubType = "shield"
	EQPSubTypeBody EQPSubType = "body"
	EQPSubTypeHead EQPSubType = "head"
)

type JsonGameEqp struct {
	Id      int `json:"eqpId"`
	Type    EQPType `json:"type"`
	SubType EQPSubType `json:"subType"`
	Name    string `json:"name"`
}

//Fitを得る
func (eqp JsonGameEqp) getFit() int {
	fit := 0
	switch eqp.Type {
	case EQPTypeWeapon:
		switch eqp.SubType {
		case EQPSubTypeMelee:
			fit = 20;
		case EQPSubTypeFire:
			fit = 35;
		case EQPSubTypeArea:
			fit = 81;
		case EQPSubTypeArrow:
			fit = 30;
		case EQPSubTypeInstant:
		}
	case EQPTypeRepair:
	case EQPTypeHeal:
		switch eqp.SubType {
		case EQPSubTypeFire:
			fit = 30;
		case EQPSubTypeArea:
			fit = 82;
		}
	case EQPTypeAccesory:
		switch eqp.SubType {
		case EQPSubTypeShield:
			fit = 15;
		case EQPSubTypeBody:
			fit = 10;
		case EQPSubTypeHead:
			fit = 10;
		}
	}
	return fit
}

// Jsonからパースする
func CreateGameEqps(filePath string) map[string]JsonGameEqp {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		Dlogln("Read Error: ", err)
	}

	var jsonGameEqps map[string]JsonGameEqp

	json_err := json.Unmarshal(file, &jsonGameEqps)
	if json_err != nil {
		Dlogln("Format Error: ", json_err)
	}

//	Dlog("%+v\n", jsonGameEqps)
	//Scan2()

	G_jsonGameEqps = jsonGameEqps
	InitGlobals()

	return jsonGameEqps
}

func InitGlobals() {
	for _, eqp := range G_jsonGameEqps {
		switch eqp.Type {
		case EQPTypeWeapon:{
			switch eqp.SubType {
			case EQPSubTypeMelee :
				G_jsonGameMelees = append(G_jsonGameMelees, eqp)
			case EQPSubTypeFire:
				G_jsonGameFires = append(G_jsonGameFires, eqp)
			case EQPSubTypeArea:
				G_jsonGameFireAreas = append(G_jsonGameFireAreas, eqp)
			case EQPSubTypeArrow:
				G_jsonGameArrows = append(G_jsonGameArrows, eqp)
			}
		}
		case EQPTypeHeal:{
			switch eqp.SubType {
			case EQPSubTypeInstant:
				G_jsonGameHeals = append(G_jsonGameHeals, eqp)
			case EQPSubTypeFire:
				G_jsonGameHeals = append(G_jsonGameHeals, eqp)
			case EQPSubTypeArea:
				G_jsonGameHealAreas = append(G_jsonGameHealAreas, eqp)
			}
		}
		case EQPTypeRepair:{
		}
		case EQPTypeAccesory:{
			switch eqp.SubType {
			case EQPSubTypeShield:
				G_jsonGameShields = append(G_jsonGameShields, eqp)
			case EQPSubTypeBody:
				G_jsonGameBodys = append(G_jsonGameBodys, eqp)
			}
		}
		default:{
		}
		}
	}
}

var G_jsonGameEqps map[string]JsonGameEqp
var G_jsonGameMelees []JsonGameEqp
var G_jsonGameArrows []JsonGameEqp
var G_jsonGameFires []JsonGameEqp
var G_jsonGameFireAreas []JsonGameEqp
var G_jsonGameHeals []JsonGameEqp
var G_jsonGameHealAreas []JsonGameEqp

var G_jsonGameShields []JsonGameEqp
var G_jsonGameBodys []JsonGameEqp

//ランダムなMeleeを返す
func PickUpRandomMelee(enemySample Enemy) JsonGameEqp {
	//TODO:モンスターなら体当たりを返す
	return G_jsonGameMelees[lottery.GetRandomInt(0, len(G_jsonGameMelees) - 1)]
}

func PickUpRandomHeal() JsonGameEqp {
	return G_jsonGameHeals[lottery.GetRandomInt(0, len(G_jsonGameHeals) - 1)]
}

func PickUpRandomRanged(enemySample Enemy) JsonGameEqp {
	if enemySample.characterId == CharacterIdArcher {
		return G_jsonGameArrows[lottery.GetRandomInt(0, len(G_jsonGameArrows) - 1)]
	}
	return G_jsonGameFires[lottery.GetRandomInt(0, len(G_jsonGameFires) - 1)]
}

func PickUpRandomRangedAOE(enemySample Enemy) JsonGameEqp {
	return G_jsonGameFireAreas[lottery.GetRandomInt(0, len(G_jsonGameFireAreas) - 1)]
}

func PickUpRandomShield(EnemySample Enemy) JsonGameEqp {
	return G_jsonGameShields[lottery.GetRandomInt(0, len(G_jsonGameShields) - 1)]
}

func PickUpRandomBody(EnemySample Enemy) JsonGameEqp {
	return G_jsonGameBodys[lottery.GetRandomInt(0, len(G_jsonGameBodys) - 1)]
}

//敵種別に対応するEQPを返す
func PickUpEQPWithSampleMainEnemy(enemySample Enemy) JsonGameEqp {
	switch enemySample.fixedRole{
	case RoleTank:
		return PickUpRandomMelee(enemySample)
	case RoleHealer:
		return PickUpRandomHeal()
	case RoleDpsMelee:
		return PickUpRandomMelee(enemySample)
	case RoleDpsRanged:
		return PickUpRandomRanged(enemySample)
	case RoleDpsAoe:
		return PickUpRandomRangedAOE(enemySample)
	case RoleBuff:
		return PickUpRandomRanged(enemySample)
	case RoleDeBuff:
		return PickUpRandomRanged(enemySample)
	}
	return PickUpRandomMelee(enemySample)
}
