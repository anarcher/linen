package linen

import "testing"

func TestConfigParse(t *testing.T) {
    c,err := NewConfig("./linen.example.yaml")
    if err != nil {
        t.Error(err)
    }
    if c.SourceDir != "./" {
        t.Errorf("c.Sourcedir don't matched %s,%s",c.SourceDir,"./")
    }
    if c.Params["subtitle"] == "" {
        t.Errorf("c.Params.subtitle is wrong. %s",c.Params)
    }
}
