# calc-bonus  
  
## 概要  
  
賞与の計算が行えるプログラム  
プログラムは、対話式シェルに実行できる  
  
```
go run .\calc.go
```  
  
## 計算式  
  
式は、自身の勤務する企業のものであり、汎用性はない。  
このリポジトリがプライベートなのもそのためである。  
  
> 賞与 = (基本給比例分 + 成果反映分 + 専門職手当) * 出勤率  
> 基本給比例分 = 基本給 * 基本給比例分月数  
> 成果反映分 = 評語に基づく係数 × 係数単価  
> 専門職手当 = 専門職手当 * 1 month  
> 出勤率 = 1 - ( (欠勤日数 + (遅早H / 4)) / (期間中の稼働日数) )   

