package service

import (
	"fmt"
	"log"
	"sync"
)

// Builder ensures that each service is initialized only once
type Builder interface {
	Build(args ...interface{}) error
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
	if err := bb.Build(args...); err != nil {
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

// Build is default implementation
func (bb *BaseBuilder) Build(args ...interface{}) error {
	panic("Not implemented!")
}
