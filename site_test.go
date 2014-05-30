package linen

import "testing"

func TestNewSite(t *testing.T) {
    config,err := NewConfig("./linen.example.yaml")
    if err != nil {
        t.Error(err)
    }

    site,err := NewSite(config)
    if err != nil {
        t.Error(site,err)
    }
}
