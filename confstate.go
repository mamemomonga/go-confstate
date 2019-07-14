package confstate

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
)

var (
	Configs            interface{} // Configs data
	States             interface{} // States data
	BaseDir            string // Base directory
	ConfigsFile        string // ConfigsFile Filename
	StatesFile         string // StatesFile Filename
	OffsetFromBin      string // Base directory offset from executable binary
	DefaultConfigsFile string // Default filename ConfigFile
	DefaultStatesFile  string // Default filename StatesFile
	Debug              bool   = false // Debug Mode
)

// GetDir Get absorute directory from excutable binary
func GetDir(p string) (string, error) {
	return filepath.Abs(filepath.Join(BaseDir, p))
}

// Init Initialize
func Init() {
	// Get relative directory from executable binary
	exe, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	BaseDir, err = filepath.Abs(filepath.Join(filepath.Dir(exe), OffsetFromBin))
	if err != nil {
		log.Fatal(err)
	}
}

// Load Initialize and Load
func Load() error {

	// Set default value if ConfigsFile is empty
	if ConfigsFile == "" {
		if c, err := GetDir(DefaultConfigsFile); err != nil {
			return err
		} else {
			ConfigsFile = c
		}
	}
	if _, err := os.Stat(ConfigsFile); !os.IsNotExist(err) {
		// Load ConfigsFile
		buf, err := ioutil.ReadFile(ConfigsFile)
		if err != nil {
			return err
		}
		err = yaml.Unmarshal(buf, Configs)
		if err != nil {
			return err
		}
		if Debug {
			log.Printf("Load: %s", ConfigsFile)
		}

	} else {
		// Create a file with default values
		buf, err := yaml.Marshal(Configs)
		if err != nil {
			return err
		}
		err = ioutil.WriteFile(ConfigsFile, buf, 0644)
		if err != nil {
			return err
		}
		if Debug {
			log.Printf("Save: %s", ConfigsFile)
		}
	}

	if c, err := GetDir(DefaultStatesFile); err != nil {
		log.Fatal(err)
	} else {
		statesFile = c
	}

	if _, err := os.Stat(statesFile); !os.IsNotExist(err) {
		// Read if there is a state file
		if err := LoadStates(); err != nil {
			return err
		}
	}

	return nil
}

// LoadStates Load StatesFile
func LoadStates() error {
	buf, err := ioutil.ReadFile(statesFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, States)
	if err != nil {
		return err
	}
	if Debug {
		log.Printf("Load: %s", statesFile)
	}
	return nil
}

// SaveStates Save StatesFile
func SaveStates() error {
	buf, err := json.MarshalIndent(States, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(statesFile, buf, 0644)
	if err != nil {
		return err
	}
	if Debug {
		log.Printf("Save: %s", statesFile)
	}
	return nil
}

