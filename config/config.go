package config

import (
	"encoding/json"
	"log"
	"os"
)

type config struct {
	CertFile       string
	KeyFile        string
	HTTPSMode      bool
	HTTPPort       string
	LocalhostMode  bool
	DomainName     string
	LoginWorks     bool
	FilesDir       string
	MathjaxDir     string
	JsCSSDir       string
	DeployBlog     bool
	TaskConfig     TaskConfig
	HashedPassword string
}

// TaskConfig stores configuration for task management database
type TaskConfig struct {
	DbURL      string
	DbUser     string
	DbPassword string
	DbType     string
}

// IsNotEmpty is opposite of IsEmpty
func (t TaskConfig) IsNotEmpty() bool {
	return !t.IsEmpty()
}

// IsEmpty checks whether taskConfig has all it's fields declared properly or not.
func (t TaskConfig) IsEmpty() bool {
	return t == TaskConfig{} || t.DbURL == "" || t.DbType == ""
}

// config stores the configuration
var conf config

const configFile string = "./config.json"

// CertFile -
func CertFile() string {
	return conf.CertFile
}

// KeyFile -
func KeyFile() string {
	return conf.KeyFile
}

// HTTPSMode -
func HTTPSMode() bool {
	return conf.HTTPSMode
}

// HTTPPort -
func HTTPPort() string {
	return conf.HTTPPort
}

// LocalhostMode -
func LocalhostMode() bool {
	return conf.LocalhostMode
}

// DomainName -
func DomainName() string {
	return conf.DomainName
}

// LoginWorks -
func LoginWorks() bool {
	return conf.LoginWorks
}

// FilesDir -
func FilesDir() string {
	return conf.FilesDir
}

// MathjaxDir -
func MathjaxDir() string {
	return conf.MathjaxDir
}

// JsCSSDir -
func JsCSSDir() string {
	return conf.JsCSSDir
}

// DeployBlog -
func DeployBlog() bool {
	return conf.DeployBlog
}

// HashedPassword -
func HashedPassword() string {
	return conf.HashedPassword
}

// TaskConfiguration returns configuration properties related to task server, which includes db details.
func TaskConfiguration() TaskConfig {
	return conf.TaskConfig
}

func init() {
	conf = initialize(configFile)
}

// initialize is there so I can test it.
// I don't know how to call init function in a test module
func initialize(configFile string) config {
	initconfigDefaults()

	file, err := os.Open(configFile)
	if err != nil {
		log.Fatal("Could not open file. Error: ", err)
	}
	defer file.Close()

	decoder := json.NewDecoder(file)
	var conf config
	err = decoder.Decode(&conf)

	if err != nil {
		log.Fatal("Error parsing json file: ", err)
	}

	log.Printf("%+v", conf)

	return conf
}

func initconfigDefaults() {
	conf.DomainName = "orakem.site"
	conf.LocalhostMode = true
	conf.LoginWorks = true
	conf.MathjaxDir = "../mathjax/"
	conf.JsCSSDir = "../js/"
	conf.FilesDir = "../files/"
	conf.DeployBlog = false
	conf.TaskConfig = TaskConfig{}
}
