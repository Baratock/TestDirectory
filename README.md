# TestDirectory

library to create temporary folders for testing in go

## Adding File

    dir := tempdirectory.New()
    defer dir.Remove() // ensure that testdirectory is removed after creation