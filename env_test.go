package aknenv

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"testing"
)

func TestVariableValue(t *testing.T) {
	SetPath("./")
	if strings.Compare(GetEnv("DB_HOST"), "localhost") != 0 {
		t.Errorf("VARIABLE VALUE 1 IS INCORRECT")
		t.Fail()
		return
	} else {
		t.Logf("VARIABLE VALUE 1 IS CORRECT")
	}
	if strings.Compare(GetEnv("DB_PORT"), "3306") != 0 {
		t.Errorf("VARIABLE VALUE 2 IS INCORRECT")
		t.Fail()
		return
	} else {
		t.Logf("VARIABLE VALUE 2 IS CORRECT")
	}
	t.Logf("SUCCESS")
}

func TestPath(t *testing.T) {
	filePath := filepath.Join(string(SetPath("./")), ".env")
	envFile, err := ioutil.ReadFile(filePath)
	if err != nil {
		t.Errorf(err.Error())
		t.Fail()
		return
	}
	t.Logf("env file content:\n%s", envFile)
	t.Logf("Success")
}

func TestLoad(t *testing.T) {
	SetPath("./")
	fmt.Println("DB_HOST", GetEnv("DB_HOST"))
	t.Logf("try")
}
