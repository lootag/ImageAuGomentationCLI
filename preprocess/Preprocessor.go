package preprocess

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
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
		classesToExclude []string,
		wg *sync.WaitGroup)
}
