module github.com/lightninglabs/aperture

go 1.15

require (
	git.schwanenlied.me/yawning/bsaes.git v0.0.0-20190320102049-26d1add596b6 // indirect
	github.com/NebulousLabs/fastrand v0.0.0-20181203155948-6fb6489aac4e // indirect
	github.com/btcsuite/btcd v0.22.0-beta.0.20220207191057-4dc4ff7963b4
	github.com/btcsuite/btcd/btcec/v2 v2.1.0
	github.com/btcsuite/btcd/btcutil v1.1.0
	github.com/btcsuite/btclog v0.0.0-20170628155309-84c8d2346e9f
	github.com/btcsuite/btcwallet/wtxmgr v1.5.0
	github.com/fortytw2/leaktest v1.3.0
	github.com/golang/protobuf v1.5.2
	github.com/grpc-ecosystem/go-grpc-prometheus v1.2.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.5.0
	github.com/jessevdk/go-flags v1.4.0
	github.com/lightninglabs/lightning-node-connect/hashmailrpc v1.0.2
	github.com/lightninglabs/lndclient v0.12.0-9
	github.com/lightningnetwork/lnd v0.14.2-beta
	github.com/lightningnetwork/lnd/cert v1.1.1
	github.com/lightningnetwork/lnd/tlv v1.0.2
	github.com/prometheus/client_golang v1.11.0
	github.com/stretchr/testify v1.7.0
	go.etcd.io/etcd/client/v3 v3.5.1
	go.etcd.io/etcd/server/v3 v3.5.1
	golang.org/x/crypto v0.0.0-20210921155107-089bfa567519
	golang.org/x/net v0.0.0-20211015210444-4f30a5c0130f
	golang.org/x/time v0.0.0-20210220033141-f8bda1e9f3ba
	google.golang.org/grpc v1.39.0
	google.golang.org/protobuf v1.27.1
	gopkg.in/macaroon.v2 v2.1.0
	gopkg.in/yaml.v2 v2.4.0
)

replace (
	github.com/lightninglabs/lndclient => github.com/lightninglabs/lndclient v1.0.1-0.20220310212109-ce9e183be306
	github.com/lightningnetwork/lnd => github.com/lightningnetwork/lnd v0.14.1-beta.0.20220309185510-262591c3331c
)

// Fix etcd token renewal issue https://github.com/etcd-io/etcd/pull/13262.
replace go.etcd.io/etcd/client/v3 => github.com/lightninglabs/etcd/client/v3 v3.5.1-retry-patch
