package desc

type NotifyKinds byte

const (
	NotifyFail         NotifyKinds = 0x01 << 0
	NotifySuccess      NotifyKinds = 0x01 << 1
	NotifyBatchFail    NotifyKinds = 0x01 << 2
	NotifyBatchSuccess NotifyKinds = 0x01 << 3
	NotifyCheckFail    NotifyKinds = 0x01 << 4
	NotifyCheckSuccess NotifyKinds = 0x01 << 5
)

type Notify struct {
	NotifyKinds
}
