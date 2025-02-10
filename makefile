#*============================================================================*#
#*=====*                          Helpers  GRPC                         *=====*#
#*============================================================================*#

.PHONY: protobuf-update-proto-files protobuf-clean protobuf-gen-all

protobuf-update-proto-files:
	@for dir in ./*; do \
		if [ -d "$$dir" ] && [ "$$dir" != "./proto-files" ]; then \
			mkdir -p "$$dir/proto"; \
			cp -r ./proto-files/* "$$dir/proto/"; \
		fi; \
	done

protobuf-gen-all:
	@echo "Generating protobuf files..."
	@cd proto-files && buf generate || { echo "Error: buf generate failed"; exit 1; }
	@echo "Copying generated files to subdirectories..."
	@for dir in ./*; do \
		if [ -d "$$dir" ] && [ "$$dir" != "./proto-files" ]; then \
			echo "Processing directory: $$dir"; \
			mkdir -p "$$dir/proto"; \
			cp -r ./proto-files/gen "$$dir/proto/" 2>/dev/null || { echo "Warning: Failed to copy files to $$dir/proto/"; }; \
		fi; \
	done
	@echo "Cleaning up temporary files..."
	@if [ -d "proto-files/gen" ]; then \
		rm -rf proto-files/gen; \
	else \
		echo "No temporary files to clean."; \
	fi

protobuf-clean:
	@rm -rf ./**/proto/gen

#*============================================================================*#
#*=====*                           Dev Docker                           *=====*#
#*============================================================================*#
.PHONY: docker-side

docker-side:
	@docker compose --profile side up