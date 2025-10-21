# ip3country-go

A zero-dependency, super small, IP address to 2-letter [country code](https://en.wikipedia.org/wiki/ISO_3166-1_alpha-2) lookup library. There are already several libraries available, but none met our requirements for binary size and lookup speed.

This project in its entirety is <500KB, compared to most alternatives out there that are in the MBs (but might provide more than just countries).

The database used in this project is compacted from [IP2Location](https://lite.ip2location.com/database/ip-country). Their DB1.LITE edition is provided under CCA, with the attribution below:

---

**NOTE**

This site or product includes IP2Location LITE data available from <a href="https://lite.ip2location.com">https://lite.ip2location.com</a>.

---

## Usage

Add the dependency to your `go.mod` file:

```
require (
	github.com/statsig-io/ip3country-go v0.2.0
)
```

```go
lookup := countrylookup.New()
// Lookup using ip4 str
countryCode, ok := lookup.LookupIp("123.45.67.8") // => 'KR'

// Lookup using numeric ip
countryCode, ok = lookup.LookupNumericIp(2066563848) // => 'KR'
```

## Accuracy

`ip3country`'s accuracy is dependent on IP2Location LITE's accuracy. In our experience, for country lookups, it is accurate enough for most applications.

IP2Location publishes accuracy reports here: https://www.ip2location.com/data-accuracy

There's also this third party report available: https://www.cl.cam.ac.uk/~nz247/publications/JSAC2011-Geolocation.pdf

## Name

All variants of "ip2country" were already taken in npm when we published the original, nodejs verion. Now, we have re-created it for use in go applications (or as a binary in other languages)

## Motivation

At [Statsig](https://www.statsig.com), we needed a library to evaluate location-based conditions on Feature Gates (Feature Flags).  We use ip3country to power country-level rollouts or A/B tests for new features.
