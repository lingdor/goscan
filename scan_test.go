package goscan

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
	"testing"
)

func TestScan(t *testing.T) {

	type TestScanInfo struct {
		title string
		input string
		wants []string
	}

	testCases := []TestScanInfo{
		TestScanInfo{
			title: "empty empty",
			input: "\n",
			wants: []string{},
		},
		TestScanInfo{
			title: "test first space and two words",
			input: "  get key\n",
			wants: []string{"get", "key"},
		},
		TestScanInfo{
			title: "test first space and two words",
			input: "  get key\n",
			wants: []string{"get", "key"},
		},
		TestScanInfo{
			title: "test double quotation marks in parameters",
			input: "  set key \t\"val\"\n",
			wants: []string{"set", "key", "val"},
		},
		TestScanInfo{
			title: "test the multilines in double quotation marks ",
			input: `set xx "import text:
		\"all men are created equal\""
		`,
			wants: []string{"set", "xx", `import text:
		"all men are created equal"`},
		},
	}

	for _, testCase := range testCases {

		t.Run(testCase.title, func(t *testing.T) {

			reader := bufio.NewReader(bytes.NewReader([]byte(testCase.input)))
			line := NewFScanner(reader)
			words, err := line.ScanWords()
			if err != nil {
				t.Error(err)
				return
			}
			result := fmt.Sprintf("%+v", words)
			want := fmt.Sprintf("%+v", testCase.wants)
			if result != want {
				fmt.Fprintf(os.Stderr, "want:%s\nresult:%s\n", want, result)
				t.Fail()
			}
		})
	}

}
