package marathon

import (
	"encoding/json"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestExternalVolumeParsing(t *testing.T) {
	input := `{
  "id": "app1",
  "instances": 1,
  "cpus": 1,
  "mem": 1024,
  "cmd": "echo yes > xxx/file && sleep 3600",
  "container": {
    "type": "MESOS",
    "volumes": [
      {
        "containerPath": "xxx",
        "mode": "rw",
        "external": {
          "provider": "csi",
          "name": "no-need",
          "options": {
            "pluginName": "nfs.csi.k8s.io",
            "capability": {
              "accessType": "mount",
              "accessMode": "MULTI_NODE_MULTI_WRITER",
              "fsType": "nfs"
            },
            "volumeContext": {
              "server": "172.16.10.137",
              "share": "/mnt"
            }
          }
        }
      }
    ]
  }
}`

	var app Application
	err := json.Unmarshal([]byte(input), &app)
	assert.NoError(t, err)
	v := *app.Container.Volumes
	assert.Equal(t, *v[0].External.Options, map[string]interface{}{
		"pluginName": "nfs.csi.k8s.io",
		"capability": map[string]interface{}{
			"accessType": "mount",
			"accessMode": "MULTI_NODE_MULTI_WRITER",
			"fsType":     "nfs",
		},
		"volumeContext": map[string]interface{}{
			"server": "172.16.10.137",
			"share":  "/mnt",
		},
	})
}
