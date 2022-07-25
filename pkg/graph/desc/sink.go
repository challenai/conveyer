package desc

type Sink struct {
	DSL   string
	Retry int
	// Database                 string
	// Table                    string
	GenerateTable            bool
	CreateIndex              bool
	CleanHistoryData         bool
	RemoveWatermarkIfSucceed bool
}
