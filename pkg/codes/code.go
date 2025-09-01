package codes

import "fmt"

var (
	CorrectCode = GetCode("739")

	CodeTaskNameMap = map[string]string{
		CorrectCode: "flag{brute_force}",
	}
)

func CheckCode(code string) (string, bool) {
	code, ok := CodeTaskNameMap[code+""]
	return code, ok
}

func GetCode(code string) string {
	return fmt.Sprintf("%s", code)
}
