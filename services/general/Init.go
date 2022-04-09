package general

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.CheckIranNationalCodeR, &CheckIranNational)
}
