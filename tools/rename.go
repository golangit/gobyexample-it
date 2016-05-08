package main

import (
    "fmt"
    "io/ioutil"
    "os"
    "regexp"
    "strings"
)

func main() {
    if len(os.Args) != 3 {
        fmt.Println(`Usage:
    go run tools/rename.go <from> <to>`)
        return
    }
    var (
        oldName   = os.Args[1]
        oldNameID = makeName(oldName)
        newName   = os.Args[2]
        newNameID = makeName(newName)
    )
    examplesRaw, err := ioutil.ReadFile("examples.txt")
    if err != nil {
        panic(err)
    }
    err = ioutil.WriteFile(
        "examples.txt",
        []byte(
            strings.Replace(string(examplesRaw), oldName+"\n", newName+"\n", 1),
        ), 0644,
    )
    if err != nil {
        panic(err)
    }
    files, err := ioutil.ReadDir("examples/" + oldNameID)
    if err != nil {
        panic(err)
    }
    for _, file := range files {
        fn := file.Name()
        if len(fn) >= len(oldNameID) && fn[:len(oldNameID)] == oldNameID {
            err = os.Rename(
                fmt.Sprintf("examples/%s/%s", oldNameID, fn),
                fmt.Sprintf("examples/%s/%s%s", oldNameID, newNameID, fn[len(oldNameID):]),
            )
            if err != nil {
                panic(err)
            }
        }
    }
    err = os.Rename("examples/"+oldNameID, "examples/"+newNameID)
    if err != nil {
        panic(err)
    }
    fmt.Println("done!")
}

var dashPat = regexp.MustCompile("\\-+")

func makeName(s string) string {
    s = strings.ToLower(s)
    s = strings.NewReplacer(
        " ", "-",
        "/", "-",
        "'", "",
    ).Replace(s)
    s = dashPat.ReplaceAllString(s, "-")
    return s
}
