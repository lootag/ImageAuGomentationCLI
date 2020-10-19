package main;

import(
	"regexp";
)

func getUserDefinedExclusions(argument string) []string{
	userDefinedExclusions := []string{};
	if argument == ""{
		return userDefinedExclusions;
	}
	exclusionRegex := regexp.MustCompile(`[A-Za-z_]+;`)
	matches := exclusionRegex.FindAllString(argument, -1);
	for matchIndex := range matches{
		userDefindeExclusion := matches[matchIndex][:len(matches[matchIndex]) - 1]
		userDefinedExclusions = append(userDefinedExclusions, userDefindeExclusion);
	}
	return userDefinedExclusions;
}