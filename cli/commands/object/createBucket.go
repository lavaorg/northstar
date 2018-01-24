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
	"github.com/verizonlabs/northstar/cli/commands"
	"github.com/verizonlabs/northstar/cli/util"
	"github.com/verizonlabs/northstar/object/client"
	"github.com/verizonlabs/northstar/object/model"
)

type CreateBucketCMD struct {
	client *client.ObjectClient
	cmd    *flag.FlagSet
	name   *string
}

func NewCreateBucket(client *client.ObjectClient) commands.Command {
	cmd := flag.NewFlagSet("object-bucket-create", flag.ExitOnError)
	name := cmd.String("name", "testbucket", "The bucket name")
	return &CreateBucketCMD{client: client, cmd: cmd, name: name}
}

func (createBucket *CreateBucketCMD) Run(args []string) error {
	createBucket.cmd.Parse(args)

	if !createBucket.cmd.Parsed() {
		return errors.New("Failed to parse cmd")
	}

	bucket := &model.Bucket{Name: *createBucket.name}
	_, err := createBucket.client.CreateBucket(util.GetAccountID(), bucket)
	if err != nil {
		return err
	}

	fmt.Printf("Bucket %s created\n", bucket.Name)
	return nil
}
