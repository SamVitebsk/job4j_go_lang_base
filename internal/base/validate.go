package base

type ValidateRequest struct {
	UserID      string
	Title       string
	Description string
}

func Validate(req *ValidateRequest) []string {
	res := make([]string, 0)

	if req == nil {
		res = append(res, "req is nil")
		return res
	}

	if req.UserID == "" {
		res = append(res, "UserID is required")
	}
	if req.Title == "" {
		res = append(res, "Title is required")
	}
	if req.Description == "" {
		res = append(res, "Description is required")
	}

	return res
}
