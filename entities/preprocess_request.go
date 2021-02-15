package entities

type PreprocessRequest struct{
    ImagePaths []string
    ImageNames []string
    AnnotationType AnnotationType
    folder string
    size int
    annotated bool
    classesToExclude string
}
