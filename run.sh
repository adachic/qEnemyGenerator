#!/bin/sh
QUEST_ID=$1

if [ $QUEST_ID = '' ] ; then
        exit
fi

go build && ./qEnemyGenerator \
        -map map/$QUEST_ID.map.json \
        -quest debug/quest.json \
        -eqp debug/eqp.json \
        -character debug/character.json \
        -questId $QUEST_ID
