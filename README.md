# checkip

CLI tool that finds out geographic location of an IP address using [GeoLite2 City database](https://dev.maxmind.com/geoip/geoip2/geolite2/). It takes care of downloading the database if it's not present or it's outdated. You need to set the `GEOIP_LICENSE_KEY` environment variable so it can download the database. Read [this](https://dev.maxmind.com/geoip/geoip2/geolite2/#Download_Access) for how to get the license key (it's free).

Installation:

```
> go install
```

Usage:

```
> geoip 1.1.1.1
city unknown, Australia, AU

> geoip $(dig +short reisinge.net)
Frankfurt am Main, Germany, DE

> geoip $(curl -s util.reisinge.net/addr)
Partizanska lupca, Slovakia, SK
```
