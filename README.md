# Post-Metrics

doin a thing

Ideally listens for a POST request with json. Passes the `metric` value onto statsd/statsite at port 8125

run with `docker-compose up`

todo:

* make sure udp sending is working, currently unknown
* add config
* add influx capability?
* accidental ddos
