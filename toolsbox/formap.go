package toolsbox

func CheckArgs[T any](args []string, originmap map[string]T) bool {
	for _, v := range args {
		if _, ok := originmap[v]; !ok {
			return false
		}
	}
	return true
}
