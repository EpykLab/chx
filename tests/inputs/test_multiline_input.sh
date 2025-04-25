#!/bin/bash


function test_ips() {
    cat tests/inputs/ip/ips | chx ip aipdb --format
    cat tests/inputs/ip/ips | chx ip alienvault --format

    # This should not be run...never. never run this unless you would like
    # to be rate limited
    ## -- cat tests/inputs/ip/ips | chx ip crowdsec --format
}

function test_domains() {
    # Test domains
    cat tests/inputs/dns/domains | chx domain alienvault --format
}

function test_hashes() {
    # Test hashes
    cat tests/inputs/dns/hashes | chx hash vthash --format
    cat tests/inputs/dns/hashes | chx hash alienvault --format
}

case $1 in

  ip)
    test_ips
    ;;

  domain)
      test_domains
      ;;

  hash)
      test_hashes
  ;;

  *)
    printf "unknown command \n"
    ;;
esac
