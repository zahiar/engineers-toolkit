package create

import (
	"../../../client/jenkins"
	"../../../input"
	"fmt"
	"github.com/urfave/cli"
	"os"
	"strings"
)

func execute(projectName string) string {
	projectCode := strings.ReplaceAll(strings.ToLower(projectName), " ", "-")

	jenkinsClient, err := jenkins.GetClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	jobConfig := jenkins.JobConfig{
		KeepDependencies: false,
		SCM: jenkins.SCM{
			Class: "hudson.scm.NullSCM",
		},
		CanRoam:                          true,
		Disabled:                         false,
		BlockBuildWhenDownstreamBuilding: false,
		BlockBuildWhenUpstreamBuilding:   false,
		Triggers: jenkins.Triggers{
			Class: "vector",
		},
		ConcurrentBuild: false,
	}
	jobConfigXml, _ := jenkins.MarshallJobConfigToXML(jobConfig)

	jenkinsClient.CreateFolder(projectCode)
	_, err = jenkinsClient.CreateJobInFolder(jobConfigXml, "new-job-name", projectCode)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return fmt.Sprintf("http://localhost:8080/job/%s", projectCode)
}

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create a new Jenkins Pipeline",
		Action: func(c *cli.Context) error {
			projectName := input.GetInputAsString("Project Name (E.G. My App)")

			create := input.GetInputAsBool(">> Create Jenkins Pipeline? (yes/no)", "yes")
			if create == true {
				createdPipeline := execute(projectName)

				fmt.Printf("Created: %s\n", createdPipeline)
			} else {
				fmt.Println("Cancelled creating Jenkins Pipeline")
			}

			return nil
		},
	}
}
