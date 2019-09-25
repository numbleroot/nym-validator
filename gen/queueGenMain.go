// queue_gen.go - template for creating infinite channel/queues.
// Copyright (C) 2019  Jedrzej Stuczynski.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package jobqueue is implemented based on eapache's infinite channel template:
package main

import (
	"flag"
	"os"
	"path/filepath"
	"text/template"
	"time"
)

type queueGen struct {
	Name           string // name of the queue, ex. JobQueue
	Type           string // full type that is being put on queue, ex. *jobpacket.JobPacket
	TypeName       string // name of the type being put on queue, ex. JobPacket
	TypeImportPath string // import path to the struct put on queue, ex.
	// github.com/nymtech/nym-validator/crypto/coconut/concurrency/jobpacket
}

func getOutputFileName() string {
	callee := os.Getenv("GOFILE")
	if len(callee) < 3 {
		return ""
	}
	return callee[:len(callee)-3] + "_queueGen.go"
}

// main is called via go generate from a specific file. It then fills the template defined below with the type specific
// information in order to define queue of specific type based on the eapache's infinite channel.
func main() {
	name := flag.String("name", "", "name of the queue, ex. JobQueue")
	fullType := flag.String("type", "", "full type that is being put on queue, ex. *jobpacket.JobPacket")
	typeName := flag.String("typeName", "", "name of the type being put on queue, ex. JobPacket")
	typeImportPath := flag.String(
		"typeImportPath",
		"",
		"import path to the struct put on queue, ex. github.com/nymtech/nym-validator/crypto/coconut/concurrency/jobpacket")

	flag.Parse()

	queueGen := queueGen{
		Name:           *name,
		Type:           *fullType,
		TypeName:       *typeName,
		TypeImportPath: *typeImportPath,
	}

	funcMap := template.FuncMap{
		"now":      time.Now,
		"fileName": getOutputFileName,
		"getenv":   os.Getenv,
	}

	tmpl := `// Code generated by queueGen; DO NOT EDIT.

// {{ fileName }} - Queue implementation for putting seemingly infinite number of {{ .TypeName }}s onto a channel.
// Copyright (C) 2019-{{ now.Year }}  Jedrzej Stuczynski.
//
// This program is free software: you can redistribute it and/or modify
// it under the terms of the GNU Affero General Public License as
// published by the Free Software Foundation, either version 3 of the
// License, or (at your option) any later version.
//
// This program is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
// GNU Affero General Public License for more details.
//
// You should have received a copy of the GNU Affero General Public License
// along with this program.  If not, see <http://www.gnu.org/licenses/>.

// Package {{ getenv "GOPACKAGE" }} is implemented based on eapache's infinite channel template:
// https://github.com/eapache/channels/blob/master/infinite_channel.go
// As explained by the author in documentation: https://godoc.org/github.com/eapache/channels due to Go's type system
// limitations direct import of his library is impractical.
// This reimplementation allows more type safety due to having more strict channel type than an empty interface.
// On top of better, even though slightly slower, queue as now it is thread-safe which is crucial as the queue
// can be read by multiple workers.
package {{ getenv "GOPACKAGE" }}

import (
	"fmt"

	"{{ .TypeImportPath }}"
	"github.com/Workiva/go-datastructures/queue"
)

// minQueueLen is the smallest capacity that queue may have.
const minQueueLen = 16

// {{ .Name }} represents a seemingly 'infinite' (depending on available memory) FIFO queue working on {{ .TypeName }} items.
type {{ .Name }} struct {
	input  chan {{ .Type }}
	output chan {{ .Type }}
	length chan int64
	buffer *queue.Queue
}

// In returns input channel for writing {{ .TypeName }}s.
func (ic *{{ .Name }}) In() chan<- {{ .Type }} {
	return ic.input
}

// Out returns output channel for reading {{ .TypeName }}s.
func (ic *{{ .Name }}) Out() <-chan {{ .Type }} {
	return ic.output
}

// Len returns number of elements in the queue.
func (ic *{{ .Name }}) Len() int64 {
	return <-ic.length
}

// Close closes the input channel, however, output and hence the goroutine will be open until the queue is exhausted.
func (ic *{{ .Name }}) Close() {
	close(ic.input)
}

func (ic *{{ .Name }}) infiniteBuffer() {
	var input, output chan {{ .Type }}
	var next {{ .Type }}
	input = ic.input

	for input != nil || output != nil {
		select {
		case elem, open := <-input:
			if open {
				err := ic.buffer.Put(elem)
				if err != nil {
					// there is nothing more we can do,
					// if we don't panic, all workers will block trying to read/write to closed channels.
					// Moreover the error should never be returned during normal operations.
					panic(fmt.Sprintf("The {{ .Name }} is in invalid state: %v", err))
				}
			} else {
				input = nil
			}
		case output <- next:
			_, err := ic.buffer.Get(1)
			// same rationale as with previous panics
			if err != nil {
				panic(fmt.Sprintf("The {{ .Name }} is in invalid state: %v", err))
			}
		case ic.length <- ic.buffer.Len():
		}

		if ic.buffer.Len() > 0 {
			output = ic.output
			nextT, err := ic.buffer.Peek()

			if err != nil {
				// Error can only happen if either the queue is already disposed what similarly to input should have
				// never happened accidentally and then there's nothing we can really do or if there are no items in the
				// queue which is also impossible due to explicit check.
				panic(fmt.Sprintf("The {{ .Name }} is in invalid state: %v", err))
			}
			next = nextT.({{ .Type }})

		} else {
			output = nil
			next = nil
		}
	}

	close(ic.output)
	close(ic.length)
}

// New creates a new instance of a {{ .Name }}.
func New() *{{ .Name }} {
	ic := &{{ .Name }}{
		input:  make(chan {{ .Type }}),
		output: make(chan {{ .Type }}),
		length: make(chan int64),
		buffer: queue.New(minQueueLen),
	}

	go ic.infiniteBuffer()
	return ic
}
`

	t := template.Must(template.New("queue").Funcs(funcMap).Parse(tmpl))

	// As per documentation: The generator is run in the package's source directory.
	cwd, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	outF := filepath.Join(cwd, getOutputFileName())

	f, err := os.OpenFile(outF, os.O_WRONLY|os.O_CREATE, 0664)
	if err != nil {
		panic(err)
	}

	if err := t.Execute(f, queueGen); err != nil {
		panic(err)
	}
}
