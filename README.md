# AMO Formatters 

Formatters package create validation response like:
```
{
    "message": "Unprocessable Entity",
    "errors": [
        {
            "attribute": "email",
            "validator": {
                "name": "email"
            }
        },
        {
            "attribute": "password",
            "validator": {
                "name": "required"
            }
        }
    ]
}
```


## Install

```bash
go get github.com/amomama/amo_formatters
```

## Examples

### In fiber

```go
// Handler code
request := security.NewAuthenticateRequest(ctx, sh.validator)
err := request.Validate()

if err != nil {
    code, response := formatter.ValidationResponse(err)
    return ctx.Status(code).JSON(response)
}

```

### With net.http

```go
// Handler code
func (handler *MenuHandler) Create(w http.ResponseWriter, r *http.Request) {
    err = handler.validator.Validate(menuDTO)
    if err != nil {
        code, response := formatter.ValidationResponse(err)
        w.WriteHeader(code)

		data, _ := := json.Marshal(response)
		w.Write(data)
        return
    }
	
	...
}

```
