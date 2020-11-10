package scan

import (
	"io/ioutil"
)

func(this ScanningService) getAnnotationPaths(folderToScan string) []string {
	annotationsToRead := []string{}
	root := folderToScan + "/Annotations"
	fileInfos, err := ioutil.ReadDir(root)
	if err != nil {
		panic(err)
	}
	for fileInfoIndex := range fileInfos {
		annotationsToRead = append(annotationsToRead, folderToScan+"/Annotations/"+fileInfos[fileInfoIndex].Name())
	}
	return annotationsToRead
}