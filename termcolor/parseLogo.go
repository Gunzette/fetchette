package termcolor

import (
	"encoding/json"
	"os"
	"path"
	"strconv"
	"strings"
)

func ParseColorString(str string, colList [][3]int) (string, error) {
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
			return "", err
		}

		activeCol := colList[colIndex]

		res += FgColor(part[3:], &activeCol)
	}

	return res, nil
}

type JSONLogoStruct struct {
	File   string
	Colors [][3]int
}

func generateLogoDataFromJSONInfo(filename string) (JSONLogoStruct, error) {
	exPath, err := os.Executable()
	if err != nil {
		return JSONLogoStruct{}, err
	}

	content, err := os.ReadFile(path.Dir(exPath) + "/ascii/" + filename)
	if err != nil {
		return JSONLogoStruct{}, err
	}

	var data JSONLogoStruct
	err = json.Unmarshal(content, &data)
	if err != nil {
		return JSONLogoStruct{}, err
	}

	return data, nil
}

func GetColoredLogoString(filename string) (string, error) {
	jsonData, err := generateLogoDataFromJSONInfo(filename)
	if err != nil {
		return "", err
	}

	exPath, err := os.Executable()
	if err != nil {
		return "", err
	}

	logoString, err := os.ReadFile(path.Dir(exPath) + "/ascii/" + jsonData.File)
	if err != nil {
		return "", err
	}

	coloredString, err := ParseColorString(string(logoString), jsonData.Colors)
	if err != nil {
		return "", err
	}

	return coloredString, nil
}
