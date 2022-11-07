package output

import (
	"encoding/json"
	"log"
	"time"
)

var OutputStruct Output

type Output struct {
	Logging Logging `yaml:"logging" json:"logging"`
	Info    Info   `yaml:"info" json:"info"`
}

type Info struct {
	AppName string `yaml:"appname,omitempty" json:"appname,omitempty"`
	Version string `yaml:"version,omitempty" json:"version,omitempty"`
}

type Logging struct {
	Lines []Line `yaml:"line" json:"line"`
}

type Line struct {
	Timestamp string `yaml:"timestamp,omitempty" json:"timestamp,omitempty"`
	Command   string `yaml:"command,omitempty" json:"command,omitempty"`
	Type      string `yaml:"type,omitempty" json:"type,omitempty"`
	Message   string `yaml:"message,omitempty" json:"message,omitempty"`
}

const (
	LogTypeError   = "ERROR"
	LogTypeWarning = "WARNING"
	LogTypeInfo    = "INFO"
	LogTypeDebug   = "DEBUG"
)

func (outObj *Output) AddLoggingLine(logType, command, message string) {
	line := Line{
		Timestamp: time.Now().String(),
		Type:      logType,
		Command:   command,
		Message:   message,
	}
	outObj.Logging.Lines = append(outObj.Logging.Lines, line)
}

func (outObj *Output) PrintLogging() {
		marshalledOutput, err := json.Marshal(outObj)
		if err != nil {
			log.Printf("failed output marshal: %s", err)
		}
		log.Printf("%s", string(marshalledOutput))
}
