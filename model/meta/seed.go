package meta

import (
	"encoding/json"
	"io/ioutil"
	"log"
)

func ReadContentFromJson(fileName string) UMLMeta {
	data, err := ioutil.ReadFile(fileName)
	content := UMLMeta{}
	if nil != err {
		log.Panic(err.Error())
	} else {
		err = json.Unmarshal(data, &content)
	}

	return content
}

var SystemMeta *UMLMeta
var DefualtAuthServiceMeta *UMLMeta

func init() {
	content := ReadContentFromJson("./seeds/system-meta.json")
	SystemMeta = &content

	authContent := ReadContentFromJson("./seeds/auth-meta.json")
	DefualtAuthServiceMeta = &authContent
}
