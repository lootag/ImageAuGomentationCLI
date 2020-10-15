package annotationWriters

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

type AnnotationWriter interface {
	Write(annotation entities.Annotation, augmentationWaitGroup *sync.WaitGroup)
}
