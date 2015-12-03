[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-collector-libvirt.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-collector-libvirt )
# snap collector plugin - libvirt

1. [Getting Started](#getting-started)
  * [System Requirements](#system-requirements)
  * [Installation](#installation)
  * [Configuration and Usage](configuration-and-usage)
2. [Documentation](#documentation)
  * [Collected Metrics](#collected-metrics)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3. [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License](#license)
6. [Acknowledgements](#acknowledgements)

## Getting Started

### System Requirements

Linux/*BSD system with libvirt installed

### Installation

#### Install libvirt and libvirt-dev package

For Ubuntu > 14.04
```
apt-get install libvirt-bin libvirt-dev
```
For Fedora >= 21 / CentOS >= 7 / RedHat >= 7
```
yum install libvirt libvirt-devel
```

#### Compile plugin
```
make
```

## Documentation

### Examples

By default the plugin is using qemu:///system uri. To monitor external
systems, you can pass the uri parameter to the snapd deamon configuration.
An example configuration file can be found in example directory.


### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Data Type | Description
----------|-----------|-----------------------
/libvirt/{domain_name}/disk/{device_name}/wrreq| uint64|Write Requests
/libvirt/{domain_name}/disk/{device_name}/rdreq| uint64|Read Requests
/libvirt/{domain_name}/disk/{device_name}/wrbytes| uint64|Write Bytes
/libvirt/{domain_name}/disk/{device_name}/rdbytes| uint64|Read Bytes
/libvirt/{domain_name}/mem/mem| uint64|Amount of memory specified on domain creation
/libvirt/{domain_name}/mem/swap_in| uint64|Amount of memory swapped in
/libvirt/{domain_name}/mem/swap_out| uint64|Amount of memory swapped out
/libvirt/{domain_name}/mem/major_fault| uint64|Number of major faults
/libvirt/{domain_name}/mem/minor_fault| uint64|Number of minor faults
/libvirt/{domain_name}/mem/free| uint64|Total amount of free memory
/libvirt/{domain_name}/mem/max| uint64|Total amount of memory
/libvirt/{domain_name}/cpu/cputime| uint64|Cputime ( all vcpus )
/libvirt/{domain_name}/cpu/vcpu/{vcpu_nr}/cputime| uint64|Cputime for one vcpu
/libvirt/{domain_name}/net/{interface_name}/rxbytes| uint64|Bytes received
/libvirt/{domain_name}/net/{interface_name}/rxpackets| uint64|Packets received
/libvirt/{domain_name}/net/{interface_name}/rxerrs| uint64|Errors on receive
/libvirt/{domain_name}/net/{interface_name}/rxdrop| uint64|Drops on receive
/libvirt/{domain_name}/net/{interface_name}/txbytes| uint64|Bytes transmitted
/libvirt/{domain_name}/net/{interface_name}/txpackets| uint64|Packets transmitted
/libvirt/{domain_name}/net/{interface_name}/txerrs| uint64|Errors on transmit
/libvirt/{domain_name}/net/{interface_name}/txdrop| uint64|Drops on transmit

**_IMPORTANT_**: not all hypervisors expose all these metrics. Please check
your hypervisor or libvirt documentation.

### Community Support
This repository is one of **many** plugins in the **snap framework**: a powerful telemetry agent framework.
The full project is at https://github.com/intelsdi-x/snap.

### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-libvirt/issues).

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-libvirt/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-libvirt/pulls).

## Community Support
This repository is one of **many** plugins in **snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support)

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
This is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

* Author: [Marcin Spoczynski](https://github.com/sandlbn/)

## Thank You
And **thank you!** Your contribution, through code and participation, is incredibly important to us.
