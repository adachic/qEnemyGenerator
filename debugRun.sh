#!/bin/sh
QUEST_ID=1

go build && ./qEnemyGenerator \
        -map map/5cd86000-4926-4cdb-a6b0-c0f9a6628383.json \
        -quest debug/quest.json \
        -eqp debug/eqp.json \
        -character debug/character.json \
        -questId $QUEST_ID
