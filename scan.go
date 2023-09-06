package goscan

import (
	"bufio"
	"errors"
	"io"
	"os"
	"strings"

	"github.com/lingdor/goscan/utils"
)

var ErrToEndNotEmpty = errors.New("to end is not empty")

type Scanner interface {
	Scan() (str string, err error)
	ScanWords() ([]string, error)
	ReadToEnd() (str string, err error)
	ToEnd() error
	//todo ScanVars([]interface{}) error
}

type LineScanner struct {
	Scanner
	reader io.Reader
	isEnd  bool
}

func (c *LineScanner) Scan() (str string, err error) {

	if c.isEnd {
		return "", io.EOF
	}

	var bs []byte = make([]byte, 1)
	var result []byte = make([]byte, 0, 10)
	var n int
	startChar := byte('\n')
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
			c.isEnd = true
			break
		}

		if len(result) == 0 && utils.IsSpace(char) {
			continue // ignore the start space
		}
		if startChar == '\n' {
			if char == startChar {
				// 结束
				c.isEnd = true
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
func (c *LineScanner) ScanWords() (words []string, err error) {

	words = make([]string, 0, 2)
	var str string
	for {
		if str, err = c.Scan(); err != nil {
			return
		}
		if str != "" {
			words = append(words, str)
		}
		if c.isEnd {
			return
		}
	}
}
func (c *LineScanner) ReadToEnd() (str string, err error) {
	if c.isEnd {
		return "", nil
	}
	reader := bufio.NewReader(c.reader)
	var bytes []byte
	if bytes, err = reader.ReadBytes('\n'); err != nil {

		str = ""
		return
	}
	c.isEnd = true
	str = strings.TrimSpace(string(bytes))
	return
}

func (c *LineScanner) ToEnd() (err error) {
	var str string
	str, err = c.ReadToEnd()
	if str != "" {
		err = ErrToEndNotEmpty
	}
	return
}

// NewScanner to read stdin line content for scan operation
func NewLineScanner() Scanner {
	return NewFLineScanner(os.Stdin)
}

// NewFScanner to read Reader line content for scan operation
func NewFLineScanner(reader io.Reader) (line Scanner) {
	return &LineScanner{
		reader: reader,
		isEnd:  false,
	}
}
