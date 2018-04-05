package aknenv

import (
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Env struct {
	Path string
}

func (e *Env) Set(str string) *Env {
	e.Path = str
	return e
}

func (e *Env) GetEnv(str string) string {
	var val string
	if len(getSystemVariable(str)) > 0 {
		return getSystemVariable(str)
	}
	filePath := filepath.Join(e.Path, ".env")
	env, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(fmt.Sprintf("Variable %s is not found anywhere", str))
		return val
	}
	arr := envToArray(string(env))
	for _, v := range arr {
		slice := sliceValue(v)
		if len(slice) == 2 {
			if strings.Compare(slice[0], str) == 0 {
				val = slice[1]
			}
		}
	}
	return val
}

func GetEnv(str string, path ...string) string {
	var val string
	if len(getSystemVariable(str)) > 0 {
		return getSystemVariable(str)
	}
	p := strings.Join(path, "")
	var filePath string
	if len(p) > 0 {
		filePath = filepath.Join(p, ".env")
	} else {
		filePath = filepath.Join("./", ".env")
	}
	env, err := ioutil.ReadFile(filePath)
	if err != nil {
		log.Println(fmt.Sprintf("Variable %s is not found anywhere", str))
		return val
	}
	arr := envToArray(string(env))
	for _, v := range arr {
		slice := sliceValue(v)
		if len(slice) == 2 {
			if strings.Compare(slice[0], str) == 0 {
				val = slice[1]
			}
		}
	}
	return val
}

func envToArray(env string) []string {
	arr := strings.Split(env, "\n")
	return arr
}

func getSystemVariable(str string) string {
	val := os.Getenv(str)
	return val
}

func sliceValue(str string) []string {
	arr := make([]string, 2)
	arr = strings.SplitN(str, `=`, 2)
	if len(arr) == 2 {
		// fmt.Println(arr)
		removeWhitespace(&arr[0])
		removeCarriage(&arr[1])
	}
	return arr
}

func removeWhitespace(str *string) {
	*str = strings.Replace(*str, " ", "", -1)
}

func removeCarriage(str *string) {
	*str = strings.Replace(*str, "\r", "", -1)
}
