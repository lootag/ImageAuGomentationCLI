package flip; 

import(
	"image";
	"fmt";
	"sync";
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type FlipService struct{

}


func (flipService FlipService) Augment(toAugment chan image.Image, 
									   wg *sync.WaitGroup, 
									   augmented chan image.Image,
									   options entities.Options){
	defer wg.Done();
	flip();
}

func flip(){
	fmt.Println("Flip");
}