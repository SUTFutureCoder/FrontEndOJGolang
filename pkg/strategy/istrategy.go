package startegy

const (
	WsJudgerResultCallBack = "JudgerResultCallBack"
	WsGetSubmitListByLabId = "SubmitListByLabId"
	WsTicker = "Ticker"
)

type istrategy interface {
	execute(v ...interface{})
}

type strategyDto struct {
	StraType string `json:"type"`
	Context  interface{} `json:"-"`
	Data     interface{} `jsson:"data"`
}

func strategyFactory(strategyType string, strategy *Strategy) istrategy {
	r := strategyDto{
		StraType: strategyType,
		Context: strategy.Context,
	}
	switch strategyType {
	case WsJudgerResultCallBack:
		return &result{r}
	case WsTicker:
		return &ticker{r}
	case WsGetSubmitListByLabId:
		return &submitList{r}
	}
	return nil
}