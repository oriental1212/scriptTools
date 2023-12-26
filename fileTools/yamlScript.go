package fileTools

import (
	"fmt"
	"github.com/oriental1212/scriptTools/sqlTools"
	"gopkg.in/yaml.v2"
	"os"
	"strings"
)

type Database struct {
	Table  string `yaml:"table"`
	Action string `yaml:"action"`
}

type Info struct {
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Url      string `yaml:"url"`
}

type Config struct {
	SQLInfo     Info                `yaml:"info"`
	SQLDatabase map[string]Database `yaml:"databaseList"`
}

func YamlScript() {
	file, err := os.ReadFile("./SqlScript.yaml")
	if err != nil {
		panic(err)
	}
	var config Config
	err = yaml.Unmarshal(file, &config)
	if err != nil {
		panic(err)
	}

	for name, database := range config.SQLDatabase {
		fmt.Println("当前运行数据库:", name)
		tableList := strings.Split(database.Table, ",")
		actionList := strings.Split(database.Action, ",")
		resultList := make([]string, 0, len(actionList)*len(tableList))
		for _, table := range tableList {
			for _, action := range actionList {
				newString := strings.Replace(action, "{tableName}", table, 1)
				resultList = append(resultList, newString)
			}
		}
		db := sqlTools.OpenDataBase(config.SQLInfo.Username, config.SQLInfo.Password, config.SQLInfo.Url, name)
		for _, sqlStatement := range resultList {
			fmt.Printf("本次执行SQL: %s", sqlStatement)
			result := db.Exec(sqlStatement)
			if result.Error != nil {
				fmt.Println(" 结果:sql执行失败，错误原因:", result.Error)
			} else {
				fmt.Println(" 结果:sql执行成功")
			}
		}
	}
}
