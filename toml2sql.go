package main

import (
	"fmt"
	"os"

	"github.com/BurntSushi/toml"
)

type Query struct {
	Measurement string   `toml:"measurement"`
	Sql         string   `toml:"sql"`
	Tags        []string `toml:"tags"`
	Fields      []string `toml:"fields"`
	Pivot       bool     `toml:"pivot"`
	PivotKey    string   `toml:"pivot_key"`
	Enable      bool     `toml:"enable"`
}

type QueryMap struct {
	Title   string  `toml:Title"`
	Queries []Query `toml:"sql_metric"`
}

func main() {
	var queryMap QueryMap
	infileName := "query.toml"
	outfileName := "query.sql"

	argCount := len(os.Args) - 1
	args := os.Args

	if argCount > 0 {
		infileName = args[1]
	} 
	if argCount > 1 {
		outfileName = args[2]
	}

	_, err := toml.DecodeFile(infileName, &queryMap)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	outfile, err := os.OpenFile(outfileName, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}
	defer outfile.Close()

	for _, queries := range queryMap.Queries {
		sql := queries.Sql + ";\n\n"
		_, err = outfile.WriteString(sql)
	}
}

