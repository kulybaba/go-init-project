package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"strings"
)

type Env struct {
	projectName string
	dbDriver    string
	dbUser      string
	dbPassword  string
	dbHost      string
	dbPort      string
	dbName      string
}

func (env *Env) SetProjectName(projectName string) {
	env.projectName = projectName
}

func (env *Env) SetDbDriver(dbDriver string) {
	env.dbDriver = dbDriver
}

func (env *Env) SetDbUser(dbUser string) {
	env.dbUser = dbUser
}

func (env *Env) SetDbPassword(dbPassword string) {
	env.dbPassword = dbPassword
}

func (env *Env) SetDbHost(dbHost string) {
	env.dbHost = dbHost
}

func (env *Env) SetDbPort(dbPort string) {
	env.dbPort = dbPort
}

func (env *Env) SetDbName(dbName string) {
	env.dbName = dbName
}

func (env Env) CreateEnv() {
	file, err := os.Create("./.env")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	_, err = file.Write([]byte(initData["projectName"]["fieldName"] + "=" + strings.Replace(env.projectName, "\n", "", -1) + "\n"))
	if err != nil {
		log.Fatal(err)
	}
	_, err = file.Write([]byte(initData["dbDriver"]["fieldName"] + "=" + strings.Replace(env.dbDriver, "\n", "", -1) + "\n"))
	if err != nil {
		log.Fatal(err)
	}
	dbUrl := fmt.Sprintf("DATABASE_URL=%v:%v@tcp(%v:%v)/%v",
		strings.Replace(env.dbUser, "\n", "", -1),
		strings.Replace(env.dbPassword, "\n", "", -1),
		strings.Replace(env.dbHost, "\n", "", -1),
		strings.Replace(env.dbPort, "\n", "", -1),
		strings.Replace(env.dbName, "\n", "", -1) + "\n",
	)
	_, err = file.Write([]byte(dbUrl))
	if err != nil {
		log.Fatal(err)
	}
}

var initData = map[string]map[string]string{
	"projectName": {
		"message":   "Enter project name",
		"default":   "Project",
		"fieldName": "PROJECT_NAME",
	},
	"dbDriver": {
		"message":   "Enter DB driver (\"mysql\" or \"postgres\")",
		"default":   "mysql",
		"fieldName": "DATABASE",
	},
	"dbUser": {
		"message":   "Enter DB username",
		"default":   "username",
		"fieldName": "DB_USER",
	},
	"dbPassword": {
		"message":   "Enter DB password",
		"default":   "password",
		"fieldName": "DB_PASSWORD",
	},
	"dbHost": {
		"message":   "Enter DB host (default: localhost)",
		"default":   "localhost",
		"fieldName": "DB_HOST",
	},
	"dbPort": {
		"message":   "Enter DB port (default: \"3306\")",
		"default":   "3306",
		"fieldName": "DB_PORT",
	},
	"dbName": {
		"message":   "Enter DB name",
		"default":   "db_name",
		"fieldName": "DB_NAME",
	},
}

func main() {
	fmt.Println("*  *  *  *  *  *  *  *")
	fmt.Println("*                    *")
	fmt.Println("* GO Basic Framework *")
	fmt.Println("*                    *")
	fmt.Println("*  *  *  *  *  *  *  *")
	fmt.Println()
	fmt.Println("Init project:")

	env := InitEnv()
	env.CreateEnv()

	fmt.Println()
	fmt.Println("[OK] The project successfully initialized!!!")
}

func InitEnv() (env Env) {
	env = Env{}
	reader := bufio.NewReader(os.Stdin)
	i := 1
	for index, value := range initData {
		fmt.Printf("%v. %s -> ", i, value["message"])
		text, err := reader.ReadString('\n')
		if err != nil {
			log.Fatal(err)
		}
		if len(text) == 1 {
			reflect.ValueOf(&env).MethodByName("Set" + strings.Title(index)).Call([]reflect.Value{reflect.ValueOf(value["default"])})
		} else {
			reflect.ValueOf(&env).MethodByName("Set" + strings.Title(index)).Call([]reflect.Value{reflect.ValueOf(text)})
		}
		i++
	}
	return
}
