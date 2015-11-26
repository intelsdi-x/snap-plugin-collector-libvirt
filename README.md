<!--
http://www.apache.org/licenses/LICENSE-2.0.txt


Copyright 2015 Intel Corporation

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
-->
[![Build Status](https://api.travis-ci.com/intelsdi-x/snap-plugin-collector-libvirt.svg?token=FhmCtm9AdqhSXoSbqxo2&branch=master)](https://travis-ci.com/intelsdi-x/snap-plugin-collector-libvirt )
# Plugin - snap libvirt collector

1. [Getting Started](#getting-started)
2. [Documentation](#documentation)
  * [Examples](#examples)
  * [Roadmap](#roadmap)
3.  [Community Support](#community-support)
4. [Contributing](#contributing)
5. [License and Authors](#license-and-authors)
6. [Thank You](#thank-you)

## Getting Started


### Install libvirt and libvirt-dev package

* Ubuntu > 14.04
```
apt-get install libvirt-bin libvirt-dev
```
* Fedora >= 21 / CentOS >= 7 / RedHat >= 7
```
yum install libvirt libvirt-devel
```

### Compile plugin
```
make
```

## Documentation

### Examples

By default the plugin is using qemu:///system uri. To monitor external
systems, you can pass the uri parameter to the snapd deamon configuration.
An example configuration file can be found in example directory.


### Exposed metrics
Using the libvirt plugin you can collect the following metrics from the libvirt domain

*  /libvirt/disk/{device_name}/wrreq - Write Requests
*  /libvirt/disk/{device_name}/rdreq - Read Requests
*  /libvirt/disk/{device_name}/wrbytes - Write Bytes
*  /libvirt/disk/{device_name}/rdbytes - Read Bytes
*  /libvirt/mem/mem - Amount of memory specified on domain creation
*  /libvirt/mem/swap_in - Amount of memory swapped in
*  /libvirt/mem/swap_out - Amount of memory swapped out
*  /libvirt/mem/major_fault - Number of major faults
*  /libvirt/mem/minor_fault - Number of minor faults
*  /libvirt/mem/free - Total amount of free memory
*  /libvirt/mem/max - Total amount of memory
*  /libvirt/cpu/cputime - Cputime ( all vcpus )
*  /libvirt/cpu/vcpu/{vcpu_nr}/cputime - Cputime for one vcpu
*  /libvirt/net/{interface_name}/rxbytes - Bytes received
*  /libvirt/net/{interface_name}/rxpackets - Packets received
*  /libvirt/net/{interface_name}/rxerrs - Errors on receive
*  /libvirt/net/{interface_name}/rxdrop - Drops on receive
*  /libvirt/net/{interface_name}/txbytes - Bytes transmitted
*  /libvirt/net/{interface_name}/txpackets - Packets transmitted
*  /libvirt/net/{interface_name}/txerrs - Errors on transmit
*  /libvirt/net/{interface_name}/txdrop - Drops on transmit

**_IMPORTANT_**: not all hypervisors expose all these metrics. Please check
your hypervisor or libvirt documentation.

### Community Support
This repository is one of **many** plugins in the **snap framework**: a powerful telemetry agent framework.
The full project is at https://github.com/intelsdi-x/snap.

### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-libvirt/issues).

## Contributing
We love contributions! :heart_eyes:

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License and Authors
This is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

* Author: [Marcin Spoczynski](https://github.com/sandlbn/)

## Thank You
And **thank you!** Your contribution is incredibly important to us.

