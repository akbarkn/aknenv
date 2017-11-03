package aknenv

import (
	"strings"
	"testing"
)

func TestVariableValue(t *testing.T) {
	if strings.Compare(GetEnv("DB_HOST"), "localhost") != 0 {
		t.Errorf("VARIABLE VALUE 1 IS INCORRECT, SHOULD BE %s", GetEnv("DB_HOST"))
		t.Fail()
		return
	} else {
		t.Logf("VARIABLE VALUE 1 IS CORRECT")
	}
	if strings.Compare(GetEnv("DB_PORT"), "3306") != 0 {
		t.Errorf("VARIABLE VALUE 2 IS INCORRECT, SHOULD BE %s", GetEnv("DB_PORT"))
		t.Fail()
		return
	} else {
		t.Logf("VARIABLE VALUE 2 IS CORRECT")
	}
	t.Logf("SUCCESS")
}

func TestVariableValue2(t *testing.T) {
	if strings.Compare(GetEnv("DB_HOST", "./test"), "192.168.1.2") != 0 {
		t.Errorf("VARIABLE VALUE 1 IS INCORRECT, SHOULD BE %s", GetEnv("DB_HOST", "./test"))
		t.Fail()
		return
	} else {
		t.Logf("VARIABLE VALUE 1 IS CORRECT")
	}
	if strings.Compare(GetEnv("DB_PORT", "./test"), "3307") != 0 {
		t.Errorf("VARIABLE VALUE 2 IS INCORRECT, SHOULD BE %s", GetEnv("DB_PORT", "./test"))
		t.Fail()
		return
	} else {
		t.Logf("VARIABLE VALUE 2 IS CORRECT")
	}
	t.Logf("SUCCESS")
}
