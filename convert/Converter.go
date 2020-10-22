package convert

import (
	"sync"

	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type Converter interface {
	ConvertToJPG(imagesToConvert chan entities.ImageInfo,
		wg *sync.WaitGroup)
	ConvertToText(annotationsToConvert chan entities.Annotation,
		wg *sync.WaitGroup,
		annotationType entities.AnnotationType)
}
