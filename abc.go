package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
	"strings"
)

func main() {
	ReadFile("in.txt", "out.txt")
	ReadFile2("in2.txt", "out2.txt")
}
func ReadFile(filename, outName string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	file2, err := os.OpenFile(outName, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
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
		_, err = file2.WriteString(s + "\n")
		if err != nil {
			log.Println(err)
		}

	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func ReadFile2(filename, outName string) {
	file, err := os.Open(filename)
	if err != nil {
		log.Fatal(err)
	}
	file2, err := os.OpenFile(outName, os.O_WRONLY, os.ModeAppend)
	if err != nil {
		log.Fatal(err)
	}
	defer file2.Close()
	defer file.Close()
	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		s := scanner.Text()

		if Replace4(&s) {

		}
		_, err = file2.WriteString(s)
		if err != nil {
			log.Println(err)
		}
	}
	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}
func Replace4(str *string) bool {
	reg := regexp.MustCompile(`[*]([a-zA-Z0-9.]+)\s?`)
	s := reg.FindString(*str)
	if s != "" {
		tpl := `
if %s!=nil{
%s
}`
		s = strings.TrimSpace(s)
		s = s[1:]
		*str = fmt.Sprintf(tpl, s, *str)
		return true
	}

	return false
}
func Replace(str *string) bool {
	reg := regexp.MustCompile(`[a-zA-Z0-9]+\s?\{`)
	s := reg.FindString(*str)
	//fmt.Println("---", *str, "---")
	if s != "" {
		*str = reg.ReplaceAllString(*str, `message ${0} `)
		return true
	}
	return false
}

func Replace1(str *string, int2 *int) bool {
	reg := regexp.MustCompile(`([a-zA-Z0-9]+)\s\[\]([a-zA-Z0-9]+).*`)
	if reg.FindString(*str) != "" {
		*str = reg.ReplaceAllString(*str, fmt.Sprintf("repeated ${2} ${1} =%d;", *int2))
		*int2++
		return true
	}
	return false
}
func Replace2(str *string, int2 *int) bool {
	reg := regexp.MustCompile(`([a-zA-Z0-9]+)\s([a-zA-Z0-9]+).*`)
	if reg.FindString(*str) != "" {
		*str = reg.ReplaceAllString(*str, fmt.Sprintf(" ${2} ${1} =%d;", *int2))
		*int2++
		return true
	}
	return false
}
