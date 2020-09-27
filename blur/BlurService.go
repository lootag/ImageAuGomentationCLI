package blur;

import(
	"image";
	"fmt";
	"sync";
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type BlurService struct{

}

func (blurService BlurService) Augment(toAugment chan image.Image, 
									   wg *sync.WaitGroup,
									   augmented chan image.Image,
									   options entities.Options) {
	defer wg.Done();
	blur();
}

func blur(){
	fmt.Println("Blur");
}
