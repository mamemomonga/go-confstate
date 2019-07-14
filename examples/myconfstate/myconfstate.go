// Example myconfstate
package myconfstate

import(
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
	cs.ConfigsFile = cf // ConfigsFile Filename
	cs.OffsetFromBin = ".." // Base directory offset from executable binary
	cs.DefaultConfigsFile = "configs.yaml" // Configs File default filename
	cs.DefaultStatesFile =  "states.json" // States File default filename
	cs.Debug = true // Debug mode

	// Initalize Configs
	cs.Configs = &Configs{
		Key1: "Value1",
		Key2: "Value2",
		Key3: "Value3",
		Users: []string{"user1","user2","user3"},
	}
	// Initalize States
	confstate.States = &States{
		Passwords: make(map[string]string),
	}

	return confstate.Load()
}

// C Accessor for Configs
func C() *Configs {
	return confstate.Configs.(*Configs)
}

// S Accessor for States
func S() *States {
	return confstate.States.(*States)
}

// GetDir Accessor for GetDir(Not required)
func GetDir(p string) (string, error) {
	return confstate.GetDir(p)
}

// Save States file
func Save() error {
	return confstate.SaveStates()
}

