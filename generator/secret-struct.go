package main

// Secret stores the code, the code's creation time and the delay (seconds)
type Secret struct {
	Code      string `json:"code"`
	Timestamp int64  `json:"timestamp"` // nanosecond
	Delay     int    `json:"delay"`     // second
}
