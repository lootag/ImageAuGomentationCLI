package exclusion;

import(
	"sync";
)

func getAnnotationPathsFromImageNamesWorker(imageName string, 
	annotationPaths chan string,
	preprocessWaitGroup *sync.WaitGroup){
	defer (*preprocessWaitGroup).Done();
	annotationPath := "./Annotations/" + imageName[:len(imageName) - 3] + "xml";
	annotationPaths <- annotationPath;
}
