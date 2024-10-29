package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"path"
	"strings"
	"unicode/utf8"

	"github.com/gunzette/fetchette/modules"
	"github.com/gunzette/fetchette/termcolor"
)

const LOGOGAP int = 2

type JsonStruct struct {
	TextColors [3][3]int
}

func getTextColors(filename string) [3][3]int {
	exPath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}

	content, err := os.ReadFile(path.Dir(exPath) + "/ascii/" + filename)
	if err != nil {
		log.Fatal("Error while opening file: ", err)
	}

	var payload JsonStruct
	err = json.Unmarshal(content, &payload)
	if err != nil {
		log.Fatal("Error while umarshalling json: ", err)
	}
	return payload.TextColors
}

func displayFetch(colorsFile string, moduleOutputs []string) error {
	// Get logo string
	logo, err := termcolor.GetColoredLogoString(colorsFile)
	if err != nil {
		return err
	}

	outputLen := len(moduleOutputs)

	textColors := getTextColors(colorsFile)

	// Add to each line of logo
	for i, line := range strings.Split(logo, "\n") {
		if line == "" {
			continue
		}

		if i < outputLen {
			moduleOutput, err := termcolor.ParseColorString(moduleOutputs[i], textColors[:])
			if err != nil {
				return err
			}

			fmt.Println(
				line,
				strings.Repeat(" ", LOGOGAP),
				moduleOutput,
			)
		} else {
			fmt.Println(line)
		}
	}
	return nil
}

func main() {
	// Check arguments
	if len(os.Args) == 1 {
		log.Fatal("Not enough arguments")
	}

	userAtHost := modules.GetUserAtHost()

	err := displayFetch(os.Args[1]+"Colors.json", []string{userAtHost, strings.Repeat("-", utf8.RuneCountInString(userAtHost)), modules.GetOS(), modules.GetKernel()})
	if err != nil {
		log.Fatal("Error while displaying fetch: ", err)
	}
}
