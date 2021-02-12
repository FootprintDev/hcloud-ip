# Hetzner Cloud Floating IP Assigner
A CLI utility to assign a floating IP to the Server executing it.

## Requirements
+ Hostname equals Cloud Server name
+ Read/Write Cloud API Key
+ Name of the floating IP Address

## Usage:

```bash
root@NGB-VoIP-0:~# ./hcloud-ip -ip VoIP -key [Cloud API Key]
```