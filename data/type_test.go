package data

import (
	"bytes"
	"encoding/json"
	"testing"
)

func TestUserCopy(t *testing.T) {
	displayName := "testDisplay"
	email := "testEmail"
	enterprise := "testEnterprise"
	userType := 1
	password := "testPassword"
	status := 1
	user1 := User{
		ID:          "testID",
		DisplayName: &displayName,
		Email:       &email,
		Enterprise:  &enterprise,
		Type:        &userType,
		Password:    &password,
		Status:      &status,
	}

	user2 := User{}
	user2.CopyValue(user1)
	user1Byte, _ := json.Marshal(user1)
	user2Byte, _ := json.Marshal(user2)
	if !bytes.Equal(user1Byte, user2Byte) {
		t.Errorf("Different after copy: \nuser1:[%s]\nuser2:[%s]", user1Byte, user2Byte)
	}
}

func TestUserActive(t *testing.T) {
	var status int
	user := User{
		Status: &status,
	}

	status = 0
	if user.IsActive() {
		t.Errorf("User is active only when Status is 1")
	}
	status = 1
	if !user.IsActive() {
		t.Errorf("User should be active when Status is 1")
	}
}

func TestUserJWTToken(t *testing.T) {
	displayName := "testDisplay"
	email := "testEmail"
	enterprise := "testEnterprise"
	userType := 1
	password := "testPassword"
	status := 1
	user1 := User{
		ID:          "testID",
		DisplayName: &displayName,
		Email:       &email,
		Enterprise:  &enterprise,
		Type:        &userType,
		Password:    &password,
		Status:      &status,
	}
	user2 := User{}

	token, _ := user1.GenerateToken()
	user2.SetValueWithToken(token)
	user1Byte, _ := json.Marshal(user1)
	user2Byte, _ := json.Marshal(user2)
	if !bytes.Equal(user1Byte, user2Byte) {
		t.Errorf("Different after use token: \nuser1:[%s]\nuser2:[%s]", user1Byte, user2Byte)
	}
}

func TestUserValid(t *testing.T) {
	displayName := "testDisplay"
	email := "testEmail"
	enterprise := "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
	userType := 1
	password := "00001111222233334444555566667777"
	status := 1
	user := User{
		ID:          "testID",
		DisplayName: &displayName,
		Email:       &email,
		Enterprise:  &enterprise,
		Type:        &userType,
		Password:    &password,
		Status:      &status,
	}

	if !user.IsValid() {
		t.Errorf("User should be valid in orig case")
	}

	email = ""
	if user.IsValid() {
		t.Errorf("User should be invalid when email is empty")
	}

	email = "testEmail"
	enterprise = "abc"
	if user.IsValid() {
		t.Errorf("User should be invalid when enterprise is not uuid")
	}

	enterprise = "aaaaaaaa-aaaa-aaaa-aaaa-aaaaaaaaaaaa"
	password = "123"
	if user.IsValid() {
		t.Errorf("User should be invalid when password is not md5 hash")
	}
}
