# go-google-vision-api
Access commands to the Google Cloud Vision API

# Use(example)

Get Binary(Latest) google-vision : [https://github.com/fkei/go-google-vision-api/releases/latest](https://github.com/fkei/go-google-vision-api/releases/latest)

```
$ ./google-vision -api-key=XXXXXXXXXX -path=test.png
----
1:
{"labelAnnotations":[{"description":"black cat","mid":"/m/03dj64","score":0.88111246},{"description":"silhouette","mid":"/m/03thgk","score":0.87085229},{"description":"teapot","mid":"/m/01fh4r","score":0.52110529}],"safeSearchAnnotation":{"adult":"VERY_UNLIKELY","medical":"VERY_UNLIKELY","spoof":"VERY_UNLIKELY","violence":"UNLIKELY"}}
```

### Support Features types

- Type: "FACE_DETECTION"
- Type: "LANDMARK_DETECTION"
- Type: "LOGO_DETECTION",
- Type: "LABEL_DETECTION",
- Type: "TEXT_DETECTION",
- Type: "SAFE_SEARCH_DETECTION",
- Type: "IMAGE_PROPERTIES",
- MaxResults: 10


### Help

```
$ ./google-vision --help
Usage of ./google-vision:
  -api-key="": Google API KEY
  -path="": Image file path
```

# Build

```
$ make build
```

# Source Run

```
$ make run
```
