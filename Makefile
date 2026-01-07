SERVICES = auth-service account-service transaction-service notification-service

run:
	@echo "Use make -C <service> run"

build:
	for s in $(SERVICES); do \
		make -C $$s build; \
	done

docker:
	for s in $(SERVICES); do \
		make -C $$s docker; \
	done
