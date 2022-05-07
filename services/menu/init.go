package menu

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle(env.MenuR, &Duty)
}
