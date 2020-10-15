package preprocess;


import(
	"sync";
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

func xmlWorker(originalHeight int, 
	originalWidth int,
	newSize int,
	originalBoundingBox entities.BoundingBox,
	rotateWaitGroup *sync.WaitGroup){
	defer (*rotateWaitGroup).Done();
}