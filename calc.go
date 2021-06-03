package main

//========================================================================
// 賞与の計算を行うためのプログラム
// 賞与の計算方法は以下
// 賞与 = (基本給比例分 + 成果反映分 + 専門職手当) * 出勤率
// 基本給比例分 = 基本給 * 基本給比例分月数
// 成果反映分 = 評語に基づく係数 × 係数単価
// 専門職手当 = 専門職手当 * 1 month
// 出勤率 = 1 - ( (欠勤日数 + (遅早H / 4)) / (期間中の稼働日数) )
//========================================================================

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
)

type Strings struct {
	Messages []string `json:"messages"`
	Types []string `json:"types"`
}

func main() {

	displayStrings := Readfile("./config/strings.json")

	base := Question(displayStrings.Messages[0], "int")
	month := Question(displayStrings.Messages[1], "float")
	x := base * month

	coefficient := Question(displayStrings.Messages[2], "int")
	unitPrice := Question(displayStrings.Messages[3], "int")
	y := coefficient * unitPrice

	z := Question(displayStrings.Messages[4], "int")

	// 賞与
	result := int(x + y + z)

	// 結果表示
	fmt.Printf("あなたの賞与は、%d円 です。\n", result)
	fmt.Printf("内訳 %s:%d円, %s:%d円, %s:%d円\n", displayStrings.Types[0], int(x), displayStrings.Types[1], int(y), displayStrings.Types[2], int(z))
}

func Question(q string, typestring string) float64 {

	fmt.Printf("%s: ", q)

	integer := 0
	float := 0.0
	var err error

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {

		text := scanner.Text()

		if typestring == "int" {
			// 整数値に変換
			integer, err = strconv.Atoi(text)
		} else if typestring == "float" {
			// 小数値に変換
			float, err = strconv.ParseFloat(scanner.Text(), 64)
		} else {
			fmt.Println("int, floatのタイプを選択する必要があります")
			os.Exit(1)
		}

		if err != nil {
			fmt.Println("数値に変換できる値を入力してください")
			fmt.Printf("%s: ", q)
		} else {
			break
		}
	}

	if err := scanner.Err(); err != nil {
		panic(err)
	}

	if typestring == "int" {
		return float64(integer)
	}
	return float
}

func Readfile(f string) Strings {

	// jsonファイル読み込み
	raw, err := ioutil.ReadFile(f)
    if err != nil {
        fmt.Println(err.Error())
        os.Exit(1)
    }

	var data Strings
	json.Unmarshal(raw, &data)
	if err != nil {
        fmt.Println(err.Error())
		os.Exit(1)
    }

	return data
}