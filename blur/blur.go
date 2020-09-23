package blur;

import(
	
	"fmt";
	"sync";
)

type BlurService struct{

}

func (blurService BlurService) Augment(wg *sync.WaitGroup) {
	defer wg.Done();
	blur();
}

func blur(){
	fmt.Println("Blur");
}
