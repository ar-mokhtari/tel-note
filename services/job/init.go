package job

import (
	"github.com/ar-mokhtari/tel-note/env"
	"net/http"
)

func Init() {
	http.Handle(env.ListOfJobR, &GetJob)
	http.Handle(env.DeleteJobByIdR, &DeleteJob)
	http.Handle(env.EditJobByIdR, &EditJob)
	http.Handle(env.FindJobR, &FindJobID)
	http.Handle(env.InsertNewJobR, &NewJob)
}
