# The TinyKV Scheduler
### Build

```
make scheduler
```

It builds the binary of  `tinyscheduler-server` to `bin` dir.

### Run

Under the binary dir, run the following commands:

```
mkdir -p data
```

```
./tinyscheduler-server  --advertise-peer-urls=http://[listen-peer-urls]:2380  --advertise-client-urls=http://[listen-client-urls]:2379
```

