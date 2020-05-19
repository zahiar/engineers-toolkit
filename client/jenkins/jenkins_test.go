package jenkins

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMarshallJobConfigToXML(t *testing.T) {
	jobConfig := JobConfig{
		KeepDependencies: false,
		SCM: SCM{
			Class: "hudson.scm.NullSCM",
		},
		CanRoam:                          true,
		Disabled:                         false,
		BlockBuildWhenDownstreamBuilding: false,
		BlockBuildWhenUpstreamBuilding:   false,
		Triggers: Triggers{
			Class: "vector",
		},
		ConcurrentBuild: false,
	}
	resultXml, _ := MarshallJobConfigToXML(jobConfig)

	expectedXml := "<?xml version=\"1.0\" encoding=\"UTF-8\"?>\n<project><actions></actions><description></description>" +
		"<keepDependencies>false</keepDependencies><properties></properties><scm class=\"hudson.scm.NullSCM\"></scm>" +
		"<canRoam>true</canRoam><disabled>false</disabled><blockBuildWhenDownstreamBuilding>false</blockBuildWhenDownstreamBuilding>" +
		"<blockBuildWhenUpstreamBuilding>false</blockBuildWhenUpstreamBuilding><triggers class=\"vector\"></triggers>" +
		"<concurrentBuild>false</concurrentBuild><builders></builders><publishers></publishers><buildWrappers></buildWrappers></project>"
	assert.Equal(t, expectedXml, resultXml)
}

func TestGetClient(t *testing.T) {
	client, _ := GetClient()

	assert.Equal(t, "http://localhost:8080", client.Server)
}
