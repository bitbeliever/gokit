package main

import (
	"fmt"
	"github.com/iancoleman/strcase"
	"golang.design/x/clipboard"
	"log"
	"os"
	"strings"
)

var typeMapper = map[string]string{
	"String":  "string",
	"Boolean": "bool",
}

func main() {
	str := string(clipboard.Read(clipboard.FmtText))
	fmt.Println(str, "\n\n")

	//var result bytes.Buffer
	var res string
	lines := wash(strings.Split(str, "\n"))
	for _, line := range lines {
		if line == "" {
			continue
		}

		words := strings.Fields(line)
		if len(words) < 3 {
			fmt.Println("ERROR: len < 3", words)
			return
		}
		res += fmt.Sprintf("%s  %s %s//", strings.Title(words[0]), typeMapper[words[1]], jsonTag(words[0]))
		for _, w := range words[2:] {
			res += w
		}
		res += "\n"
	}

	res += "\n"
	fmt.Println(res)
	f, err := os.OpenFile("fields.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal(err)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	_, err = f.Write([]byte(res))
	if err != nil {
		fmt.Println(err)
	}
}

func jsonTag(s string) string {
	return fmt.Sprintf("`json:\"%s\"`", strcase.ToLowerCamel(s))
}

func wash(lines []string) []string {
	var result []string
	for _, line := range lines {
		if line == "" {
			continue
		}

		// 不包含类型信息, 此行添加到上一行
		var flag bool
		for typ := range typeMapper {
			if strings.Contains(line, typ) {
				flag = true
				break
			}
		}
		if !flag {
			result[len(result)-1] += "  " + strings.TrimSuffix(line, "\n")
		} else {
			result = append(result, strings.TrimSuffix(line, "\n"))
		}
	}

	return result
}
