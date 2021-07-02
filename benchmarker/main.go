package main

import (
	"go/parser"
	"go/token"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
)

func main(){
	dirFile, err := os.Open(".")
	if err != nil {
		panic(err)
	}
	defer dirFile.Close()
	info, err := dirFile.Stat()
	if err != nil {
		panic(err)
	}
	if !info.IsDir() {
		panic("Path is not a directory: .")
	}

	pkgs, err := parser.ParseDir(token.NewFileSet(), ".", func(info fs.FileInfo) bool {return filepath.Ext(info.Name()) == ".go"}, 0)
	for _, e := range pkgs{
		e.Scope.
	}
	if err != nil {
		panic(err)
	}
	testing.Benchmark()
}
