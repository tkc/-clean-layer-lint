## clean-layer-lint

![test](https://github.com/tkc/clean-layer-lint/workflows/test/badge.svg?branch=master)

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


## Install

```bash
$ go install git@github.com:tkc/clean-layer-lint.git
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

## License

MIT ✨


