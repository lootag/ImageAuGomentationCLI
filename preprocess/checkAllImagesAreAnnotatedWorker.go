package preprocess

import (
	"fmt"
	"os"
	"sync"
)

func checkAllImagesAreAnnotatedWorker(fileName string,
	preprocessWaitGroup *sync.WaitGroup,
	checkedAnnotations chan string) {
	defer (*preprocessWaitGroup).Done()
	xmlToBeChecked := "./Annotations/" + fileName[:len(fileName)-3] + "xml"
	_, err := os.Stat(xmlToBeChecked)
	if err == nil {
		checkedAnnotations <- xmlToBeChecked
	} else{
		fmt.Println("The image " + fileName + " is not annotated. Ignoring.")
	}
}
