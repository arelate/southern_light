package vangogh_integration

const (
	RatingUnknown  = "Not Rated"
	RatingPositive = "Positive"
	RatingMixed    = "Mixed"
	RatingNegative = "Negative"
)

func RatingDesc(ri int64) string {
	rd := RatingUnknown
	if ri >= 70 {
		rd = RatingPositive
	} else if ri >= 40 {
		rd = RatingMixed
	} else if ri > 0 {
		rd = RatingNegative
	}
	return rd
}
