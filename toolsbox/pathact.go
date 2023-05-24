package toolsbox

import "strings"

// type Dir struct {
// 	name   string
// 	child  *Dir
// 	parent *Dir
// }

// func (s *Dir) String() string {
// 	if s.parent == nil {
// 		return s.name
// 	}
// 	return s.String() + "/" + s.name
// }
// func ToDir(fullpath string) *Dir {
// 	patharr := strings.Split(fullpath, "/")
// 	if len(patharr) > 2 {
// 		root := &Dir{name: patharr[0]}
// 		var cur *Dir
// 		cur = root
// 		for i := 1; i < len(patharr)-1; i++ {
// 			cur.child = &Dir{name: patharr[i], parent: cur}
// 			cur = cur.child
// 		}
// 		return root
// 	} else if len(patharr) == 2 {
// 		return &Dir{name: patharr[0]}
// 	} else {
// 		return nil
// 	}
// }

// func GetChildDir(origin *Dir, level int) *Dir {
// 	var cur *Dir = origin
// 	for cur.child != nil && level > 0 {
// 		cur = cur.child
// 		level--
// 	}
// 	return cur
// }

func GetBaseName(path string) string {
	patharr := strings.Split(path, "/")
	if len(patharr) > 0 {
		return patharr[len(patharr)-1]
	}
	return ""
}
func GetDirPath(path string) string {
	patharr := strings.Split(path, "/")
	if len(patharr) > 1 {
		patharr = patharr[:len(patharr)-1]
	}
	return strings.Join(patharr, "/")
}
