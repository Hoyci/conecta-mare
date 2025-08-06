package communities

import (
	"conecta-mare-server/pkg/httphelpers"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

var (
	instance *communitiesHandler
	Once     sync.Once
)

func NewHandler(categoriesService CommunitiesService) *communitiesHandler {
	Once.Do(
		func() {
			instance = &communitiesHandler{
				communitiesService: categoriesService,
			}
		},
	)

	return instance
}

func (h communitiesHandler) RegisterRoutes(r *chi.Mux) {
	r.Route(
		"/api/v1/communities", func(r chi.Router) {
			// Public
			r.Get("/", h.handleGetCommunities)
		},
	)
}

func (h communitiesHandler) handleGetCommunities(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	communities, err := h.communitiesService.GetCommunities(ctx)
	if err != nil {
		httphelpers.WriteJSON(w, err.Code, err.Err)
		return
	}

	httphelpers.WriteJSON(w, http.StatusOK, map[string]any{"communities": communities})
}
