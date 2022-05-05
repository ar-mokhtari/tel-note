package reporter

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init() {
	http.Handle(env.ContactReportR, &ContactReport)
}
