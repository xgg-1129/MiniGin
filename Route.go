package MiniGin

type Route struct {
	handleFunMap map[string]handleFun
}

func getRoute() *Route{
	return &Route{handleFunMap: make(map[string]handleFun)}
}
