package startegy

type Strategy struct {
	Context interface{}
}

func (strategy *Strategy) ExecStrategy(s string, d interface{}) {
	if istrategy := strategyFactory(s, strategy); istrategy != nil {
		istrategy.execute(d)
		return
	}
}
