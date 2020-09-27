package preprocess;
import(
	"image";
	"sync";
)

type IPreprocessingService interface{
	Preprocess(images *[]string, toAugment chan image.Image, wg *sync.WaitGroup, size int);
}

