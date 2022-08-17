package tasks

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var data = make(map[string]int)
var code = make([]string, 0)
var pos int = 0

func parseBlock(blockData map[string]int) {
	for pos < len(code) {
		line := code[pos]

		if line == "}" {
			pos++
			return
		}

		if line == "{" {
			pos++
			parseBlock(blockData)
			continue
		}

		args := strings.Split(line, "=")
		i, err := strconv.Atoi(args[1])
		if err == nil {
			blockData[args[0]] = i
		} else {
			// check both blockData and data
			// blockData overrides value in data if blockData has value at given key
			// if not, value is taken from data

			_, ok := blockData[args[1]] // check if blockData has value of variable which value
			if ok {                     // we try to assign
				blockData[args[0]] = blockData[args[1]] // blockData overrides values of data for a time
				fmt.Println(blockData[args[0]])         // block exists
			} else {
				blockData[args[0]] = data[args[1]] // WARNING: Returns zero if doesn't exist ???
				fmt.Println(data[args[1]])
			}
		}

		pos++
	}
}

func T4() {
	// read line
	// if not {} split by "="
	// add value to map[string]interface{} (is this type good ???)

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		if scanner.Text() == "" {
			break
		}
		code = append(code, scanner.Text())
	}

	for pos < len(code) {
		if code[pos] == "{" { // evaluating block
			pos++ // so that we won't try splitting "{"
			blockData := make(map[string]int)
			parseBlock(blockData) // should impicitly set pos to correct one
			continue
		}

		args := strings.Split(code[pos], "=")
		i, err := strconv.Atoi(args[1])
		if err == nil {
			data[args[0]] = i
		} else {
			data[args[0]] = data[args[1]] // WARNING: Returns zero if doesn't exist ???
			fmt.Println(data[args[1]])
		}

		pos++
	}
}
