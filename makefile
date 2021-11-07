build:
	cd tracker && go build . && mv tracker ../dist
run_production:
	export GIN_MODE=release && ./dist/tracker
