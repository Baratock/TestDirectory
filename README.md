# TestDirectory

library to create temporary folders for testing in go.
TestDirectories give you the ability to assert created files for existance and content
Assertions are currently based on the golang [testify-package](https://github.com/stretchr/testify)

## Adding File

    dir := tempfiles.New()
    defer dir.Remove() // ensure that testdirectory is removed after creation
    
    
