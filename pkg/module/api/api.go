package api

import "github.com/aquasecurity/trivy/pkg/module/serialize"

const Version = 1

type Module interface {
	Version() int
	Name() string
}

type Analyzer interface {
	RequiredFiles() []string
	Analyze(filePath string) (*serialize.AnalysisResult, error)
}

type PostScanner interface {
	PostScan(serialize.Results) serialize.Results
}
