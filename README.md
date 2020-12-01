# File Name Editor-GO

**Add string to the end of the file name**  

```
go run main.go -p folder/ -ae <string> -ext .<extension_name>
```

**Add string to the beginning of the file name**  

```
go run main.go -p folder/ -ab <string> -ext .<extension_name>
```

**Rename a file**   

```
go run main.go -p folder/ -f <filename> -rn <new file name> 
```

**Show All Files**

```
go run main.go -p folder2/ -ls
```
**Undo Last Change**  
```
go run main.go -u
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
        -rn, --rename <new file name>           new file name
        -u, --undo                              undo last change
        -ls										show all files
Common Options: 
        -help                                   show help message
```
### Examples     
---
* Add string to the end of the file name  

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
* Rename a file  

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
### LICENSE  
The MIT License (MIT) - see [LICENSE.md](https://github.com/recep/add-string-to-filename/blob/master/LICENSE.md) for more details

