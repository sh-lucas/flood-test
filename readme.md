## Flood Testing your API
This is a stupidly simple cli project written in golang.
It does not have any dependency other than golang, and should be easily instalable.

Do `flood-test <ip> <workers>`. Each worker is 1 req/s to the URL.
Also, with golang, your CLI should not get overloaded doing way too many requests in a single thread.

## Installing
bash```
go install github.com/sh-lucas/flood-test
```