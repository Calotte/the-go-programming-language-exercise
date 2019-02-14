package du1

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"sync"
	"time"
)

var sema = make(chan struct{}, 20)
var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

func dirents(dir string) []os.FileInfo {
	select {
	case sema <- struct{}{}:
	case <-done:
			return nil
	}
	defer func() { <-sema }()
	entries, err := ioutil.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "du1:%v\n", err)
		return nil
	}
	return entries
}

func walkDir(dir string, fileSz chan<- int64) {
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			subdir := filepath.Join(dir, entry.Name())
			walkDir(subdir, fileSz)
		} else {
			fileSz <- entry.Size()
		}
	}
}

func walkDir2(dir string, n *sync.WaitGroup, fileSz chan<- int64) {
	defer n.Done()
	if cancelled() {
		return
	}
	for _, entry := range dirents(dir) {
		if entry.IsDir() {
			n.Add(1)
			subdir := filepath.Join(dir, entry.Name())
			go walkDir2(subdir, n, fileSz)
		} else {
			fileSz <- entry.Size()
		}
	}
}

func printDiskUsage(nfiles, nbyte int64) {
	fmt.Printf("%d files,%.fGB\n", nfiles, float64(nbyte/1e9))
}

func Du2() {
	go func() {
		os.Stdin.Read(make([]byte, 1))
		close(done)
	}()

	var verbose = flag.Bool("v", false, "show verbose progress messages")
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	var tick <-chan time.Time
	if *verbose {
		tick = time.Tick(500 * time.Millisecond)
	}
	fileSizes := make(chan int64)
	var n sync.WaitGroup
	for _, root := range roots {
		n.Add(1)
		go walkDir2(root, &n, fileSizes)
	}
	go func() {
		n.Wait()
		close(fileSizes)
	}()

	var nfiles, nbytes int64
loop:
	for {
		select {
		case <-done:
			for range fileSizes {
			}
			return
		case sz, ok := <-fileSizes:
			if !ok {
				break loop
			}
			nfiles++
			nbytes += sz
		case <-tick:
			printDiskUsage(nfiles, nbytes)
		}
	}
	printDiskUsage(nfiles, nbytes)
}

func Du1() {
	flag.Parse()
	roots := flag.Args()
	if len(roots) == 0 {
		roots = []string{"."}
	}
	fileSizes := make(chan int64)
	go func() {
		for _, root := range roots {
			walkDir(root, fileSizes)
		}
		close(fileSizes)
	}()
	var nfiles, nbytes int64
	for size := range fileSizes {
		nfiles++
		nbytes += size
	}
	printDiskUsage(nfiles, nbytes)
}
