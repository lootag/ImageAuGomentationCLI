package annotationReaders

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type AnnotationReader interface {
	Read(annotationPath string, inputAnnotations chan entities.Annotation, wg *sync.WaitGroup)
	ReadSync(annotationPath string) entities.Annotation
}
