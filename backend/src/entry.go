package main

import "persistence"

func main() {
	persistence.CreateDatabase()
	persistence.GetCustomers()
	persistence.GetFeatures()
	persistence.GetCustomerFeatures(2)
	persistence.UpdateFeature(1, true)
}
