module github.com/blocktree/whitecoin-adapter

go 1.12

require (
	github.com/astaxie/beego v1.12.0
	github.com/blocktree/bitshares-adapter v1.0.5
	github.com/blocktree/go-owaddress v1.0.5
	github.com/blocktree/go-owcdrivers v1.1.21
	github.com/blocktree/go-owcrypt v1.0.4
	github.com/blocktree/openwallet v1.5.4
	github.com/btcsuite/btcd v0.0.0-20190315201642-aa6e0f35703c
	github.com/btcsuite/btcutil v0.0.0-20190316010144-3ac1210f4b38
	github.com/denkhaus/logging v0.0.0-20180714213349-14bfb935047c
	github.com/emirpasic/gods v1.12.0
	github.com/imroc/req v0.2.3
	github.com/juju/errors v0.0.0-20181118221551-089d3ea4e4d5
	github.com/mr-tron/base58 v1.1.1
	github.com/pkg/errors v0.8.1
	github.com/pquerna/ffjson v0.0.0-20181028064349-e517b90714f7
	github.com/shiena/ansicolor v0.0.0-20151119151921-a422bbe96644 // indirect
	github.com/shopspring/decimal v0.0.0-20191009025716-f1972eb1d1f5
	github.com/stretchr/testify v1.4.0
	github.com/tidwall/gjson v1.3.4
	golang.org/x/crypto v0.0.0-20190701094942-4def268fd1a4

)

replace github.com/imroc/req => github.com/blocktree/req v0.2.5
