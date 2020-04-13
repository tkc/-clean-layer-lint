## clean-layer-lint
![test](https://github.com/tkc/clean-layer-lint/workflows/test/badge.svg)

```                                                                 
 _____ _                __                       __    _     _   
|     | |___ ___ ___   |  |   ___ _ _ ___ ___   |  |  |_|___| |_ 
|   --| | -_| .'|   |  |  |__| .'| | | -_|  _|  |  |__| |   |  _|
|_____|_|___|__,|_|_|  |_____|__,|_  |___|_|    |_____|_|_|_|_|  
                                 |___|                           
                                                          
```                                                


Clean architecture validator for go, like a The Dependency Rule and interaction between packages in your Go projects.
Checks illegal imported layer. layer order depend on `clean-layer.json` setteing.

![image](https://user-images.githubusercontent.com/181991/79132738-d273ab00-7de5-11ea-818f-1df4f93a782f.png)

## TODO

- [ ]  example
- [ ]  Specify files to ignore

## Install

```bash
$ go get -u github.com/tkc/clean-layer-lint
```


## Prepare clean-layer.json

```json
{
   "path": "github.com/tkc/clean-layer-lint/src",
   "order":[
       "domain",
       "usecase",
       "interfaces",
       "infrastructure"
    ],
    "ignore":[]
 }
```

`note:json format`

|  key  |  description  |
| :----: | :---- |
|  path  |  go.mod module name + your layder directory path  |
|  order  |  your layers order  |
|  ignore  |  your ignore layoers   |

## Run 
```bash
$ clean-layer-lint ./...
```

### GitHub Actions Setting
[lint.yml](https://github.com/tkc/clean-layer-lint/blob/master/.github/workflows/lint.yml)

```yml
on: [push, pull_request]
name: lint clean architecture
jobs:
  test:
    strategy:
      matrix:
        go-version: [1.13.x]
        platform: [ubuntu-latest, macos-latest, windows-latest]
    runs-on: ${{ matrix.platform }}
    steps:
    - name: Install Go
      if: success()
      uses: actions/setup-go@v1
      with:
        go-version: ${{ matrix.go-version }}
    - name: Setup env
      run: |
        echo "::set-env name=GOPATH::$(go env GOPATH)"
        echo "::set-env name=GOBIN::$(go env GOPATH)/bin"
        echo "::add-path::$(go env GOPATH)/bin"
    - name: install
      run: go get -u github.com/tkc/clean-layer-lint   
    - name: Checkout code
      uses: actions/checkout@v1
    - name: lint clean architecture
      run: clean-layer-lint ./...
```    

## License

MIT ✨


