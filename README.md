[![Build Status](https://travis-ci.org/jreisinger/checkip.svg?branch=master)](https://travis-ci.org/jreisinger/checkip)

# checkip

`checkip` is a CLI tool that finds out information about an IP address (the output is [colored](https://reisinge.net/blog/2021-01-15-check-ip-address) in real terminal):

```
$ checkip 218.92.0.158
AS          4134, 218.92.0.0 - 218.93.8.255, CHINANET-BACKBONE No.31,Jin-rong Street
AbuseIPDB   reported abusive 5414 times with 100% confidence (chinatelecom.com.cn)
DNS         lookup 218.92.0.158: nodename nor servname provided, or not known
Geolocation Qingdao, China, CN
IPsum       found on 5 blacklists
OTX         threat score 3 (seen 2019-12-01 - 2020-03-24)
Shodan      OS unknown, 1 open port: 22 (OpenSSH, 5.3)
ThreatCrowd voted malicious/harmless by equal number of users
VirusTotal  79 harmless, 2 suspicious, 6 malicious analysis results
```

## Installation

Download the latest [release](https://github.com/jreisinger/checkip/releases)
for your operating system and architecture. Copy it to your `bin` folder (or
some other folder on your `PATH`) and make it executable.

The same spelled out in Bash:

```
export SYS=linux # or darwin
export ARCH=amd64
export REPO=checkip
export REPOURL=https://github.com/jreisinger/$REPO
curl -L $REPOURL/releases/latest/download/$REPO-$SYS-$ARCH -o $HOME/bin/$REPO
chmod u+x $HOME/bin/$REPO
```

## Config File

For some checks (see below) to work you need to register and get a
LICENSE/API key. Then create a `$HOME/.checkip.yaml` using your editor of
choice. Provide your API/license keys using the following template:

```
ABUSEIPDB_API_KEY: aaaaaaaabbbbbbbbccccccccddddddddeeeeeeeeffffffff11111111222222223333333344444444
GEOIP_LICENSE_KEY: abcdef1234567890
VIRUSTOTAL_API_KEY: aaaaaaaabbbbbbbbccccccccddddddddeeeeeeeeffffffff1111111122222222
SHODAN_API_KEY: aaaabbbbccccddddeeeeffff11112222
```

You can also use environment variables with the same names as in the config file.

## Features

* Easy to install since it's a single binary.
* Files necessary for some checks are automatically downloaded and updated in the background.
* Checks are done concurrently to save time.
* Output is colored to improve readability.
* You can select which checks you want to run.
* Exit code is the number of checks that say the IP address is not OK.
* It's easy to add new checks.

Currently these checks (types of information) are available:

* AS (Autonomous System) data using TSV file from [iptoasn](https://iptoasn.com/).
* [AbuseIPDB](https://www.abuseipdb.com) reports that the IP address is malicious. You need to [register](https://www.abuseipdb.com/register?plan=free) to get the API key (it's free).
* DNS names using [net.LookupAddr](https://golang.org/pkg/net/#LookupAddr) Go function.
* Geographic location using [GeoLite2 City database](https://dev.maxmind.com/geoip/geoip2/geolite2/) file. You need to [register](https://dev.maxmind.com/geoip/geoip2/geolite2/#Download_Access) to get the license key (it's free).
* Blacklists the IP address is found on according to [IPsum](https://github.com/stamparm/ipsum) file.
* Threat score from [OTX](https://otx.alienvault.com/).
* [Shodan](https://www.shodan.io/) scan data. You need to [register](https://account.shodan.io/register) to get the API key (it's free).
* [ThreatCrowd](https://www.threatcrowd.org/) voting about whether the IP address is malicious.
* [VirusTotal](https://developers.virustotal.com/v3.0/reference#ip-object) analysis results. You need to [register](https://www.virustotal.com/gui/join-us) to to get the API key (it's free).

## Development

```
vim main.go
make install # version defaults to "dev" if VERSION envvar is not set
```

When you push to GitHub Travis CI will try and build a release for you and
publish it on GitHub. Builds are done inside Docker container. To build a
release locally:

```
make release
```

Check test coverage:

```
go test -coverprofile cover.out ./...
go tool cover -html=cover.out
```
