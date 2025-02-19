package core

import (
	"fmt"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"log"
	"server/config"
	"server/global"
)

//var AppConfig *config.Config

// const configFile = "setting.yaml"
const configFile = "D:\\code\\golang\\project\\gvb\\server\\setting.yaml"

// 读取配置文件
func InitConf() {
	global.Config = &config.Config{}
	yamlConfig, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("read config file error, %v", err))
	}
	err = yaml.Unmarshal(yamlConfig, global.Config)
	if err != nil {
		log.Fatal("unmarshal config file error, %v", err)
	}
	log.Println("config yamlFile load Init success.")

	//fmt.Println(global.Config)

	//viper.SetConfigName("setting")
	//viper.SetConfigType("yaml")
	//viper.AddConfigPath("./")
	//
	//if err := viper.ReadInConfig(); err != nil {
	//	log.Fatalf("viper.ReadInConfig() error(%v)", err)
	//}
	//
	//// 初始化全局配置
	//global.Config = &config.Config{}
	//
	//// 将 YAML 配置映射到结构体
	//if err := viper.Unmarshal(global.Config); err != nil {
	//	log.Fatalf("viper.Unmarshal() error(%v)", err)
	//}

}

func SetYaml() error {
	data := global.Config
	//将结构体信息转换为yaml格式
	byteData, err := yaml.Marshal(data)
	if err != nil {
		return err
	}
	//写入文件
	err = ioutil.WriteFile(configFile, byteData, 0644)
	if err != nil {
		return err
	}
	return nil
}
