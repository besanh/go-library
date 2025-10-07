# go-library
All Golang libraries

Forcing the gopkg updating new code
```
https://pkg.go.dev/github.com/besanh/go-library@<tag name>
```

## Run mock test
1. Setup .mockery.yml
```
packages:
  github.com/besanh/go-library/<folder name>:
    interfaces:
      <interface name>:
        config:
          # directory to output the mock file
          dir: ""
          # package name for the generated mock
          pkgname: "<package name>"
          # name of the mock struct
          structname: "<mock name>"
          # filename for the generated mock
          filename: "<file name>.go"
          # format the output using goimports
          formatter: "goimports"
          # overwrite existing files if they already exist
          force-file-write: true
```

2. CMD
```
cd <destination>
mockery
```