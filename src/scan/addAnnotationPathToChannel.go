package scan

import (
	"sync"
)

func(this ScanningService) addAnnotationPathToChannel(annotationPath string,
	annotationsToRead chan string,
	getAnnotationPathsWaitGroup *sync.WaitGroup) {
	defer (*getAnnotationPathsWaitGroup).Done()
	annotationsToRead <- annotationPath
}