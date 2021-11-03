# Product Images

## Uploading REST

Empty the imagestore folder first then run:

```
curl localhost:9090/images/1/test.png -XPOST --data-binary @frap.png
```

## Uploading Multipart Form

EMpty the imagestore folder first then run:

```
curl localhost:9091/ -F 'id=1' -F 'file=@./frap.png'
```
