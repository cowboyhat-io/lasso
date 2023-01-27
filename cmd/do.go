package cmd

import (
	"context"
	"fmt"
	"github.com/google/uuid"
	"github.com/microsoft/azure-devops-go-api/azuredevops"
	"github.com/microsoft/azure-devops-go-api/azuredevops/workitemtracking"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
	"strings"
)

var (
	kind string
)

// doCmd represents some talking with ado api
var doCmd = &cobra.Command{
	Use:   "do",
	Short: "maybe make some NFRs idk do things in ADO amirite?",
	Long:  ``,
	Run: func(cmd *cobra.Command, args []string) {
		// Set up connection to ADO org
		connection := azuredevops.NewPatConnection(viper.GetString("ORG_URL"), viper.GetString("PERSONAL_ACCESS_TOKEN"))

		ctx := context.Background()
		workItemClient, err := workitemtracking.NewClient(ctx, connection)
		if err != nil {
			log.Fatalf("Err %v", err)
		}
		// set up arg
		if strings.ToLower(kind) == "auto" {
			queryID, err := uuid.Parse(viper.GetString("MOVERS_LEAVERS_QUERY_ID"))
			if err != nil {
				fmt.Printf("ERR %v", err)
			}
			// replace _ with var to do something with
			_, err = workItemClient.QueryById(ctx, workitemtracking.QueryByIdArgs{&queryID, nil, nil, nil, nil})
			if err != nil {
				log.Fatalf("ERROR: %v", err)
			}

		}

	},
}

func init() {
	rootCmd.AddCommand(doCmd)
	doCmd.Flags().StringVar(&kind, "kind", "manual", "Runs either the manual or auto report for movers and leavers for TecdunderMifflinulse")
}
