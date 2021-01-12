package manifest

import (
	"fmt"
	"io"
	"io/ioutil"
	"strings"
	"sync"
	"time"

	"github.com/sampointer/digg/fetcher"
)

type manifest struct {
	docs []string
	lock sync.Mutex
}

var m manifest

func init() {
	m.update()
	m.cron()
}

//GetManifest is a thread safe Getter for the manifest document
func GetManifest() []io.Reader {
	var docs []io.Reader

	m.lock.Lock()
	for _, doc := range m.docs {
		docs = append(docs, strings.NewReader(doc))
	}
	m.lock.Unlock()
	return docs
}

func (m *manifest) cron() {
	ticker := time.NewTicker(1 * time.Hour)
	go func() {
		for {
			select {
			case <-ticker.C:
				fmt.Println("update of Google manifest begins")
				m.update()
				fmt.Println("update of Google manifest ends")
			}
		}
	}()
}

func (m *manifest) update() {
	var newDocs []string

	docs, err := fetcher.Fetch()
	if err != nil {
		fmt.Println("failed to update Google manifest")
		return
	}

	for _, doc := range docs {
		b, err := ioutil.ReadAll(doc)
		if err != nil {
			fmt.Println("failed to read Google manifest during update")
			return
		}
		newDocs = append(newDocs, string(b))
	}

	m.lock.Lock()
	m.docs = newDocs
	m.lock.Unlock()
}
