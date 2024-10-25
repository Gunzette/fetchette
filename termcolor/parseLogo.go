package termcolor

import (
	"encoding/json"
	"log"
	"os"
	"path"
	"strconv"
	"strings"
)

func ParseColorString(str string, colList [][3]int) string {
	strParts := strings.Split(str, "{")

	var res string

	for i, part := range strParts {
		if i == 0 {
			res += part
			continue
		}

		partRune := []rune(part)

		if string(partRune[0:2]) == "-1" {
			res += ResetAttrib + part[3:]
			continue
		}

		colIndex, err := strconv.Atoi(string(partRune[0:2]))

		if err != nil {
			log.Fatal(err)
		}

		activeCol := colList[colIndex]

		res += FgColor(part[3:], &activeCol)
	}

	return res
}

type JSONLogoStruct struct {
	File   string
	Colors [][3]int
}

func generateLogoDataFromJSONInfo(filename string) (string, [][3]int) {
	exPath, err := os.Executable()
	if err != nil {
		log.Fatal("Error while getting executable path: ", err)
	}

	content, err := os.ReadFile(path.Dir(exPath) + "/ascii/" + filename)
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}

	var data JSONLogoStruct
	err = json.Unmarshal(content, &data)
	if err != nil {
		log.Fatal("Error while parsing JSON: ", err)
	}

	return data.File, data.Colors
}

func GetColoredLogoString(filename string) string {
	file, colors := generateLogoDataFromJSONInfo(filename)

	exPath, err := os.Executable()
	if err != nil {
		log.Fatal("Error while getting executable path: ", err)
	}

	logoString, err := os.ReadFile(path.Dir(exPath) + "/ascii/" + file)
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}

	coloredString := ParseColorString(string(logoString), colors)

	return coloredString
}
