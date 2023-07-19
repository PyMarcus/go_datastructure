package main

import (
	"bufio"
	"fmt"
	"strings"
	"log"
	"os"
)

var counter map[string]int = make(map[string]int)

// i saw this in the linkedin, so i want implement this algorithm
func readFile(fileName string) {
	file, err := os.Open(fileName)

	if err != nil {
		log.Fatalln(err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		row := scanner.Text()
		countWords(row)
	}
	show(&counter)
}

func countWords(row string){
	words := strings.Split(row, " ")

	if value, exist := counter[words[0]]; exist == false{
		counter[words[0]] = 1
	}else{
		counter[words[0]] = value + 1
	}

	if value, exist := counter[words[1]]; exist == false{
		counter[words[1]] = 1
	}else{
		counter[words[1]] = value + 1
	}
}

func show(words *map[string]int){
	for k, v := range *words{
		log.Println(k, v)
	}
}

func main(){
	curDir, _ := os.Getwd()
	fileName := os.Args
	readFile(fmt.Sprintf("%s\\count_words\\%s", curDir, fileName[1]))
}