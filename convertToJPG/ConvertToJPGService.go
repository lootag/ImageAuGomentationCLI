package convertToJPG;

import(
	"fmt";
	"sync";
	"os";
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type ConvertToJPGService struct{

}

func (convertToJPGSerivce ConvertToJPGService) Convert(toConvert chan entities.ImageInfo,
	mainWaitGroup *sync.WaitGroup, 
	action string,
	fileNames *[]string){
	defer (*mainWaitGroup).Done();
	fmt.Println("converting " + action);
	var wg sync.WaitGroup;
	err := os.Mkdir("./AugmentedImages", 0755);
	if err != nil{
	}
	for image := range toConvert{
		wg.Add(1);
		go convertToJPGWorker(image, &wg);
	}
	wg.Wait();
}