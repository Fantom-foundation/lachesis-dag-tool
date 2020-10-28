# Lachesis DAG tool

DAG export, analyze and convert tools.
See `lachesis-dag-tool help`.


## Run Neo4j graph-db:

 - `make neo4j` runs neo4j docker container on "bolt://localhost:7687" with `./docker/neo4j-data/` volume;
 - `make neo4j-sql` opens Neo4j Cypher shell;
 - `make stop` stops docker containers;


## Load DAG into Neo4j db

 - run Neo4j db;
 - `lachesis-dag-tool import --datadir=${LACHESIS_DATADIR}`;


## Read DAG from Neo4j db

 - run Neo4j db;
 - load DAG into Neo4j;
 - `lachesis-dag-tool read`;


## Read DAG from KV db

 - `lachesis-dag-tool read1 --datadir=${LACHESIS_DATADIR}`;
 - compare performance with Neo4j db;
