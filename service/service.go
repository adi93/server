// package service provides various services on domain objects.
//
// To register a new service, do the following:
//
// 1. Define an interface for the service. Also define a unique name for the service.
// Make sure that it does not clash with an existing name.
//
// 2. Create a concrete type, containing the repositories required for the service to work.
// Ideally, one service should have access to a single repo. Implement the interface by this type.
//
// 3. Create a builder struct, which embeds a pointer to BaseBuilder. Also, implement the
//	buildInternal(args ...interface{}) error
// method, in which repos are taken as input, and wired to the service.
//
// 3. Declare an init() method. In that method, create a serviceBuilder with
// the unique name created above, and add it to the "Initializers" map, defined in service.go
//
// 4. Finally, provide a
//	func Initialize<ServiceName>(<input repositories as arguments>) error
// method, which is the public interface to main.go
package service

import (
	"fmt"
	"log"
	"sync"
)

// Builder ensures that each service is initialized only once
type Builder interface {
	buildInternal(args ...interface{}) error
	Name() string
	lock()
	unlock()
	isInitialized() bool
	initialize()
}

// Initializers store the builders of each service
var Initializers map[string]Builder = make(map[string]Builder, 0)

// BaseBuilder forms the base of all builders.
type BaseBuilder struct {
	name        string
	initialized bool
	sync.Mutex
}

// NewBaseBuilder returns a base builder
func NewBaseBuilder(name string, initialized bool) BaseBuilder {
	return BaseBuilder{name, initialized, sync.Mutex{}}
}

func build(bb Builder, args ...interface{}) error {
	log.Printf("")
	log.Printf("Building %s service", bb.Name())
	bb.lock()
	defer bb.unlock()
	if bb.isInitialized() == true {
		return fmt.Errorf("Initializing %s again! Not allowed", bb.Name())
	}
	if err := bb.buildInternal(args...); err != nil {
		log.Printf("Error building %s service: %s", bb.Name(), err.Error())
		return err
	}
	bb.initialize()
	log.Printf("Successfully built %s service", bb.Name())
	return nil
}

func (bb *BaseBuilder) lock() {
	log.Printf("Acquiring lock for %s", bb.Name())
	bb.Lock()
	log.Printf("Acquired lock for %s", bb.Name())
}

func (bb *BaseBuilder) unlock() {
	log.Printf("Releasing lock for %s", bb.Name())
	bb.Unlock()
	log.Printf("Released lock for %s", bb.Name())
}

func (bb *BaseBuilder) isInitialized() bool {
	return bb.initialized
}
func (bb *BaseBuilder) initialize() {
	bb.initialized = true
}

// Name is default implementation
func (bb *BaseBuilder) Name() string {
	return bb.name
}

// buildInternal is default implementation, to be overridden by the struct embedding *BaseBuilder
func (bb *BaseBuilder) buildInternal(args ...interface{}) error {
	panic("Not implemented!")
}
