package common

type (
	Category struct {
		ID   string `json:"id"`
		Name string `json:"name"`
		Icon string `json:"icon"`
	}

	Subcategory struct {
		ID         string `json:"id"`
		Name       string `json:"name"`
		CategoryID string `json:"category_id,omitempty"`
	}

	CategoryWithSubcats struct {
		Category
		Subcategories []Subcategory `json:"subcategories"`
	}

	CategoryWithUserCount struct {
		Category
		UserCount int `json:"user_count"`
	}
)
