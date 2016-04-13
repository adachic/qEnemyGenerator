#!/bin/sh
MIN_QUEST_ID=$1
MAX_QUEST_ID=$2

if [ $MIN_QUEST_ID = '' ] ; then
        exit
fi

if [ $MAX_QUEST_ID = '' ] ; then
        exit
fi

for i in `seq $MIN_QUEST_ID $MAX_EUQST_ID`
do
   ./run.sh $i
done
