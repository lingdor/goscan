package strs

func IsEmptyChar(bt byte) bool {
	return bt == '\t' || bt == ' ' || bt == '\r' || bt == '\n' || bt == byte(0)
}

func IsSpace(bt byte) bool {
	return bt == '\t' || bt == ' '
}

func EscapeWord(char byte) byte {

	switch char {
	case 'n':
		return '\n'
	case 'r':
		return '\r'
	case 'a':
		return '\a'
	case 'b':
		return '\b'
	case 't':
		return '\t'
	case 'v':
		return '\v'
	case '0':
		return byte(0)
	default:
		return char
	}

}
