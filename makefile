# ==============================================================================
# Variables
SPLIT = $(subst -, ,$@)
DAY = $(word 2, $(SPLIT))

# ==============================================================================
# Build, add permissions, move to executables. It can be used for creating executable on shell level
build:
	go build -o aocgen

add-permissions:
	chmod +x aocgen

move:
	sudo mv aocgen /usr/local/bin/aocgen

deploy: build add-permissions move

# ==============================================================================
# Generate new day

newday-%:
	go run main.go $(DAY)