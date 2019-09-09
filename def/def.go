package def

import "sync"

type DictionaryLoad struct {
	RwLock     sync.RWMutex
	UpdateTime int64
}

type Config struct {
	ConfigPath    string
	LogPath       string
	LogLevel      string
	WordPath string
	TermDepth int
	UpdateInterval int64
	HttpPort int
}

type FormatTerm struct {
	Term string
	Pos string
	Frequency int
}
