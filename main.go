package main

import (
	"fmt"
	"friendly-potato/integrations"
	"log"
)

func main() {
	fmt.Println("teste")
	err := integrations.InitAPI("IQjVxP6VgH67usDcznZ6O6NGtOyJdBvyood9WMvS")

	if err != nil {
		log.Fatalf("Fatal on auth %v", err)
	}

	zones, err := integrations.ListZones()
	fmt.Printf("%v", zones)
	if err != nil {
		log.Fatal(err)
	}

	zones, err = integrations.ListZones()
	if err != nil {
		log.Fatal(err)
	}

	for _, z := range zones {
		fmt.Printf("zone: %v id: %v\n", z.Name, z.Id)
	}

	newZone := integrations.Zone{"idtest", " tutorialedge.net", []integrations.Record{}}
	zone, err := integrations.CreateZone(newZone)
	if err != nil {
		log.Fatal(err)
	}

	err = integrations.DeleteZone(zone)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("%v\n", zone)

}
