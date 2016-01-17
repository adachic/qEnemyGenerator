package main
import "github.com/adachic/lottery"

//敵情報
type Enemy struct {
	characterId CharacterId
	role        Role
}

type Role int
const (
	RoleTank Role = iota + 1
	RoleHealer
	RoleDpsMelee
	RoleDpsRanged
	RoleDpsAoe
	RoleBuff
	RoleDeBuff
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

//敵情報一覧の生成
func CreateEnemySamples() []Enemy {
	enemies := []Enemy{}
	//TODO json対応
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleTank})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleDpsMelee})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleDpsRanged})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleDpsAoe})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleHealer})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleBuff})
	enemies = append(enemies, Enemy{characterId:CharacterIdSword, role:RoleDeBuff})
	return enemies
}

//ロールに対応したサンプルを抽出
func PickUpRandomSampleWithRole(enemiesSample []Enemy, role Role) Enemy{
	filterdEnemy := []Enemy{}
	for i := 0; i < len(enemiesSample); i++ {
		enemy := enemiesSample[i]
		if enemy.role != role {
			continue
		}
		//TODO ここきてない
		filterdEnemy = append(filterdEnemy, enemy)
	}

	filterdEnemyNum := len(filterdEnemy)
	idx := lottery.GetRandomInt(0, filterdEnemyNum)
	return filterdEnemy[idx]
}

//Fitを返す
func (enemy Enemy)getFit() int{
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
