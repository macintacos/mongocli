// Copyright 2020 MongoDB Inc
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package users

import (
	"errors"

	"github.com/mongodb/mongocli/internal/cli"
	"github.com/mongodb/mongocli/internal/config"
	"github.com/mongodb/mongocli/internal/flag"
	"github.com/mongodb/mongocli/internal/store"
	"github.com/mongodb/mongocli/internal/usage"
	"github.com/spf13/cobra"
)

const listTemplate = `ID	FIRST NAME	LAST NAME	USERNAME
{{.ID}}	{{.FirstName}}	{{.LastName}}	{{.Username}}
`

type DescribeOpts struct {
	cli.GlobalOpts
	cli.OutputOpts
	store    store.UsersDescriber
	username string
	ID       string
}

func (opts *DescribeOpts) init() error {
	var err error
	opts.store, err = store.New(config.Default())
	return err
}

func (opts *DescribeOpts) Run() error {
	var r interface{}
	var err error

	if opts.username != "" {
		r, err = opts.store.UserByName(opts.username)
	}

	if opts.ID != "" {
		r, err = opts.store.UserByID(opts.ID)
	}

	if err != nil {
		return err
	}

	return opts.Print(r)
}

func (opts *DescribeOpts) validate() error {
	if opts.ID == "" && opts.username == "" {
		return errors.New("must supply one of 'ID' or 'username'")
	}

	if opts.ID != "" && opts.username != "" {
		return errors.New("cannot supply both 'ID' and 'username'")
	}

	return nil
}

// mongocli iam project(s) user(s) describe --id ID --username USERNAME
func DescribeBuilder() *cobra.Command {
	opts := &DescribeOpts{}
	cmd := &cobra.Command{
		Use:     "describe",
		Aliases: []string{"get"},
		Short:   describeIAMUser,
		PreRunE: func(cmd *cobra.Command, args []string) error {
			return opts.PreRunE(
				opts.init,
				opts.InitOutput(cmd.OutOrStdout(), listTemplate),
				opts.validate,
			)
		},
		RunE: func(cmd *cobra.Command, args []string) error {
			return opts.Run()
		},
	}

	cmd.Flags().StringVar(&opts.username, flag.Username, "", usage.Username)
	cmd.Flags().StringVar(&opts.ID, flag.ID, "", usage.UserID)

	cmd.Flags().StringVarP(&opts.Output, flag.Output, flag.OutputShort, "", usage.FormatOut)

	return cmd
}
