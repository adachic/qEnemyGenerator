package main

func EnemiesWithZone(creteriaEvaluationPerSlice int, zones []JsonZone) []EnemyAppear {

	//個体をランダムN個生成する
	//それぞれの適応度計算する

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

type Enemy struct {
	characterId CharacterId
}

