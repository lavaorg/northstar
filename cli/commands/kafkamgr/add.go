/*
Copyright (C) 2017 Verizon. All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package kafkamngr

import (
	"flag"

	"errors"
	"fmt"
	"github.com/lavaorg/northstar/cli/commands"
	"github.com/lavaorg/northstar/kafkamgr"
)

type AddTopicsCmd struct {
	client      *kafkamgr.KafkaMngrClient
	cmd         *flag.FlagSet
	name        *string
	partitions  *int
	replication *int
}

func NewAddTopics(client *kafkamgr.KafkaMngrClient) commands.Command {
	cmd := flag.NewFlagSet("topics-add", flag.ExitOnError)
	name := cmd.String("name", "test", "The topic name")
	partitions := cmd.Int("partitions", 1, "The number of partitions")
	replication := cmd.Int("replication", 3, "The replication factor")
	return &AddTopicsCmd{client: client,
		cmd:         cmd,
		name:        name,
		partitions:  partitions,
		replication: replication}
}

func (output *AddTopicsCmd) Run(args []string) error {
	output.cmd.Parse(args)

	if !output.cmd.Parsed() {
		return errors.New("Failed to parse cmd")
	}

	topic := &kafkamgr.Topic{Name: *output.name,
		Partitions:  *output.partitions,
		Replication: *output.replication}

	err := output.client.CreateTopic(topic)
	if err != nil {
		return err
	}

	fmt.Println("Topic created")
	return nil
}
