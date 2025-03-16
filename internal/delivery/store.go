package delivery

import (
	"encoding/json"
	"github.com/Pxe2k/kaspi-task/pkg"
	"io"
	"net/http"
)

func (h *Handler) store(w http.ResponseWriter, r *http.Request) {
	body, err := io.ReadAll(r.Body)
	if err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	var in StoreRequest
	if err = json.Unmarshal(body, &in); err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	if err = in.Validate(); err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	if err = h.uc.Store(r.Context(), in); err != nil {
		pkg.ToErr(w, http.StatusBadRequest, err)
		return
	}

	pkg.ToJSON(w, http.StatusOK, map[string]string{
		"status": "success",
	})
}
