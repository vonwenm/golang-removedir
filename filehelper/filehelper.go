package filehelper

import (
	"os"
	"fmt"
)
  
// exists returns whether the given file or directory exists or not
func Exists(path string) bool {
    _, err := os.Stat(path)
    if err == nil { return true }
    if os.IsNotExist(err) { return false }
    return false
}

// remove directory
func RemoveDir(path string) bool {

      err := os.RemoveAll(path)

      if err != nil {
          fmt.Println(err)
          return false
      } else {
		  return true
	  }
 }
