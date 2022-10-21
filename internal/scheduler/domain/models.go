package domain

type Scheduler struct {
	OrderName   string `firestore:"order_name"`
	KitchenName string `firestore:"kitchen_name"`
	DroneName   string `firestore:"drone_name"`
}
