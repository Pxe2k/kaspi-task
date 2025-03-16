package delivery

import (
	"github.com/Pxe2k/kaspi-task/internal/domain"
	"github.com/Pxe2k/kaspi-task/pkg"
	"github.com/gorilla/mux"
	"net/http"
)

func (h *Handler) find(w http.ResponseWriter, r *http.Request) {
	in := FindRequest{
		Name: mux.Vars(r)["name"],
	}

	if err := in.validate(); err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	persons, err := h.uc.Find(r.Context(), in.Name)
	if err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	pkg.ToJSON(w, http.StatusOK, convToResp(persons))
}

func convToResp(persons []domain.Person) FindResponse {
	p := make([]personsResponse, 0, len(persons))
	for _, person := range persons {
		p = append(p, personsResponse{
			DocumentID: person.DocumentID,
			Name:       person.Name,
			Phone:      person.Phone,
		})
	}

	return FindResponse{
		Persons: p,
	}
}
