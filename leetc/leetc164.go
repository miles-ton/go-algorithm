package leetc

import "strings"

func compareVersion(version1 string, version2 string) int {

	vs1, vs2 := strings.Split(version1, "."), strings.Split(version2, ".")
	i := 0
	for ; i < len(vs1) && i < len(vs2); i++ {
		if strings.Compare(trimPreZero(vs1[i]), trimPreZero(vs2[i])) == 0 {
			continue
		} else {
			return compDigitStr(trimPreZero(vs1[i]), trimPreZero(vs2[i]))
		}
	}
	for ; i < len(vs1); i++ {
		if trimPreZero(vs1[i]) == "0" {
			continue
		} else {
			return 1
		}
	}

	for ; i < len(vs2); i++ {
		if trimPreZero(vs2[i]) == "0" {
			continue
		} else {
			return -1
		}
	}
	return 0
}

func trimPreZero(rawStr string) string {
	i := 0
	for ; i < len(rawStr)-1; i++ {
		if rawStr[i] == uint8('0') {
			continue
		} else {
			break
		}
	}
	return rawStr[i:]
}

func compDigitStr(str1, str2 string) int {
	if len(str1) > len(str2) {
		return 1
	} else if len(str1) < len(str2) {
		return -1
	}
	for i := 0; i < len(str1); i++ {
		if str1[i] == str2[i] {
			continue
		} else if str1[i] > str2[i] {
			return 1
		} else {
			return -1
		}
	}
	return 0
}
