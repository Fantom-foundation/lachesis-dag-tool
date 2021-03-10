# Lachesis DAG tool: txsgen

Transactions and API-calls generator for stress and performance testing of lachesis net.
It uses generated accounts. Command `./txsgen fakeaccs` generates fakenet accounts.
See `txsgen help` for details.


## API-calls:

 - `./txsgen calls` deploys test Ballot contract, and generates a lot of calls to it then;


## Amount transfers:

 - `./txsgen transfers` generates a lot of transfer transactions;


## Configuration:

Use `./txsgen.toml` config file.
