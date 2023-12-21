package protondb_integration

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
	switch s.Tier {
	case "pending":
		return s.ProvisionalTier
	default:
		return s.Tier
	}
}

func (s *Summary) GetConfidence() string {
	return s.Confidence
}
