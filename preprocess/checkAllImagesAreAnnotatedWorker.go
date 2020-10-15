package preprocess;

import(
	"sync";
	"os";
)

func checkAllImagesAreAnnotatedWorker(fileName string, 
	preprocessWaitGroup *sync.WaitGroup, 
	checkedAnnotations chan string){
	defer (*preprocessWaitGroup).Done();
	xmlToBeChecked := "./Annotations/" + fileName[:len(fileName) - 3] + "xml"
	_, err := os.Stat(xmlToBeChecked);
	if err != nil{
		panic("The image " + fileName + " is not annotated. Exiting");
	}
	checkedAnnotations <- xmlToBeChecked;
}