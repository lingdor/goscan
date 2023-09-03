package goscan

import (
	"goscan/strs"
	"io"
	"os"
)

type ScanedLine interface {
	Scan() (str string, end bool, err error)
	ScanWords() ([]string, error)
	//todo ScanVars([]interface{}) error
}

type ReaderScanedLine struct {
	ScanedLine
	reader io.Reader
}

func (c ReaderScanedLine) Scan() (str string, end bool, err error) {

	var bs []byte = make([]byte, 1)
	var result []byte = make([]byte, 0, 10)
	var n int
	startChar := byte('\n')
	end = false
	for {
		if n, err = c.reader.Read(bs); err != nil {
			str = ""
			return
		} else if n != len(bs) {
			break
			//end
		}
		char := bs[0]
		if len(result) == 0 && (char == '"' || char == '\'') {
			startChar = char
			continue
		} else if len(result) == 0 && char == '\n' {
			// blank input
			end = true
			break
		}

		if len(result) == 0 && strs.IsSpace(char) {
			continue // ignore the start space
		}
		if startChar == '\n' {
			if char == startChar {
				// 结束
				end = true
				break
			} else if strs.IsEmptyChar(char) {
				break
			}
		} else {
			if char == '\\' {
				var next []byte = make([]byte, 1)
				if _, err = c.reader.Read(next); err != nil {
					return
				}
				char = strs.EscapeWord(next[0])
			} else if startChar == char {
				if len(result) == 0 || result[len(result)-1] != '\\' {
					break
				}
			}
		}
		result = append(result, char)
	}
	str = string(result)
	err = nil
	return
}
func (c ReaderScanedLine) ScanWords() (words []string, err error) {

	words = make([]string, 0, 2)
	var end bool
	var str string
	for {
		if str, end, err = c.Scan(); err != nil {
			return
		}
		if str != "" {
			words = append(words, str)
		}
		if end {
			return
		}
	}
}

// NewScanStd to read stdin line content for scan operation
func NewScanStd() (ScanedLine, error) {
	return NewFScan(os.Stdin)
}

// NewFScan to read Reader line content for scan operation
func NewFScan(reader io.Reader) (line ScanedLine, err error) {
	line = &ReaderScanedLine{
		reader: reader,
	}
	return
}
