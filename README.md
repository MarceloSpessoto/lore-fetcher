# lore-fetcher

Lore-fetcher is a tool for integrating a mailing list with a Kernel CI pipeline.

It tracks the most recent patches of a given development mailing list, passes the collected
patches to a CI pipeline and returns the results to given mail thread.

Thus, it can be broken down into three phases:
1. The active monitoring of the mailing list.
2. The interaction with the CI pipeline, sending the patches to it and receiving the testing result.
3. The submission of given result to original thread in the mailing list.
