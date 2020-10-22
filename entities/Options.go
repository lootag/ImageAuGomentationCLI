package entities

type Options struct {
	Size                  int
	Direction             Direction
	Sigma                 float64
	Annotated             bool
	BatchSize             int
	InAnnotationType      AnnotationType
	OutAnnotationType     AnnotationType
	ExclusionThreshold    int
	UserDefinedExclusions []string
}
