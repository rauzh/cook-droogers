test:
	rm -rf allure-results
	export ALLURE_OUTPUT_PATH="/Users/rauzh/Desktop/cook_droogers/PPO_DB/cook-droogers/backend" && go test -shuffle on ./... --race --parallel 8
	cp environment.properties allure-results

allure:
	cp -R allure-reports/history allure-results
	rm -rf allure-reports
	allure generate allure-results -o allure-reports
	allure serve allure-results -p 4000

report: test allure

.PHONY: all
all: test allure report
