package graphql

type Origin struct {
	Device string
	Lon    float64
	Lat    float64
}

type Telemetry struct {
	From Origin
	When string
}

type Temperature struct {
	Telemetry
	Temp float64
}
