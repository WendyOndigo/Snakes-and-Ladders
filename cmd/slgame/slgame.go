package main

import (
	slgame "github.com/b6luong/Snakes-and-Ladders/components"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

func check(e error) {
	if e != nil && e != io.EOF {
		panic(e)
	}
}

func runGameUsingStdIn() {
	reader := bufio.NewReader(os.Stdin)
	var input string
	config := ""

	input, err := reader.ReadString('\n')
	check(err)
	for err != io.EOF {
		config += input
		input, err = reader.ReadString('\n')
		check(err)
	}
	game := slgame.ReadFrom(config)
	fmt.Println(slgame.Print(game))
}

func runGameUsingJsonIO(infile string, outfile string) {
	jsonRawData, err := ioutil.ReadFile(infile)
	check(err)

	var jsonData map[string]interface{}
	parseError := json.Unmarshal(jsonRawData, &jsonData)
	check(parseError)

	author := jsonData["author"].(string)

	jsonOutData := make(map[string]interface{})
	jsonOutData["requestedBy"] = author

	resultSet := []map[string]string{}

	for _, i := range jsonData["configurations"].([]interface{}) {
		obj := i.(map[string]interface{})
		title := obj["title"].(string)
		config := obj["commands"].(string)
		result := make(map[string]string)
		result["title"] = title
		result["result"] = slgame.Print(slgame.ReadFrom(config))
		resultSet = append(resultSet, result)
		//fmt.Println("title:", title)
		//fmt.Println("config:", config)
		//fmt.Println(slgame.Print(slgame.ReadFrom(config)))
	}

	jsonOutData["resultSet"] = resultSet
	jsonOutRaw, marshErr := json.Marshal(jsonOutData)
	check(marshErr)
	writeErr := ioutil.WriteFile(outfile, jsonOutRaw, 0644)
	check(writeErr)
}

func runGameUsingJsonIn(infile string) {
	jsonRawData, err := ioutil.ReadFile(infile)
	check(err)

	var jsonData map[string]interface{}
	parseError := json.Unmarshal(jsonRawData, &jsonData)
	check(parseError)

	for _, i := range jsonData["configurations"].([]interface{}) {
		obj := i.(map[string]interface{})
		//title := obj["title"].(string)
		config := obj["commands"].(string)
		//fmt.Println("title:", title)
		//fmt.Println("config:", config)
		fmt.Println(slgame.Print(slgame.ReadFrom(config)))
	}
}

func main() {

	jsonIn := false
	jsonInFile := ""
	jsonOut := false
	jsonOutFile := ""
	pwd, pathErr := os.Getwd()
	check(pathErr)
	/*
		fmt.Println(os.Args)
		fmt.Println(pwd + "/")
	*/
	for _, arg := range os.Args {
		if strings.Contains(arg, "-json-in=") {
			jsonIn = true
			jsonInFile = pwd + "/" + arg[9:]
		} else if strings.Contains(arg, "-json-out=") {
			jsonOut = true
			jsonOutFile = pwd + "/" + arg[10:]
		}
	}

	/*fmt.Println("jsonIn:", jsonIn)
	fmt.Println("jsonInFile:", jsonInFile)
	fmt.Println("jsonOut:", jsonOut)
	fmt.Println("jsonOutFile:", jsonOutFile)
	*/
	if jsonIn && jsonOut {
		runGameUsingJsonIO(jsonInFile, jsonOutFile)
	} else if !jsonIn && jsonOut {
		//runGameUsingJsonOut(jsonOutFile)
	} else if jsonIn && !jsonOut {
		runGameUsingJsonIn(jsonInFile)
	} else {
		runGameUsingStdIn()
	}
}
