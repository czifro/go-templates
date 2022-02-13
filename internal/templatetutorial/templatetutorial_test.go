package templatetutorial

import (
	"bytes"
	"testing"
)

func TestLoader(t *testing.T) {
	var tt = []struct {
		testName     string
		templateFile TemplateFile
		data         interface{}
		expect       string
		success      bool
	}{
		{"Example01_NoData", ex01, nil, "Hi <no value>\n\nYou are are welcome to this tutorial", true},
		{"Example01_EmptyString", ex01, "", "Hi \n\nYou are are welcome to this tutorial", true},
		{"Example01_String", ex01, "Sam", "Hi Sam\n\nYou are are welcome to this tutorial", true},
	}
	for _, tc := range tt {
		t.Run(tc.testName, func(t *testing.T) {
			var buf bytes.Buffer
			err := loadAndRender(tc.templateFile, tc.data, &buf)
			if (err == nil) != tc.success {
				t.Fatalf("expected success == %v; got error: %v", tc.success, err)
			}
			if tc.expect != buf.String() {
				t.Fatalf("expected result == %q; got: %q", tc.expect, buf.String())
			}
		})
	}
}
