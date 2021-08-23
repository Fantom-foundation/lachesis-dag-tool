# Opera tool: DAG-reader

DAG-reader subscribes to a running opera node API and exports new event to the db. See `dagreader help`.


## Run Neo4j db:

 - `make neo4j` runs neo4j docker container on "bolt://localhost:7687" with `./docker/neo4j-data/` volume;
 - `make neo4j-sql` opens Neo4j Cypher shell;
 - `make stop` stops docker containers;


## Load DAG into Neo4j db

 - run Neo4j db first;
 - from go-opera node: `dagreader [--api=ws://127.0.0.1:4500] [--dagstart=1] saveto [--neo4j=bolt://localhost:7687]`;

Use 'dagstart' param to skip genesis blocks (4564024 for mainnet).


## Read DAG from Neo4j db

Field 'role' hints event consensus role (atropos or not).
Role which ends with "*" means that event is detected but not found in the node datadir.

 - run Neo4j db;
 - load DAG into Neo4j;
 - open Cypher console shell `make neo4j-sql`;
 - query example: find event ancestors
```
@neo4j> MATCH (p:Event {id: "11:83:c4fa5e79b127a9140705d0b78f1c42886cd9cffc167a46d8"})-[:PARENT]->(s:Event) RETURN DISTINCT s.id;
+----------------------------------------------------------+
| s.id                                                     |
+----------------------------------------------------------+
| "11:83:001ca14eff90a3c22e9293f020f946f67544e51f868c3077" |
| "11:78:a45de96510b9695338d42a9a660037b649f473e8c1707d8c" |
+----------------------------------------------------------+
2 rows available after 69 ms, consumed after another 51 ms
:exit
```