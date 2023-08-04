# ngx_mruby_webp_debug

Environment to reproduce that an error occurs in the stream in HTTP2 when using `mruby_output_header_filter` directive

### e2e blowser debugging

```
docker-compose up
```

and access to [https://localhost:10443](https://localhost:10443)

### e2e golang client debugging

It is useful to check client-side behavior using debugging tools such as delve.

```
docker-compose up -d

cd go-http-cli/
go run main.go
```

### gdb debugging

first `docker-compose up`, and

```
# debug master process
docker exec --privileged -it mruby_http2_debug_ngx_1 gdb -p 1

# debug worker process
docker exec -it ngx_mruby_http2_debug_ngx_1 ps aux # check process
docker exec --privileged -it ngx_mruby_http2_debug_ngx_1 gdb -p 7 # example
```
