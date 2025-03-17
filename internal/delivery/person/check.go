package person

import (
	"github.com/Pxe2k/kaspi-task/pkg"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) check(w http.ResponseWriter, r *http.Request) {
	in := CheckRequest{
		DocumentID: mux.Vars(r)["document_id"],
	}

	if err := in.validate(); err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	gender, birthDate, err := h.uc.Check(r.Context(), in.DocumentID)
	if err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	pkg.ToJSON(w, http.StatusOK, CheckResponse{
		Gender:      gender,
		DateOfBirth: birthDate,
	})
}
