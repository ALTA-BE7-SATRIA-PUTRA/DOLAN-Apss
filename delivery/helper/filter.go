package helper

import (
	"strings"
)

func FilterComment(comment string) (string, error) {
	find := []string{"anjing", "babi", "asu", "celeng", "pelacur", "lonte"}
	replace := "*"
	newText := comment
	for i, txt := range find {
		newText = strings.ReplaceAll(newText, txt, strings.Repeat(replace, len(find[i])))
	}

	return newText, nil
}
