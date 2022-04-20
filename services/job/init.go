package job

import (
	"net/http"
	"tel-note/env"
)

func Init() {
	http.Handle(env.ListOfJobR, &GetJob)
	http.Handle(env.DeleteJobByIdR, &DeleteJob)
	http.Handle(env.EditJobByIdR, &EditJob)
	http.Handle(env.FindJobR, &FindJobID)
	http.Handle(env.InsertNewJobR, &NewJob)
}
