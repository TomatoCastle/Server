package session

import (
	"sync"
	"time"
)
var mu sync.Mutex
var sessions = make (map[string]Session)


type Session struct{
	dbid int64
	crateAt time.Time
}

