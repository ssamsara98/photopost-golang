package lib

// FiberLogger logger for gin framework [subbed from main logger]
type FiberLogger struct {
	*Logger
}

// Writer interface implementation for gin-framework
func (l FiberLogger) Write(p []byte) (n int, err error) {
	str := string(p)
	size := len(p)
	l.Info(str)
	return size, nil
}
