# lore-fetcher

## Introduction

`lore-fetcher` is a project that aims to ease the process of managing multiple CI-Tron CI farms. It also provides some tools and features to simplify and
automate kernel CI pipelines, such as:
+ Triggering new jobs as new patches are sent to kernel development mailing lists;
+ Providing a simple admin TUI to manage farms;
+ Providing a REST API for dashboard frontends;

## How to run

Inside the repository root directory, run:

```
docker compose up -d database
```

to start the postgres database container.

Then, `cd` to the `lore-fetcher` directory and run:

```
go run cmd/main.go
```

to start the application.
