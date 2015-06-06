package model

import (
	"testing"
)

func TestSkillz(t *testing.T) {
	gopher := &Gopher{Skillz: []string{"Unicycle", "Topiary"}}
	gopher.Zap()
	if len(gopher.Skillz) != 0 {
		t.Errorf("Gopher should have no skillz after zap - %+v", gopher.Skillz)
	}
}

//This test doesnt check that Skillz was emptied
func TestKapow(t *testing.T) {
	gopher := &Gopher{Skillz: []string{"Unicycle", "Topiary"}}
	err := gopher.Kapow()
	if err != nil {
		t.Errorf("Kapow returned an error - %+v", err)
	}
}
