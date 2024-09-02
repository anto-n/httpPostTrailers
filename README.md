# httpPostTrailers

## EnvoyProxy configuration

In order to enable Trailer headers processing by EnvoyProxy enable it under HttpConnectionManager
```
- name: envoy.filters.network.http_connection_manager
  typed_config:
    "@type": type.googleapis.com/envoy.extensions.filters.network.http_connection_manager.v3.HttpConnectionManager
    http_protocol_options:
      enable_trailers: true
```


## Run 
Execute below command
```
go run main.go
```
