package controller

/*
// Create
func save[T ml.IEntity](c IController, w RW, item T, name string) (err error) {
	item.SetId(0)
	if err = c.Db().Create(&item).Error; err != nil {
		hp.Error(w, "Failed to save "+name, hp.StatusInternalServerError)
		log.Printf("Failed to save %s: %v", name, err)
		return
	}
	log.Printf("Saved %s (%v)", name, item)
	return
}

// Read
func find[T any](c IController, w RW, items *[]T, name string) (err error) {
	if err = c.Db().Find(&items).Error; err != nil {
		hp.Error(w, "Failed to retrieve items", hp.StatusInternalServerError)
		log.Printf("Failed to retrieve items: %v", err)
		return
	}
	log.Printf("Found %d %ss", len(*items), name)
	return
}
*/
