package job

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.ListOfJobR, &GetJob)
}
