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
	Configs            interface{}         // Configs data
	States             interface{}         // States data
	BaseDir            string              // Base directory
	ConfigsFile        string              // ConfigsFile Filename
	StatesFile         string              // StatesFile Filename
	OffsetFromBin      string              // Base directory offset from executable binary
	DefaultConfigsFile string              // Default filename ConfigFile
	DefaultStatesFile  string              // Default filename StatesFile
	Debug              bool        = false // Debug Mode
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

// LoadConfigs Initialize and Load
func LoadConfigs() error {

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

	return nil
}

// LoadStates Load StatesFile
func LoadStates() error {

	if StatesFile == "" {
		if c, err := GetDir(DefaultStatesFile); err != nil {
			log.Fatal(err)
		} else {
			StatesFile = c
		}
	}

	// Skip if StatesFile not exist
	if _, err := os.Stat(StatesFile); os.IsNotExist(err) {
		return nil
	}

	buf, err := ioutil.ReadFile(StatesFile)
	if err != nil {
		return err
	}
	err = json.Unmarshal(buf, States)
	if err != nil {
		return err
	}
	if Debug {
		log.Printf("Load: %s", StatesFile)
	}
	return nil
}

// SaveStates Save StatesFile
func SaveStates() error {
	buf, err := json.MarshalIndent(States, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(StatesFile, buf, 0644)
	if err != nil {
		return err
	}
	if Debug {
		log.Printf("Save: %s", StatesFile)
	}
	return nil
}
