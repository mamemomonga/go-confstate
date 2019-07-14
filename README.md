# go-confstate

Set in YAML, save in JSON

[![GoDoc](https://godoc.org/github.com/mamemomonga/go-confstate?status.svg)](https://godoc.org/github.com/mamemomonga/go-confstate)

# 設定と状態

* YAMLで設定、JSONで状態を保存する
* それぞれのパスを未指定の場合は、実行バイナリの相対位置から決定される
* このサンプルでは、ユーザごとにパスワードを生成する

# INSTALL

	go get -u -v github.com/mamemomonga/go-confstate

# Examples

main.go

	package main
	
	import (
		"github.com/davecgh/go-spew/spew"
		confstate "github.com/mamemomonga/go-confstate"
		"github.com/sethvargo/go-password/password"
		"log"
		"flag"
	)
	
	// Configs file structure
	type Configs struct {
		StatesFile string   `yaml:"states_file"`
		Key1       string   `yaml:"key_1"`
		Key2       string   `yaml:"key_2"`
		Key3       string   `yaml:"key_3"`
		Users      []string `yaml:"users"`
	}
	
	// States file structure
	type States struct {
		Passwords map[string]string `json:"passwords"`
	}
	
	// Load Initalize and load
	func LoadConfigs(cf string) error {
	
		confstate.ConfigsFile = cf // ConfigsFile Filename
	
		confstate.DefaultConfigsFile = "etc/configs.yaml" // Configs File default filename
	
		confstate.DefaultStatesFile =  "etc/states.json" // States File default filename
	
		confstate.DefaultBaseDirType = confstate.DBTBin // Basedir is offset from executable binary
		confstate.OffsetFromBin = "." // Base directory offset from executable binary
		// If you do not specify ConfigsFile, using "go run" will cause problems due to the location of the execution binary.
		// ConfigsFileを指定しない場合、"go run"を使用すると実行バイナリの位置の関係で問題が発生します。
	
		// confstate.DefaultBaseDirType = confstate.DBTHome // DBTHome: Basedir is Home Directory
	
		// confstate.DefaultBaseDirType = confstate.DBTWork // DBTWork: Basedir is Work Directory
	
		confstate.Debug = true // Debug mode
	
		// Initalize Configs
		confstate.Configs = &Configs{
			Key1: "Value1",
			Key2: "Value2",
			Key3: "Value3",
			Users: []string{"user1","user2","user3"},
		}
		// Initalize States
		confstate.States = &States{
			Passwords: make(map[string]string),
		}
	
		if err := confstate.LoadConfigs(); err != nil {
			return err
		}
		if err := confstate.LoadStates(); err != nil {
			return err
		}
		return nil
	}
	
	// C Accessor for Configs
	func C() *Configs {
		return confstate.Configs.(*Configs)
	}
	
	// S Accessor for States
	func S() *States {
		return confstate.States.(*States)
	}
	
	func main() {
	
		c := flag.String("c","","ConfigsFile")
		flag.Parse()
	
		// Initalize and load configs and states
		if err := LoadConfigs(*c); err != nil {
			log.Fatal(err)
		}
	
		// Set a password for a user whose password is not set
		for _, name := range C().Users {
			if _, ok := S().Passwords[name]; !ok {
				pwd, err := password.Generate(16, 2, 1, false, false)
				if err != nil {
					log.Fatal(err)
				}
				S().Passwords[name] = pwd
			}
		}
	
		// Save StatesFile
		if err := confstate.SaveStates(); err != nil {
			log.Fatal(err)
		}
	
		// Dump it
		spew.Dump(S())
	}


run

	$ go build -o example ./main.go
	$ ./example

# LICENSE

MIT

