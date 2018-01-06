package util

import "testing"

type testCase struct {
	input  string
	output bool
}

var validTestCase = []testCase{
	testCase{input: "", output: false},
	testCase{input: "abc", output: true},
}

func TestValidString(t *testing.T) {
	if IsValidString(nil) {
		t.Error("IsValidString(nil) should be false")
	}
	for _, test := range validTestCase {
		if IsValidString(&test.input) != test.output {
			t.Errorf("IsValidString(%v) should be %v", test.input, test.output)
		}
	}
}

var validUUIDCase = []testCase{
	testCase{input: "", output: false},
	testCase{input: "abcd", output: false},
	testCase{input: "0f7b4143-f0ae-11e7-bd86-0242ac120003", output: true},
	testCase{input: "a0f7b4143-f0ae-11e7-bd86-0242ac120003b", output: false},
	testCase{input: "a0f7b4143-f0ae-11e7-bd86-0242ac120003", output: false},
	testCase{input: "0f7b4143-f0ae-11e7-bd86-0242ac120003b", output: false},
}

func TestValidUUID(t *testing.T) {
	for _, test := range validUUIDCase {
		if IsValidUUID(test.input) != test.output {
			t.Errorf("IsValidString(%v) should be %v", test.input, test.output)
		}
	}
}
