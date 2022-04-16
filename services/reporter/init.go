package reporter

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.ContactReportR, &ContactReport)
}
