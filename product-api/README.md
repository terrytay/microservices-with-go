## To recompile go-swagger:

Install binary from https://github.com/go-swagger/go-swagger in your OS (unix)

Then run the command:

```bash
make swagger
```

Each time swagger.yaml is configured, please manually go into info and add title.

## To regenerate the swagger HTTP clientL:

Run the following command in the dir `sdk`

```bash
swagger generate client -f ../swagger.yaml -A product-api
```
