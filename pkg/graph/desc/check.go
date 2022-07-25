package desc

type CheckLevel byte

const (
	CheckLevelNone CheckLevel = iota
	CheckLevelCount
	CheckLevelSampling
	CheckLevelEachRow
)

type Check struct {
	Level string
}
