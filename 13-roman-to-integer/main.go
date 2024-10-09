package main

func romanToInt(s string) int {
	romanNumeralMap := map[string]int{
		"I":  1,
		"V":  5,
		"X":  10,
		"L":  50,
		"C":  100,
		"D":  500,
		"M":  1000,
		"IV": 4,
		"IX": 9,
		"XL": 40,
		"XC": 90,
		"CD": 400,
		"CM": 900,
	}

	var t, lv int
	for idx := len(s) - 1; idx >= 0; idx-- {
		currentChar := string(s[idx])
		value := romanNumeralMap[currentChar]
		if lv > 0 && romanNumeralMap[currentChar+string(s[idx+1])] != 0 {
			t -= value
		} else {
			t += value
		}
		lv = value
	}
	return t
}
