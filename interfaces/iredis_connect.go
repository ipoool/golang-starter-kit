package interfaces

// IRedis - Redis interfaces
type IRedis interface {
	Connect()
	Ping() bool
}
