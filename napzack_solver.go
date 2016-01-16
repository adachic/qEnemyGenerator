package main
import (
	"github.com/adachic/lottery"
)

//遺伝的アルゴリズムのパラメータ
type GeneEnvironment struct {
	NumPerAge        int //世代あたりの個体数N
	Zones            []JsonZone
	QuestEnvironment QuestEnvironment
	EnemySamples     []Enemy
}

//個体とその遺伝子
type GeneUnit struct {
	GenericEnemyNum      int             //敵の数
	GenericEnemyLPT_Num  int             //ライトパーティの数
	GenericEnemyFPT_Num  int             //フルパーティーの数
	GenericEnemyIndv_Num int             //単体の数
	GenericUnitEnemies   []GeneUnitEnemy //スライスあたりの敵一覧

	Fit                  int             //適応度
}

//敵個体
type GeneUnitEnemy struct {
	enemy Enemy
	eqp   JsonGameEqp
	zone  JsonZone
}

//Fitを返す
func (geneUnitEnemy GeneUnitEnemy)getFit() int {
	fit := 0
	fit += geneUnitEnemy.enemy.getFit()
	fit += geneUnitEnemy.zone.getFit()
	fit += geneUnitEnemy.eqp.getFit()
	return fit
}

type PTType int
const (
	PTTypeTHD PTType = 1 + iota
	PTTypeTHDD
	PTTypeTHDDB
	PTTypeTTHHDDDD
	PTTypeTTHHDDDB
)

func EnemiesWithZone(creteriaEvaluationPerSlice int, zones []JsonZone, questEnvironment QuestEnvironment) []EnemyAppear {
	geneEnvironment := CreateGeneEnvironment(zones, questEnvironment);

	//個体をランダムN個生成する
	//それぞれの適応度計算する
//	geneUnitsPerAge := CreateRundomsWithGeneUnitsPerAge(geneEnvironment);
	{
		//世代操作開始
		{
			//次のいずれかを行う
			//A.個体を2つ選択し、交差

			//B.突然変異

			//C.次世代にそのまま追加

		}
		//次世代が一定になっていれば次世代を対象として世代操作開始に戻る

	}


	return nil
}

//ランダムなZoneを返す
func (geneEnvironment GeneEnvironment) choiceRandomZone() JsonZone {
	zoneNum := len(geneEnvironment.Zones)
	idx := lottery.GetRandomInt(0, zoneNum - 1)
	return geneEnvironment.Zones[idx]
}

//ランダムなRoleを返す
func GetRoleRandom() Role {
	return Role(lottery.GetRandomInt(int(RoleTank), int(RoleDeBuff)))
}

//ランダムなDPSRoleを返す
func GetRoleRandomDPS() Role {
	return Role(lottery.GetRandomInt(int(RoleDpsMelee), int(RoleDpsAoe)))
}

//ランダムなBuffRoleを返す
func GetRoleRandomBuff() Role {
	return Role(lottery.GetRandomInt(int(RoleBuff), int(RoleDeBuff)))
}

//指定Roleのランダムな敵を生成
func CreateRundomEnemyWithType(geneEnvironment GeneEnvironment, role Role, zone JsonZone) GeneUnitEnemy {
	geneUnitEnemy := GeneUnitEnemy{}
	//種別が決定
	geneUnitEnemy.enemy = PickUpRandomSampleWithRole(geneEnvironment.EnemySamples, role)

	//EQP
	geneUnitEnemy.eqp = PickUpEQPWithSampleEnemy(geneUnitEnemy.enemy)

	//Zone
	geneUnitEnemy.zone = zone

	return geneUnitEnemy
}

//ランダムなフルパーティーの生成
func CreateRundomEnemyFPT(geneEnvironment GeneEnvironment) GenericEnemyPT {
	genericEnemyFpt := GenericEnemyPT{}
	genericEnemyFpt.LptType = PTType(lottery.GetRandomInt(int(PTTypeTTHHDDDD), int(PTTypeTTHHDDDB)))
	genericEnemyFpt.Zone = geneEnvironment.choiceRandomZone()
	enemies := []GeneUnitEnemy{}
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleTank, genericEnemyFpt.Zone))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleTank, genericEnemyFpt.Zone))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleHealer, genericEnemyFpt.Zone))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleHealer, genericEnemyFpt.Zone))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyFpt.Zone))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyFpt.Zone))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyFpt.Zone))
	switch genericEnemyFpt.LptType {
	case PTTypeTTHHDDDD:
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyFpt.Zone))
	case PTTypeTTHHDDDB:
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomBuff(), genericEnemyFpt.Zone))
	}
	genericEnemyFpt.Enemies = enemies
	return genericEnemyFpt
}

//ランダムなソロ敵を生成
func CreateRundomEnemyIndv(geneEnvironment GeneEnvironment) GeneUnitEnemy {
	role := GetRoleRandom()
	zone := geneEnvironment.choiceRandomZone()
	genericEnemy := CreateRundomEnemyWithType(geneEnvironment, role, zone)
	return genericEnemy
}

//ランダムなライトパーティの生成
func CreateRundomEnemyLPT(geneEnvironment GeneEnvironment) GenericEnemyPT {
	genericEnemyLpt := GenericEnemyPT{}
	genericEnemyLpt.LptType = PTType(lottery.GetRandomInt(int(PTTypeTHD), int(PTTypeTHDDB)))
	genericEnemyLpt.Zone = geneEnvironment.choiceRandomZone()
	enemies := []GeneUnitEnemy{}
	switch genericEnemyLpt.LptType {
	case PTTypeTHD:
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleTank, genericEnemyLpt.Zone))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleHealer, genericEnemyLpt.Zone))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyLpt.Zone))
	case PTTypeTHDD:
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleTank, genericEnemyLpt.Zone))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleHealer, genericEnemyLpt.Zone))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyLpt.Zone))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyLpt.Zone))
	case PTTypeTHDDB:
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleTank, genericEnemyLpt.Zone))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleHealer, genericEnemyLpt.Zone))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyLpt.Zone))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyLpt.Zone))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomBuff(), genericEnemyLpt.Zone))
	}
	genericEnemyLpt.Enemies = enemies
	return genericEnemyLpt
}

//TTHHDDDD

//パーティ
type GenericEnemyPT struct {
	LptType PTType //パーティの種別
	Zone    JsonZone
	Enemies []GeneUnitEnemy
}


//ランダムな個体を生成
func CreateRandomGeneUnit(geneEnvironment GeneEnvironment) *GeneUnit {
	geneUnit := &GeneUnit{}
	geneUnit.GenericEnemyNum = lottery.GetRandomInt(10, 80)
	geneUnit.GenericEnemyLPT_Num = lottery.GetRandomInt(0, 20) //3-5人パーティ
	geneUnit.GenericEnemyFPT_Num = lottery.GetRandomInt(0, 10) //8人パーティ

	geneUnitEnemies := []GeneUnitEnemy{}

	//敵生成
	for i := 0; i < geneUnit.GenericEnemyLPT_Num; i++ {
		genericEnemyLPT := CreateRundomEnemyLPT(geneEnvironment)
		for j := 0; j < len(genericEnemyLPT.Enemies); j++ {
			geneUnitEnemies = append(geneUnitEnemies, genericEnemyLPT.Enemies[j])
		}
	}
	for i := 0; i < geneUnit.GenericEnemyFPT_Num; i++ {
		genericEnemyFPT := CreateRundomEnemyFPT(geneEnvironment)
		for j := 0; j < len(genericEnemyFPT.Enemies); j++ {
			geneUnitEnemies = append(geneUnitEnemies, genericEnemyFPT.Enemies[j])
		}
	}
	geneUnit.GenericEnemyIndv_Num = geneUnit.GenericEnemyNum - len(geneUnitEnemies)
	for i := 0; i < geneUnit.GenericEnemyIndv_Num; i++ {
		indvEnemy := CreateRundomEnemyIndv(geneEnvironment)
		geneUnitEnemies = append(geneUnitEnemies, indvEnemy)
	}

	geneUnit.GenericUnitEnemies = geneUnitEnemies

	geneUnit.calcFit()
	return geneUnit
}

//Fitの算出
func (geneUnit *GeneUnit) calcFit() {
	fit := 0
	for _, enemy := range geneUnit.GenericUnitEnemies {
		fit += enemy.getFit()
	}
	geneUnit.Fit = fit
}

//個体をランダムN個生成する
func CreateRundomsWithGeneUnitsPerAge(geneEnvironment GeneEnvironment) []*GeneUnit {
	geneUnitsPerAge := []*GeneUnit{}

	for i := 0; i < geneEnvironment.NumPerAge; i++ {
		geneUnit := CreateRandomGeneUnit(geneEnvironment)
		geneUnitsPerAge = append(geneUnitsPerAge, geneUnit)
	}

	return geneUnitsPerAge
}

//アルゴリズムで使う変数を一括で作る
func CreateGeneEnvironment(zones []JsonZone, questEnvironment QuestEnvironment) GeneEnvironment {
	dst := []JsonZone{}
	copy(dst, zones)
	return GeneEnvironment{
		NumPerAge: 10,
		Zones: dst,
		QuestEnvironment: questEnvironment,
		EnemySamples: CreateEnemySamples(),
	}
}


