module github.com/cloudquery/cloudquery/plugins/destination/mysql

go 1.19

require (
	github.com/apache/arrow/go/v13 v13.0.0-20230630125530-5a06b2ec2a8e
	github.com/cloudquery/plugin-sdk/v4 v4.8.1-rc1
	github.com/go-sql-driver/mysql v1.7.1
	github.com/google/uuid v1.3.0
	github.com/rs/zerolog v1.29.1
	golang.org/x/exp v0.0.0-20230626212559-97b1e661b5df
)

// TODO: remove once all updates are merged
replace github.com/apache/arrow/go/v13 => github.com/cloudquery/arrow/go/v13 v13.0.0-20230710001530-a2a76ebbb85f

require (
	github.com/cloudquery/plugin-pb-go v1.6.0 // indirect
	github.com/cloudquery/plugin-sdk/v2 v2.7.0 // indirect
	github.com/davecgh/go-spew v1.1.1 // indirect
	github.com/getsentry/sentry-go v0.22.0 // indirect
	github.com/ghodss/yaml v1.0.0 // indirect
	github.com/goccy/go-json v0.10.2 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/flatbuffers v23.5.26+incompatible // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/providers/zerolog/v2 v2.0.0-rc.3 // indirect
	github.com/grpc-ecosystem/go-grpc-middleware/v2 v2.0.0-rc.3 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/klauspost/compress v1.16.6 // indirect
	github.com/klauspost/cpuid/v2 v2.2.5 // indirect
	github.com/mattn/go-colorable v0.1.13 // indirect
	github.com/mattn/go-isatty v0.0.19 // indirect
	github.com/pierrec/lz4/v4 v4.1.18 // indirect
	github.com/pmezard/go-difflib v1.0.0 // indirect
	github.com/spf13/cobra v1.7.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/stretchr/testify v1.8.4 // indirect
	github.com/thoas/go-funk v0.9.3 // indirect
	github.com/zeebo/xxh3 v1.0.2 // indirect
	golang.org/x/mod v0.11.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sync v0.3.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	golang.org/x/tools v0.10.0 // indirect
	golang.org/x/xerrors v0.0.0-20220907171357-04be3eba64a2 // indirect
	google.golang.org/genproto/googleapis/rpc v0.0.0-20230629202037-9506855d4529 // indirect
	google.golang.org/grpc v1.56.1 // indirect
	google.golang.org/protobuf v1.31.0 // indirect
	gopkg.in/yaml.v2 v2.4.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)
