package graphql

type Origin struct {
	Device string
	Lon    float64
	Lat    float64
}

type Telemetry interface {
}

type Temperature struct {
	From Origin
	When string
	Temp float64
}
