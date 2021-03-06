package main

import (
	"testing"

	"github.com/anz-bank/sysl/pkg/parse"
	"github.com/sirupsen/logrus"
	"github.com/spf13/afero"
	"github.com/stretchr/testify/assert"

	"github.com/anz-bank/sysl/pkg/datamodeldiagram"
)

func TestDataModel(t *testing.T) {
	t.Parallel()
	p := &datamodelCmd{}
	p.Output = "whatever.svg"
	p.Project = "Project"
	p.ClassFormat = "%(classname)"
	p.Direct = false
	fs := afero.NewOsFs()
	filename := testDir + "multiple-app-datamodel.sysl"
	m, err := parse.NewParser().Parse(filename, fs)
	if err != nil {
		panic(err)
	}
	outmap, err := datamodeldiagram.GenerateDataModels(&p.CmdContextParamDatagen, m, logrus.New())
	assert.Nil(t, err, "Generating the data diagrams failed")
	err = p.GenerateFromMap(outmap, fs)
	assert.Nil(t, err, "Generating the data diagrams failed")

	expected := map[string]string{"whatever.svg": `@startuml
	''''''''''''''''''''''''''''''''''''''''''
	''                                      ''
	''  AUTOGENERATED CODE -- DO NOT EDIT!  ''
	''                                      ''
	''''''''''''''''''''''''''''''''''''''''''
	
	class "App.User" as _0 << (D,orchid) >> {
	+ beep : **Server.ftgyhb**
	+ weuiyfgwihe : **AnotherApp.AnotherType**
	+ whatever : int
	}
	class "Server.ftgyhb" as _1 << (D,orchid) >> {
	+ blah : int
	}
	_0 *-- "1..1 " _1
	@enduml
`}
	assert.Equal(t, outmap, expected)
}
