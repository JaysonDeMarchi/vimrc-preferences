package main

import (
    "fmt"
    "os"
    "os/user"
    "io/ioutil"
    "log"
    "path"
    "runtime"
)

func checkForErrors(err error) {
    if err != nil {
        log.Fatal(err)
    }
}

func getHomePath() string {
    usr, err := user.Current()
    checkForErrors(err)
    return usr.HomeDir
}

func getCurrentPath() string {
    _, filePath, _, ok := runtime.Caller(0)
    if !ok {
        log.Fatal("Error: Not possible to recover information for the current path\n")
    }
    return path.Dir(filePath)
}

func getFile(fileName string) *os.File {
    file, err := os.Open(fileName)
    checkForErrors(err)
    return file
}

func getContent(file *os.File) []byte {
    text, err := ioutil.ReadAll(file)
    checkForErrors(err)
    return text
}

func getPermissions(filePath string) os.FileMode {
    info, err := os.Stat(filePath)
    checkForErrors(err)
    mode := info.Mode()
    return mode
}

func main () {
    const VimrcFileName = ".vimrc"
    var currentVimrcPath string = getHomePath() + "/" + VimrcFileName
    updatedVimrc := getFile(getCurrentPath() + "/" + VimrcFileName)
    defer updatedVimrc.Close()
    updatedContent := getContent(updatedVimrc)
    permissions := getPermissions(currentVimrcPath)

    err := ioutil.WriteFile(currentVimrcPath, updatedContent, permissions)
    checkForErrors(err)
    fmt.Printf("Updated: %s\n", currentVimrcPath)
}
