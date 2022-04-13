package menu

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.MenuR, &Duty)
}
