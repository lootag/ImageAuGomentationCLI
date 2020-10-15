package convert;

import(
	"sync";
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type Converter interface{
	ConvertToJPG(toConvert chan entities.ImageInfo, 
		wg *sync.WaitGroup, 
		action string,
		fileNames *[]string);
	ConvertToText(toConvert chan entities.Annotation,
		wg *sync.WaitGroup,
		annotationType entities.AnnotationType)
}