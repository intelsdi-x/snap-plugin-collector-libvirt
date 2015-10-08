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
[![Build Status](https://magnum.travis-ci.com/intelsdi-x/pulse.svg?token=2ujsxEpZo1issFyVWX29&branch=master)](https://magnum.travis-ci.com/intelsdi-x/pulse-plugin-collector-libvirt )
## Pulse Libvirt Collector Plugin

# Description
Collect following metrics from libvirt domain

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


# Assumptions
* "libvirt" - libvirt daemon is installed and started.

# Configuration

By default plugin is using qemu:///system uri. To monitor external
system, you can pass uri parameter to the pulsed deamon configuration.
Example configuration file can be find in examples directory.

## Building

```
make
```

## Testing

```
export TEST=unit
make test
```

