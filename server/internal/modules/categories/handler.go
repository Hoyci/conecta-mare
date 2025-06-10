package categories

import (
	"conecta-mare-server/pkg/httphelpers"
	"net/http"
	"sync"

	"github.com/go-chi/chi/v5"
)

var (
	instance *categoryHandler
	Once     sync.Once
)

func NewHandler(categoriesService CategoriesService) *categoryHandler {
	Once.Do(
		func() {
			instance = &categoryHandler{
				categoriesService: categoriesService,
			}
		},
	)

	return instance
}

func (h categoryHandler) RegisterRoutes(r *chi.Mux) {
	// m := middlewares.NewWithAuth(h.accessKey)
	r.Route(
		"/api/v1/categories", func(r chi.Router) {
			// Public
			r.Get("/", h.handleGetCategories)
		},
	)
}

func (h categoryHandler) handleGetCategories(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	includeSubcats := r.URL.Query().Get("include") == "subcategories"

	categories, err := h.categoriesService.GetCategories(ctx, includeSubcats)
	if err != nil {
		httphelpers.WriteJSON(w, err.Code, err.Err)
	}

	httphelpers.WriteJSON(w, http.StatusOK, categories)
}
