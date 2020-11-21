# File Name Editor-GO

Add string to the end of the file name
```
go run main.go -path folder/ -ae [string] -ext .[extension_name]
```

Add string to the beginning of the file name
```
go run main.go -path folder/ -ab [string] -ext .[extension_name]
```

Rename   
```
go run main.go -path folder/ -f [filename] -rn [new file name] 
```
### Examples     
---
**Before**  

```bash
folder/
├── doc-1.txt
└── sw.json
```

**Run**  

```
go run main.go -path folder/ -ae -api -ext .json
```
**After**  

```bash
folder/
├── doc-1.txt
└── sw-api.json *
```
---
**Run**  
```
go run main.go -path folder/ -f doc-1.txt -rn doc-2
```
```bash
folder/
├── doc-2.txt *
└── sw-api.json
```

---
### TODO

- [x] add string to the end of the file name  

- [x] add string to the beginning of the file name  

- [x] extension filter   

- [ ] undo last change  

- [ ] command line print design  

- [ ] flag options helper  
- [x] rename file 

  