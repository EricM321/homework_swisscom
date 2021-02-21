package main

import "persistence"

func main() {
	persistence.CreateDatabase()
	persistence.GetCustomers()
	persistence.GetFeatures()
	persistence.GetCustomerFeatures()
	persistence.UpdateCustomer(1, "project_manager")
	persistence.UpdateFeature(1, true)
}
