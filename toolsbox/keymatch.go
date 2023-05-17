package toolsbox

func KeyMatch(tocmp []byte, cmplist []byte) (lang int, content []byte) {
	if len(tocmp) == 0 || len(cmplist) == 0 {
		return
	}
	i := 0
	j := 0
	endpos := 0
	for j < len(cmplist) {
		for i < len(tocmp) && tocmp[i] != cmplist[j] {
			i++
		}
		if i >= len(tocmp) {
			i = endpos
		} else {
			//if match,flush end position
			endpos = i + 1
			lang++
			content = append(content, tocmp[i])
		}
		j++
	}
	return
}
