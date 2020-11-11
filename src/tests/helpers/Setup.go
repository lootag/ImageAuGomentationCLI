package helpers

import (
	"os"
	"path/filepath"

	"github.com/lootag/ImageAuGomentationCLI/commons"
)

func Setup(defaultFolder bool, annotationType string) {
	targetImagesFolder := filepath.Join("resources", annotationType, "Images")
	targetAnnotationsFolder := filepath.Join("resources", annotationType, "Annotations")
	if defaultFolder {
		_, err := os.Stat("Images")
		if err != nil {
			os.Mkdir("Images", 0755)
		}
		commons.CopyDirectory(targetImagesFolder, "Images")
		_, err = os.Stat("Annotations")
		if err != nil {
			os.Mkdir("Annotations", 0755)
		}
		commons.CopyDirectory(targetAnnotationsFolder, "Annotations")
	}
	os.Remove("AugmentedImages/*.*")
	os.Remove("AugmentedAnnotations/*.*")
}