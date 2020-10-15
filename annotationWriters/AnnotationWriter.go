package annotationWriters;

import(
	"sync";
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type AnnotationWriter interface{
	Write(annotation entities.Annotation, augmentationWaitGroup *sync.WaitGroup);
}