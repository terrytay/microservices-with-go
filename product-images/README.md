To test file upload:

Empty the imagestore folder first then:

```
curl localhost:9090/images/1/test.png -XPOST --data-binary @frap.png
```
