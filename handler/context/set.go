/*
Copyright © 2019 Portworx

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
package context

import (
	"fmt"

	pxcmd "github.com/portworx/pxc/cmd"
	"github.com/portworx/pxc/pkg/commander"
	"github.com/portworx/pxc/pkg/contextconfig"
	"github.com/portworx/pxc/pkg/util"
	"github.com/spf13/cobra"
)

var contextSetCmd *cobra.Command

var _ = commander.RegisterCommandVar(func() {
	contextSetCmd = &cobra.Command{
		Use:     "set [NAME]",
		Aliases: []string{"use"},
		Example: "$ px context set mynewcontext",
		Short:   "Set the current context configuration",
		Args: func(cmd *cobra.Command, args []string) error {
			if len(args) < 1 {
				return fmt.Errorf("Must supply a name for context")
			}
			return nil
		},
		Long: ``,
		RunE: contextSetExec,
	}
})

var _ = commander.RegisterCommandInit(func() {
	pxcmd.ContextAddCommand(contextSetCmd)
})

func SetAddCommand(cmd *cobra.Command) {
	contextSetCmd.AddCommand(cmd)
}

func contextSetExec(cmd *cobra.Command, args []string) error {
	name := args[0]
	contextManager, err := contextconfig.NewContextManager(pxcmd.GetConfigFile())
	if err != nil {
		return err
	}

	if err := contextManager.SetCurrent(name); err != nil {
		return err
	}

	util.Printf("%s is now the current context", name)

	return nil
}
