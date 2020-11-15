package helpers

import (
	"os"
)

func Destroy() {
	if _, err := os.Stat("Images"); err == nil {
		os.RemoveAll("Images")
	}

	if _, err := os.Stat("Annotations"); err == nil {
		os.RemoveAll("Annotations")
	}

	if _, err := os.Stat("AugmentedImages"); err == nil {
		os.RemoveAll("AugmentedImages")
	}

	if _, err := os.Stat("AugmentedAnnotations"); err == nil {
		os.RemoveAll("AugmentedAnnotations")
	}
}