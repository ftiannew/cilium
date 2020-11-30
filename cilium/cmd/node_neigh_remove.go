// Copyright 2020 Authors of Cilium
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package cmd

import (
	"net"

	"github.com/cilium/cilium/api/v1/client/daemon"
	"github.com/cilium/cilium/api/v1/models"

	"github.com/spf13/cobra"
)

var nodeNeighRemoveCmd = &cobra.Command{
	Use:     "remove <neigh name> <neigh IP>",
	Aliases: []string{"rm"},
	Short:   "Remove node as a neighbor from current node's neighbor table",
	Example: `cilium node neigh remove "node1" 10.10.10.10`,
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) != 2 {
			Usagef(cmd, "Missing node name and/or node IP")
		}

		if len(args[0]) == 0 {
			Fatalf("Invalid node name, cannot be empty\n")
		}

		var ip net.IP
		if ip = net.ParseIP(args[1]); ip == nil {
			Fatalf("Invalid IP address %q\n", args[1])
		}

		if _, err := client.Daemon.DeleteClusterNodesNeigh(
			daemon.NewDeleteClusterNodesNeighParams().WithRequest(&models.NodeNeighRequest{
				Name: args[0],
				IP:   args[1],
			}),
		); err != nil {
			Fatalf("Unable to remove %q from neighbor table: %v\n", ip, err)
		}
	},
}

func init() {
	nodeNeighCmd.AddCommand(nodeNeighRemoveCmd)
}
