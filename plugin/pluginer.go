package plugin

import (
	"encoding/xml"
	"errors"
	"io/ioutil"
	"plugin"
	"strings"

	"github.com/oswaldoooo/octools/toolsbox"
)

type Pluginer struct {
	classpath string
	confinfo  *cnfinfo
	coremap   map[string]func(*plugin.Plugin, ...any) error //scan the plugin's function
}
type cnfinfo struct {
	XMLName  xml.Name      `xml:"plugin_info"`
	Plugins  []plugin_info `xml:"plugin"`
	RootPath string        `xml:"rootpath"`
}
type plugin_info struct {
	XMLName   xml.Name `xml:"plugin"`
	ClassName string   `xml:"classname,attr"`
	Name      string   `xml:"name,attr"`
}

/*
create pluginers,and read from configure file to cache
*/
func CreatePluginer(classpath string) (pluginer *Pluginer, err error) {
	content, err := ioutil.ReadFile(classpath)
	if err == nil {
		cnf := new(cnfinfo)
		err = xml.Unmarshal(content, cnf)
		if err == nil {
			if cnf.RootPath[len(cnf.RootPath)-1] != '/' {
				//the path is ...../testpath,it'wll read to ...../testpath/
				cnf.RootPath += "/"
			}
			pluginer = &Pluginer{classpath: classpath, confinfo: cnf, coremap: make(map[string]func(*plugin.Plugin, ...any) error)}
		}
	}
	return
}

/*
put the lookup function into pluginer maker,lookup all pluginer
*/
func (s *Pluginer) LookUpAll() (allerr []error) {
	allerr = []error{}
	for _, plugininfo := range s.confinfo.Plugins {
		if lookupfunc, ok := s.coremap[plugininfo.ClassName]; ok {
			//if lookup function existed,do it
			pluginer, err := toolsbox.ScanPluginByName(plugininfo.Name, s.confinfo.RootPath+"plugin")
			if err == nil {
				//only to simple pack
				err = lookupfunc(pluginer)
			}
			if err != nil {
				allerr = append(allerr, err)
			}
		}
	}
	return
}

/*
complex pack,put args into the target classname that need args
*/
func (s *Pluginer) Lookup(classname string, args ...any) (err error) {
	pluginname := ""
	if _, ok := s.coremap[classname]; !ok {
		//if classname not in coremap,return error
		err = errors.New(classname + " is dont existed in coremap")
		return
	}
	usefunc := s.coremap[classname]
	for i, j := 0, len(s.confinfo.Plugins)-1; ; {
		if j < i {
			break
		}
		if s.confinfo.Plugins[i].ClassName == classname {
			pluginname = s.confinfo.Plugins[i].Name
			break
		} else if s.confinfo.Plugins[j].ClassName == classname {
			pluginname = s.confinfo.Plugins[j].Name
			break
		}
		i++
		j--
	}
	if len(pluginname) > 0 && !strings.ContainsRune(pluginname, ' ') {
		pluginer, err := toolsbox.ScanPluginByName(pluginname, s.confinfo.RootPath+"plugin")
		if err == nil {
			usefunc(pluginer, args...)
		}
	} else {
		err = errors.New("not find plugin file that classname is " + classname)
	}
	return
}
