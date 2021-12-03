# monitor go file changes and show coverage
watch_coverage:
	nodemon -e go -x "godotenv -f .env_test go test -coverprofile=coverage.out ./... && go tool cover -func=coverage.out"

# monitor the unit tests on change
watch_unit:
	air -c etc/air_test.conf

# runs ui locally and monitors
watch_local_ui:
	npm run dev

# runs app locally and monitors
watch_local_app:
	air -c docker/app/air.conf
