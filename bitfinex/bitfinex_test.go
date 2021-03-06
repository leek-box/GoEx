package bitfinex

import (
	"net/http"
	"testing"

	"github.com/leek-box/GoEx"
)

var bfx = New(http.DefaultClient, "", "")

func TestBitfinex_GetTicker(t *testing.T) {
	ticker, _ := bfx.GetTicker(goex.ETH_BTC)
	t.Log(ticker)
}

func TestBitfinex_GetDepth(t *testing.T) {
	dep, _ := bfx.GetDepth(2, goex.ETH_BTC)
	t.Log(dep.AskList)
	t.Log(dep.BidList)
}
