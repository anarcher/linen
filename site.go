package linen

import (
    "os"
    "fmt"
//    "github.com/spacemonkeygo/errors"
    "github.com/stretchr/powerwalk"
)

type Site struct {
    config Config
}

func NewSite(config Config) (Site, error) {
    site := Site{config: config}
	return site, nil
}

func (site *Site) LoadPages() error {
    err := powerwalk.Walk(site.config.SourceDir,site.walkForFileLoad)
    if err != nil {
        return err
    }
    return nil
}

func (site *Site) walkForFileLoad(path string, info os.FileInfo, err error) error {
    fmt.Println(path)
    return nil
}
