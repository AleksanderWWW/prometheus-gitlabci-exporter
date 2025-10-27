package exporter


type Metrics struct {
	PipelineCount int32
}

func GetMetrics(config exporterConfig) (*Metrics, error) {
	return &Metrics{
		PipelineCount: 42,
	}, nil
}
