# qEnemyGenerator

qEnemyGeneratorは、ステージ情報自動生成ツールです。

遺伝的アルゴリズムで、敵の種類・装備・レベル・配置・出てくるタイミングを決定します。

作者のadachicは、従来、エクセルで敵情報を調整していましたが、このツールを用いることにより、クエストの難易度を数値で設定するだけで最適なステージ情報を生成します。

主な機能：

- 敵のパーティ生成（タンク・ヒーラー・DPSといった変性）を適正に配置
- あらゆるマップ形状に対応

# 入力

- マップ情報(qMapBuilderかqMapEditorで出力したもの)
- クエスト情報
- アイテム・装備情報
- キャラクタ基本情報

# 出力

- 敵配置情報(json形式)

# 使い方

```
go build && ./qEnemyGenerator \
        -map map/$QUEST_ID.map.json \
        -quest debug/quest.json \
        -eqp debug/eqp.json \
        -character debug/character.json \
        -questId $QUEST_ID
```

