package session

//https://astaxie.gitbooks.io/build-web-application-with-golang/ja/06.2.htmlを参考にさせていただいています。
import "sync"

type Manager struct{
	cookiName string
	lock sync.Mutex
	provider Provider
	liveingTime int64
}

type Provider interface {
	SessionInit(sid string) (Session,error)//sessionの初期化を実装か
	SessionRead(sid string) (Session, error)//sidが示すセッションを変数を返すもしなければ新しく作る
	SessionDestroy(sid string) error//sessionを削除
	SessionGarbageCollection(maxLifeTime int64)//時間切れのセッションをお掃除
}

type Session interface {
	Set(key, value interface{}) error //set session value
	Get(key interface{}) interface{}  //get session value
	Delete(key interface{}) error     //delete session value
	SessionID() string                //back current sessionID
}
