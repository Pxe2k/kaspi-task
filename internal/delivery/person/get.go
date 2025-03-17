package person

import (
	"github.com/Pxe2k/kaspi-task/pkg"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) get(w http.ResponseWriter, r *http.Request) {
	in := GetRequest{
		DocumentID: mux.Vars(r)["document_id"],
	}

	if err := in.validate(); err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	person, err := h.uc.Get(r.Context(), in.DocumentID)
	if err != nil {
		pkg.ToErr(w, http.StatusInternalServerError, err)
		return
	}

	pkg.ToJSON(w, http.StatusOK, GetResponse{
		DocumentID: person.DocumentID,
		Name:       person.Name,
		Phone:      person.Phone,
	})
}
