package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

func main() {
	fmt.Println("Updating README.md")

	contents := getContents(".")

	readmeContent := fmt.Sprintf("# etc\n\n## Table of contents\n\n%s", printContents(contents, 0))

	err := ioutil.WriteFile("README.md", []byte(readmeContent), 0644)
	panicOnError(err)
}

const (
	directoryContentType string = "directory"
	fileContentType      string = "file"
)

type Content struct {
	Name     string    `json:"name"`
	Href     string    `json:"href"`
	Type     string    `json:"type"`
	Contents []Content `json:"contents,omitempty"`
}

func (c Content) isDir() bool {
	return c.Type == directoryContentType
}

func getContents(parentDir string) []Content {
	files, err := ioutil.ReadDir(parentDir)
	panicOnError(err)

	contents := []Content{}

	for _, file := range files {
		if strings.HasPrefix(file.Name(), ".") || file.Name() == "README.md" || file.Name() == "_config.yml" || file.Name() == "obsidian_templates" {
			continue
		}

		path := parentDir + "/" + file.Name()

		name := strings.Replace(file.Name(), "./", "", 1)
		if strings.HasSuffix(name, ".md") {
			name = strings.Replace(file.Name(), ".md", "", 1)
		}

		content := Content{
			Name: name,
			Href: path,
			Type: fileContentType,
		}

		if file.IsDir() {
			contents := getContents(path)
			content.Contents = contents
			content.Type = directoryContentType
		}

		contents = append(contents, content)
	}

	return contents
}

func printContents(contents []Content, level int) string {
	out := ""

	for _, content := range contents {
		indent := ""
		for i := 0; i < level; i++ {
			indent += "  "
		}

		line := indent + "- "
		if content.isDir() {
			line += content.Name
		} else {
			line += fmt.Sprintf("[%s](%s)", content.Name, content.Href)
		}
		line += "\n"

		out += line

		if content.isDir() {
			out += printContents(content.Contents, level+1)
		}
	}

	return out
}

func panicOnError(err error) {
	if err != nil {
		panic(err)
	}
}
