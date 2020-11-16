package entities

type TwitterFeed struct {
	ID           string `json:"id_str"`
	Text         string `json:"text"`
	RetweetCount int64  `json:"retweet_count"`
}
