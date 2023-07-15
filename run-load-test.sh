docker-compose -f docker-compose.loadtest.yaml up -d influxdb grafana
echo "--------------------------------------------------------------------------------------"
echo "Load testing with Grafana dashboard http://localhost:3000/d/k6/k6-load-testing-results"
echo "--------------------------------------------------------------------------------------"
docker-compose -f docker-compose.loadtest.yaml run --rm k6 run /scripts/k6.js