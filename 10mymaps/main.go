package main

import "fmt"

func main() {
	fmt.Println("This is example of Maps")
	cloud_provider := make(map[string]string)

	cloud_provider["AWS"] = "Amazon Web Services"
	cloud_provider["AZ"] = "Azure"
	cloud_provider["GCP"] = "Google Cloud Provider"

	fmt.Println(" Cloud providers are ", cloud_provider)

	//delete element from map
	delete(cloud_provider, "AZ")

	// Looping over map
	for key, value := range cloud_provider {
		fmt.Printf("%v stands for %v\n", key, value)
	}
}
