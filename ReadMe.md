# ReadMe

## Layout
* https://github.com/golang-standards/project-layout

### CSV
Goes in `./configs/csv/server.csv`

or optionally specify custom path `-csv-in`

### Output Folder
Defaults to `./output_dir`

or optionally specify custom path `-outdir`

## Run

To generate new config files
```shell
$ ./prom-agent-config -init -csv-in configs/csv/AzMachines.csv -outdir myfolder
```

# Build Windows exe
`OOS=windows go build ./cmd/promconfig/prom-agent-config.go`

# ToDo

## Update from git patches?

```shell
git format-patch e5450e6bb0e030545605f2172b11c6fb54cac21c...925071850721e8ad65dcb2deeba3c19fc68715cb --stdout > ../prom-agent-config/git-patches/1commit.patch
```

* Need to parse Git Diff lines
* CSV reader on a string/
* Only care about + (Adds)