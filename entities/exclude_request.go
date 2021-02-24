package entities

type ExcludeRequest struct {
	ExclusionThreshold    int
	UserDefinedExclusions []string
	ImageNames            []string
	Folder                string
	AnnotationType        AnnotationType
}
