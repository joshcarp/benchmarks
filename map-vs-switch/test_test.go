package map_vs_switch

import (
	"testing"
)

func BenchmarkOptionSwitch(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionSwitch("key12pufn98asdf9sfd8n")
	}
}

func BenchmarkOptionMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionMap("key12pufn98asdf9sfd8n")
	}
}

func BenchmarkOptionGlobalMap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		OptionMapGlobal("key12pufn98asdf9sfd8n")
	}
}

func OptionSwitch(s string) string {
	switch s {
	case "key0pufn98asdf9sfd8n":
		return "val0pufn98asdf9sfd8n"
	case "key1pufn98asdf9sfd8n":
		return "val1pufn98asdf9sfd8n"
	case "key2pufn98asdf9sfd8n":
		return "val2pufn98asdf9sfd8n"
	case "key3pufn98asdf9sfd8n":
		return "val3pufn98asdf9sfd8n"
	case "key4pufn98asdf9sfd8n":
		return "val4pufn98asdf9sfd8n"
	case "key5pufn98asdf9sfd8n":
		return "val5pufn98asdf9sfd8n"
	case "key6pufn98asdf9sfd8n":
		return "val6pufn98asdf9sfd8n"
	case "key7pufn98asdf9sfd8n":
		return "val7pufn98asdf9sfd8n"
	case "key8pufn98asdf9sfd8n":
		return "val8pufn98asdf9sfd8n"
	case "key9pufn98asdf9sfd8n":
		return "val9pufn98asdf9sfd8n"
	case "key1pufn98asdf9sfd8n0":
		return "val1pufn98asdf9sfd8n0"
	case "key1pufn98asdf9sfd8n1":
		return "val1pufn98asdf9sfd8n1"
	case "key1pufn98asdf9sfd8n2":
		return "val1pufn98asdf9sfd8n2"
	case "key1pufn98asdf9sfd8n3":
		return "val1pufn98asdf9sfd8n3"
	case "key1pufn98asdf9sfd8n4":
		return "val1pufn98asdf9sfd8n4"
	case "key1pufn98asdf9sfd8n5":
		return "val1pufn98asdf9sfd8n5"
	case "key1pufn98asdf9sfd8n6":
		return "val1pufn98asdf9sfd8n6"
	case "key1pufn98asdf9sfd8n7":
		return "val1pufn98asdf9sfd8n7"
	case "key1pufn98asdf9sfd8n8":
		return "val1pufn98asdf9sfd8n8"
	case "key1pufn98asdf9sfd8n9":
		return "val1pufn98asdf9sfd8n9"
	case "key2pufn98asdf9sfd8n0":
		return "val2pufn98asdf9sfd8n0"
	case "key2pufn98asdf9sfd8n1":
		return "val2pufn98asdf9sfd8n1"
	case "key2pufn98asdf9sfd8n2":
		return "val2pufn98asdf9sfd8n2"
	case "key2pufn98asdf9sfd8n3":
		return "val2pufn98asdf9sfd8n3"
	case "key2pufn98asdf9sfd8n4":
		return "val2pufn98asdf9sfd8n4"
	case "key2pufn98asdf9sfd8n5":
		return "val2pufn98asdf9sfd8n5"
	case "key2pufn98asdf9sfd8n6":
		return "val2pufn98asdf9sfd8n6"
	case "key2pufn98asdf9sfd8n7":
		return "val2pufn98asdf9sfd8n7"
	case "key2pufn98asdf9sfd8n8":
		return "val2pufn98asdf9sfd8n8"
	case "key2pufn98asdf9sfd8n9":
		return "val2pufn98asdf9sfd8n9"
	default:
		return ""
	}
}

func OptionMap(s string) string {
	return map[string]string{
		"key0pufn98asdf9sfd8n":  "val0pufn98asdf9sfd8n",
		"key1pufn98asdf9sfd8n":  "val1pufn98asdf9sfd8n",
		"key2pufn98asdf9sfd8n":  "val2pufn98asdf9sfd8n",
		"key3pufn98asdf9sfd8n":  "val3pufn98asdf9sfd8n",
		"key4pufn98asdf9sfd8n":  "val4pufn98asdf9sfd8n",
		"key5pufn98asdf9sfd8n":  "val5pufn98asdf9sfd8n",
		"key6pufn98asdf9sfd8n":  "val6pufn98asdf9sfd8n",
		"key7pufn98asdf9sfd8n":  "val7pufn98asdf9sfd8n",
		"key8pufn98asdf9sfd8n":  "val8pufn98asdf9sfd8n",
		"key9pufn98asdf9sfd8n":  "val9pufn98asdf9sfd8n",
		"key10pufn98asdf9sfd8n": "val10pufn98asdf9sfd8n",
		"key11pufn98asdf9sfd8n": "val11pufn98asdf9sfd8n",
		"key12pufn98asdf9sfd8n": "val12pufn98asdf9sfd8n",
		"key13pufn98asdf9sfd8n": "val13pufn98asdf9sfd8n",
		"key14pufn98asdf9sfd8n": "val14pufn98asdf9sfd8n",
		"key15pufn98asdf9sfd8n": "val15pufn98asdf9sfd8n",
		"key16pufn98asdf9sfd8n": "val16pufn98asdf9sfd8n",
		"key17pufn98asdf9sfd8n": "val17pufn98asdf9sfd8n",
		"key18pufn98asdf9sfd8n": "val18pufn98asdf9sfd8n",
		"key19pufn98asdf9sfd8n": "val19pufn98asdf9sfd8n",
		"key20pufn98asdf9sfd8n": "val20pufn98asdf9sfd8n",
		"key21pufn98asdf9sfd8n": "val21pufn98asdf9sfd8n",
		"key22pufn98asdf9sfd8n": "val22pufn98asdf9sfd8n",
		"key23pufn98asdf9sfd8n": "val23pufn98asdf9sfd8n",
		"key24pufn98asdf9sfd8n": "val24pufn98asdf9sfd8n",
		"key25pufn98asdf9sfd8n": "val25pufn98asdf9sfd8n",
		"key26pufn98asdf9sfd8n": "val26pufn98asdf9sfd8n",
		"key27pufn98asdf9sfd8n": "val27pufn98asdf9sfd8n",
		"key28pufn98asdf9sfd8n": "val28pufn98asdf9sfd8n",
		"key29pufn98asdf9sfd8n": "val29pufn98asdf9sfd8n",
	}[s]
}

var globalmap = map[string]string{
	"key0pufn98asdf9sfd8n":  "val0pufn98asdf9sfd8n",
	"key1pufn98asdf9sfd8n":  "val1pufn98asdf9sfd8n",
	"key2pufn98asdf9sfd8n":  "val2pufn98asdf9sfd8n",
	"key3pufn98asdf9sfd8n":  "val3pufn98asdf9sfd8n",
	"key4pufn98asdf9sfd8n":  "val4pufn98asdf9sfd8n",
	"key5pufn98asdf9sfd8n":  "val5pufn98asdf9sfd8n",
	"key6pufn98asdf9sfd8n":  "val6pufn98asdf9sfd8n",
	"key7pufn98asdf9sfd8n":  "val7pufn98asdf9sfd8n",
	"key8pufn98asdf9sfd8n":  "val8pufn98asdf9sfd8n",
	"key9pufn98asdf9sfd8n":  "val9pufn98asdf9sfd8n",
	"key10pufn98asdf9sfd8n": "val10pufn98asdf9sfd8n",
	"key11pufn98asdf9sfd8n": "val11pufn98asdf9sfd8n",
	"key12pufn98asdf9sfd8n": "val12pufn98asdf9sfd8n",
	"key13pufn98asdf9sfd8n": "val13pufn98asdf9sfd8n",
	"key14pufn98asdf9sfd8n": "val14pufn98asdf9sfd8n",
	"key15pufn98asdf9sfd8n": "val15pufn98asdf9sfd8n",
	"key16pufn98asdf9sfd8n": "val16pufn98asdf9sfd8n",
	"key17pufn98asdf9sfd8n": "val17pufn98asdf9sfd8n",
	"key18pufn98asdf9sfd8n": "val18pufn98asdf9sfd8n",
	"key19pufn98asdf9sfd8n": "val19pufn98asdf9sfd8n",
	"key20pufn98asdf9sfd8n": "val20pufn98asdf9sfd8n",
	"key21pufn98asdf9sfd8n": "val21pufn98asdf9sfd8n",
	"key22pufn98asdf9sfd8n": "val22pufn98asdf9sfd8n",
	"key23pufn98asdf9sfd8n": "val23pufn98asdf9sfd8n",
	"key24pufn98asdf9sfd8n": "val24pufn98asdf9sfd8n",
	"key25pufn98asdf9sfd8n": "val25pufn98asdf9sfd8n",
	"key26pufn98asdf9sfd8n": "val26pufn98asdf9sfd8n",
	"key27pufn98asdf9sfd8n": "val27pufn98asdf9sfd8n",
	"key28pufn98asdf9sfd8n": "val28pufn98asdf9sfd8n",
	"key29pufn98asdf9sfd8n": "val29pufn98asdf9sfd8n",
}

func OptionMapGlobal(s string) string {
	return globalmap[s]
}
