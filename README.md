# nb (NetBox client)

## Usage:
Basic console NetBox client. This tool much faster than web ui to retrive few lines of information.

You can use filtering in query after primary reuest.
```
nb dev "lab?&role=rr"
```
Which mean find in dcim devices which name conatain "lab" and role is "rr"

## Current section to search:

 - ip /api/ipam/ip-addresses/
 - pref /api/ipam/prefixes/
 - agg /api/ipam/aggregates/
 - dev /api/dcim/devices/
 - vm /api/virtualization/virtual-machines/

## TODO:

- [ ] Move api endpoints and its aliases to confiuration file
- [ ] Add default hostname to config file
- [ ] Add auth if there will be demand
