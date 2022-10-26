package http

type GetHealthCheckResponse struct {
	Status string `json:"status"`
}

type DependenciesResponse struct {
	MongoDB string `json:"mongodb"`
}

type GetInfoResponse struct {
	Dependencies DependenciesResponse `json:"dependencies"`
}
