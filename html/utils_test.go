package html

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"
	"gokogiri/help"
	"testing"
)

func badOutput(actual string, expected string) {
	fmt.Printf("Got:\n[%v]\n", actual)
	fmt.Printf("Expected:\n[%v]\n", expected)
}

func getTestData(name string) (input []byte, output []byte, error string) {
	var errorMessage string
	offset := "\t"
	inputFile := filepath.Join(name, "input.txt")

	input, err := ioutil.ReadFile(inputFile)

	if err != nil {
		errorMessage += fmt.Sprintf("%vCouldn't read test (%v) input:\n%v\n", offset, name, offset+err.Error())
	}

	output, err = ioutil.ReadFile(filepath.Join(name, "output.txt"))

	if err != nil {
		errorMessage += fmt.Sprintf("%vCouldn't read test (%v) output:\n%v\n", offset, name, offset+err.Error())
	}

	return input, output, errorMessage
}

func collectTests(suite string) (names []string, error string) {
	testPath := filepath.Join("tests", suite)
	entries, err := ioutil.ReadDir(testPath)

	if err != nil {
		return nil, fmt.Sprintf("Couldn't read tests:\n%v\n", err.Error())
	}

	for _, entry := range entries {
		if strings.HasPrefix(entry.Name(), "_") || strings.HasPrefix(entry.Name(), ".") {
			continue
		}

		if entry.IsDir() {
			names = append(names, filepath.Join(testPath, entry.Name()))
		}
	}

	return
}

func CheckXmlMemoryLeaks(t *testing.T) {
	help.LibxmlCleanUpParser()
	if ! help.LibxmlCheckMemoryLeak() {
		t.Errorf("Memory leaks: %d!!!", help.LibxmlGetMemoryAllocation())
		help.LibxmlReportMemoryLeak()
	}
}

