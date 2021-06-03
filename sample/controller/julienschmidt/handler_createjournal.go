package julienschmidt

import (
	"accounting/application/apperror"
	"accounting/infrastructure/log"
	"accounting/infrastructure/util"
	"accounting/usecase/createjournal"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/julienschmidt/httprouter"
)

// createJournalHandler ...
func (r *Controller) createJournalHandler(inputPort createjournal.Inport) httprouter.Handle {

	return func(w http.ResponseWriter, r *http.Request, ps httprouter.Params) {

		// for accessing query params /createjournal?id=123
		// r.URL.Query().Get("id")

		ctx := log.Context(r.Context(), "createjournal")

		var req createjournal.InportRequest
		if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
			newErr := apperror.FailUnmarshalResponseBodyError
			log.Error(ctx, err.Error())
			http.Error(w, util.MustJSON(NewErrorResponse(newErr)), http.StatusBadRequest)
			return
		}

		log.Info(ctx, util.MustJSON(req))

		res, err := inputPort.Execute(ctx, req)
		if err != nil {
			log.Error(ctx, err.Error())
			http.Error(w, util.MustJSON(NewErrorResponse(err)), http.StatusBadRequest)
			return
		}

		log.Info(ctx, util.MustJSON(res))
		fmt.Fprint(w, NewSuccessResponse(res))

	}
}