package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

func main() {
	ReadFile("in.txt")
}
func ReadFile(filename string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		s := scanner.Text()

		if Replace(&s) {
			i = 1
		} else if Replace1(&s, &i) {

		} else if Replace2(&s, &i) {

		}
		fmt.Println(s)
		//break
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func Replace(str *string) bool {
	reg := regexp.MustCompile(`[a-zA-Z]+\s\{`)
	s := reg.FindString(*str)
	if s != "" {
		*str = reg.ReplaceAllString(*str, `message ${0} `)
		return true
	}
	return false
}

func Replace1(str *string, int2 *int) bool {
	reg := regexp.MustCompile(`([a-zA-Z]+)\s\[\]([a-zA-Z]+).*`)
	if reg.FindString(*str) != "" {
		*str = reg.ReplaceAllString(*str, fmt.Sprintf("repeated ${2} ${1} =%d;", *int2))
		*int2++
		return true
	}
	return false
}
func Replace2(str *string, int2 *int) bool {
	reg := regexp.MustCompile(`([a-zA-Z]+)\s([a-zA-Z0-9]+).*`)
	if reg.FindString(*str) != "" {
		*str = reg.ReplaceAllString(*str, fmt.Sprintf(" ${2} ${1} =%d;", *int2))
		*int2++
		return true
	}
	return false
}
