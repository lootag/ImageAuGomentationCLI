package main;

import(
	"sync";
)

type IAugmentationService interface{
	Augment(wg *sync.WaitGroup);
}

