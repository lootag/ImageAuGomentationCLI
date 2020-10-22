package main

import (
	"regexp"
)

func getUserDefinedExclusions(argument string) []string {
	userDefinedExclusions := []string{}
	if argument == "" {
		return userDefinedExclusions
	}
	classesRegex := regexp.MustCompile(`[A-Za-z_]+;`)
	matches := classesRegex.FindAllString(argument, -1)
	for matchIndex := range matches {
		userDefindeExclusion := matches[matchIndex][:len(matches[matchIndex])-1]
		userDefinedExclusions = append(userDefinedExclusions, userDefindeExclusion)
	}
	return userDefinedExclusions
}
