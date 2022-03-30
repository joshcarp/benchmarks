package map_vs_switch

import (
	"testing"
)

func BenchmarkOptionSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionSwitch("12pufn98asdf9sfd8n")
	}
}

func BenchmarkOptionMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionMap("12pufn98asdf9sfd8n")
	}
}

func BenchmarkOptionGlobalMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionMapGlobal("12pufn98asdf9sfd8n")
	}
}

func OptionSwitch(s string) string {
	switch s {
	case "0pufn98asdf9sfd8n":
		return "0pufn98asdf9sfd8n"
	case "1pufn98asdf9sfd8n":
		return "1pufn98asdf9sfd8n"
	case "2pufn98asdf9sfd8n":
		return "2pufn98asdf9sfd8n"
	case "3pufn98asdf9sfd8n":
		return "3pufn98asdf9sfd8n"
	case "4pufn98asdf9sfd8n":
		return "4pufn98asdf9sfd8n"
	case "5pufn98asdf9sfd8n":
		return "5pufn98asdf9sfd8n"
	case "6pufn98asdf9sfd8n":
		return "6pufn98asdf9sfd8n"
	case "7pufn98asdf9sfd8n":
		return "7pufn98asdf9sfd8n"
	case "8pufn98asdf9sfd8n":
		return "8pufn98asdf9sfd8n"
	case "9pufn98asdf9sfd8n":
		return "9pufn98asdf9sfd8n"
	case "1pufn98asdf9sfd8n0":
		return "1pufn98asdf9sfd8n0"
	case "1pufn98asdf9sfd8n1":
		return "1pufn98asdf9sfd8n1"
	case "1pufn98asdf9sfd8n2":
		return "1pufn98asdf9sfd8n2"
	case "1pufn98asdf9sfd8n3":
		return "1pufn98asdf9sfd8n3"
	case "1pufn98asdf9sfd8n4":
		return "1pufn98asdf9sfd8n4"
	case "1pufn98asdf9sfd8n5":
		return "1pufn98asdf9sfd8n5"
	case "1pufn98asdf9sfd8n6":
		return "1pufn98asdf9sfd8n6"
	case "1pufn98asdf9sfd8n7":
		return "1pufn98asdf9sfd8n7"
	case "1pufn98asdf9sfd8n8":
		return "1pufn98asdf9sfd8n8"
	case "1pufn98asdf9sfd8n9":
		return "1pufn98asdf9sfd8n9"
	case "2pufn98asdf9sfd8n0":
		return "2pufn98asdf9sfd8n0"
	case "2pufn98asdf9sfd8n1":
		return "2pufn98asdf9sfd8n1"
	case "2pufn98asdf9sfd8n2":
		return "2pufn98asdf9sfd8n2"
	case "2pufn98asdf9sfd8n3":
		return "2pufn98asdf9sfd8n3"
	case "2pufn98asdf9sfd8n4":
		return "2pufn98asdf9sfd8n4"
	case "2pufn98asdf9sfd8n5":
		return "2pufn98asdf9sfd8n5"
	case "2pufn98asdf9sfd8n6":
		return "2pufn98asdf9sfd8n6"
	case "2pufn98asdf9sfd8n7":
		return "2pufn98asdf9sfd8n7"
	case "2pufn98asdf9sfd8n8":
		return "2pufn98asdf9sfd8n8"
	case "2pufn98asdf9sfd8n9":
		return "2pufn98asdf9sfd8n9"
	default:
		return ""
	}
}

func OptionMap(s string) string {
	return map[string]string{
		"0pufn98asdf9sfd8n":  "0pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n":  "1pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n":  "2pufn98asdf9sfd8n",
		"3pufn98asdf9sfd8n":  "3pufn98asdf9sfd8n",
		"4pufn98asdf9sfd8n":  "4pufn98asdf9sfd8n",
		"5pufn98asdf9sfd8n":  "5pufn98asdf9sfd8n",
		"6pufn98asdf9sfd8n":  "6pufn98asdf9sfd8n",
		"7pufn98asdf9sfd8n":  "7pufn98asdf9sfd8n",
		"8pufn98asdf9sfd8n":  "8pufn98asdf9sfd8n",
		"9pufn98asdf9sfd8n":  "9pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n0": "10pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n1": "11pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n2": "12pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n3": "13pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n4": "14pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n5": "15pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n6": "16pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n7": "17pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n8": "18pufn98asdf9sfd8n",
		"1pufn98asdf9sfd8n9": "19pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n0": "20pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n1": "21pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n2": "22pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n3": "23pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n4": "24pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n5": "25pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n6": "26pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n7": "27pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n8": "28pufn98asdf9sfd8n",
		"2pufn98asdf9sfd8n9": "29pufn98asdf9sfd8n",
	}[s]
}

var globalmap = map[string]string{
	"0pufn98asdf9sfd8n":  "0pufn98asdf9sfd8n",
	"1pufn98asdf9sfd8n":  "1pufn98asdf9sfd8n",
	"2pufn98asdf9sfd8n":  "2pufn98asdf9sfd8n",
	"3pufn98asdf9sfd8n":  "3pufn98asdf9sfd8n",
	"4pufn98asdf9sfd8n":  "4pufn98asdf9sfd8n",
	"5pufn98asdf9sfd8n":  "5pufn98asdf9sfd8n",
	"6pufn98asdf9sfd8n":  "6pufn98asdf9sfd8n",
	"7pufn98asdf9sfd8n":  "7pufn98asdf9sfd8n",
	"8pufn98asdf9sfd8n":  "8pufn98asdf9sfd8n",
	"9pufn98asdf9sfd8n":  "9pufn98asdf9sfd8n",
	"10pufn98asdf9sfd8n": "10pufn98asdf9sfd8n",
	"11pufn98asdf9sfd8n": "11pufn98asdf9sfd8n",
	"12pufn98asdf9sfd8n": "12pufn98asdf9sfd8n",
	"13pufn98asdf9sfd8n": "13pufn98asdf9sfd8n",
	"14pufn98asdf9sfd8n": "14pufn98asdf9sfd8n",
	"15pufn98asdf9sfd8n": "15pufn98asdf9sfd8n",
	"16pufn98asdf9sfd8n": "16pufn98asdf9sfd8n",
	"17pufn98asdf9sfd8n": "17pufn98asdf9sfd8n",
	"18pufn98asdf9sfd8n": "18pufn98asdf9sfd8n",
	"19pufn98asdf9sfd8n": "19pufn98asdf9sfd8n",
	"20pufn98asdf9sfd8n": "20pufn98asdf9sfd8n",
	"21pufn98asdf9sfd8n": "21pufn98asdf9sfd8n",
	"22pufn98asdf9sfd8n": "22pufn98asdf9sfd8n",
	"23pufn98asdf9sfd8n": "23pufn98asdf9sfd8n",
	"24pufn98asdf9sfd8n": "24pufn98asdf9sfd8n",
	"25pufn98asdf9sfd8n": "25pufn98asdf9sfd8n",
	"26pufn98asdf9sfd8n": "26pufn98asdf9sfd8n",
	"27pufn98asdf9sfd8n": "27pufn98asdf9sfd8n",
	"28pufn98asdf9sfd8n": "28pufn98asdf9sfd8n",
	"29pufn98asdf9sfd8n": "29pufn98asdf9sfd8n",
}

func OptionMapGlobal(s string) string {
	return globalmap[s]
}
