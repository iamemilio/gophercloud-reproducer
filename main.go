package main

import (
	"fmt"

	"github.com/gophercloud/utils/openstack/clientconfig"
)

func main() {
	opts := clientconfig.ClientOpts{}

	cloud, err := clientconfig.GetCloudFromYAML(&opts)
	if err != nil {
		fmt.Printf("Error: %v", err)
	} else {
		fmt.Printf("Cloud:\n%+v\n", cloud)
	}
}
