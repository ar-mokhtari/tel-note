package job

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init(mux *http.ServeMux) {
	mux.Handle(env.ListOfJobR, &GetJob)
	mux.Handle(env.DeleteJobByIdR, &DeleteJob)
	mux.Handle(env.EditJobByIdR, &EditJob)
	mux.Handle(env.FindJobR, &FindJobID)
	mux.Handle(env.InsertNewJobR, &NewJob)
}
