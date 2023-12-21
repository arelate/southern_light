package protondb_integration

import "github.com/arelate/southern_light"

type Summary struct {
	BestReportedTier string  `json:"bestReportedTier"`
	Confidence       string  `json:"confidence"`
	ProvisionalTier  string  `json:"provisionalTier"`
	Score            float64 `json:"score"`
	Tier             string  `json:"tier"`
	Total            int     `json:"total"`
	TrendingTier     string  `json:"trendingTier"`
}

func (s *Summary) String() string {
	tier := ""
	switch s.Tier {
	case "pending":
		tier = s.ProvisionalTier
	default:
		tier = s.Tier
	}
	return southern_light.FirstToUpper(tier)
}

type ConfidenceGetter interface {
	GetConfidence() string
}

func (s *Summary) GetConfidence() string {
	return southern_light.FirstToUpper(s.Confidence)
}
