package constant

// Key value pair struct
type KeyValue struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type ReserveErrorCode int

type ErrorDetails struct {
	Key   string      `json:"key"`
	Value interface{} `json:"value"`
}
