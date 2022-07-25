package notify

type Notifier interface {
	Notify(msg string) error
}
