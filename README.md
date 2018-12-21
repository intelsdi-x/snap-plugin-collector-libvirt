# DISCONTINUATION OF PROJECT 

**This project will no longer be maintained by Intel.  Intel will not provide or guarantee development of or support for this project, including but not limited to, maintenance, bug fixes, new releases or updates.  Patches to this project are no longer accepted by Intel. If you have an ongoing need to use this project, are interested in independently developing it, or would like to maintain patches for the community, please create your own fork of the project.**


[![Build Status](https://api.travis-ci.org/intelsdi-x/snap-plugin-collector-libvirt.svg)](https://travis-ci.org/intelsdi-x/snap-plugin-collector-libvirt )
[![Go Report Card](http://goreportcard.com/badge/intelsdi-x/snap-plugin-collector-libvirt)](http://goreportcard.com/report/intelsdi-x/snap-plugin-collector-libvirt)
# Snap collector plugin - libvirt

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

For MacOS
```
brew install libvirt
```
For Ubuntu > 14.04
```
apt-get install libvirt-bin libvirt-dev
```
For Fedora >= 21 / CentOS >= 7 / RedHat >= 7
```
yum install libvirt libvirt-devel
```

#### Download the plugin binary:

You can get the pre-built binaries for your OS and architecture from the plugin's [GitHub Releases](https://github.com/intelsdi-x/snap-plugin-collector-libvirt/releases) page. Download the plugin from the latest release and load it into `snapteld` (`/opt/snap/plugins` is the default location for Snap packages).

#### To build the plugin binary:

Fork https://github.com/intelsdi-x/snap-plugin-collector-libvirt
Clone repo into `$GOPATH/src/github.com/intelsdi-x/`:

```
$ git clone https://github.com/<yourGithubID>/snap-plugin-collector-libvirt.git
```

Build the Snap libvirt plugin by running make within the cloned repo:
```
$ make
```
This builds the plugin in `./build/`

## Documentation

### Examples

For quick plugin test using simple VM, you can go through steps below:
1. Install QEMU and libvirt:
```
sudo apt install qemu-kvm libvirt-bin virtinst qemu-system
```
2. Download Ubuntu cloud image:
```
wget https://cloud-images.ubuntu.com/xenial/current/xenial-server-cloudimg-amd64-disk1.img
```
3. Create VM:
```
virt-install --name=test_vm --arch=x86_64 --vcpus=1 --ram=512 --os-type=linux --hvm --connect=qemu:///system --network bridge:br0 --disk path=xenial-server-cloudimg-amd64-disk1.img,size=20 --boot hd --accelerate --vnc --noautoconsole --keymap=es
```
4. Load Snap libvirt collector plugin and create task:
```
wget https://raw.githubusercontent.com/intelsdi-x/snap-plugin-collector-libvirt/master/example/tasks/task-example.yaml
snaptel plugin load snap-plugin-collector-libvirt
snaptel task create -t task-example.yaml
```
5. Check task output:
```
snaptel task list
snaptel task watch <Task ID>
```
6. Remove test VM:
```
virsh destroy test_vm
virsh undefine test_vm
```

By default the plugin is using qemu:///system uri. To monitor external
systems, you can pass the uri parameter to the snapteld deamon configuration.
An example configuration file can be found in example directory.

### Collected Metrics
This plugin has the ability to gather the following metrics:

Namespace | Data Type | Description
----------|-----------|-----------------------
/intel/libvirt/{domain_name}/disk/{device_name}/wrreq| int64|Write Requests
/intel/libvirt/{domain_name}/disk/{device_name}/rdreq| int64|Read Requests
/intel/libvirt/{domain_name}/disk/{device_name}/wrbytes| int64|Write Bytes
/intel/libvirt/{domain_name}/disk/{device_name}/rdbytes| int64|Read Bytes
/intel/libvirt/{domain_name}/memory/mem| int64|Amount of memory specified on domain creation
/intel/libvirt/{domain_name}/memory/swap_in| int64|Amount of memory swapped in
/intel/libvirt/{domain_name}/memory/swap_out| int64|Amount of memory swapped out
/intel/libvirt/{domain_name}/memory/major_fault| int64|Number of major faults
/intel/libvirt/{domain_name}/memory/min_fault| int64|Number of minor faults
/intel/libvirt/{domain_name}/memory/max| int64|Total amount of memory
/intel/libvirt/{domain_name}/memory/unused| int64| Amount of memory left unused by the system 
/intel/libvirt/{domain_name}/memory/available| int64| Amount of usable memory
/intel/libvirt/{domain_name}/memory/actual_balloon| int64| Current balloon value 
/intel/libvirt/{domain_name}/memory/rss| int64| Resident Set Size of the process running the domain
/intel/libvirt/{domain_name}/memory/nr| int64| Number of statistics supported by the interface
/intel/libvirt/{domain_name}/cpu/cputime| int64|Cputime ( all vcpus )
/intel/libvirt/{domain_name}/cpu/cputime/{vcpu_nr}| int64|Cputime for one vcpu (not supported on qemu without kvm)
/intel/libvirt/{domain_name}/network/{interface_name}/rxbytes| int64|Bytes received
/intel/libvirt/{domain_name}/network/{interface_name}/rxpackets| int64|Packets received
/intel/libvirt/{domain_name}/network/{interface_name}/rxerrs| int64|Errors on receive
/intel/libvirt/{domain_name}/network/{interface_name}/rxdrop| int64|Drops on receive
/intel/libvirt/{domain_name}/network/{interface_name}/txbytes| int64|Bytes transmitted
/intel/libvirt/{domain_name}/network/{interface_name}/txpackets| int64|Packets transmitted
/intel/libvirt/{domain_name}/network/{interface_name}/txerrs| int64|Errors on transmit
/intel/libvirt/{domain_name}/network/{interface_name}/txdrop| int64|Drops on transmit

**_IMPORTANT_**: not all hypervisors expose all these metrics. Please check
your hypervisor or libvirt documentation.

### Roadmap
As we launch this plugin, we do not have any outstanding requirements for the next release. If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-libvirt/issues).

If you have a feature request, please add it as an [issue](https://github.com/intelsdi-x/snap-plugin-collector-libvirt/issues/new) and/or submit a [pull request](https://github.com/intelsdi-x/snap-plugin-collector-libvirt/pulls).

## Community Support
This repository is one of **many** plugins in **Snap**, a powerful telemetry framework. See the full project at http://github.com/intelsdi-x/snap.

To reach out to other users, head to the [main framework](https://github.com/intelsdi-x/snap#community-support).

## Contributing
We love contributions!

There's more than one way to give back, from examples to blogs to code updates. See our recommended process in [CONTRIBUTING.md](CONTRIBUTING.md).

## License
[Snap](http://github.com/intelsdi-x/snap), along with this plugin, is an Open Source software released under the Apache 2.0 [License](LICENSE).

## Acknowledgements
This is Open Source software released under the Apache 2.0 License. Please see the [LICENSE](LICENSE) file for full license details.

* Author: [Marcin Spoczynski](https://github.com/sandlbn/)

This software has been contributed by MIKELANGELO, a Horizon 2020 project co-funded by the European Union. https://www.mikelangelo-project.eu/
## Thank You
And **thank you!** Your contribution, through code and participation, is incredibly important to us.
