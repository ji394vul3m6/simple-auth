package util

import "testing"

type stringValidTestCase struct {
	input  string
	output bool
}

var validTestCase = []stringValidTestCase{
	stringValidTestCase{"", false},
	stringValidTestCase{"abc", true},
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

var validUUIDCase = []stringValidTestCase{
	stringValidTestCase{"", false},
	stringValidTestCase{"abcd", false},
	stringValidTestCase{"0f7b4143-f0ae-11e7-bd86-0242ac120003", true},
	stringValidTestCase{"a0f7b4143-f0ae-11e7-bd86-0242ac120003b", false},
	stringValidTestCase{"a0f7b4143-f0ae-11e7-bd86-0242ac120003", false},
	stringValidTestCase{"0f7b4143-f0ae-11e7-bd86-0242ac120003b", false},
}

func TestValidUUID(t *testing.T) {
	for _, test := range validUUIDCase {
		if IsValidUUID(test.input) != test.output {
			t.Errorf("IsValidString(%v) should be %v", test.input, test.output)
		}
	}
}

var validMD5Case = []stringValidTestCase{
	stringValidTestCase{"", false},
	stringValidTestCase{"abcd", false},
	stringValidTestCase{"202cb962ac59075b964b07152d234b70", true},
	stringValidTestCase{"202cb962ac59075b964b07152d234b7!", false},
	stringValidTestCase{"202cb962ac59075b964b07152d234b7-", false},
	stringValidTestCase{"202cb962ac59075b964b07152d234b7", false},
}

//
func TestValidMD5(t *testing.T) {
	for _, test := range validMD5Case {
		if IsValidMD5(test.input) != test.output {
			t.Errorf("IsValidString(%v) should be %v", test.input, test.output)
		}
	}
}

type checkContainTestCase struct {
	check     interface{}
	container []interface{}
	output    bool
}

var inSliceCase = []checkContainTestCase{
	checkContainTestCase{
		1, []interface{}{1, 2, 3}, true,
	},
	checkContainTestCase{
		1, []interface{}{"1", "a", "b"}, false,
	},
	checkContainTestCase{
		"abc", []interface{}{1, "abc", "b"}, true,
	},
}

func TestInSlice(t *testing.T) {
	for _, test := range inSliceCase {
		if IsInSlice(test.check, test.container) != test.output {
			t.Errorf("IsInSlice(%v, %#v) should be %v", test.check, test.container, test.output)
		}
	}
}
