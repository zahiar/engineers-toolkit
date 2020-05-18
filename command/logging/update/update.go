package update

import (
	grayloga "../../../client/graylog"
	"../../../input"
	"context"
	"fmt"
	"github.com/suzuki-shunsuke/go-graylog/graylog/graylog"
	"github.com/urfave/cli"
	"os"
)

func execute(streamId string, streamName string, streamDescription string) string {
	ctx := context.Background()

	graylogClient, err := grayloga.GetClient()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	stream := graylog.Stream{
		ID:          streamId,
		Title:       streamName,
		Description: streamDescription,
	}
	if _, err := graylogClient.UpdateStream(ctx, &stream); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	return fmt.Sprintf("http://localhost:9000/streams/%s", stream.ID)
}

func GetCommand() *cli.Command {
	return &cli.Command{
		Name:  "update",
		Usage: "Update a Graylog Stream",
		Action: func(c *cli.Context) error {
			streamId := input.GetInputAsString("Stream to update ID (E.G. 5ec1cb602ab79c0012c7da54)")
			newStreamName := input.GetInputAsString("Stream Name")
			newStreamDescription := input.GetInputAsString("Stream Description")

			create := input.GetInputAsBool(">> Update Graylog Stream? (yes/no)", "yes")
			if create == true {
				created := execute(streamId, newStreamName, newStreamDescription)
				fmt.Printf("Updated: %s\n", created)
			} else {
				fmt.Println("Cancelled updating Graylog Stream")
			}

			return nil
		},
	}
}
