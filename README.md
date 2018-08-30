### How to run it:

```bash
make run
# or
make build && ./jumpcloud
# or
make install && jumpcloud
```

### Test endpoints:

**/hash** endpoint:
```bash
curl -i --data "password=abc" localhost:8080/hash
```

**/stats** endpoint:
```bash
curl -i localhost:8080/stats
```

**/shutdown** endpoint:
```bash
curl -i localhost:8080/shutdown
```

#### Test hashing package:

```bash
cd hashing
go test -v .
```

##### Improvements (only if the code grows more):
 - Move each handler to a new package with its own context
 - Move *stats* logic to a new package
 - Move middlewares to a new package
 - More tests
 
