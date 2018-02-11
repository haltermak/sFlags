package sFlags

import (
	"fmt"
	"strconv"
	"strings"
)

func CreateFlags(input string) (string, map[string]string, error) {
	temp := strings.Fields(input)
	var err error
	var output map[string]string
	output = make(map[string]string)
	for i := 1; i < len(temp)-1; {
		if strings.HasPrefix(temp[i], "-") {
			if i+1 == len(temp) || strings.HasPrefix(temp[i+1], "-") {
				output[temp[i]] = "true"
				i++
			} else {
				output[temp[i]] = temp[i+1]
				i += 2
			}
		}
	}
	fmt.Println(output)
	return temp[0], output, err
}

func FlagToInt(fM map[string]string, key string) (int, error) {
	var err error
	temp, err := strconv.Atoi(fM[key])
	return temp, err
}

func FlagToString(fM map[string]string, key string) (string, error) {
	var err error
	return fM[key], err
}

func FlagToBool(fM map[string]string, key string) (bool, error) {
	var err error
	var temp bool
	temp, err = strconv.ParseBool(fM[key])
	return temp, err
}
