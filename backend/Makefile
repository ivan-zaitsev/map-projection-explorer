migrate:
	docker run --net=host --rm \
	-v ${CURDIR}/config/sql:/liquibase/changelog liquibase/liquibase:latest update \
	--changelog-file=changelog/changelog.xml --classpath=/liquibase/changelog --url=$$LIQUIBASE_DATABASE_URI
