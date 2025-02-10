package ports

type DbPort interface {
	CreateOrder(id string) (string, error)
	UpdateOrder(id string) (string, error)
}
