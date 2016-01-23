package main
import (
	"github.com/adachic/lottery"
	"io/ioutil"
	"fmt"
	"encoding/json"
	"gopkg.in/go-pp/pp.v2"
)


type Role int
const (
	RoleTank  Role = iota
	RoleHealer
	RoleDpsMelee
	RoleDpsRanged
	RoleDpsAoe
	RoleBuff
	RoleDeBuff
)

type CharacterId int
const (
	CharacterIdShield CharacterId = iota
	CharacterIdSword
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
	CharacterIdDemon
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

//敵情報
type Enemy struct {
	characterId CharacterId
	roles       []Role

	fit         int
	enemyJson   EnemyJson
}

//クエスト
type EnemyJson struct {
	Str         int
	Int         int
	Vit         int
	Agi         int
	Dex         int
	Sense       int
	Luk         int
	Sum         int

	Plain       int
	Cave        int
	Remains     int
	Poisonswamp int
	Fire        int
	Snow        int
	Jozen       int
	Castle      int
	Role1       string
	Role2       string
	Role3       string
}

func GetRole(roleString string) Role {
	switch roleString{
	case "tank":
		return RoleTank
	case "healer":
		return RoleHealer
	case "dps-melee":
		return RoleDpsMelee
	case "dps-ranged":
		return RoleDpsRanged
	case "dps-aoe":
		return RoleDpsAoe
	case "buff":
		return RoleBuff
	case "debuff":
		return RoleBuff
	}
	return RoleTank
}

func (enemyJson EnemyJson)getFit() int {
	return enemyJson.Sum
}

func (enemyJson EnemyJson)getRoles() []Role {
	var roles  []Role
	if (enemyJson.Role1 != "") {
		role := GetRole(enemyJson.Role1)
		roles = append(roles, role)
	}
	if (enemyJson.Role2 != "") {
		role := GetRole(enemyJson.Role2)
		roles = append(roles, role)
	}
	if (enemyJson.Role3 != "") {
		role := GetRole(enemyJson.Role3)
		roles = append(roles, role)
	}
	return roles
}

// Jsonからパースする
func CreateEnemySamplesJ(filePath string) map[string]EnemyJson {
	var jsonGameEnemies map[string]EnemyJson

	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Println("Read Error: ", err)
	}

	json_err := json.Unmarshal(file, &jsonGameEnemies)
	if json_err != nil {
		fmt.Println("Format Error: ", json_err)
	}

	pp.Printf("%+v\n", jsonGameEnemies)

	Scan()

	return jsonGameEnemies
}

func CreateEnemy(key string, enemyJson EnemyJson) Enemy {
	var characterId CharacterId
	switch key {
	case "shield":
		characterId = CharacterIdShield
	case "sword":
		characterId = CharacterIdSword
	case "archer":
		characterId = CharacterIdArcher
	case "mage":
		characterId = CharacterIdMage
	case "healer":
		characterId = CharacterIdHealer
	case "thief":
		characterId = CharacterIdThief
	case "warlock":
		characterId = CharacterIdWarlock
	case "ninja":
		characterId = CharacterIdNinja

	case "slimeB":
		characterId = CharacterIdSlimeB
	case "slimeR":
		characterId = CharacterIdSlimeR
	case "slimeY":
		characterId = CharacterIdSlimeY
	case "slimeG":
		characterId = CharacterIdSlimeG
	case "slimeD":
		characterId = CharacterIdSlimeD
	case "skeleton":
		characterId = CharacterIdSkeleton
	case "penguin":
		characterId = CharacterIdPenguin
	case "demon":
		characterId = CharacterIdDemon
	case "goblin":
		characterId = CharacterIdGoblin
	case "goblinC":
		characterId = CharacterIdGoblinC
	case "lizard":
		characterId = CharacterIdLizard
	case "lizardP":
		characterId = CharacterIdLizardP
	case "lizardC":
		characterId = CharacterIdLizardC
	case "frog":
		characterId = CharacterIdFrog
	case "frogP":
		characterId = CharacterIdFrogP
	case "bat":
		characterId = CharacterIdBat
	case "batI":
		characterId = CharacterIdBatI
	case "batP":
		characterId = CharacterIdBatP
	case "ghost":
		characterId = CharacterIdGhost
	case "ghostI":
		characterId = CharacterIdGhostI
	case "spore":
		characterId = CharacterIdSpore
	case "sporeP":
		characterId = CharacterIdSporeP
	case "sporeC":
		characterId = CharacterIdSporeC
	case "necR":
		characterId = CharacterIdNecR
	case "necP":
		characterId = CharacterIdNecP
	case "necD":
		characterId = CharacterIdNecD
	case "witch":
		characterId = CharacterIdWitch
	case "witchV":
		characterId = CharacterIdWitchV
	case "witchS":
		characterId = CharacterIdWitchS
	case "gigantI":
		characterId = CharacterIdGigantI
	case "gigantP":
		characterId = CharacterIdGigantP
	case "gigantC":
		characterId = CharacterIdGigantC
	case "skeletonW":
		characterId = CharacterIdSkeletonW
	}

	return Enemy{
		characterId:characterId,
		roles:enemyJson.getRoles(),
		fit:enemyJson.getFit(),
		enemyJson:enemyJson,
	}
}

//敵情報一覧の生成
func CreateEnemySamples() []Enemy {
	enemies := []Enemy{}
	enemyMap := CreateEnemySamplesJ("./debug/character.json")
	for key, enemy := range enemyMap {
		if len(enemy.getRoles()) == 0{
			continue
		}
		enemies = append(enemies, CreateEnemy(key, enemy))
	}
	/*
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleTank})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleDpsMelee})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleDpsRanged})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleDpsAoe})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleHealer})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleBuff})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleDeBuff})
	*/
	return enemies
}

//ロールに対応したサンプルを抽出
func PickUpRandomSampleWithRole(enemiesSample []Enemy, role Role) Enemy {
	filterdEnemy := []Enemy{}
	for i := 0; i < len(enemiesSample); i++ {
		enemy := enemiesSample[i]

		containedRole := false
		for _, role_ := range enemy.roles {
			if role_ == role{
				containedRole = true
			}
		}
		if containedRole {
			continue
		}
		filterdEnemy = append(filterdEnemy, enemy)
	}
	filterdEnemyNum := len(filterdEnemy)
	idx := lottery.GetRandomInt(0, filterdEnemyNum)
	return filterdEnemy[idx]
}

//Fitを返す
func (enemy Enemy)getFit() int {
	//TODO,パラメータのちゃんとした値の算出
	switch enemy.characterId {
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
	return 30
}

//所持しているRoleをランダムに選択
func (enemy Enemy) pickUpRandomRole() Role{
	rolesCount := len(enemy.roles)
	if rolesCount == 1 {
		return enemy.roles[0]
	}
	idx := lottery.GetRandomInt(0, rolesCount)
	return enemy.roles[idx]
}

