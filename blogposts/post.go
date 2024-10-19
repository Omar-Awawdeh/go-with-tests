package blogposts

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"strings"
)

const (
	titleSeperator       = "Title: "
	descriptionSeperator = "Description: "
	tagsSeperator        = "Tags: "
)

type Post struct {
	Title, Description, Body string
	Tags                     []string
}

func newPost(postFile io.Reader) (Post, error) {
	scanner := bufio.NewScanner(postFile)

	readMetaLine := func(tagName string) string {
		scanner.Scan()
		return strings.TrimPrefix(scanner.Text(), tagName)
	}

	return Post{
		Title:       readMetaLine(titleSeperator),
		Description: readMetaLine(descriptionSeperator),
		Tags:        strings.Split(readMetaLine(tagsSeperator), ", "),
		Body:        readBody(scanner),
	}, nil
}

func readBody(scanner *bufio.Scanner) string {
	scanner.Scan() // Ignore line containing "---"

	buf := bytes.Buffer{}
	for scanner.Scan() {
		fmt.Fprintln(&buf, scanner.Text())
	}
	body := strings.TrimSuffix(buf.String(), "\n")
	return body
}
