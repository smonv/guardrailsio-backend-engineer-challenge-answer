package beca

type Finding struct {
	FindingType string           `json:"type"`
	RuleId      string           `json:"ruleId"`
	Location    *FindingLocation `json:"location"`
	Metadata    *FindingMetadata `json:"metadata"`
}

type FindingLocation struct {
	Path      string                    `json:"path"`
	Positions *FindingLocationPositions `json:"positions"`
}

type FindingLocationPositions struct {
	Begin *FindingLocationPositionsBegin `json:"begin"`
}

type FindingLocationPositionsBegin struct {
	Line int `json:"line"`
}

type FindingMetadata struct {
	Description string `json:"description"`
	Severity    string `json:"severity"`
}
