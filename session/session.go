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
	Set(key, value interface{}) error //valueをセット
	Get(key interface{}) interface{}  //valueを取得
	Delete(key interface{}) error     //kyeにもとずいて削除
	SessionID() string                //sidを取得
}
