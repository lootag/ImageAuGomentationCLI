package preprocess

import (
	"regexp"
)

func(this PreprocessingService) renameAnnotation(annotationName string) string {
	extensionRegex := regexp.MustCompile(`\.[a-z]+$`)
	matches := extensionRegex.FindAllString(annotationName, -1)
	extension := matches[0]
	numberOfCharactersToDelete := len(extension) - 1
	newAnnotationName := "resized" + annotationName[:len(annotationName)-numberOfCharactersToDelete] + "xml"
	return newAnnotationName
}