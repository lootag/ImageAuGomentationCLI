package annotationReaders

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type AnnotationReader interface {
	Read(annotationPath string, inputAnnotations chan entities.Annotation, wg *sync.WaitGroup) //Async
	ReadSync(annotationPath string) entities.Annotation
}
