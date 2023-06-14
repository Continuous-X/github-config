module github-config

go 1.19

require (
	github.com/google/go-github/v48 v48.2.0
	github.com/mitchellh/go-homedir v1.1.0
	github.com/spf13/cobra v1.6.1
	github.com/spf13/viper v1.15.0
	golang.org/x/exp v0.0.0-20221230185412-738e83a70c30
	golang.org/x/oauth2 v0.9.0
	gopkg.in/yaml.v2 v2.4.0
	k8s.io/apimachinery v0.27.0
)

require (
	github.com/frankban/quicktest v1.14.4 // indirect
	github.com/fsnotify/fsnotify v1.6.0 // indirect
	github.com/golang/protobuf v1.5.3 // indirect
	github.com/google/go-querystring v1.1.0 // indirect
	github.com/hashicorp/hcl v1.0.0 // indirect
	github.com/inconshreveable/mousetrap v1.1.0 // indirect
	github.com/magiconair/properties v1.8.7 // indirect
	github.com/mitchellh/mapstructure v1.5.0 // indirect
	github.com/pelletier/go-toml/v2 v2.0.6 // indirect
	github.com/spf13/afero v1.9.3 // indirect
	github.com/spf13/cast v1.5.0 // indirect
	github.com/spf13/jwalterweatherman v1.1.0 // indirect
	github.com/spf13/pflag v1.0.5 // indirect
	github.com/subosito/gotenv v1.4.2 // indirect
	golang.org/x/crypto v0.10.0 // indirect
	golang.org/x/net v0.11.0 // indirect
	golang.org/x/sys v0.9.0 // indirect
	golang.org/x/text v0.10.0 // indirect
	google.golang.org/appengine v1.6.7 // indirect
	google.golang.org/protobuf v1.28.1 // indirect
	gopkg.in/check.v1 v1.0.0-20201130134442-10cb98267c6c // indirect
	gopkg.in/ini.v1 v1.67.0 // indirect
	gopkg.in/yaml.v3 v3.0.1 // indirect
)

exclude (
	cloud.google.com/go/storage v1.0.0
	cloud.google.com/go/storage v1.10.0
	cloud.google.com/go/storage v1.14.0
	cloud.google.com/go/storage v1.5.0
	cloud.google.com/go/storage v1.6.0
	github.com/prometheus/client_model v0.0.0-20190812154241-14fe0d1b01d4
	go.etcd.io/etcd/client/v2 v2.305.5
)
