package util

import (
	"testing"
)

func TestJWTfunction(t *testing.T) {
	test := map[string]string{
		"a": "123",
		"b": "456",
	}
	ret, _ := GetJWTTokenWithCustomInfo(test)
	resolve, _ := ResolveJWTToken(ret)
	resolveObj := resolve.(map[string]interface{})
	if resolveObj["a"] != "123" || resolveObj["b"] != "456" {
		t.Errorf("Incorrect resolve jwt string, except: [%#v], get: [%#v]", test, resolve)
	}
}
