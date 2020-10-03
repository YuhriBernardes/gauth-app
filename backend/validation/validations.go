package validation

func RequiredString(s string) (valid bool) {

	if s == "" {
		return false
	}

	return true

}
