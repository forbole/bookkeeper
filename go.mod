module github.com/forbole/bookkeeper

go 1.17

require (
	github.com/cosmos/cosmos-sdk v0.44.0
	github.com/scorredoira/email v0.0.0-20191107070024-dc7b732c55da
	github.com/spf13/cobra v1.4.0
	github.com/superoo7/go-gecko v1.0.0
	google.golang.org/grpc v1.45.0
)

require (
	github.com/inconshreveable/mousetrap v1.0.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
)

replace github.com/gogo/protobuf => github.com/regen-network/protobuf v1.3.3-alpha.regen.1

replace github.com/cosmos/cosmos-sdk => github.com/desmos-labs/cosmos-sdk v0.43.0-alpha1.0.20211102084520-683147efd235

replace github.com/tendermint/tendermint => github.com/forbole/tendermint v0.34.13-0.20210820072129-a2a4af55563d
