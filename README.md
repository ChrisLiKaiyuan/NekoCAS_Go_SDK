# NekoCAS_Go_SDK
NekoCAS SDK for Go

## Installation
```
go get -u github.com/NekoWheel/NekoCAS_Go_SDK
```

## Begin to use
```go
// Input the NekoCAS URL and the service secret.
cas := NekoCAS.New("https://cas.n3ko.co", "vNOZpKdqnUYcztBjUhvvPLpeYCIIBVev")

// Validate the ticket.
user, err := cas.Validate("ST-oadwZVbq5lUy151InUCC6UHDLI2l586k")
if err != nil {
    log.Println(err)
    return
} else {
    log.Println(user.Token)
}
```