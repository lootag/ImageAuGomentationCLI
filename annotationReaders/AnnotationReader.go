package annotationReaders

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

type AnnotationReader interface {
	Read(filePath string, inputAnnotations chan entities.Annotation, wg *sync.WaitGroup)
}
