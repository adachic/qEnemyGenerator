package main

import (
	"fmt"
	"io/ioutil"
	"encoding/json"
	"os"
	"encoding/csv"
	"strconv"
	"io"
)

func ConvCsvJson(jtocmode int, jsonFilePath string, csvFilePath string) {
	if (jtocmode == 1) {
		var jsonStub JsonStub
		//Json読み込み
		{
			file, err := ioutil.ReadFile(jsonFilePath)

			json_err := json.Unmarshal(file, &jsonStub)
			if err != nil {
				fmt.Printf("Format Error: ", json_err)
			}
		}

		//CSV出力
		{
			// fmt.Println(jsonStub.EnemyAppears);

			file, _ := os.Create(csvFilePath)
			writer := csv.NewWriter(file) //utf8
			{
				var columns []string
				columns = append(columns, "Id");
				columns = append(columns, "SId");
				columns = append(columns, "CId");
				columns = append(columns, "CIdStr");
				columns = append(columns, "Lv");

				columns = append(columns, "MEqpId");
				columns = append(columns, "MEqpLv");
				columns = append(columns, "MEqpType");
				columns = append(columns, "MEqpSubType");
				columns = append(columns, "MEqpName");

				columns = append(columns, "S1EqpId");
				columns = append(columns, "S1EqpType");
				columns = append(columns, "S1EqpSubType");
				columns = append(columns, "S1EqpName");

				columns = append(columns, "S2EqpId");
				columns = append(columns, "S2EqpType");
				columns = append(columns, "S2EqpSubType");
				columns = append(columns, "S2EqpName");

				columns = append(columns, "S3EqpId");
				columns = append(columns, "S3EqpType");
				columns = append(columns, "S3EqpSubType");
				columns = append(columns, "S3EqpName");

				columns = append(columns, "AI");
				columns = append(columns, "AppearTime");

				columns = append(columns, "ZoneId");
				columns = append(columns, "pos1");
				columns = append(columns, "pos2");
				columns = append(columns, "pos3");
				columns = append(columns, "pos4");
				columns = append(columns, "pos5");
				columns = append(columns, "pos6");
				writer.Write(columns);
			}

			for _, enemyAppear := range jsonStub.EnemyAppears {
				var columns []string
				columns = append(columns, strconv.Itoa(enemyAppear.Id))
				columns = append(columns, strconv.Itoa(enemyAppear.Sample.Id))
				columns = append(columns, strconv.Itoa(int(enemyAppear.Sample.CharacterId)))
				columns = append(columns, enemyAppear.Sample.CharacterIdStr)
				columns = append(columns, strconv.Itoa(enemyAppear.Sample.UnitLevel))

				columns = append(columns, strconv.Itoa(enemyAppear.Sample.MainEqp.Id))
				columns = append(columns, strconv.Itoa(enemyAppear.Sample.MainEqpLevel))
				columns = append(columns, string(enemyAppear.Sample.MainEqp.Type))
				columns = append(columns, string(enemyAppear.Sample.MainEqp.SubType))
				columns = append(columns, enemyAppear.Sample.MainEqp.Name)

				columns = append(columns, strconv.Itoa(enemyAppear.Sample.SubEqp1.Id))
				columns = append(columns, string(enemyAppear.Sample.SubEqp1.Type))
				columns = append(columns, string(enemyAppear.Sample.SubEqp1.SubType))
				columns = append(columns, enemyAppear.Sample.SubEqp1.Name)

				columns = append(columns, strconv.Itoa(enemyAppear.Sample.SubEqp2.Id))
				columns = append(columns, string(enemyAppear.Sample.SubEqp2.Type))
				columns = append(columns, string(enemyAppear.Sample.SubEqp2.SubType))
				columns = append(columns, enemyAppear.Sample.SubEqp2.Name)

				columns = append(columns, strconv.Itoa(enemyAppear.Sample.SubEqp3.Id))
				columns = append(columns, string(enemyAppear.Sample.SubEqp3.Type))
				columns = append(columns, string(enemyAppear.Sample.SubEqp3.SubType))
				columns = append(columns, enemyAppear.Sample.SubEqp3.Name)

				columns = append(columns, string(enemyAppear.AIType))
				columns = append(columns, strconv.Itoa(enemyAppear.AppearTime))

				columns = append(columns, strconv.Itoa(enemyAppear.Zone.Id))
				columns = append(columns, strconv.Itoa(enemyAppear.Zone.Pos1))
				columns = append(columns, strconv.Itoa(enemyAppear.Zone.Pos2))
				columns = append(columns, strconv.Itoa(enemyAppear.Zone.Pos3))
				columns = append(columns, strconv.Itoa(enemyAppear.Zone.Pos4))
				columns = append(columns, strconv.Itoa(enemyAppear.Zone.Pos5))
				columns = append(columns, strconv.Itoa(enemyAppear.Zone.Pos6))

				writer.Write(columns)
			}
			writer.Flush()
		}

		return;
	}
	var jsonStub JsonStub

	//CSV読み込み
	{
		var enemyAppears []*EnemyAppear
		file, _ := os.Open(csvFilePath)
		reader := csv.NewReader(file)
		for {
			record, err := reader.Read()
			if err == io.EOF {
				break
			} else {
				fmt.Println("read error")
			}
			id , _:= strconv.Atoi(record[0])
			if(id == 0){
				continue
			}
			sid, _ := strconv.Atoi(record[1])
			cid, _ := strconv.Atoi(record[2])
			ulv, _ := strconv.Atoi(record[4])
			eqpId, _ := strconv.Atoi(record[5])
			s1Id, _ := strconv.Atoi(record[10])
			s2Id, _ := strconv.Atoi(record[14])
			s3Id, _ := strconv.Atoi(record[18])
			appearTime, _ := strconv.Atoi(record[23])
			zoneId, _ := strconv.Atoi(record[24])
			pos1, _ := strconv.Atoi(record[25])
			pos2, _ := strconv.Atoi(record[26])
			pos3, _ := strconv.Atoi(record[27])
			pos4, _ := strconv.Atoi(record[28])
			pos5, _ := strconv.Atoi(record[29])
			pos6, _ := strconv.Atoi(record[30])


			enemyAppears = append(enemyAppears,
				&EnemyAppear{
					Id:id,
					Sample:&EnemySample{
						Id: sid,
						CharacterId: CharacterId(cid),
						CharacterIdStr: record[3],
						UnitLevel: ulv,
						MainEqp: JsonGameEqp{
							Id:eqpId,
							Type:G_jsonGameEqps[strconv.Itoa(eqpId)].Type,
							SubType:G_jsonGameEqps[strconv.Itoa(eqpId)].SubType,
							Name:G_jsonGameEqps[strconv.Itoa(eqpId)].Name,
						},
						SubEqp1: JsonGameEqp{
							Id:s1Id,
							Type:G_jsonGameEqps[strconv.Itoa(s1Id)].Type,
							SubType:G_jsonGameEqps[strconv.Itoa(s1Id)].SubType,
							Name:G_jsonGameEqps[strconv.Itoa(s1Id)].Name,
						},
						SubEqp2: JsonGameEqp{
							Id:s2Id,
							Type:G_jsonGameEqps[strconv.Itoa(s2Id)].Type,
							SubType:G_jsonGameEqps[strconv.Itoa(s2Id)].SubType,
							Name:G_jsonGameEqps[strconv.Itoa(s2Id)].Name,
						},
						SubEqp3: JsonGameEqp{
							Id:s3Id,
							Type:G_jsonGameEqps[strconv.Itoa(s3Id)].Type,
							SubType:G_jsonGameEqps[strconv.Itoa(s3Id)].SubType,
							Name:G_jsonGameEqps[strconv.Itoa(s3Id)].Name,
						},

					},
					AIType: AIType(record[22]),
					AppearTime: appearTime,
					Zone:JsonZone{
						Id:zoneId,
						Pos1:pos1,
						Pos2:pos2,
						Pos3:pos3,
						Pos4:pos4,
						Pos5:pos5,
						Pos6:pos6,
					},
				})
		}
		jsonStub.EnemyAppears = enemyAppears

	}

	//Json出力
	{
		bytes, json_err := json.Marshal(jsonStub)
		if json_err != nil {
			fmt.Println("Json Encode Error: ", json_err)
		}

		//	fmt.Printf("bytes:%+v\n", string(bytes))

		file, err := os.Create(jsonFilePath)
		_, err = file.Write(bytes)
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}


