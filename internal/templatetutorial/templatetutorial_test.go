package templatetutorial

import (
	"bytes"
	"testing"
)

type TestCase struct {
	testName     string
	templateName string
	data         []interface{}
	expect       string
	success      bool
}

func TestLoader(t *testing.T) {
	var tt = []TestCase{
		{"FakeExample", "doesnt_exist.tpl", []interface{}{nil}, "", false},
		{"Example01_NoData", "ex01.tpl", []interface{}{nil}, "Hi <no value>\n\nYou are are welcome to this tutorial\n", true},
		{"Example01_EmptyString", "ex01.tpl", []interface{}{""}, "Hi \n\nYou are are welcome to this tutorial\n", true},
		{"Example01_String", "ex01.tpl", []interface{}{"Samus"}, "Hi Samus\n\nYou are are welcome to this tutorial\n", true},
		{"Example02_NoData", "ex02/*", []interface{}{nil, nil}, "Hi <no value>\n\nThis is ex02/temp01\nHi <no value>\n\nThis is ex02/temp02\n", true},
		{"Example02_SingleString", "ex02/*", []interface{}{"Scorpion"}, "Hi Scorpion\n\nThis is ex02/temp01\n", true},
		{"Example02_MultiString", "ex02/*", []interface{}{"Samus", "Metroid"}, "Hi Samus\n\nThis is ex02/temp01\nHi Metroid\n\nThis is ex02/temp02\n", true},
	}
	subject := New("")
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			var buf bytes.Buffer
			if err := subject.LoadTemplate(tc.templateName, tc.data); (err == nil) != tc.success {
				t.Fatalf("expected success == %v; got err: %v", tc.success, err)
			}
			if err := subject.ExecuteExample(tc.templateName, &buf); (err == nil) != tc.success {
				t.Fatalf("expected success == %v; got err: %v", tc.success, err)
			}
			if tc.expect != buf.String() {
				t.Fatalf("expected result == %q; got: %q", tc.expect, buf.String())
			}
		})
	}
}
