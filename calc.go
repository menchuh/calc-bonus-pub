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
	"fmt"
	"os"
	"strconv"
)

func main() {

	base := Question("基本給はいくらですか？", "int")
	month := Question("基本給比例分月数はいくつですか？", "float")
	// 基本給比例分
	x := base * month

	coefficient := Question("評語に基づく係数はいくつですか？", "int")
	unitPrice := Question("係数単価はいくらですか？", "int")
	// 成果反映分
	y := coefficient * unitPrice

	// 専門職手当
	z := Question("専門職手当はいくらですか？", "int")

	// 欠勤

  // 賞与
	result := int(x + y + z)

	// 結果表示
	fmt.Printf("あなたの賞与は、%d円 です。\n", result)
	fmt.Printf("内訳 基本給比例分:%d円, 成果反映文:%d円, 専門職手当:%d円\n", int(x), int(y), int(z))
}

func Question(q string, typestring string) float64 {

	fmt.Printf("%s: ", q)

	if typestring != "int" && typestring != "float" {
		fmt.Println("int, floatのタイプを選択する必要があります")
  	os.Exit(1)
	}

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