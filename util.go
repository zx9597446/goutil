//util functions in golang
//drop util.go to your source directory
//then modify package name
package util

import (
	"encoding/json"
	"log"
	"strconv"
)

func panicIf(err error) {
	if err != nil {
		panic(err)
	}
}

func logIf(err error) {
	if err != nil {
		log.Println(err)
	}
}

func panicfIf(err error, format string, v ...interface{}) {
	if err != nil {
		log.Panicf(format, v...)
	}
}

func atoi(a string) int {
	i, err := strconv.Atoi(a)
	panicIf(err)
	return i
}

func itoa(i int) string {
	return strconv.Itoa(i)
}

func mustJSON(v interface{}) []byte {
	data, err := json.Marshal(v)
	panicIf(err)
	return data
}

func mustFromJSON(data []byte, v interface{}) {
	err := json.Unmarshal(data, v)
	panicIf(err)
}
