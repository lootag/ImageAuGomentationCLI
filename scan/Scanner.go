package scan

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type Scanner interface {
	Scan(annotationType entities.AnnotationType, folderToScan string)
}
