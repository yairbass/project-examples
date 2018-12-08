package demo

import (
	"fmt"
	"os"

	log "github.com/sirupsen/logrus"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetLevel(log.DebugLevel)
}
func DecorateString(arg string) string {
	log.WithField("arg", arg).Info("myFunction was called!")
	return fmt.Sprintf("%s - by JFrog", arg)
}
