package augmentation;

import(
	"sync";
)

type IAugmentationService interface{
	Augment(wg *sync.WaitGroup);
}

