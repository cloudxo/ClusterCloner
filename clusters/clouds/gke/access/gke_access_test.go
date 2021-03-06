package access

import (
	"clustercloner/clusters/machinetypes"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestParseMachineType(t *testing.T) {
	machineType := "e2-highcpu-8"
	mt, err := GetMachineTypes().Get(machineType)
	assert.Nil(t, err)
	if mt.Name != machineType {
		t.Fatal(mt.Name)
	}
	if mt.CPU != 8 {
		t.Fatal(mt.CPU)

	}
	if mt.RAMMB != 8000 {
		t.Fatal(mt.RAMMB)
	}
}
func TestParseMachineType2(t *testing.T) {
	name := "n1-ultramem-40"
	mt, err := GetMachineTypes().Get(name)
	assert.Nil(t, err)
	if mt.Name != name {
		t.Fatal(mt.Name)
	}
	if mt.CPU != 40 {
		t.Fatal(mt.CPU)
	}
	if mt.RAMMB != 961000 {
		t.Fatal(mt.RAMMB)
	}
}
func TestParseMissingMachineType2(t *testing.T) {
	name := "xx-xx-40"
	mt, err := gkeMachineTypes.Get(name)
	assert.NotNil(t, err)
	zero := machinetypes.MachineType{}
	if mt != zero {
		t.Fatalf("expect failure with %s", mt.Name)
	}
}
