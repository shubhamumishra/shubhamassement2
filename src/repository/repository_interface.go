package repository

/***This interface is common for all the repository files****/
type MysqlRepository interface {
	PlaceOrder(req interface{}) error
	CancelOrder(orderId int) error
	StoreItem(req interface{}) error
}
