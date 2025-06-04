package valueobjects

type Role string

const (
	Client       Role = "client"
	Professional Role = "professional"
)

func (r Role) IsValid() bool {
	switch r {
	case Client, Professional:
		return true
	}

	return false
}
