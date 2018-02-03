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

package datasets

import (
	"flag"

	"errors"
	"fmt"
	"github.com/lavaorg/northstar/cli/commands"
	"github.com/lavaorg/northstar/cli/util"
	"github.com/lavaorg/northstar/data/datasets/client"
)

type DeleteDatasetCmd struct {
	client *client.DatasetsClient
	cmd    *flag.FlagSet
	id     *string
}

func NewDeleteDataset(client *client.DatasetsClient) commands.Command {
	cmd := flag.NewFlagSet("datasets-delete", flag.ExitOnError)
	id := cmd.String("id", "", "The dataset id")

	return &DeleteDatasetCmd{client: client, cmd: cmd, id: id}
}

func (d *DeleteDatasetCmd) Run(args []string) error {
	d.cmd.Parse(args)

	if !d.cmd.Parsed() {
		return errors.New("Failed to parse cmd")
	}

	if *d.id == "" {
		return errors.New("Please set an id using -id.")
	}

	err := d.client.DeleteDataset(util.GetAccountID(), *d.id)
	if err != nil {
		return err
	}

	fmt.Println("Dataset deleted")
	return nil
}
