# lore-fetcher

## Introduction

`lore-fetcher` is a tool for integrating a mailing list with a Kernel CI pipeline.

It has a range of different features which can be used to help in different sections of a CI pipeline
in the context of pre-submission testing of patches, e.g.:
+ It has a `fetch` feature, that can track the most recents patches from a given mailing list and
then trigger a Jenkins pipeline;
+ It can send build results for mailing lists with the `send` option;
+ It can prepare an environment for testing the patch (i.e., clone the kernel subtree and automatically apply
the tested patch with `b4`) with `apply`

## How to install

There is a very simple install script `install.sh` that compiles the Go source code and places it in `/usr/bin`.
Run it with 

```
    sudo ./install.sh
```

## How to start using

The core feature of `lore-fetcher` is its ability to track mailing lists. You can try it with the `--fetch` flag.


The `lore-fetcher --fetch` feature basically activates a daemon mode where `lore-fetcher` will be listening
to the given mailing list for new patches. Once it finds new patches, it can trigger a Jenkins pipeline.
The jenkins Pipeline must be configured with a `PATCH` parameter, so it can get the recent patch and start
a new build with it.

Notice it is also important to provide essential information, such as the mailing list to track. This can be done
with additional parameter flags.

```
lore-fetcher --fetch --mailing-list=all --fetch-interval=37
```

The command above will search for new patches from the `all` mailing list every 37 seconds. 

More details about `lore-fetcher`, its use cases, and proper configuration can be found in its documentation.

You can also try a basic Jenkins CI infrastructure integrated with `lore-fetcher` by running the `jenkins-sample` docker-compose
environment with a single `docker-compose up` command. Check `jenkins-sample/README.md` for more information about
how it is configured and how `lore-fetcher` is used to aid the automation of such infrastructure.

## The Jenkins sample

As mentioned above, there is a Jenkins `docker-compose` environment running a simple jenkins server environment
attached to a `lore-fetcher` instance, receiving new patches and testing a simple `tinyconfig` compilation.

Change to directory `jenkins-sample` and run `docker-compose up -d`, and it is as simple as that. There
will be a container running the `lore-fetcher` application and sending triggers to a Jenkins container.

You can access the Jenkins server Web GUI with any browser on `localhost:8080`, and log in with the 
`admin` user with password `admin`.
