package types

type MemoryResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MemoryRequestParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
