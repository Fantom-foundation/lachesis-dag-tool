# Lachesis DAG tool: txsgen

Transactions and API-calls generator for stress and performance testing of lachesis net.

## Accounts

Txs and Contract calls are created behalf of pre-created accounts.
Accounts are stored in the keystore dir withoutn any encryption.
Commands:

 -  `./txsgen fakeaccs` creates fakenet accounts;
 -  `./txsgen initbalance` gives them money from config.Payer;

See `txsgen help` for details.


## API-calls:

 - `./txsgen calls` deploys test Ballot contract, and generates a lot of calls to it then;


## Amount transfers:

 - `./txsgen transfers` generates a lot of transfer transactions;


## Configuration:

Use `./txsgen.toml` config file.
