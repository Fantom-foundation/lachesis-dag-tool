# Lachesis DAG tool: exporter

DAG export, analyze and convert tool.
See `exporter help`.


## Run Neo4j graph-db:

 - `make neo4j` runs neo4j docker container on "bolt://localhost:7687" with `./docker/neo4j-data/` volume;
 - `make neo4j-sql` opens Neo4j Cypher shell;
 - `make stop` stops docker containers;


## Load DAG into Neo4j db

 - run Neo4j db first;
 - from go-lachesis datadir: `importer import --datadir=${LACHESIS_DATADIR} [epoch_from [epoch_to]]`;
 - from go-lachesis p2p net: `importer listen --network=main [epoch_from [epoch_to]]`;


## Read DAG from Neo4j db

 - run Neo4j db;
 - load DAG into Neo4j;
 - `importer read "0x0000000400000016f9a4c23827a98e8dfa1358a41eb79d71e889c97c973722ab"`;


## Read DAG from KV db

 - `importer read1 --datadir=${LACHESIS_DATADIR} "0x0000000400000016f9a4c23827a98e8dfa1358a41eb79d71e889c97c973722ab"`;
 - compare performance with Neo4j db;
