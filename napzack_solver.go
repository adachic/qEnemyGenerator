package main
import (
	"github.com/adachic/lottery"
	"fmt"
//	"gopkg.in/go-pp/pp.v2"
)

//遺伝的アルゴリズムのパラメータ
type GeneEnvironment struct {
	GeneNumPerAge    int //世代あたりの個体数N
	Ages             int //世代数
	Insections       int //交差点数
	Zones            []JsonZone
	QuestEnvironment QuestEnvironment
	JsonGameMap      JsonGameMap
	EnemySamples     []Enemy
}

//個体とその遺伝子
type GeneUnit struct {
	GenericEnemyNum      int              //敵の数
	GenericEnemyLPT_Num  int              //ライトパーティの数
	GenericEnemyFPT_Num  int              //フルパーティーの数
	GenericEnemyIndv_Num int              //単体の数
	GenericUnitEnemies   []*GeneUnitEnemy //スライスあたりの敵一覧

	Fit                  int              //適応度
}

func (geneUnit *GeneUnit)copy() *GeneUnit {
	clonedSlice := make([]*GeneUnitEnemy, len(geneUnit.GenericUnitEnemies))
	copy(clonedSlice, geneUnit.GenericUnitEnemies)
	return &GeneUnit{
		GenericUnitEnemies: clonedSlice,
	}
}

//敵個体
type GeneUnitEnemy struct {
	enemy   Enemy
	eqp     JsonGameEqp
	zone    JsonZone
	ptId    int //ptの全体でのユニークid
	ptCount int //ptに何人いるか
}

//Fitを返す
func (geneUnit *GeneUnit)getFit(geneEnvironment GeneEnvironment) int {
	geneUnit.calcFit(geneEnvironment)
	return geneUnit.Fit
}

func (geneUnitEnemy GeneUnitEnemy)dumpFit(geneEnvironment GeneEnvironment) {
	fit1 := geneUnitEnemy.enemy.getFit()
	fit2 := geneUnitEnemy.zone.getFit(geneUnitEnemy.enemy.fixedRole, geneEnvironment)
	fit3 := geneUnitEnemy.eqp.getFit()
	fmt.Printf("[(%d):%d/%d/%d]", fit1+fit2+fit3,fit1,fit2,fit3)
}


//Fitを返す
func (geneUnitEnemy GeneUnitEnemy)getFit(geneEnvironment GeneEnvironment) int {
	fit := 0
	fit += geneUnitEnemy.enemy.getFit()
	fit += geneUnitEnemy.zone.getFit(geneUnitEnemy.enemy.fixedRole, geneEnvironment)
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

//n点交差
func (geneUnit *GeneUnit)IntersectGeneUnitWith(other *GeneUnit, geneEnvironment GeneEnvironment) bool {
	tradedPickupedPtIds := []int{}

	err := true

	//n個のPTか、Indvを交差
	for i := 0; i < geneEnvironment.Insections; i++ {
		//まずランダムに選択
		fmt.Printf("///%+v\n", *geneUnit)
		pickupedEnemiesA := geneUnit.pickupedEnemies()

		if (len(tradedPickupedPtIds) > 0&&
		tradedPickupedPtIds[0] == pickupedEnemiesA[0].ptId) {
			//前回に交換したことあるやつだったらやりなおす
			continue
		}

		//何体いるか
		numForIntersection := len(pickupedEnemiesA)

		//もう一方からも選択
		pickupedEnemiesB, err2 := other.pickupedEnemiesWithNum(numForIntersection)
		if (err2) {
			continue
		}

		//AとBをトレード
		Trade(pickupedEnemiesA, pickupedEnemiesB)

		//トレードしたPtIdを保存
		tradedPickupedPtIds = append(tradedPickupedPtIds, pickupedEnemiesA[0].ptId)

		fmt.Println("[TRADE]tradeDone.")
		err = false
	}
	return err
}

//交換する
func Trade(geneUnitEnemiesA []*GeneUnitEnemy, geneUnitEnemiesB []*GeneUnitEnemy) {
	//	tmp := make([]*GeneUnitEnemy, len(geneUnitEnemiesA))
	tmp := []*GeneUnitEnemy{}

	fmt.Printf("[TRADE]A:%+v\n", geneUnitEnemiesA)
	fmt.Printf("[TRADE]B:%+v\n", geneUnitEnemiesB)
	//B->tmp
	for _, enemyB := range geneUnitEnemiesB {
		tmp = append(tmp, enemyB)
	}
	//A->B
	for i, enemyA := range geneUnitEnemiesA {
		geneUnitEnemiesB[i] = enemyA
	}
	//tmp->A
	//	fmt.Printf("T:%+v\n", tmp)
	//	fmt.Printf("A:%+v\n", geneUnitEnemiesA)
	for i, enemyB := range tmp {
		geneUnitEnemiesA[i] = enemyB
	}
}

//ランダムにn体のPTかIndvを選択
func (geneunit *GeneUnit)pickupedEnemiesWithNum(numForIntersection int) (geneUnitEnemies []*GeneUnitEnemy, err bool) {
	candidatePtIds := []int{}
	for ptId := 0; ptId < len(geneunit.GenericUnitEnemies); ptId++ {
		num := 0
		for _, enemy := range geneunit.GenericUnitEnemies {
			if (enemy.ptId == ptId) {
				num++
			}
		}
		if (num == numForIntersection) {
			candidatePtIds = append(candidatePtIds, ptId)
		}
	}
	if (len(candidatePtIds) == 0) {
		//numForIntersectionにあったPTはないので
		//気合で取り出す
		//		fmt.Printf("candidatePtIds:%+v,numForIntersection:%+v\n", candidatePtIds, numForIntersection)
		restNeedCount := numForIntersection
		cnt := 0
		//		fmt.Printf("aho1:restNeedCnt:%d\n", restNeedCount)
		for _, enemy := range geneunit.GenericUnitEnemies {
			//			fmt.Printf("aho1:restNeedCnt:%d, enemyPtCount:%d\n", restNeedCount, enemy.ptCount)
			if (enemy.ptCount > restNeedCount) {
				continue
			}
			cnt++
			geneUnitEnemies = append(geneUnitEnemies, enemy)
			if (cnt >= enemy.ptCount) {
				restNeedCount -= enemy.ptCount
				cnt = 0
			}
		}
		if (restNeedCount != 0) {
			fmt.Printf("gunu:%d\n", restNeedCount)
			return nil, true
		}
	}else {
		fmt.Printf("candidatePtIds:%+v,numForIntersection:%+v\n", candidatePtIds, numForIntersection)
		idx := 0
		if (len(candidatePtIds) >= 2) {
			idx = lottery.GetRandomInt(0, len(candidatePtIds) - 1)
		}
		choicedPtId := candidatePtIds[idx]
		for _, enemy := range geneunit.GenericUnitEnemies {
			if (enemy.ptId == choicedPtId) {
				geneUnitEnemies = append(geneUnitEnemies, enemy)
			}
		}
	}
	return geneUnitEnemies, false
}

//ランダムにPTかIndvを選択する
func (geneunit *GeneUnit)pickupedEnemies() (geneUnitEnemies []*GeneUnitEnemy) {
	geneUnitEnemy := geneunit.GenericUnitEnemies[lottery.GetRandomInt(0, len(geneunit.GenericUnitEnemies) - 1)]
	for _, unitEnemy := range geneunit.GenericUnitEnemies {
		if (geneUnitEnemy.ptId == unitEnemy.ptId) {
			geneUnitEnemies = append(geneUnitEnemies, unitEnemy)
		}
	}
	return geneUnitEnemies
}

//突然変異
func (geneunit *GeneUnit)MutateSuddenly(geneEnvironment GeneEnvironment) {
	pickedUpEnemies := geneunit.pickupedEnemies()
	fmt.Printf("[MUTATE]ptId:%+v, ptCount:%+v\n", pickedUpEnemies[0].ptId, pickedUpEnemies[0].ptCount)
	for _, enemy := range pickedUpEnemies {
		enemy.mutate(geneEnvironment)
	}
}

//最もFitの高い個を返す
func GetMaxFitGene(geneUnits []*GeneUnit,geneEnvironment GeneEnvironment) *GeneUnit {
	geneMaxFitUnit := &GeneUnit{Fit:0}
	for _, unit := range geneUnits {
		if geneMaxFitUnit.Fit < unit.getFit(geneEnvironment) {
			geneMaxFitUnit = unit
		}
	}
	return geneMaxFitUnit
}

func EnemiesWithZone(creteriaEvaluationPerSlice int, zones []JsonZone, questEnvironment QuestEnvironment, geneEnvironment GeneEnvironment) []EnemyAppear {

	println("=== creteriaEvaluationPerSlice:", creteriaEvaluationPerSlice)

	//次世代の器
	nextGeneUnits := make([]*GeneUnit, geneEnvironment.GeneNumPerAge)

	//個体をランダムN個生成する
	//それぞれの適応度計算する
	geneUnitsPerAge := CreateRundomsWithGeneUnitsPerAge(creteriaEvaluationPerSlice, geneEnvironment);
	for age := 0; age < geneEnvironment.Ages; age++ {
		println("===age", age)
		Scan()
		for i := 0; i < len(geneUnitsPerAge); i++ {
			//世代操作開始
			//選択
			//エリートの選択
			println("===", i, "/", len(geneUnitsPerAge))
			relottery:
			//			println("gomi")
			{
				operationType := lottery.GetRandomInt(0, 100)
				//次のいずれかを行う
				switch {
				case operationType <= 1:{
					//B.突然変異
					var src1 *GeneUnit
					var idx int
					for {
						//						println("gomi0")
						idx = lottery.GetRandomInt(0, len(geneUnitsPerAge))
						src1 = geneUnitsPerAge[idx]
						if (src1 != nil) {
							break
						}
					}
					src1.MutateSuddenly(geneEnvironment)
					nextGeneUnits[i] = src1.copy()
					geneUnitsPerAge[idx] = nil
				}
				case operationType <= 20:{
					//C.次世代にそのまま追加
					var src1 *GeneUnit
					var idx int
					for {
						//					println("gomi1")
						//idx = lottery.GetRandomInt(0, len(geneUnitsPerAge) - 1)
						idx = lottery.GetRandomInt(0, len(geneUnitsPerAge))
						/*
						fmt.Printf("[COPY]unko50--:%+v len:%d, cap:%d, idx:%d\n",
							geneUnitsPerAge, len(geneUnitsPerAge), cap(geneUnitsPerAge), idx)
							*/
						src1 = geneUnitsPerAge[idx]
						if (src1 != nil) {
							break
						}
					}
					fmt.Printf("[COPY]idxPerAge:%+v\n", idx)
					nextGeneUnits[i] = geneUnitsPerAge[idx].copy()
					geneUnitsPerAge[idx] = nil
				}
				case operationType <= 100:{
					if (i + 1 >= len(geneUnitsPerAge)) {
						goto relottery
					}else {
						//A.個体を2つ選択し、交差
						var src1 *GeneUnit
						var src2 *GeneUnit
						var idx1 int
						var idx2 int
						for {
							//							println("gomi2")
							idx1 = lottery.GetRandomInt(0, len(geneUnitsPerAge))
							src1 = geneUnitsPerAge[idx1]
							if (src1 != nil) {
								break
							}
						}
						for {
							//							println("gomi3")
							idx2 = lottery.GetRandomInt(0, len(geneUnitsPerAge))
							src2 = geneUnitsPerAge[idx2]
							if (src2 != nil) {
								break
							}
						}
						fmt.Printf("[TRADE]startTrade:\n")
						err := src1.IntersectGeneUnitWith(src2, geneEnvironment)
						if (err) {
							fmt.Printf("[TRADE]trade failed.....:\n")
							goto relottery
						}
						fmt.Printf("[TRADE]trade completed!!!:\n")
						nextGeneUnits[i] = src1.copy()
						nextGeneUnits[i + 1] = src2.copy()
						//						fmt.Printf("s1:%+v s2:%+v\n", *src1, *src2)
						//						fmt.Printf("1:%+v 2:%+v\n", *nextGeneUnits[i], *nextGeneUnits[i+1])
						geneUnitsPerAge[idx1] = nil
						geneUnitsPerAge[idx2] = nil
						i++
					}
				}
				}
			}
			//			fmt.Printf("[%+v]\n->\n[%+v]\n", geneUnitsPerAge, nextGeneUnits)
		}
		//次世代が一定になっていれば次世代を対象として世代操作開始に戻る
		geneUnitsPerAge = nextGeneUnits
		nextGeneUnits = make([]*GeneUnit, geneEnvironment.GeneNumPerAge)
		//		fmt.Printf("[next to age:%+v -> %+v\n", geneUnitsPerAge, nextGeneUnits)

		maxFitGeneUnit := GetMaxFitGene(geneUnitsPerAge, geneEnvironment)
		//		fmt.Printf("maxFitGene:%+v\n", *maxFitGeneUnit)
		fmt.Printf("[fit age:%d]%d\n", age, maxFitGeneUnit.Fit)

		Scan()
	}

	//最終世代で最もFitが高いものを選び、EnemyAppearに変換する
	maxFitGeneUnit := GetMaxFitGene(geneUnitsPerAge, geneEnvironment)
	fmt.Printf("[fit final]%d\n", maxFitGeneUnit.Fit)
	fmt.Printf("[fit final geneParam]%+v\n", *maxFitGeneUnit)
	for _, enemy := range maxFitGeneUnit.GenericUnitEnemies {
		fmt.Printf("[fit final enemy]%+v\n", enemy)
		//pp.Printf("[fit final enemy]%+v\n", enemy)
	}
	return nil
}

func Scan2() {
	ore := 1
	fmt.Scan(&ore)
}
func Scan() {
	return
	ore := 1
	fmt.Scan(&ore)
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
func CreateRundomEnemyWithType(geneEnvironment GeneEnvironment, role Role, zone JsonZone, ptId int, ptCount int) *GeneUnitEnemy {
	geneUnitEnemy := &GeneUnitEnemy{ptId: ptId, ptCount: ptCount}

	//種別が決定
	geneUnitEnemy.enemy = PickUpRandomSampleWithRole(geneEnvironment.EnemySamples, role)

	//EQP
	geneUnitEnemy.eqp = PickUpEQPWithSampleEnemy(geneUnitEnemy.enemy)

	//Zone
	geneUnitEnemy.zone = zone

	return geneUnitEnemy
}

//ランダムなフルパーティーの生成
func CreateRundomEnemyFPT(geneEnvironment GeneEnvironment, ptId int) GenericEnemyPT {
	genericEnemyFpt := GenericEnemyPT{PtId: ptId}
	genericEnemyFpt.LptType = PTType(lottery.GetRandomInt(int(PTTypeTTHHDDDD), int(PTTypeTTHHDDDB)))
	genericEnemyFpt.Zone = geneEnvironment.choiceRandomZone()
	enemies := []*GeneUnitEnemy{}
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleTank, genericEnemyFpt.Zone, ptId, 8))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleTank, genericEnemyFpt.Zone, ptId, 8))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleHealer, genericEnemyFpt.Zone, ptId, 8))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleHealer, genericEnemyFpt.Zone, ptId, 8))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyFpt.Zone, ptId, 8))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyFpt.Zone, ptId, 8))
	enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyFpt.Zone, ptId, 8))
	switch genericEnemyFpt.LptType {
	case PTTypeTTHHDDDD:
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyFpt.Zone, ptId, 8))
	case PTTypeTTHHDDDB:
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomBuff(), genericEnemyFpt.Zone, ptId, 8))
	}
	genericEnemyFpt.Enemies = enemies
	return genericEnemyFpt
}

//ランダムなソロ敵を生成
func CreateRundomEnemyIndv(geneEnvironment GeneEnvironment, ptId int) *GeneUnitEnemy {
	role := GetRoleRandom()
	zone := geneEnvironment.choiceRandomZone()
	genericEnemy := CreateRundomEnemyWithType(geneEnvironment, role, zone, ptId, 1)
	return genericEnemy
}

//突然変異
func (geneUnitEnemy *GeneUnitEnemy)mutate(geneEnvironment GeneEnvironment) {
	geneUnitEnemy = CreateRundomEnemyWithType(geneEnvironment,
		geneUnitEnemy.enemy.pickUpRandomRole(),
		geneUnitEnemy.zone,
		geneUnitEnemy.ptId,
		geneUnitEnemy.ptCount,
	)
}

//ランダムなライトパーティの生成
func CreateRundomEnemyLPT(geneEnvironment GeneEnvironment, ptId int) GenericEnemyPT {
	genericEnemyLpt := GenericEnemyPT{PtId:ptId}
	genericEnemyLpt.LptType = PTType(lottery.GetRandomInt(int(PTTypeTHD), int(PTTypeTHDDB)))
	genericEnemyLpt.Zone = geneEnvironment.choiceRandomZone()
	enemies := []*GeneUnitEnemy{}
	switch genericEnemyLpt.LptType {
	case PTTypeTHD:
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleTank, genericEnemyLpt.Zone, ptId, 3))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleHealer, genericEnemyLpt.Zone, ptId, 3))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyLpt.Zone, ptId, 3))
	case PTTypeTHDD:
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleTank, genericEnemyLpt.Zone, ptId, 4))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleHealer, genericEnemyLpt.Zone, ptId, 4))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyLpt.Zone, ptId, 4))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyLpt.Zone, ptId, 4))
	case PTTypeTHDDB:
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleTank, genericEnemyLpt.Zone, ptId, 5))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, RoleHealer, genericEnemyLpt.Zone, ptId, 5))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyLpt.Zone, ptId, 5))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomDPS(), genericEnemyLpt.Zone, ptId, 5))
		enemies = append(enemies, CreateRundomEnemyWithType(geneEnvironment, GetRoleRandomBuff(), genericEnemyLpt.Zone, ptId, 5))
	}
	genericEnemyLpt.Enemies = enemies
	return genericEnemyLpt
}

//TTHHDDDD

//パーティ
type GenericEnemyPT struct {
	PtId    int
	LptType PTType //パーティの種別
	Zone    JsonZone
	Enemies []*GeneUnitEnemy
}


//ランダムな個体を生成
func CreateRandomGeneUnit(canCreateMaxNum int, geneEnvironment GeneEnvironment, ptId *int) *GeneUnit {
	geneUnit := &GeneUnit{}
	willCreateNum := lottery.GetRandomInt(1, canCreateMaxNum)
	geneUnit.GenericEnemyNum = lottery.GetRandomInt(1, willCreateNum)
	geneUnit.GenericEnemyLPT_Num = lottery.GetRandomInt(0, willCreateNum / 3) //3-5人パーティ
	geneUnit.GenericEnemyFPT_Num = lottery.GetRandomInt(0, willCreateNum / 8) //8人パーティ

	geneUnitEnemies := []*GeneUnitEnemy{}

	//敵生成
	for i := 0; i < geneUnit.GenericEnemyLPT_Num; i++ {
		genericEnemyLPT := CreateRundomEnemyLPT(geneEnvironment, *ptId + i)
		for j := 0; j < len(genericEnemyLPT.Enemies); j++ {
			geneUnitEnemies = append(geneUnitEnemies, genericEnemyLPT.Enemies[j])
		}
		*ptId = i + *ptId
	}
	*ptId++
	for i := 0; i < geneUnit.GenericEnemyFPT_Num; i++ {
		genericEnemyFPT := CreateRundomEnemyFPT(geneEnvironment, *ptId + i)
		for j := 0; j < len(genericEnemyFPT.Enemies); j++ {
			geneUnitEnemies = append(geneUnitEnemies, genericEnemyFPT.Enemies[j])
		}
		*ptId = i + *ptId
	}
	*ptId++
	geneUnit.GenericEnemyIndv_Num = geneUnit.GenericEnemyNum - len(geneUnitEnemies)
	for i := 0; i < geneUnit.GenericEnemyIndv_Num; i++ {
		indvEnemy := CreateRundomEnemyIndv(geneEnvironment, *ptId + i)
		geneUnitEnemies = append(geneUnitEnemies, indvEnemy)
	}
	*ptId++

	geneUnit.GenericUnitEnemies = geneUnitEnemies
	geneUnit.calcFit(geneEnvironment)
	fmt.Printf("unko:geneUnitEnemies:%+v fit(%d) ", geneUnitEnemies, geneUnit.Fit)

	geneUnit.dumpEnemyFit(geneEnvironment)
	fmt.Printf("\n")

	return geneUnit
}


func (geneUnit *GeneUnit) dumpEnemyFit(geneEnvironment GeneEnvironment) {
	for _, enemy := range geneUnit.GenericUnitEnemies {
		enemy.dumpFit(geneEnvironment)
	}
}

//Fitの算出
func (geneUnit *GeneUnit) calcFit(geneEnvironment GeneEnvironment) {
	fit := 0
	for _, enemy := range geneUnit.GenericUnitEnemies {
		fit += enemy.getFit(geneEnvironment)
	}
	geneUnit.Fit = fit
}

//個体をランダムN個生成する
func CreateRundomsWithGeneUnitsPerAge(creteriaEvaluationPerSlice int, geneEnvironment GeneEnvironment) []*GeneUnit {
	geneUnitsPerAge := []*GeneUnit{}

	numPerAge := creteriaEvaluationPerSlice / geneEnvironment.QuestEnvironment.PointPerOne
	ptId := 0
	for i := 0; i < geneEnvironment.GeneNumPerAge; i++ {
		geneUnit := CreateRandomGeneUnit(numPerAge, geneEnvironment, &ptId)
		geneUnitsPerAge = append(geneUnitsPerAge, geneUnit)
	}
	fmt.Printf("[]geneUnitsPerAge:%+v\n", geneUnitsPerAge)

	return geneUnitsPerAge
}

//アルゴリズムで使う変数を一括で作る
func CreateGeneEnvironment(zones []JsonZone, questEnvironment QuestEnvironment, gameMap JsonGameMap) GeneEnvironment {
	//	dst := [len(zones)]JsonZone{}
	dst := make([]JsonZone, len(zones))
	copy(dst, zones)
	return GeneEnvironment{
		GeneNumPerAge: 10,
		Zones: dst,
		EnemySamples: CreateEnemySamples(),
		Ages: 10,
		Insections: 2,
		QuestEnvironment: questEnvironment,
		JsonGameMap:gameMap,
	}
}


