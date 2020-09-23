package rotate;

import(
	"fmt";
	"sync";
)

type RotateService struct{

}

func (rotateService RotateService) Augment(wg *sync.WaitGroup){
	defer wg.Done();
	rotate();
}

func rotate(){
	fmt.Println("Rotate");
}