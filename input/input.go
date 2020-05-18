package input

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func GetInputAsString(question string) string {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("%s: ", question)
	scanner.Scan()

	return scanner.Text()
}

func GetInputAsStringWithDefaultValue(question string, defaultAnswer string) string {
	input := GetInputAsString(fmt.Sprintf("%s [%s]", question, defaultAnswer))
	if input == "" {
		return defaultAnswer
	}

	return input
}

func GetInputAsStringArray(question string, delimiter string) []string {
	rawInput := GetInputAsString(question)
	return strings.Split(rawInput, delimiter)
}

func GetInputAsBool(question string, stringMatchForTrue string) bool {
	rawInput := GetInputAsString(question)
	return rawInput == stringMatchForTrue
}
