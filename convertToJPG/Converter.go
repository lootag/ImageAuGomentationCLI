package convertToJPG;

import(
	"sync";
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type Converter interface{
	Convert(toConvert chan entities.ImageInfo, 
		wg *sync.WaitGroup, 
		action string,
		fileNames *[]string);
}