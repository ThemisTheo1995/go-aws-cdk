package hello_test

import (
	"encoding/json"
	"testing"

	"github.com/aws/aws-cdk-go/awscdk/v2"
	"github.com/stretchr/testify/assert"
	"github.com/tidwall/gjson"

	"hello"
)

func TestHelloWorldStack(t *testing.T) {
	// GIVEN
	app := awscdk.NewApp(nil)

	// WHEN
	stack := hello.NewHelloWorldStack(app, "MyStack", nil)

	// THEN
	bytes, err := json.Marshal(app.Synth(nil).GetStackArtifact(stack.ArtifactId()).Template())
	if err != nil {
		t.Error(err)
	}

	template := gjson.ParseBytes(bytes)
	displayName := template.Get("Resources.MyTopic86869434.Properties.DisplayName").String()
	assert.Equal(t, "MyCoolTopic", displayName)
}
