package linen

import (
//    "github.com/spacemonkeygo/errors"
//    "github.com/stretchr/powerwalk"
)

type Site struct {
    config Config
}

func NewSite(config Config) (Site, error) {
    site := Site{config: config}
	return site, nil
}
