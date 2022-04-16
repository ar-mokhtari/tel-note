package sex

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.ListOfSexR, &GetSex)
}
