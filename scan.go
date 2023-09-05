package goscan

import (
	"bufio"
	"github.com/lingdor/goscan/utils"
	"io"
	"os"
	"strings"
)

type Scanner interface {
	Scan() (str string, end bool, err error)
	ScanWords() ([]string, error)
	CheckToEnd() (check bool, err error)
	//todo ScanVars([]interface{}) error
}

type ReaderScanner struct {
	Scanner
	reader io.Reader
}

func (c ReaderScanner) Scan() (str string, end bool, err error) {

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

		if len(result) == 0 && utils.IsSpace(char) {
			continue // ignore the start space
		}
		if startChar == '\n' {
			if char == startChar {
				// 结束
				end = true
				break
			} else if utils.IsEmptyChar(char) {
				break
			}
		} else {
			if char == '\\' {
				var next []byte = make([]byte, 1)
				if _, err = c.reader.Read(next); err != nil {
					return
				}
				char = utils.EscapeWord(next[0])
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
func (c ReaderScanner) ScanWords() (words []string, err error) {

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
func (c ReaderScanner) CheckToEnd() (check bool, err error) {
	reader := bufio.NewReader(c.reader)
	var bytes []byte
	if bytes, err = reader.ReadBytes('\r'); err != nil {
		check = false
		return
	}
	str := strings.TrimSpace(string(bytes))
	check = str == ""
	return
}

// NewScanner to read stdin line content for scan operation
func NewScanner() Scanner {
	return NewFScanner(os.Stdin)
}

// NewFScanner to read Reader line content for scan operation
func NewFScanner(reader io.Reader) (line Scanner) {
	return &ReaderScanner{
		reader: reader,
	}
}
