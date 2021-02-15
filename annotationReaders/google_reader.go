package annotationReaders

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

type GoogleReader struct {
}

func (this GoogleReader) Read(annotationPath string,
	inputAnnotations chan entities.Annotation,
	wg *sync.WaitGroup) {
	defer wg.Done()

}

func (this GoogleReader) ReadSync(annotationPath string) entities.Annotation {
	var annotation entities.Annotation
	return annotation
}
