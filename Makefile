# Define the default target
.DEFAULT_GOAL := help

# Environment-specific variables
STAGING_ENV_SCOPE=staging.env
QA_ENV_SCOPE=qa.env
RELEASE_ENV_SCOPE=release.env

# Help message
help:
	@echo "Usage:"
	@echo "  make run [stage|qa|rel]  Run the specified environment"
	@echo "  make down [stage|qa|rel] Stop and remove the specified environment containers"
	@echo "  make logs [stage|qa|rel] View logs for the specified environment"
	@echo "  make rebuild [stage|qa|rel] Rebuild and run the specified environment"

# Run environment
run: 
	@$(MAKE) _run ENV_SCOPE=$(subst run_,,$@)_ENV_FILE

run_stage: 
	@$(MAKE) _run ENV_SCOPE=$(STAGING_ENV_SCOPE)

run_qa:
	@$(MAKE) _run ENV_SCOPE=$(QA_ENV_SCOPE)

run_rel:
	@$(MAKE) _run ENV_SCOPE=$(RELEASE_ENV_SCOPE)

# Stop and remove containers
down: 
	@$(MAKE) _down ENV_SCOPE=$(subst down_,,$@)_ENV_FILE

down_stage: 
	@$(MAKE) _down ENV_SCOPE=$(STAGING_ENV_SCOPE)

down_qa:
	@$(MAKE) _down ENV_SCOPE=$(QA_ENV_SCOPE)

down_rel:
	@$(MAKE) _down ENV_SCOPE=$(RELEASE_ENV_SCOPE)

# View logs
logs:
	@$(MAKE) _logs ENV_SCOPE=$(subst logs_,,$@)_ENV_FILE

logs_stage:
	@$(MAKE) _logs ENV_SCOPE=$(STAGING_ENV_SCOPE)

logs_qa:
	@$(MAKE) _logs ENV_SCOPE=$(QA_ENV_SCOPE)

logs_rel:
	@$(MAKE) _logs ENV_SCOPE=$(RELEASE_ENV_SCOPE)

# Rebuild and run
rebuild: 
	@$(MAKE) _rebuild ENV_SCOPE=$(subst rebuild_,,$@)_ENV_FILE

rebuild_stage: 
	@$(MAKE) _rebuild ENV_SCOPE=$(STAGING_ENV_SCOPE)

rebuild_qa:
	@$(MAKE) _rebuild ENV_SCOPE=$(QA_ENV_SCOPE)

rebuild_rel:
	@$(MAKE) _rebuild ENV_SCOPE=$(RELEASE_ENV_SCOPE)

# Internal targets
_run:
	@echo "Running with $(ENV_SCOPE)..."
	@ENV_SCOPE=$(ENV_SCOPE) docker-compose up -d --build

_down:
	@echo "Stopping and removing containers for $(ENV_SCOPE)..."
	@ENV_SCOPE=$(ENV_SCOPE) docker-compose down

_logs:
	@echo "Viewing logs for $(ENV_SCOPE)..."
	@ENV_SCOPE=$(ENV_SCOPE) docker-compose logs -f

_rebuild:
	@echo "Rebuilding and running with $(ENV_SCOPE)..."
	@ENV_SCOPE=$(ENV_SCOPE) docker-compose down
	@ENV_SCOPE=$(ENV_SCOPE) docker-compose up -d --build
