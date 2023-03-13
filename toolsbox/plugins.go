package toolsbox

import (
	"plugin"
	"strings"
)
//scan plugin by full path
func ScanPlugin(filepath string)(pluginer *plugin.Plugin,err error){
	pluginer,err=plugin.Open(filepath)
	return
}
//scan plugin by plugin name ,input plugin name and root path of plugin
func ScanPluginByName(plugin_name,rootpath string)(pluginer *plugin.Plugin,err error){
	filename:=""
	if strings.Contains(plugin_name,".so"){
		filename=plugin_name
	}else{
		filename=plugin_name+".so"
	}
	realpath:=rootpath
	if rootpath[len(rootpath)-1:]=="/"{
		realpath+=filename
	}else{
		realpath+="/"+filename
	}
	pluginer,err=plugin.Open(realpath)
	return
}