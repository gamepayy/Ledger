# GamePayy Ledger

This ledger is a work in progress. It is not yet ready for production.

## Overview

This ledger is the backend system that manages GamePayy's platform balances and transactions. It utilizes MYSQl as the database through PlanetScale. It is written in Go and utilizes the [Gin](https://github.com/gin-gonic/gin) framework to handle HTTP requests and build the documentation. It utilizes Redis for caching.
It exposes a REST API for the GamePayy platform to interact with.

## Summary

## Diagrams
![Diagram Image Link](https://bafkreiczjo4bxt6nxqwdzyjy7orlklruzskw5du434zfayet3552ygoxpa.ipfs.w3s.link/)
![Rewards Distributor](https://bafkreieso3uu4warki36znxzs25lkbfmvjk4ghqxrpiwzldasfh3rz2sym.ipfs.w3s.link/)
![Users Database]()
![Balances Database]()
![PendingRewards Database]()

## Getting Started

## How the Rewards System Works
A user can earn rewards by completing tasks. These tasks are called  "Challenges". Each action has a set of requirements that must be met in order for the user to earn the reward. Each requirement has a specific metric and a value that must be reached.
Rewards can also be earned by competing in tournaments and matches. These are mediated by the GamePayy platform through the usage of Toornament APIs. The GamePayy platform will send a request to the ledger to create a new tournament or match. The ledger will then create a new tournament or match in the database. The GamePayy platform will then create a new tournament or match in the Toornament database.
The GamePayy platform will then receive a request from arbitrators to the ledger to update the tournament or match with the results. The ledger will then update the tournament or match in the database. The GamePayy platform will then update the tournament or match in the Toornament database.

## Rewards Distributor
