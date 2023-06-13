package backup

import (
	"github.com/filecoin-project/go-address"
)

const (
	clientPrivKey = "7b2254797065223a22736563703235366b31222c22507269766174654b6579223a223836493345562b652f7973466b6d61452f4d636346565a4930444e4b582b6449446246387475546938324d3d227d"
)

var (
	clientAddress = address.Address{}
)

func init() {
	clientAddress, _ = address.NewFromString("f1a6tcxpzz6cyp5geatbogmroy6cxcrstruhnp42y")
}
