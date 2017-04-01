package illuminator

import (
	"encoding/json"
	"sync"

	log "github.com/Sirupsen/logrus"
	MQTT "github.com/eclipse/paho.mqtt.golang"

	"github.com/pkedy/ci-illuminator/pipeline"
)

var (
	executionLock    sync.Mutex
	currentExecution *pipeline.Execution
	statusPipelines  = map[string]*pipeline.Pipeline{
		"building": StateBuilding,
		"success":  StateSuccess,
		"failure":  StateFailure,
	}
)

type Command struct {
	Status string `json:"status"`
}

func HandleNotification(client MQTT.Client, message MQTT.Message) {
	log.Debugf("Received: %s", string(message.Payload()))
	var command Command
	err := json.Unmarshal(message.Payload(), &command)
	if err != nil {
		log.Errorf("Could not parse message: %s", err)
		return
	}

	if pipeline, ok := statusPipelines[command.Status]; ok {
		executionLock.Lock()
		log.Infof("Starting pipeline %s", command.Status)
		if currentExecution != nil {
			currentExecution.Stop(false)
		}

		currentExecution = pipeline.Start()
		executionLock.Unlock()
	}
}

func StopExecutions() {
	executionLock.Lock()
	if currentExecution != nil {
		currentExecution.Stop(false)
	}
	executionLock.Unlock()
}
