unit-test:
	rm -rf allure-unit/allure-results
	export ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH_UNIT) && go test -shuffle on ./... --parallel 8 -tags=unit
	mkdir -p allure-unit/allure-results
	cp environment.properties allure-unit/allure-results
	cp -r allure-unit/allure-results/* allure-results/

unit-test-only:
	export ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH_UNIT) && go test -shuffle on ./... --parallel 8 -tags=unit

integration-test:
	rm -rf allure-integration/allure-results
	export ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH_INTEGR) && go test -shuffle on ./... --parallel 8 -tags=integration
	mkdir -p allure-integration/allure-results
	cp environment.properties allure-integration/allure-results
	cp -r allure-integration/allure-results/* allure-results/

integration-test-only:
	export ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH_INTEGR) && go test -shuffle on ./... --parallel 8 -tags=integration

e2e-test:
	mkdir -p allure-e2e/allure-results
	rm -rf allure-e2e/allure-results
	export ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH_E2E) && go test -shuffle on ./... --parallel 8 -tags=e2e
	mkdir -p allure-e2e/allure-results
	cp environment.properties allure-e2e/allure-results
	cp -r allure-e2e/allure-results/* allure-results/

e2e-test-only:
	export ALLURE_OUTPUT_PATH=$(ALLURE_OUTPUT_PATH_INTEGR) && go test -shuffle on ./... --parallel 8 -tags=e2e


allure:
	#cp -R allure-reports/history allure-results
	rm -rf allure-reports/

	allure generate allure-results -c -o allure-reports

get-report: unit-test integration-test e2e-test allure

show-report: unit-test integration-test e2e-test allure
	allure open allure-reports

.PHONY: all
all: get-report show-report