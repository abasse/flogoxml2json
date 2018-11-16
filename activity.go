package flogoxml2json

import (
	"strings"

	"github.com/TIBCOSoftware/flogo-lib/core/activity"
	"github.com/TIBCOSoftware/flogo-lib/logger"
	xj "github.com/basgys/goxml2json"
)

// activityLog is the default logger for the exec Activity
var log = logger.GetLogger("activity-tibco-flogoxml2json")

type Xml2jsonActivity struct {
	metadata *activity.Metadata
}

// NewActivity creates a new AppActivity
func NewActivity(metadata *activity.Metadata) activity.Activity {
	return &Xml2jsonActivity{metadata: metadata}
}

// Metadata returns the activity's metadata
func (a *Xml2jsonActivity) Metadata() *activity.Metadata {
	return a.metadata
}

func (a *Xml2jsonActivity) Eval(context activity.Context) (done bool, err error) {

	xmlinput, _ := context.GetInput("Input").(string)

	xml := strings.NewReader(xmlinput)
	json, err := xj.Convert(xml)
	if err != nil {
		log.Info(err)
	}

	context.SetOutput("result", json.String())

	return true, nil
}
