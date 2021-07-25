package map_vs_switch

import (
	"testing"
)

func BenchmarkOptionSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionSwitch("12")
	}
}

func BenchmarkOptionMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionMap("12")
	}
}

func BenchmarkOptionGlobalMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionMapGlobal("12")
	}
}

func OptionSwitch(s string) string {
	switch s {
	case "0":
		return "0"
	case "1":
		return "1"
	case "2":
		return "2"
	case "3":
		return "3"
	case "4":
		return "4"
	case "5":
		return "5"
	case "6":
		return "6"
	case "7":
		return "7"
	case "8":
		return "8"
	case "9":
		return "9"
	case "10":
		return "10"
	case "11":
		return "11"
	case "12":
		return "12"
	case "13":
		return "13"
	case "14":
		return "14"
	case "15":
		return "15"
	case "16":
		return "16"
	case "17":
		return "17"
	case "18":
		return "18"
	case "19":
		return "19"
	case "20":
		return "20"
	case "21":
		return "21"
	case "22":
		return "22"
	case "23":
		return "23"
	case "24":
		return "24"
	case "25":
		return "25"
	case "26":
		return "26"
	case "27":
		return "27"
	case "28":
		return "28"
	case "29":
		return "29"
	default:
		return ""
	}
}

func OptionMap(s string) string {
	return map[string]string{
		"0":  "0",
		"1":  "1",
		"2":  "2",
		"3":  "3",
		"4":  "4",
		"5":  "5",
		"6":  "6",
		"7":  "7",
		"8":  "8",
		"9":  "9",
		"10": "10",
		"11": "11",
		"12": "12",
		"13": "13",
		"14": "14",
		"15": "15",
		"16": "16",
		"17": "17",
		"18": "18",
		"19": "19",
		"20": "20",
		"21": "21",
		"22": "22",
		"23": "23",
		"24": "24",
		"25": "25",
		"26": "26",
		"27": "27",
		"28": "28",
		"29": "29",
	}[s]
}

var globalmap = map[string]string{
	"0":  "0",
	"1":  "1",
	"2":  "2",
	"3":  "3",
	"4":  "4",
	"5":  "5",
	"6":  "6",
	"7":  "7",
	"8":  "8",
	"9":  "9",
	"10": "10",
	"11": "11",
	"12": "12",
	"13": "13",
	"14": "14",
	"15": "15",
	"16": "16",
	"17": "17",
	"18": "18",
	"19": "19",
	"20": "20",
	"21": "21",
	"22": "22",
	"23": "23",
	"24": "24",
	"25": "25",
	"26": "26",
	"27": "27",
	"28": "28",
	"29": "29",
}

func OptionMapGlobal(s string) string {
	return globalmap[s]
}
