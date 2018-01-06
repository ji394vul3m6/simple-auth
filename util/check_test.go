package util

import "testing"

type testCase struct {
	input  string
	output bool
}

var validTestCase = []testCase{
	testCase{"", false},
	testCase{"abc", true},
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
	testCase{"", false},
	testCase{"abcd", false},
	testCase{"0f7b4143-f0ae-11e7-bd86-0242ac120003", true},
	testCase{"a0f7b4143-f0ae-11e7-bd86-0242ac120003b", false},
	testCase{"a0f7b4143-f0ae-11e7-bd86-0242ac120003", false},
	testCase{"0f7b4143-f0ae-11e7-bd86-0242ac120003b", false},
}

func TestValidUUID(t *testing.T) {
	for _, test := range validUUIDCase {
		if IsValidUUID(test.input) != test.output {
			t.Errorf("IsValidString(%v) should be %v", test.input, test.output)
		}
	}
}
