package customer

func DeleteCustomerById(ids ...uint) {
	for _, data := range ids {
		storageCustomer.DeleteCustomerById(data)
	}
}
