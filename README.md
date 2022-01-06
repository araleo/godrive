# GoDrive

GoDrive is a Go CLI tool written to wrap the Google Drive API. 

## Features

- [x] List all files in the drive
- [ ] List all files and folders in a pretty way (tree)
- [x] Search for files by partial name
- [x] Upload a file
- [ ] Upload to a certain folder
- [x] Download a file
- [x] Export a document from Google Docs as unformated txt

## Usage

Build:

```
go build .
```

List all files in your Drive:
```
godrive ls
```

Search for a file in your Drive (searchs for files with names that contain the query):
```
godrive search <query>
```

Upload a file:
```
godrive up <filepath>
```

Download a file:
```
godrive down <file id>
```

Export a documento from Google Docs
```
godrive doc <file id>
```
