package linen

type Config struct {
    sourcedir string
    destdir string
    baseurl string
    params map[string]string
}

func NewConfig(configfile string) (Config, error) {
    config := Config{}
    return config,nil
}
