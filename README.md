# "Hello, world!" for Sensu Go

## Table of Contents
- [Overview](#overview)
- [Usage examples](#usage-examples)
- [Configuration](#configuration)
  - [Asset registration](#asset-registration)
  - [Resource definition](#resource-definition)
- [Installation from source](#installation-from-source)
- [Additional notes](#additional-notes)
- [Contributing](#contributing)

## Overview

The helloworld plugin is an example Sensu Go check plugin.

## Usage examples

- Default usage:

  ```
  $ ./helloworld
  Hello, world!
  ```

- Optional `-w`/`--world` usage:

  ```
  $ ./helloworld --world example
  Hello, example world!
  ```

- The `-w`/`--world` flag can be set via an environment variable as well:

  ```
  $ SENSU_WORLD="environment variable" ./helloworld
  Hello, environment variable world!
  ```

### Help

```
A Sensu Go "hello world" check plugin

Usage:
  helloworld [flags]
  helloworld [command]

Available Commands:
  help        Help about any command
  version     Print the version number of this plugin

Flags:
  -h, --help           help for helloworld
  -w, --world string   The world to send greetings from (optional)

Use "helloworld [command] --help" for more information about a command.
```

## Configuration

### Asset registration

If you're using sensuctl 5.13 with Sensu Backend 5.13 or later, you can use the
following command to add this asset to your cluster:

```
$ sensuctl asset add calebhailey/helloworld
```

If you're using an earlier version of sensuctl, you can find the asset on the
[Bonsai Asset Index][https://bonsai.sensu.io/assets/calebhailey/helloworld].

### Resource definition

```yml
---
type: CheckConfig
api_version: core/v2
metadata:
  name: helloworld
spec:
  command: helloworld --world example
  publish: true
  interval: 30
  timeout: 5
  subscriptions:
  - helloworld
  runtime_assets:
  - calebhailey/helloworld
```

## Installation from source

The preferred way of installing and deploying this plugin is via [Sensu
Assets][10]. If you would like to compile and install the plugin from source or
contribute to it, download the latest version or create an executable script
from this source.

Clone this repository and run the following command from the local path of the
`helloworld` repository:

```
$ go build
```

## Contributing

For more information about contributing to this plugin, see [Contributing][1].

[1]: https://github.com/sensu/sensu-go/blob/master/CONTRIBUTING.md
[2]: https://github.com/sensu-community/sensu-plugin-sdk
[3]: https://github.com/sensu-plugins/community/blob/master/PLUGIN_STYLEGUIDE.md
[4]: https://github.com/sensu-community/check-plugin-template/blob/master/.github/workflows/release.yml
[5]: https://github.com/sensu-community/check-plugin-template/actions
[6]: https://docs.sensu.io/sensu-go/latest/reference/checks/
[7]: https://github.com/sensu-community/check-plugin-template/blob/master/main.go
[8]: https://bonsai.sensu.io/
[9]: https://github.com/sensu-community/sensu-plugin-tool
[10]: https://docs.sensu.io/sensu-go/latest/reference/assets/
