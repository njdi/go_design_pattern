package singleton

type EagerSingleton struct {
}

var eagerInstance = &EagerSingleton{}

func GetEagerInstance() *EagerSingleton {
	return eagerInstance
}
