package convert

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

type Converter interface {
	ConvertToJPG(toConvert chan entities.ImageInfo,
		wg *sync.WaitGroup,
		action string,
		fileNames *[]string)
	ConvertToText(toConvert chan entities.Annotation,
		wg *sync.WaitGroup,
		annotationType entities.AnnotationType)
}
