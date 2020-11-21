# File Editor
Add string to the end of the file name
```
go run main.go -path=folder/ -ae=[string] -ext=.[extension_name]
```

Add string to the beginning of the file name
```
go run main.go -path=folder/ -ab=[string] -ext=.[extension_name]
```

### Examples  

**BEFORE**  

```bash
folder/
├── doc-1.txt
└── sw.json
```

**RUN**  

```
go run main.go -path=folder/ -ae=-api -ext=.json
```
**AFTER**  

```bash
folder/
├── doc-1.txt
└── sw-api.json
```


### TODO

- [x] add string to the end of the file name  

- [x] add string to the beginning of the file name  

- [x] extension filter  

- [ ] undo last change  

- [ ] command line print design  

- [ ] flag options helper  

  