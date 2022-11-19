package main

import (
	"fmt"

	"github.com/cedricblondeau/world-cup-2022-cli-dashboard/data/worldcupjson"
)

func main() {
	client := worldcupjson.NewClient()
	fmt.Println(client.Matches())
	fmt.Println(client.GroupTables())
}
