# Post-Metrics

For when you need a REST endpoint router instead of straight UDP for statsd.

run with `docker-compose up`

accepts:

```json
{ "metric": "your.series.name:1|c" }
{
    "metrics": ["your.first.series:1|c",
              "your.second.series:1|c"
}
```

todo:

* add config
* add influx capability?
* accidental ddos
