package rotate

import (
	"github.com/lootag/ImageAuGomentationCLI/entities"
	"sync"
)

func xmlWorker(originalHeight int,
	originalWidth int,
	newSize int,
	originalBoundingBox entities.BoundingBox,
	rotateWaitGroup *sync.WaitGroup) {
	defer (*rotateWaitGroup).Done()
}
