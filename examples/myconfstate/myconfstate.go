// Example myconfstate
package myconfstate

import (
	cs "github.com/mamemomonga/go-confstate"
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
func Load(cf string) error {
	cs.ConfigsFile        = cf        // ConfigsFile Filename
	cs.DefaultBaseDirType = cs.DBTBin // Basedir is offset from executable binary
	cs.OffsetFromBin      = ".."      // Base directory offset from executable binary
	cs.DefaultConfigsFile = "etc/configs.yaml" // Configs File default filename
	cs.DefaultStatesFile  = "etc/states.json"  // States File default filename
	cs.Debug = true                   // Debug mode

	// Initalize Configs
	cs.Configs = &Configs{
		Key1:  "Value1",
		Key2:  "Value2",
		Key3:  "Value3",
		Users: []string{"user1", "user2", "user3"},
	}
	// Initalize States
	cs.States = &States{
		Passwords: make(map[string]string),
	}

	if err := cs.LoadConfigs(); err != nil {
		return err
	}

	// Sets the StatesFile if the states_file has been set
	if v := C().StatesFile; v != "" {
		cs.StatesFile = v
	}

	if err := cs.LoadStates(); err != nil {
		return err
	}
	return nil
}

// C Accessor for Configs
func C() *Configs {
	return cs.Configs.(*Configs)
}

// S Accessor for States
func S() *States {
	return cs.States.(*States)
}

// GetDir Accessor for GetDir(Not required)
func GetDir(p string) (string, error) {
	return cs.GetDir(p)
}

// Save States file
func Save() error {
	return cs.SaveStates()
}
