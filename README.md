Command

# Build
``` bash
docker build . -t dev/http-proxy-api
```

``` bash
curl http://splendidfreshclearbirds.neverssl.com:8080/online/ -H "Host: splendidfreshclearbirds.neverssl.com" --resolve splendidfreshclearbirds.neverssl.com:8080:127.0.0.1
```

``` bash
curl http://ezoic.com:8080/online/ -H "Host: ezoic.com" --resolve ezoic.com:8080:127.0.0.1
```