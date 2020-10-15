package annotationReaders;

import(
	"sync";
	"github.com/lootag/ImageAuGomentationCLI/entities";
)

type AnnotationReader interface{
	Read(filePath string, inputAnnotations chan entities.Annotation, wg *sync.WaitGroup);
}