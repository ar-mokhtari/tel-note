package general

func ISNil(input []string) []string {
	if input == nil {
		return []string{""}
	}
	return input
}
