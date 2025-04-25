#!/usr/bin/env bash

# IPs

./chx ip crowdsec 39.98.85.230

./chx ip crowdsec 39.98.85.230 --format true

./chx ip alienvault 39.98.85.230

./chx ip alienvault 39.98.85.230 --format true

./chx ip aipdb 39.98.85.230

./chx ip aipdb 39.98.85.230 --format true


# Domains

./chx domain alienvault cloudpoddle.com

./chx domain alienvault cloudpoddle.com --format true

# Hashes

chx hash vthash e37e8e14b6d79046d407ad22b6ce812e5fe8a4b3213320608fb0c1a9ce4fb98a

chx hash vthash e37e8e14b6d79046d407ad22b6ce812e5fe8a4b3213320608fb0c1a9ce4fb98a --format true
