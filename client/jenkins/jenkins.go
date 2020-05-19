package jenkins

import (
	"encoding/xml"
	"github.com/bndr/gojenkins"
)

type SCM struct {
	Class string `xml:"class,attr"`
}
type Triggers struct {
	Class string `xml:"class,attr"`
}

type JobConfig struct {
	XMLName                          xml.Name `xml:"project"`
	Actions                          string   `xml:"actions"`
	Description                      string   `xml:"description"`
	KeepDependencies                 bool     `xml:"keepDependencies"`
	Properties                       string   `xml:"properties"`
	SCM                              SCM      `xml:"scm"`
	CanRoam                          bool     `xml:"canRoam"`
	Disabled                         bool     `xml:"disabled"`
	BlockBuildWhenDownstreamBuilding bool     `xml:"blockBuildWhenDownstreamBuilding"`
	BlockBuildWhenUpstreamBuilding   bool     `xml:"blockBuildWhenUpstreamBuilding"`
	Triggers                         Triggers `xml:"triggers"`
	ConcurrentBuild                  bool     `xml:"concurrentBuild"`
	Builders                         string   `xml:"builders"`
	Publishers                       string   `xml:"publishers"`
	BuildWrappers                    string   `xml:"buildWrappers"`
}

func MarshallJobConfigToXML(config JobConfig) (string, error) {
	configByteArray, err := xml.Marshal(config)
	return xml.Header + string(configByteArray), err
}

func GetClient() (*gojenkins.Jenkins, error) {
	jenkins := gojenkins.CreateJenkins(
		nil,
		"http://localhost:8080",
		"admin",
		"admin",
	)
	_, err := jenkins.Init()

	return jenkins, err
}
