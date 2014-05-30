package linen

import (
//    "github.com/spacemonkeygo/errors"
//    "github.com/stretchr/powerwalk"
)

type Site struct {
}

func NewSite(config Config) (Site, error) {
	site := Site{}
	return site, nil
}
