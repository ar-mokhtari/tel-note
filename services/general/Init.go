package general

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init() {
	http.Handle(env.CheckIranNationalCodeR, &CheckIranNational)
}
