// main.go
package main

import (
    "github.com/dorsJfrog/testvendor/mypackage"
    "github.com/sirupsen/logrus"
)

func main() {
    logrus.Info("Starting the dummy Go package...")
    mypackage.Hello("World")
}
