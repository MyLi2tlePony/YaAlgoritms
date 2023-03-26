package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strings"
)

var regStartTag = regexp.MustCompile(`^</*[a-z]+><`)
var regEndTag = regexp.MustCompile(`></*[a-z]+>$`)

var regStartNormTag = regexp.MustCompile(`^</*[a-z]+>`)
var regEndNormTag = regexp.MustCompile(`</*[a-z]+>$`)

var regOpenTag = regexp.MustCompile(`^<[a-z]+>$`)
var regCloseTag = regexp.MustCompile(`^</[a-z]+>$`)

var regTextTag = regexp.MustCompile(`[a-z]+`)

type Tag struct {
	tag   string
	index int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	line, _ := reader.ReadString('\n')

	line = strings.TrimRight(line, "\r\n")
	tags := make([]string, 0)

	for len(line) > 0 {
		if regStartTag.MatchString(line) {
			tag := regStartTag.FindString(line)
			tag = tag[:len(tag)-1]
			line = line[len(tag):]

			tags = append(tags, tag)
		} else {
			break
		}
	}

	endsTags := make([]string, 0)

	for len(line) > 0 {
		if regEndTag.MatchString(line) {
			tag := regEndTag.FindString(line)
			tag = tag[1:]
			line = line[:len(line)-len(tag)]

			endsTags = append(endsTags, tag)
		} else {
			break
		}
	}

	if (len(endsTags)+len(tags))%2 == 0 {
		for len(line) > 0 {
			if regStartNormTag.MatchString(line) {
				tag := regStartNormTag.FindString(line)
				line = line[len(tag):]

				tags = append(tags, tag)
			} else {
				break
			}
		}

		for len(line) > 0 {
			if regEndNormTag.MatchString(line) {
				tag := regEndNormTag.FindString(line)
				line = line[:len(line)-len(tag)]

				endsTags = append(endsTags, tag)
			} else {
				break
			}
		}
	}

	tags = append(tags, line)

	for i := len(endsTags) - 1; i >= 0; i-- {
		tags = append(tags, endsTags[i])
	}

	checkStack := make([]Tag, 0)
	for i := range tags {
		if len(checkStack) != 0 && openAndClose(checkStack[len(checkStack)-1].tag, tags[i]) {
			checkStack = checkStack[:len(checkStack)-1]
		} else {
			checkStack = append(checkStack, Tag{tags[i], i})
		}
	}

	for len(checkStack) > 0 && openAndClose(checkStack[0].tag, checkStack[len(checkStack)-1].tag) {
		checkStack = checkStack[1 : len(checkStack)-1]
	}

	repairTags(checkStack)

	for i := range checkStack {
		tags[checkStack[i].index] = checkStack[i].tag
	}

	for i := range tags {
		fmt.Print(tags[i])
	}
}

func repairTags(tags []Tag) {
	if len(tags) != 2 {
		return
	}

	if regOpenTag.MatchString(tags[0].tag) {
		tags[1].tag = "</" + regTextTag.FindString(tags[0].tag) + ">"
		return
	}

	if regCloseTag.MatchString(tags[1].tag) {
		tags[0].tag = "<" + regTextTag.FindString(tags[1].tag) + ">"
		return
	}
}

func openAndClose(tag1, tag2 string) bool {
	if !regOpenTag.MatchString(tag1) || !regCloseTag.MatchString(tag2) {
		return false
	}

	if regTextTag.FindString(tag1) != regTextTag.FindString(tag2) {
		return false
	}

	return true
}
