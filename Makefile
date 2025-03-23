.PHONY: lore-fetcher
lore-fetcher:
	docker compose build --no-cache
	docker compose up
