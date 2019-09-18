package confstate_test

import (
	"testing"
	confstate "github.com/mamemomonga/go-confstate"
	"os"
//	"log"
//	"flag"
)

var skipFlag bool = false

// Configs file structure
type Configs struct {
	StatesFile string   `yaml:"states_file"`
	Key1       string   `yaml:"key_1"`
	Key2       string   `yaml:"key_2"`
	Key3       string   `yaml:"key_3"`
}

// States file structure
type States struct {
	Key4      string   `yaml:"key_4"`
}

// C Accessor for Configs
func C() *Configs {
	return confstate.Configs.(*Configs)
}

// S Accessor for States
func S() *States {
	return confstate.States.(*States)
}

func Test01Init(t *testing.T) {
	if skipFlag {
		t.SkipNow()
	}
	confstate.DefaultConfigsFile = "test.yaml"
	confstate.DefaultStatesFile =  "test.json"
	confstate.DefaultBaseDirType = confstate.DBTWork
	confstate.Debug = true
	confstate.Configs = &Configs{
		Key1: "Value1",
		Key2: "Value2",
		Key3: "Value3",
	}
	confstate.States = &States{
		Key4: "Value1",
	}
	if err := confstate.LoadConfigs(); err != nil {
		skipFlag = true
		t.Fatal(err)
	}
	if err := confstate.LoadStates(); err != nil {
		skipFlag = true
		t.Fatal(err)
	}
}

func Test02SaveState(t *testing.T) {
	if skipFlag {
		t.SkipNow()
	}
	if err := confstate.SaveStates(); err != nil {
		skipFlag = true
		t.Fatal(err)
	}
}

func Test03NewFile(t *testing.T) {
	if skipFlag {
		t.SkipNow()
	}
	if !confstate.NewConfigsFile {
		skipFlag = true
		t.Fatal("ConfigsFile not New File")
	}
	if !confstate.NewStatesFile {
		skipFlag = true
		t.Fatal("StatesFile not New File")
	}
}

func Test04RemoveFile(t *testing.T) {
	if skipFlag {
		t.SkipNow()
	}
	t.Logf("Remove: %s", confstate.DefaultConfigsFile)
	if err := os.Remove(confstate.DefaultConfigsFile); err != nil {
		skipFlag = true
		t.Fatal(err)
	}
	t.Logf("Remove: %s", confstate.DefaultStatesFile)
	if err := os.Remove(confstate.DefaultStatesFile); err != nil {
		skipFlag = true
		t.Fatal(err)
	}
}

