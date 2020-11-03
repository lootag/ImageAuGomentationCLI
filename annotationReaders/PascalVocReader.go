package annotationReaders

import (
	"encoding/xml"
	"io/ioutil"
	"os"
	"sync"
	"regexp"

	"github.com/lootag/ImageAuGomentationCLI/annotationDtos"
	"github.com/lootag/ImageAuGomentationCLI/entities"
)

type PascalVocReader struct {
}

//In the Read function I replace the "FileName" composing the image's actual name retrieved from fs, because the xml's FileName value could be different file's name
func (pascalVocReader PascalVocReader) Read(annotationPath string,
	inputAnnotations chan entities.Annotation,
	aumentationWaitGroup *sync.WaitGroup) {
	defer (*aumentationWaitGroup).Done()
	var xmlAnnotation annotationDtos.PascalVoc
	xmlFile, err := os.Open(annotationPath)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &xmlAnnotation)
	var annotation entities.Annotation
	annotation.Classes = []string{}
	annotation.BoundingBoxes = []entities.BoundingBox{}
	for objectIndex := 0; objectIndex < len(xmlAnnotation.Objects); objectIndex++ {
		var boundingBox entities.BoundingBox
		boundingBox.Xmin = xmlAnnotation.Objects[objectIndex].BndBox.Xmin
		boundingBox.Ymin = xmlAnnotation.Objects[objectIndex].BndBox.Ymin
		boundingBox.Xmax = xmlAnnotation.Objects[objectIndex].BndBox.Xmax
		boundingBox.Ymax = xmlAnnotation.Objects[objectIndex].BndBox.Ymax
		annotation.Classes = append(annotation.Classes, xmlAnnotation.Objects[objectIndex].Name)
		annotation.BoundingBoxes = append(annotation.BoundingBoxes, boundingBox)
	}
	annotation.FileName = annotationPath[14:len(annotationPath) - 3] + getImageExtension(xmlAnnotation.FileName);
	annotation.Width = xmlAnnotation.Size.Width
	annotation.Height = xmlAnnotation.Size.Height
	annotation.Depth = xmlAnnotation.Size.Depth
	inputAnnotations <- annotation
}

func (pascalVocReader PascalVocReader) ReadSync(annotationPath string) entities.Annotation {
	var xmlAnnotation annotationDtos.PascalVoc
	xmlFile, err := os.Open(annotationPath)
	if err != nil {
		panic(err)
	}
	defer xmlFile.Close()
	byteValue, _ := ioutil.ReadAll(xmlFile)
	xml.Unmarshal(byteValue, &xmlAnnotation)
	var annotation entities.Annotation
	annotation.Classes = []string{}
	annotation.BoundingBoxes = []entities.BoundingBox{}
	for objectIndex := 0; objectIndex < len(xmlAnnotation.Objects); objectIndex++ {
		var boundingBox entities.BoundingBox
		boundingBox.Xmin = xmlAnnotation.Objects[objectIndex].BndBox.Xmin
		boundingBox.Ymin = xmlAnnotation.Objects[objectIndex].BndBox.Ymin
		boundingBox.Xmax = xmlAnnotation.Objects[objectIndex].BndBox.Xmax
		boundingBox.Ymax = xmlAnnotation.Objects[objectIndex].BndBox.Ymax
		annotation.Classes = append(annotation.Classes, xmlAnnotation.Objects[objectIndex].Name)
		annotation.BoundingBoxes = append(annotation.BoundingBoxes, boundingBox)
	}
	annotation.FileName = annotationPath[14:len(annotationPath) - 3] + getImageExtension(xmlAnnotation.FileName);
	annotation.Width = xmlAnnotation.Size.Width
	annotation.Height = xmlAnnotation.Size.Height
	annotation.Depth = xmlAnnotation.Size.Depth
	return annotation
}

func getImageExtension(fileName string) string{
	extensionRegex := regexp.MustCompile(`\.[a-z]+$`);
	matches := extensionRegex.FindAllString(fileName, -1);
	extensionWithDot := matches[0];
	extension := extensionWithDot[1:len(extensionWithDot)]
	return extension;
}




