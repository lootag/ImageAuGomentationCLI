package preprocess

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type Preprocessor interface {
	Preprocess(imagePaths []string,
		imageNames []string,
		resizedImages chan entities.ImageInfo,
		resizedImagesCopy chan entities.ImageInfo,
		resizedAnnotations chan entities.Annotation,
		resizedAnnotationsCopy chan entities.Annotation,
		annotationType entities.AnnotationType,
		size int,
		annotated bool,
		classesToExclude []string,
		wg *sync.WaitGroup)
}
