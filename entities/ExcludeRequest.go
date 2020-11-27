package entities

type ExcludeRequest struct{
    ExclusionThreshold int
    UserDefinedExclusion []string
    ImageNames []string
    Folder string
    AnnotationType AnnotationType
}
