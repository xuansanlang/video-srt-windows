package tool

import (
	"strings"
)


func ChineseNumberToLowercaseLength(s string) int {
	if (!IsChineseNumber(s)) {
		return 0
	}
	s = strings.TrimSpace(s)
	zhTexts := strings.Split(s , "")
	zhTextsLens := len(zhTexts)

	numberPosi := true
	maxBaseNumber := 1

	for i:=0; i<zhTextsLens; i++ {
		if numberPosi == false {
			switch zhTexts[i] {
			case "十":
				maxBaseNumber = 2
				break
			case "百":
				maxBaseNumber = 3
				break
			case "千":
				maxBaseNumber = 4
				break
			case "万":
				maxBaseNumber = 5
				break
			}
			break
		}
		numberPosi = !numberPosi
	}
	return maxBaseNumber
}


func IsChineseNumber(text string) bool {
	text = strings.TrimSpace(text)
	if text == "" {
		return false
	}

	zhTexts := strings.Split(text , "")
	zhTextsLens := len(zhTexts)

	numberPosi := true
	for i:=0; i<zhTextsLens; i++ {
		if !ValiChineseNumberChar(zhTexts[i] , !numberPosi) {
			return false
		}

		numberPosi = !numberPosi
	}
	return true
}


func ValiChineseNumberChar(s string , unit bool) bool {
	zh_number := []string{"一","二","两","三","四","五","六","七","八","九","十"}
	zh_unit := []string{"十","百","千","万","亿"}

	if unit {
		for _,v := range zh_unit {
			if v == s {
				return true
			}
		}
		return false
	}

	for _,v := range zh_number {
		if v == s {
			return true
		}
	}
	return false
}
