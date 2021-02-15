/*
This file is part of ImageAuGomentationCLI.

ImageAuGomentationCLI is free software: you can redistribute it and/or modify
it under the terms of the GNU General Public License as published by
the Free Software Foundation, either version 2 of the License, or
(at your option) any later version.

ImageAuGomentationCLI is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU General Public License for more details.

You should have received a copy of the GNU General Public License
along with ImageAuGomentationCLI.  If not, see <https://www.gnu.org/licenses/>.
*/
package commons

import (
	"regexp"
)

func GetSemicolonSeparatedValues(argument string) []string {
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
