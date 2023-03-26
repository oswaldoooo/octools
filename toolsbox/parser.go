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

// advance parser method
func ParseListBySymbol(path, symbol string) (res map[string]string, err error) {
	f, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, err
	} else if len(f) < 3 {
		return make(map[string]string), nil
	}
	content := string(f)
	basicarr := strings.Split(content, "\n")
	res = make(map[string]string)
	for _, v := range basicarr {
		if len(v) > 2 {
			resarr := strings.Split(v, symbol)
			if len(resarr) == 2 {
				//name=path
				res[resarr[0]] = resarr[1]
			}
		}
	}
	return
}

// advanced method for format the map to file
func FormatListBySymbol(origin map[string]string, path, symbol string) (bool, error) {
	recordmsg := ""
	for k, v := range origin {
		recordmsg += k + symbol + v + "\n"
	}
	err := ioutil.WriteFile(path, []byte(recordmsg), 0666)
	if err != nil {
		return false, err
	}
	return true, err
}

// parse the bytes array to the target map,split by custom symbol freely
func ParseListUltra(content []byte, symbol string) (res map[string]string, err error) {
	contentstr := string(content)
	basicarr := strings.Split(contentstr, "\n")
	res = make(map[string]string)
	for _, v := range basicarr {
		if len(v) > 2 {
			resarr := strings.Split(v, symbol)
			if len(resarr) == 2 {
				//name=path
				res[resarr[0]] = resarr[1]
			}
		}
	}
	return
}

// foramt the map to string and custom spilt sybol freely
func FormatListUltra(origin map[string]string, symbol string) (res string) {
	res = ""
	for k, v := range origin {
		res += k + symbol + v + "\n"
	}
	return
}
