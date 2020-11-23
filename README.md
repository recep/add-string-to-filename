# File Name Editor-GO

**Add string to the end of the file name**  

```
go run main.go -p folder/ -ae <string> -ext .<extension_name>
```

**Add string to the beginning of the file name**  
```
go run main.go -p folder/ -ab <string> -ext .<extension_name>
```

**Rename**   
```
go run main.go -p folder/ -f <filename> -rn <new file name> 
```

**Show All Files**  

```
go run main.go -p folder2/ -ls
```
**Show Help Message**  
```
go run main.go -h
```
```bash
Usage: File Name Editor-GO
Options:
        -p, --path <path>                       path your folder 
        -ae <string>                            Add string to the end of the file name
        -ab <string>                            Add string to the beginning of the file name
        -f, --file <file name>                  specific file 
        -r, --rename <new file name>            new file name
        -u, --undo                              undo last change
Common Options: 
        -help                                   show help
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
go run main.go -p folder/ -ae -api -ext .json
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
go run main.go -p folder/ -f doc-1.txt -rn doc-2
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

- [x] command line print design  

- [x] flag options helper  
- [x] rename file 

  