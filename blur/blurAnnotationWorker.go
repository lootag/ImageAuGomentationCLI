package blur

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

func blurAnnotationWorker(annotationToBlur entities.Annotation,
	blurredAnnotations chan entities.Annotation,
	blurWaitGroup *sync.WaitGroup) {
	defer (*blurWaitGroup).Done()
	blurredAnnotation := annotationToBlur
	blurredAnnotation.NewName = "blur" + blurredAnnotation.FileName[:len(blurredAnnotation.FileName)-3] + "xml"
	blurredAnnotations <- blurredAnnotation

}
