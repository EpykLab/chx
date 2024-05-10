<h1 align="center">
  <img src="assets/ipchecklogo.png" alt="ipcheck" width="200px">
  <br>
</h1>

<h4 align="center">An IP Investigation Tool for the CLI</h4>

<p align="center">
<!-- Go report card -->
<a href="https://goreportcard.com/report/github.com/epyklab/ipcheck"><img src="https://goreportcard.com/badge/github.com/epyklab/ipcheck"></a>
<!-- Current Release -->
<a href="https://github.com/epyklab/ipcheck/releases"><img src="https://img.shields.io/github/release/epyklab/ipcheck"></a>
</p>

## Summary

IPCheck is a go application that faciliates investigation of IP addresses in the terminal so as to avoid the need to context switch when you are already working in the terminal. For now, IPCheck only makes use of AbuseIP DB, but future plan include integrations into into Crowdsec, and OTX.

## Installation

The easiest way to install ipcheck is go:

```bash
go install github.com/EpykLab/ipcheck@latest
```

Prebuilt binaries are also avaiable for download in the [realease section](https://github.com/epyklab/ipcheck/releases).

## Setup

The only requirement for setup is configuring the AbuseIP API key. If you don't have one already, you will need to set up a free AbuseIP account. Once you have one, export the API key as an environmental variable with the following:

```bash
export abuseipdbkey=<apikey>
```

A convenient way to have this always available is to add it you your .bashrc/.zshrc file, like so:

```bash
echo "export abuseipdbkey=<apikey>" >> ~/.zshrc
```

## Usage

Using ipcheck is simple. simple run the following:

```bash
ipcheck <ip addr>
```

This will return a json object that can be piped into `jq` for further manipulation.

Ipcheck can also return data within a give time span. For example, you can retrieve the reported information for an IP address over the 30 days. This is passed in as the second argument to ipcheck but is not required. For example:

```bash
ipcheck 170.205.29.2 20

```

This will return all information for `170.205.29.2` over the past 20 days. If no
time space is specifed, ipcheck defaults to ruturning information over the past 90 calander days.
