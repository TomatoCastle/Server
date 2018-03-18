package session

import (
	"sync"
	"time"
	"net/http"
	"fmt"
)
var lock sync.Mutex
var sessions = make (map[string]map[string]interface{})

/*
type Session struct{
	userid int64
	crateAt time.Time
	otehr
}*/
/*
type Session interface {

}*/

func SessionGet(w http.ResponseWriter, r *http.Request) (sess map[string]interface{}) {
	lock.Lock()
	defer lock.Unlock()
	if c, err := r.Cookie("go-session");c != nil{
		sess = sessions[c.Value]
	}else if err != nil || c == nil {
		sess = make(map[string]interface{})
	}
	return sess
}

func SessionSet(w http.ResponseWriter, r *http.Request, sessionColum string, value map[string]interface{}) error{
//todo セッションをほぞんするの中身を追加
	return nil
}
