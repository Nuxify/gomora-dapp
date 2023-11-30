package types

type CreateNFTLogSetGreeting struct {
	Greeting  string `json:"greeting"`
	Timestamp uint64 `json:"timestamp"`
	IsFromWS  bool   `json:"-"`
}
