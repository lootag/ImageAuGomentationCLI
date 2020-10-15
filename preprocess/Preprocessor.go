package preprocess

import (
	"sync"
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type Preprocessor interface {
	Preprocess(images *[]string, 
		fileNames *[]string,
		toAugment chan entities.ImageInfo,
		toAugmentCopy chan entities.ImageInfo,
		resizedAnnotations chan entities.Annotation,
		resizedAnnotationsCopy chan entities.Annotation,
		annotationType entities.AnnotationType,
		size int,
		xmls bool,
		wg *sync.WaitGroup)
}
