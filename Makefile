# my.id Music Streaming

BUILD_DIR = build
TEST =
UNIT_TEST_BUILD_DIR = $(BUILD_DIR)/unit-test
UNIT_TEST_PKG = \
	./core/codec \
	./core/ipc

INTEGRATION_TEST_PKG =

MOCK_DIR = \
	./core/crypto/mock \
	./core/ipc/mock

.PHONY: all
all:
# Run all
#
	@echo 'Nothing'

.PHONY: mock
mock:
# Run mockery
#
	@rm -rf $(MOCK_DIR)
	@mockery

.PHONY: unit-test
unit-test:
# Run all test cases
#
	@mkdir -p $(UNIT_TEST_BUILD_DIR)
	@go clean -testcache
	@go test \
		-v \
		-coverprofile=coverage.out \
		-outputdir $(UNIT_TEST_BUILD_DIR) \
		-run '$(TEST)' \
		$(UNIT_TEST_PKG)

.PHONY: unit-test
unit-test-coverage:
# Run all test cases coverage
#
	@go tool cover -html=$(UNIT_TEST_BUILD_DIR)/coverage.out

.PHONY: integration-test
integration-test:
# Run all test cases
#
	@go clean -testcache
	@go test \
		-v \
		-run '$(TEST)' \
		$(INTEGRATION_TEST_PKG)

.PHONY: clean
clean:
# Remove build
#
	@rm -rf build
