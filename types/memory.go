package types

type Memory struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MemoryResponse struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}

type MemoryRequestParams struct {
	Key   string `json:"key"`
	Value string `json:"value"`
}
