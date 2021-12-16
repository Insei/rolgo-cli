package cmd

import (
	"encoding/json"
	"fmt"
	"github.com/Insei/rolgo"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	ProjectID    string
	RentName     string
	Model        string
	Manufacturer string
	RentId       string
	IpxeUrl      string
)

var rentalCmd = &cobra.Command{
	Use:   "rental",
	Short: "Manage device rentals",
	Long:  `Manage device rentals`,
}

var createRentCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a new device rental",
	Long:  `Create a new device rental`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := rolgo.NewClient()
		if err != nil {
			log.Fatalf(err.Error())
		}
		r := new(rolgo.DeviceRentCreateRequest)
		r.Model = Model
		r.Manufacturer = Manufacturer
		r.Name = RentName
		r.IpxeUrl = IpxeUrl
		rent, err := client.Rents.Create(ProjectID, r)
		if err == nil {
			fmt.Println(rent.Id)
		} else {
			log.Fatalf("unable to create device rent: %v", err)
		}

	},
}

var getRentCmd = &cobra.Command{
	Use:   "get",
	Short: "Get the device rental",
	Long:  `Get the device rental`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := rolgo.NewClient()
		if err != nil {
			log.Fatalf(err.Error())
		}
		rent, err := client.Rents.Get(ProjectID, RentId)
		if err != nil {
			log.Fatalf(err.Error())
		}
		rentJson, err := json.Marshal(rent)
		if err != nil {
			log.Fatalf(err.Error())
		}
		fmt.Println(string(rentJson))
	},
}

var closeRentCmd = &cobra.Command{
	Use:   "close",
	Short: "Close the device rental",
	Long:  `Close the device rental`,
	Run: func(cmd *cobra.Command, args []string) {
		client, err := rolgo.NewClient()
		if err != nil {
			log.Fatalf(err.Error())
		}
		err = client.Rents.Release(ProjectID, RentId)
		if err != nil {
			log.Fatalf(err.Error())
		}
	},
}

func init() {
	// rental -> create
	createRentCmd.Flags().StringVarP(&RentName, "name", "n", "", "rental name")
	createRentCmd.Flags().StringVar(&Model, "model", "", "device model")
	createRentCmd.Flags().StringVarP(&Manufacturer, "manufacturer", "m", "", "device manufacturer")
	createRentCmd.Flags().StringVarP(&ProjectID, "project-id", "p", "", "project id")
	createRentCmd.Flags().StringVarP(&IpxeUrl, "ipxe-cfg-url", "i", "", "url to IPXE cfg file")
	_ = createRentCmd.MarkFlagRequired("project-id")
	_ = createRentCmd.MarkFlagRequired("name")
	_ = createRentCmd.MarkFlagRequired("model")
	_ = createRentCmd.MarkFlagRequired("manufacturer")
	_ = createRentCmd.MarkFlagRequired("ipxe-cfg-url")
	rentalCmd.AddCommand(createRentCmd)
	// rental -> close
	closeRentCmd.Flags().StringVarP(&ProjectID, "project-id", "p", "", "project id")
	closeRentCmd.Flags().StringVarP(&RentId, "id", "i", "", "rental id")
	_ = closeRentCmd.MarkFlagRequired("id")
	_ = closeRentCmd.MarkFlagRequired("project-id")
	rentalCmd.AddCommand(closeRentCmd)
	// rental -> get
	getRentCmd.Flags().StringVarP(&ProjectID, "project-id", "p", "", "project id")
	getRentCmd.Flags().StringVarP(&RentId, "id", "i", "", "rental id")
	_ = getRentCmd.MarkFlagRequired("id")
	_ = getRentCmd.MarkFlagRequired("project-id")
	rentalCmd.AddCommand(getRentCmd)
	// rental
	rootCmd.AddCommand(rentalCmd)
}
