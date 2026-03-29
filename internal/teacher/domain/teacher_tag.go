package domain

type TagSentiment string

const (
	TagSentimentPositive TagSentiment = "positive"
	TagSentimentNeutral  TagSentiment = "neutral"
	TagSentimentNegative TagSentiment = "negative"
)

type Tag struct {
	ID        int
	Name      string
	Sentiment TagSentiment
}
