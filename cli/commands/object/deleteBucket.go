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

package object

import (
	"flag"

	"errors"
	"fmt"
	"github.com/lavaorg/northstar/cli/commands"
	"github.com/lavaorg/northstar/cli/util"
	"github.com/lavaorg/northstar/object/client"
)

type DeleteBucketCMD struct {
	client *client.ObjectClient
	cmd    *flag.FlagSet
	name   *string
}

func NewDeleteBucket(client *client.ObjectClient) commands.Command {
	cmd := flag.NewFlagSet("object-bucket-delete", flag.ExitOnError)
	name := cmd.String("name", "testbucket", "The bucket name")

	return &DeleteBucketCMD{client: client,
		cmd:  cmd,
		name: name}
}

func (deleteBucket *DeleteBucketCMD) Run(args []string) error {
	deleteBucket.cmd.Parse(args)

	if !deleteBucket.cmd.Parsed() {
		return errors.New("Failed to parse cmd")
	}

	err := deleteBucket.client.DeleteBucket(util.GetAccountID(), *deleteBucket.name)
	if err != nil {
		return err
	}

	fmt.Printf("Bucket %s deleted\n", *deleteBucket.name)
	return nil
}
