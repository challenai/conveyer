package desc

type Sink struct {
	Kind  string
	DSL   string
	Retry int
	// Database                 string
	Table                    string
	GenerateTable            bool
	CreateIndex              bool
	CleanHistoryData         bool
	RemoveWatermarkIfSucceed bool
	Extra                    any
}
