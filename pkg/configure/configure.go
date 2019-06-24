package configure

import (
	"fmt"
	"github.com/dicf/go-api/pkg/file"
	"gopkg.in/yaml.v2"
	"log"
)

var Cfg = &Configure{}

func Setup() {
	cfg_byte, err := file.ReadFile("./conf/app.yaml")
	if err != nil {
		fmt.Println(err.Error())
	}
	//fmt.Println(string(cfg_byte))

	err = yaml.Unmarshal(cfg_byte, &Cfg)
	if err != nil {
		log.Println(err.Error())
	}

}
