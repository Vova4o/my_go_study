//--Summary:
//  Create a program to display server status. The server statuses are
//  defined as constants, and the servers are represented as strings
//  in the `servers` slice.
//
//--Requirements:

package main

import "fmt"

const (
	Online      = 0
	Offline     = 1
	Maintenance = 2
	Retired     = 3
)

//* Create a function to print server status displaying:
//  - number of servers
//  - number of servers for each status (Online, Offline, Maintenance, Retired)

func printServerStatus(prop map[string]int) {
	fmt.Println("There are", len(prop), "servers.")
	// for key, val := range prop {
	// 	fmt.Println("Server:", key, "status:", val)
	// }

	stats := make(map[int]int)
	for _, status := range prop {
		switch status {
		case Online:
			stats[Online] += 1
		case Offline:
			stats[Offline] += 1
		case Maintenance:
			stats[Maintenance] += 1
		case Retired:
			stats[Retired] += 1
		default:
			panic("unhadled server status")
		}
	}
	fmt.Println(stats[Online], "servers are online")
	fmt.Println(stats[Offline], "servers are offline")
	fmt.Println(stats[Maintenance], "servers are undergoing maintanace")
	fmt.Println(stats[Retired], "servers are retired")

}

func globalStatus(element []string, status int, prop map[string]int) {
	for _, key := range element {
		prop[key] = status
	}
}

func main() {
	servers := []string{"darkstar", "aiur", "omicron", "w359", "baseline"}

	//* Create a map using the server names as the key and the server status
	//  as the value
	//* Set all of the server statuses to `Online` when creating the map
	serverMap := make(map[string]int)

	// for _, val := range servers {
	// 	serverMap[val] = Online
	// }
	globalStatus(servers, Online, serverMap)
	fmt.Println(serverMap)

	//* After creating the map, perform the following actions:
	//  - call display server info function
	fmt.Println("\nInitial state:")
	printServerStatus(serverMap)

	//  - change server status of `darkstar` to `Retired`
	//  - change server status of `aiur` to `Offline`
	serverMap["darkstar"] = Retired
	serverMap["aiur"] = Offline

	//  - call display server info function
	fmt.Println("\nSome changes:")

	printServerStatus(serverMap)

	//  - change server status of all servers to `Maintenance`
	//  - call display server info function

	globalStatus(servers, Maintenance, serverMap)
	fmt.Println("\nDid few ajustments:")
	printServerStatus(serverMap)

	// serverMap["aiur"] = 6
	// fmt.Println("\nUnhandled status:")
	// printServerStatus(serverMap)

}
