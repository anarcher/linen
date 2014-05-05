package linen

import "fmt"
import "os"
import "path"
import "encoding/json"
import "gopkg.in/yaml.v1"
import "github.com/BurntSushi/toml"
import "io/ioutil"

type Config struct {
    ConfigFile string
    SourceDir string
    DestDir string
    BaseUrl string
    Title   string
    Params map[string]string
}

func NewConfig(configfile string) (Config, error) {
    config := Config{
                ConfigFile: configfile,
                SourceDir: ".",
                DestDir : "public/",
                Params : make(map[string]string),
    }
    config.ReadConfigFile()
    return config,nil
}

func (c *Config) ReadConfigFile() {
    file, err := ioutil.ReadFile(c.ConfigFile)
    if err == nil {
        switch path.Ext(c.ConfigFile) {
        case ".yaml":
            if err := yaml.Unmarshal(file,&c); err != nil {
                fmt.Printf("Error parse file: %s",err)
                os.Exit(1)
            }
        case ".json":
            if err := json.Unmarshal(file,&c); err != nil {
                fmt.Printf("Error parse file: %s",err)
                os.Exit(1)
            }
        case ".toml":
            if _,err := toml.Decode(string(file),&c); err != nil {
                fmt.Printf("Error parse file: %s",err)
                os.Exit(1)
            }
        }
    }
}
