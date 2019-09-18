package confstate

import (
	"encoding/json"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"github.com/mitchellh/go-homedir"
)

const (
	DBTWork = iota // DefaultBaseDir is WorkDir
	DBTHome        // DefaultBaseDir is HomeDir
	DBTBin         // DefaultBaseDir is Offset from excutable binary
)

var (
	Configs            interface{}   // Configs data
	States             interface{}   // States data
	BaseDir            string        // Base directory
	ConfigsFile        string        // ConfigsFile Filename
	StatesFile         string        // StatesFile Filename
	OffsetFromBin      string        // Base directory offset from executable binary
	DefaultConfigsFile string        // Default filename ConfigFile (relative path)
	DefaultStatesFile  string        // Default filename StatesFile (relative path)
	DefaultBaseDirType int = DBTBin  // Basedir get offset from executable binary
    NewConfigsFile     bool = false  // Create New ConfigsFile
    NewStatesFile      bool = false  // Create New StatesFile
	Debug              bool = false  // Debug Mode
)

// GetDir Get absorute directory from excutable binary
func GetDir(p string) (string, error) {
	return filepath.Abs(filepath.Join(BaseDir, p))
}

// LoadConfigs Initialize and Load
func LoadConfigs() error {

	// set BaseDir
	if BaseDir == "" {
		switch DefaultBaseDirType {
		case DBTBin:
			// Get relative directory from executable binary
			exe, err := os.Executable()
			if err != nil {
				log.Fatal(err)
			}
			b, err := filepath.Abs(filepath.Join(filepath.Dir(exe), OffsetFromBin))
			if err != nil {
				log.Fatal(err)
			}
			BaseDir = b
		case DBTHome:
			b,err := homedir.Dir()
			if err != nil {
				log.Fatal(err)
			}
			BaseDir = b

		case DBTWork:
			b,err := os.Getwd()
			if err != nil {
				log.Fatal(err)
			}
			BaseDir = b
		}
	}

	// Set default value if ConfigsFile is empty
	if ConfigsFile == "" {
		if c, err := GetDir(DefaultConfigsFile); err != nil {
			return err
		} else {
			ConfigsFile = c
		}
	}

	// create directory
	if err := createPath(ConfigsFile); err != nil {
		return err
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
		NewConfigsFile = true
		if Debug {
			log.Printf("Save: %s", ConfigsFile)
		}
	}

	return nil
}

// LoadStates Load StatesFile
func LoadStates() error {

	if err := createPath(StatesFile); err != nil {
		return err
	}

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

	fileExist := true
	if _, err := os.Stat(StatesFile); os.IsNotExist(err) {
		fileExist = false
	}
	buf, err := json.MarshalIndent(States, "", "\t")
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(StatesFile, buf, 0644)
	if err != nil {
		return err
	}
	if !fileExist {
		NewStatesFile = true
	}
	if Debug {
		log.Printf("Save: %s", StatesFile)
	}
	return nil
}


// create directory if not exists
func createPath(n string) error {
	d := filepath.Dir(n)
	if _, err := os.Stat(d); os.IsNotExist(err) {
		if err := os.MkdirAll(d,0755); err != nil {
			return err
		}
	}
	return nil
}

