package aknenv

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

// type Config struct {
// 	Path string
// }

// func SetPath(path string) *Config {
// 	return &Config{
// 		Path: path,
// 	}
// }

// func getPath() string {
// 	c := &Config{}
// 	return c.Path
// }

// func (c *Config) GetEnv(str string) string {
// 	if len(getSystemVariable(str)) > 0 {
// 		return getSystemVariable(str)
// 	}
// 	filePath := filepath.Join(c.Path, ".env")
// 	env, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	arr := envToArray(string(env))
// 	for _, v := range arr {
// 		slice := sliceValue(v)
// 		if strings.Compare(slice[0], str) == 0 {
// 			// if err := os.Setenv(slice[0], slice[1]); err != nil {
// 			// 	panic(err.Error())
// 			// }
// 			return slice[1]
// 		}
// 	}
// 	return ""
// }

type EnvPath string

var ep EnvPath

func SetPath(path string) {
	var pt EnvPath
	pt = EnvPath(path)
	ep.Pathing(pt)
}
func (e *EnvPath) Pathing(str EnvPath) {
	*e = str
}

func GetPath() string {
	return string(ep)
}

// func (ep EnvPath) GetEnv(str string) string {
// 	if len(getSystemVariable(str)) > 0 {
// 		return getSystemVariable(str)
// 	}
// 	filePath := filepath.Join(string(ep), ".env")
// 	env, err := ioutil.ReadFile(filePath)
// 	if err != nil {
// 		panic(err.Error())
// 	}
// 	arr := envToArray(string(env))
// 	for _, v := range arr {
// 		slice := sliceValue(v)
// 		if strings.Compare(slice[0], str) == 0 {
// 			// if err := os.Setenv(slice[0], slice[1]); err != nil {
// 			// 	panic(err.Error())
// 			// }
// 			return slice[1]
// 		}
// 	}
// 	return ""
// }

func GetEnv(str string) string {
	if len(getSystemVariable(str)) > 0 {
		return getSystemVariable(str)
	}
	// wd, _ := os.Getwd()
	// filePath := filepath.Join(wd, ".env")
	filePath := filepath.Join(string(ep), ".env")
	env, err := ioutil.ReadFile(filePath)
	if err != nil {
		panic(err.Error())
	}
	arr := envToArray(string(env))
	for _, v := range arr {
		slice := sliceValue(v)
		if strings.Compare(slice[0], str) == 0 {
			// if err := os.Setenv(slice[0], slice[1]); err != nil {
			// 	panic(err.Error())
			// }
			return slice[1]
		}
	}
	return ""
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
	removeWhitespace(&arr[0])
	removeCarriage(&arr[1])
	return arr
}

func removeWhitespace(str *string) {
	*str = strings.Replace(*str, " ", "", -1)
}

func removeCarriage(str *string) {
	*str = strings.Replace(*str, "\r", "", -1)
}
