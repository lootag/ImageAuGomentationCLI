package main;

import(
	"sync";
)

var wg sync.WaitGroup;

func main(){
	/*
	for argumentIndex := 1; argumentIndex < len(os.Args); argumentIndex++{
		augmentation, err := AugmentationsFactory(os.Args[argumentIndex]);
		if err == nil{
			wg.Add(1);
			go augmentation.Augment(&wg);
		} else {
			panic("There was a critical exception, exiting.")
		}
	}
	wg.Wait()
	*/
}