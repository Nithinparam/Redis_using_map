# Redis Using Map and Struct

## In memory data structure using maps and struct

###  SetData
*   Sets the key value pair into the json file

###  GetData
*   GetData gets the value of a specified key
###  UpadteData
*   UpdateData updates the data based on command  `inc` or `dec` i.e, Increment and Decrement
###  DeleteData
*   DeleteData deletes the data of specified key

##  Output:
```Go
>   go run main.go -cmd set -k a -v 1
>   go run main.go -cmd get -k a
>   go run main.go -cmd inc -k a
>   go run main.go -cmd dec -k a
>   go run main.go -cmd del -k a
```
