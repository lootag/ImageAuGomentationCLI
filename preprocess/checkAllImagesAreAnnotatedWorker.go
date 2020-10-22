package preprocess

import (
	"fmt"
	"os"
	"sync"
)

func checkAllImagesAreAnnotatedWorker(imageName string,
	preprocessWaitGroup *sync.WaitGroup,
	checkedAnnotations chan string) {
	defer (*preprocessWaitGroup).Done()
	xmlToBeChecked := "./Annotations/" + imageName[:len(imageName)-3] + "xml"
	_, err := os.Stat(xmlToBeChecked)
	if err == nil {
		checkedAnnotations <- xmlToBeChecked
	} else {
		fmt.Println("The image " + imageName + " is not annotated. Ignoring.")

	}
}
