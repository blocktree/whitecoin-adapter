module github.com/blocktree/whitecoin-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.2
	github.com/blocktree/go-owaddress v1.1.11
	github.com/blocktree/go-owcdrivers v1.2.0
	github.com/blocktree/go-owcrypt v1.1.1
	github.com/blocktree/openwallet/v2 v2.0.4
	github.com/btcsuite/btcd v0.20.1-beta
	github.com/btcsuite/btcutil v0.0.0-20191219182022-e17c9730c422
	github.com/denkhaus/logging v0.0.0-20180714213349-14bfb935047c
	github.com/emirpasic/gods v1.12.0
	github.com/imroc/req v0.2.4
	github.com/juju/errors v0.0.0-20181118221551-089d3ea4e4d5
	github.com/pkg/errors v0.8.1
	github.com/pquerna/ffjson v0.0.0-20181028064349-e517b90714f7
	github.com/shopspring/decimal v0.0.0-20200105231215-408a2507e114
	github.com/stretchr/testify v1.4.0
	github.com/tidwall/gjson v1.3.5
	golang.org/x/crypto v0.0.0-20191227163750-53104e6ec876

)

replace github.com/imroc/req => github.com/blocktree/req v0.2.5
