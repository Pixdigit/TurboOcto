package turboOcto

import (
	"encoding/csv"
	"os"

	"github.com/pkg/errors"
	"gitlab.com/Pixdigit/simpleSerialization"
)

var conf map[string]string = map[string]string{}

func initializeEnvironment() (err error) {

	ok, err := pathExists("./conf");	if err != nil {return errors.Wrap(err, "could not check for configuration")}
	if ok {
		ok, err := pathExists("./conf/last.csv");	if err != nil {return errors.Wrap(err, "failed to check for last configuration")}
		if ok {
			err := LoadConf("last");	if err != nil {return errors.Wrap(err, "could not load last configuration")}
		} else {
			err := LoadDefaultConf();	if err != nil {return errors.Wrap(err, "could not initialize environment")}
		}
	} else {
		err := LoadDefaultConf();	if err != nil {return errors.Wrap(err, "could not initialize environment")}
	}

	return nil
}
func LoadDefaultConf() error {
	ok, err := pathExists("./conf/default.csv");	if err != nil {return errors.Wrap(err, "failed to check for default configuration")}
	if ok {
		err := LoadConf("default");	if err != nil {return errors.Wrap(err, "could not load default configuration")}
	} else {
		//Default Configuration
		err := AddConf("updateOnRefresh", true);	if err != nil {return errors.Wrap(err, "could not set default configuration")}
		err = AddConf("fullscreen", true);	if err != nil {return errors.Wrap(err, "could not set default configuration")}
		err = AddConf("allowFrameSkipping", true);	if err != nil {return errors.Wrap(err, "could not set default configuration")}
		err = AddConf("spriteTimerMode", USE_FRAME_COUNT);	if err != nil {return errors.Wrap(err, "could not set default configuration")}
	}
	return nil
}

func GetConf(confName string) (interface{}, error) {
	config, ok := conf[confName]
	if !ok {
		return nil, errors.New("configuration " + confName + " does not exist")
	}
	value, err := simpleSerialization.Deserialize(config);	if err != nil {return nil, errors.New("invalid configuration value \"" + config + "\"")}
	return value, nil
}

func SetConf(confName string, confValue interface{}) error {
	config, ok := conf[confName]
	if !ok {
		return errors.New("configuration " + confName + " does not exist")
	}
	newConfig, err := simpleSerialization.Serialize(confValue);	if err != nil {return errors.Wrap(err, "could not change configuration")}
	oldVarType, err := simpleSerialization.TypeOfSerialized(config);	if err != nil {return errors.Wrap(err, "could not check var type")}
	newVarType, err := simpleSerialization.TypeOfSerialized(config);	if err != nil {return errors.Wrap(err, "could not check var type")}

	if oldVarType != newVarType {
		return errors.New("configuration must have same type")
	}
	conf[confName] = newConfig
	return nil
}

func AddConf(confName string, initConfValue interface{}) error {
	newConfig, err := simpleSerialization.Serialize(initConfValue);	if err != nil {return errors.Wrap(err, "could not simpleSerialization.Serialize initial conf value")}
	conf[confName] = newConfig
	return nil
}

func DelConf(confName string) error {
	_, ok := conf[confName]
	if !ok {
		return errors.New("could not delete config: does not exist")
	}
	delete(conf, confName)
	return nil
}

func SaveConf(filename string) error {
	var data [][]string
	for k, v := range conf {
		configuration := []string{k, v}
		data = append(data, configuration)
	}

	ok, err := pathExists("./conf/");	if err != nil {return errors.Wrap(err, "could not check wether conf folder exists")}
	if !ok {
		err := os.Mkdir("./conf", os.ModePerm);	if err != nil {return errors.Wrap(err, "could not create conf folder")}
	}

	file, err := os.Create("./conf/" + filename + ".csv")
	defer file.Close();	if err != nil {return errors.Wrap(err, "could not open file "+filename+" to save configuration")}
	w := csv.NewWriter(file)
	err = w.WriteAll(data);	if err != nil {return errors.Wrap(err, "could not write configuration to file"+filename)}

	return nil
}

func LoadConf(filename string) error {
	file, err := os.Open("./conf/" + filename + ".csv")
	defer file.Close();	if err != nil {return errors.Wrap(err, "could not open file "+filename+" to load configuration")}
	r := csv.NewReader(file)
	r.FieldsPerRecord = 2
	data, err := r.ReadAll();	if err != nil {return errors.Wrap(err, "could not read configuration from file "+filename)}

	for _, configuration := range data {
		//test if value can be deserialized == valid value
		_, err := simpleSerialization.Deserialize(configuration[1]);	if err != nil {return errors.Wrap(err, "could not load conf file \""+filename+"\"")}
		conf[configuration[0]] = configuration[1]
	}

	return nil
}
