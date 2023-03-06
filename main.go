package main

import (
	"encoding/json"
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type fileStructure struct {
	Title    string `json:"title,omitempty"`
	Body     string `json:"body,omitempty"`
	AltTitle string `json:"alt-title,omitempty"`
}

var (
	file   string
	folder string
)

func init() {
	fileFlag := flag.String("file", "files.json", "The file to read from")
	folderFlag := flag.String("folder", "", "The folder to put the files into")
	flag.Parse()

	file = *fileFlag
	folder = *folderFlag
}

func main() {
	err := checkExist(file, "file")
	if err != nil {
		panic(err)
	}
	err = checkExist(folder, "folder")
	if err != nil {
		panic(err)
	}
	fSlice := readFile(file)
	for _, file := range fSlice {
		createFile(file)
	}
}

func checkExist(str, kind string) error {
	if str == "" {
		return nil
	}
	stats, err := os.Stat(str)
	if err != nil {
		return errors.New("failed to vertify the " + kind)
	}
	switch kind {
	case "folder":
		if !stats.IsDir() {
			return errors.New("needs to be a directory, provided value was not")
		}
	}
	return nil
}

func readFile(str string) (fSlice []fileStructure) {
	data, err := ioutil.ReadFile(str)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(data, &fSlice)
	if err != nil {
		panic(err)
	}
	return fSlice
}

func createFile(f fileStructure) {
	fileTitle := strings.ReplaceAll(strings.ToLower(f.Title), " ", "-")
	contents := fmt.Sprintf("---\ntitle: \"%s\"\n---\n\n%s", f.Title, f.Body)
	var err error
	if f.AltTitle != "" {
		err = os.WriteFile(fmt.Sprintf("%s/%s", folder, f.AltTitle), []byte(contents), 0644)
	} else {
		err = os.WriteFile(fmt.Sprintf("%s/%s.md", folder, fileTitle), []byte(contents), 0644)
	}
	if err != nil {
		panic(err)
	}
}
