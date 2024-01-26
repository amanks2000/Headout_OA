package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func ConvertDataToInt(data string) int {
	result, err := strconv.Atoi(data)

	if err != nil {
		return -1
	}

	return result
}

func NumberValidator(data string) bool {
	pattern := `^[1-9]\d*$`

	re := regexp.MustCompile(pattern)

	return re.MatchString(data)
}

func precompute_data() {
	cache = make(map[string]string)
	folderPath := "tmp/data"

	err := filepath.Walk(folderPath, visit)
	if err != nil {
		fmt.Printf("Error: %s\n", err)
		return
	}
}

func visit(path string, info os.FileInfo, err error) error {
	if err != nil {
		fmt.Println(err)
		return nil
	}
	if !info.IsDir() {
		content, err := os.ReadFile(path)
		if err != nil {
			fmt.Println(err)
		}
		fileName, err := extractFilename(path)
		if err != nil {
			fmt.Println(err)
		}
		temp := ""
		cnt := 1
		for i := 0; i < len(content); i++ {
			c := content[i]
			if c == '\n' {
				t := fileName + "_" + strconv.Itoa(cnt)
				cache[t] = temp
				cnt++
				temp = ""
				continue
			}
			temp += string(c)
		}
		if temp != "" {
			cache[fileName+"_"+strconv.Itoa(cnt)] = temp
		}
		t := fileName + "_len"
		cache[t] = strconv.Itoa(cnt)
	}
	return nil
}

func extractFilename(path string) (string, error) {
	re := regexp.MustCompile(`([^/]+)\.\w+$`)
	matches := re.FindStringSubmatch(path)

	if len(matches) > 1 {
		return matches[1], nil
	}

	return "", fmt.Errorf("no filename found in path: %s", path)
}
