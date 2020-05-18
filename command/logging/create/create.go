package create

import (
	grayloga "../../../client/graylog"
	"../../../input"
	"context"
	"fmt"
	"github.com/suzuki-shunsuke/go-graylog/graylog/graylog"
	"github.com/urfave/cli"
	"os"
)

const DEFAULT_GRAYLOG_INDEXSET_PREFIX = "graylog"

func execute(streamName string, streamDescription string, streamIndexSetPrefix string) string {
	ctx := context.Background()

	graylogClient, err := grayloga.GetClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	indexSets, _, _, _, err := graylogClient.GetIndexSets(ctx, 0, 0, false)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	indexSetId := ""
	for _, indexSet := range indexSets {
		if indexSet.IndexPrefix == streamIndexSetPrefix {
			indexSetId = indexSet.ID
			break
		}
	}

	if indexSetId == "" {
		fmt.Sprintf("Unable to find Graylog IndexSet with IndexPrefix: %s\n", streamIndexSetPrefix)
		os.Exit(1)
	}

	stream := graylog.Stream{
		Title:       streamName,
		Description: streamDescription,
		IndexSetID:  indexSetId,
		Disabled:    false,
	}
	if _, err := graylogClient.CreateStream(ctx, &stream); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return fmt.Sprintf("http://localhost:9000/streams/%s", stream.ID)
}

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "create",
		Usage: "Create a new Graylog Stream",
		Action: func(c *cli.Context) error {
			streamName := input.GetInputAsString("Stream Name")
			streamDescription := input.GetInputAsString("Stream Description")
			streamIndexSetPrefix := input.GetInputAsStringWithDefaultValue("IndexSetPrefix this stream will belong to", DEFAULT_GRAYLOG_INDEXSET_PREFIX)

			create := input.GetInputAsBool(">> Create Graylog Stream? (yes/no)", "yes")
			if create == true {
				created := execute(streamName, streamDescription, streamIndexSetPrefix)
				fmt.Printf("Created: %s\n", created)
			} else {
				fmt.Println("Cancelled creating Graylog Stream")
			}

			return nil
		},
	}
}
