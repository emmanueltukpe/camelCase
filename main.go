package main

import (
	"fmt"
	"regexp"
	"strings"
)

type arguments struct {
	pascalCase bool
	upperCase bool
}

func camelCase (word string, conditions arguments) string {
	pascalCase := conditions.pascalCase
	upperCase := conditions.upperCase
	reg := regexp.MustCompile("[^a-zA-Z0-9]+")
	whiteSpaceWord := reg.ReplaceAllString(word, " ")
	trimmed := strings.Trim(whiteSpaceWord, " ")
	var bArr []string
	var aArr []string
	var cArr []string
	//var aArr []string
	if upperCase {
		bArr = strings.Split(trimmed, " ")
		if len(bArr) > 1 {
			bArr[1] = strings.ToUpper(bArr[1])
		}
		aArr = strings.Split(strings.Join(bArr, " "), "")
	} else {
		aArr = strings.Split(trimmed, "")
	}
	
	for i := 0; i < len(aArr); i++ {
		element := aArr[i]
		if element == " " {
			cArr = append(aArr[:i], strings.ToUpper(aArr[i+1]))
			aArr = append(cArr, aArr[i+2:]...)
			} 
			
		}
		
		camel := strings.Join(aArr, "")
	upChar := strings.ToUpper(string(camel[0]))
	lowChar := strings.ToLower(string(camel[0]))
	if pascalCase {
		if lowChar == string(camel[0]){
			arr := strings.Split(camel, "")
			arr[0] = strings.ToUpper(arr[0])
			camelFix  := strings.Join(arr, "")
			return camelFix
		}
		return camel
	}else {
		if upChar == string(camel[0]){
			arr := strings.Split(camel, "")
			arr[0] = strings.ToLower(arr[0])
			camelFix  := strings.Join(arr, "")
			return camelFix
		}
		return camel
	}

}

func main() {
	fmt.Println(camelCase("foo-bar", arguments{pascalCase: false, upperCase: false}))
	fmt.Println(camelCase("foo_bar", arguments{pascalCase: false, upperCase: false}))
	fmt.Println(camelCase("Foo-bar", arguments{pascalCase: false, upperCase: false}))
	fmt.Println(camelCase("Foo-bar", arguments{pascalCase: true, upperCase: false}))
	fmt.Println(camelCase("--foo.bar", arguments{pascalCase: false, upperCase: false}))
	fmt.Println(camelCase("Foo-BAR", arguments{pascalCase: false, upperCase: true}))
	fmt.Println(camelCase("foobar", arguments{pascalCase: false, upperCase: false}))
	fmt.Println(camelCase("fooBAR", arguments{pascalCase: true, upperCase: true}))
	fmt.Println(camelCase("foo bar", arguments{pascalCase: false, upperCase: false}))
}
