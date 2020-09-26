package preprocess;
import(
	"sync";
)

type IPreprocessingService interface{
	Preprocess(images *[]string, toAugment chan string, wg *sync.WaitGroup);
}

