# prometheus-client

Sample prometheus test environment with a number of test clients and grafana(for visualising prometheus metrics)

## Documents

- [prometheus.io](https://prometheus.io/)
- [grafana.com](https://grafana.com/)

## Docker images
- [prometheus](https://hub.docker.com/r/prom/prometheus/)
- [grafana](https://hub.docker.com/r/grafana/grafana/)

## Getting started

- clone repository
- `$ docker-compose build`
- `$ docker-compose up`
- Access to prometheus `http://localhost:9090`
- Access to grafana `http://localhost:3000`
- Set up grafana data source `http://localhost:3000`

## endpoints

- Prometheus(`http://localhost:9090`)
- Grafana(`http://localhost:3000`)

You can login grafana as admin.

ID: admin

Pass: admin
