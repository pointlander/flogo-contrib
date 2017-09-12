package rest

import (
	"context"
	"flag"
	"fmt"

	"github.com/TIBCOSoftware/flogo-lib/core/action"
	"github.com/TIBCOSoftware/flogo-lib/core/trigger"
	"github.com/TIBCOSoftware/flogo-lib/logger"
)


// log is the default package logger
var log = logger.GetLogger("trigger-tibco-cli")

var singleton *CliTrigger

// CliTrigger CLI trigger struct
type CliTrigger struct {
	metadata *trigger.Metadata
	runner   action.Runner
	config   *trigger.Config
	actions  []*actionInfo
	defAction *actionInfo
}

type actionInfo struct {
	actionId string
	invoke   bool
}

//NewFactory create a new Trigger factory
func NewFactory(md *trigger.Metadata) trigger.Factory {
	return &CliTriggerFactory{metadata: md}
}

// CliTriggerFactory CLI Trigger factory
type CliTriggerFactory struct {
	metadata *trigger.Metadata
}

//New Creates a new trigger instance for a given id
func (t *CliTriggerFactory) New(config *trigger.Config) trigger.Trigger {
	singleton := &CliTrigger{metadata: t.metadata, config: config}

	return singleton
}

// Metadata implements trigger.Trigger.Metadata
func (t *CliTrigger) Metadata() *trigger.Metadata {
	return t.metadata
}

func (t *CliTrigger) Init(runner action.Runner) {

	if t.config.Settings == nil {
		panic(fmt.Sprintf("No Settings found for trigger '%s'", t.config.Id))
	}

	if len(t.config.Handlers) == 0 {
		panic(fmt.Sprintf("No Handlers found for trigger '%s'", t.config.Id))
	}

	t.runner = runner
	hasDefault := false

	// Init handlers
	for _, handlerCfg := range t.config.Handlers {

		cmdString := "default"

		aInfo := &actionInfo{actionId:handlerCfg.ActionId, invoke:false}
		if cmd, ok := handlerCfg.Settings["command"]; ok {
			cmdString = cmd.(string)
		}

		if cmd, set := handlerCfg.Settings["default"]; set {
			if def, ok := cmd.(bool); ok && def {
				t.defAction = aInfo
				hasDefault = true
			}
		}

		t.actions = append(t.actions, aInfo)
		flag.BoolVar(aInfo.&invoke, cmdString, false, "")
	}

	if !hasDefault {
		t.defAction = t.actions[0]
	}
}

func (t *CliTrigger) Start() error {
	return nil
}

// Stop implements util.Managed.Stop
func (t *CliTrigger) Stop() error {
	return nil
}

func Invoke() (string,error) {

	flag.Parse()
	args := flag.Args()

	for _, value := range singleton.actions {

		if value.invoke {
			return singleton.Invoke(value.actionId, args)
		}
	}

	return singleton.Invoke(singleton.defAction.actionId, args)
}

func (t *CliTrigger) Invoke(actionId string, args []string) (string,error) {

	log.Infof("CLI Trigger: Invoking action '%s'", actionId)

	data := map[string]interface{}{
		"args":      args,
	}

	//todo handle error
	startAttrs, _ := t.metadata.OutputsToAttrs(data, false)

	act := action.Get(actionId)
	log.Debugf("Found action' %+x'", act)

	ctx := trigger.NewContext(context.Background(), startAttrs)
	_, replyData, err := t.runner.Run(ctx, act, actionId, nil)

	if err != nil {
		log.Debugf("CLI Trigger Error: %s", err.Error())
		return "", err
	}

	if replyData != nil {
		return replyData.(string), nil
	}

	return "", nil
}
