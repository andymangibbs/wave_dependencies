Forked version of WAVE that works

### Requirements
1. Trillian: version 1.3.2

    `go get github.com/google/trillian`

    `cd $GOPATH/src/github.com/google/trillian`

    `git checkout 10ac7e3add265b489e5d3b33c9c355850b635e88`

2. This repo


### Build
1. Build Trillian servers:

  + `server/trillian_log_erver`
  + `server/trillian_log_signer`
  + `server/trillian_map_server`

2. Build wave
  `make`

  + Build init
  + Build server


### Run
1. Start Trillian servers:
  + Reset database:
      `scripts/resetdb.sh`
  + `trillian_map_server`
  + `./trillian_log_server -rpc_endpoint 127.0.0.1:8092 -http_endpoint 127.0.0.1:8093`
  + `./trillian_log_signer -rpc_endpoint 127.0.0.1:8094 -http_endpoint 127.0.0.1:8095 -batch_size 1 --force_master`

2. Init wave:
  + `init/init`
  + mysql -u root -p test < tables.sql
  + export variables from output of init
  + export DATABASE="root:@tcp(127.0.0.1:3306)/test"

3. Start wave server:
  + `cd init; ../server`

4. Start wave auditor:
  + `cd auditor; ./auditor config.toml`

5. Run test:
  + `cd storage/testbench`
  + `go test .`
