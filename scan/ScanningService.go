package scan;

import(
	"fmt"
	"sync";
	"io/ioutil";
	"github.com/lootag/ImageAuGomentationCLI/entities";
	"github.com/lootag/ImageAuGomentationCLI/annotationReaders";
	"strconv";
)

type ScanningService struct{
	
}

func (scanningService ScanningService) Scan(annotationType entities.AnnotationType,
	folderToScan string){
	annotationsToRead := getAnnotationPaths(folderToScan);
	annotationsToGroup := readAnnotations(annotationType, annotationsToRead);
	countMap := getCountMap(annotationsToGroup);
	fmt.Println("Here's a scan of your data: ")
	for key, value := range countMap {
		fmt.Println(key + ", " + strconv.Itoa(value) + " instances");
	}
}

func getAnnotationPaths(folderToScan string) []string{
		annotationsToRead := []string{}
		root := folderToScan + "/Annotations"
		fileInfos, err := ioutil.ReadDir(root)
		if err != nil{
			panic(err);
		}
		for fileInfoIndex := range fileInfos{
			annotationsToRead = append(annotationsToRead, folderToScan + "/Annotations/" + fileInfos[fileInfoIndex].Name(), );
		}
		return annotationsToRead;
}

func readAnnotations(annotationType entities.AnnotationType,
	annotationsToRead []string) []entities.Annotation {
	annotationsToGroup := []entities.Annotation{};
	var factory annotationReaders.AnnotationReadersFactory
	annotationReader, err := factory.Create(annotationType)
	if err != nil {
		panic(err)
	}
	for annotationPathIndex := range annotationsToRead {
		annotation := annotationReader.ReadSync(annotationsToRead[annotationPathIndex]);
		annotationsToGroup = append(annotationsToGroup, annotation);
	}
	return annotationsToGroup;
}

func getCountMap(annotationsToGroup []entities.Annotation) map[string]int{
	countMap := make(map[string]int);
	for annotationIndex := range annotationsToGroup{
		for classIndex := range annotationsToGroup[annotationIndex].Classes{
			if contains(getMapKeys(countMap), annotationsToGroup[annotationIndex].Classes[classIndex]){
				countMap[annotationsToGroup[annotationIndex].Classes[classIndex]] += 1;
			} else {
				countMap[annotationsToGroup[annotationIndex].Classes[classIndex]] = 1;
			}
		}
	}
	return countMap;
}

func addAnnotationPathToChannel(annotationPath string, 
	annotationsToRead chan string, 
	getAnnotationPathsWaitGroup *sync.WaitGroup){
		defer (*getAnnotationPathsWaitGroup).Done()
		annotationsToRead <- annotationPath;
}

func contains(stringArray1 []string, toCheck string) bool{
	for stringIndex := range stringArray1{
		if stringArray1[stringIndex] == toCheck{
			return true
		}
	}
	return false;
}

func getMapKeys(stringIntMap map[string]int) []string{
	keys := []string{}
	for key, _ := range stringIntMap{
		keys = append(keys, key)
	}

	return keys;
}
