package flip; 

import(
	
	"fmt";
	"sync"
)

type FlipService struct{

}


func (flipService FlipService) Augment(wg *sync.WaitGroup){
	defer wg.Done();
	flip();
}

func flip(){
	fmt.Println("Flip");
}