package helpers

import (
	"testing"
)

func TestErrorHandler(t *testing.T) {
	ErrorHandler(nil)
}

func TestFindLoadConfigPath(t *testing.T) {
	path := "C:\\Users\\Altair\\GolandProjects\\guardrail_demo\\helpers\\"
	result := FindLoadConfigPath(path)
	println(result)

}
