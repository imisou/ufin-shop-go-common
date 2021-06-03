package runtime

import "net/http"

type RuntimeApplication struct {
	//
	engin http.Handler
}

//------------------------------------------------------------
//----------   Gin 有关 ---------------------------------------
//------------------------------------------------------------

func (r RuntimeApplication) GetGinEngin() http.Handler {
	return r.engin
}
