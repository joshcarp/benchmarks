package map_vs_switch

import "testing"

func BenchmarkOptionSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionSwitch("1")
	}
}

func BenchmarkOptionMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionMap("1")
	}
}

func OptionMap(s string) string {
	return map[string]string{"1": "1", "2": "2", "3": "3", "4": "4", "5": "5", "6": "6", "7": "7"}[s]
}

func OptionSwitch(s string) string {
	switch s {
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
	default:
		return ""
	}
}
