package toolsbox

import (
	"io/ioutil"
	"strings"
)

func ParseList(path string) (map[string]string, error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	} else if len(f) < 3 {
		return make(map[string]string), nil
	}
	content := string(f)
	basicarr := strings.Split(content, "\n")
	var namelist = make(map[string]string)
	for _, v := range basicarr {
		if len(v) > 2 {
			resarr := strings.Split(v, "=")
			if len(resarr) == 2 {
				//name=path
				namelist[resarr[0]] = resarr[1]
			}
		}
	}
	return namelist, nil
}
func FormatList(origin map[string]string, path string) (bool, error) {
	recordmsg := ""
	for k, v := range origin {
		recordmsg += k + "=" + v + "\n"
	}
	err := ioutil.WriteFile(path, []byte(recordmsg), 0666)
	if err != nil {
		return false, err
	}
	return true, err
}
