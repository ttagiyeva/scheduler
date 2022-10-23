package domain

type Scheduler struct {
	DocumentId  string `firestore:"document_id"`
	OrderName   string `firestore:"order_name"`
	KitchenName string `firestore:"kitchen_name"`
	DroneName   string `firestore:"drone_name"`
}
