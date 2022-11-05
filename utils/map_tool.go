package utils

func RemoveTopStructNew(fields map[string]string) string {
	var res string
	for _, err := range fields {
		if len(res) == 0 {
			res = err
			break
		}
	}
	return res
}
