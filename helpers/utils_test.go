package helpers

import (
	"gorm.io/gorm/utils/tests"
	"strings"
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

func TestGetDirPath(t *testing.T) {
	path := GetDirPath("load_config")
	tests.AssertEqual(t, strings.HasSuffix(path, "load_config"), true)
}

func TestHttpCode2GrpcCode(t *testing.T) {
	tests.AssertEqual(t, HttpCode2GrpcCode(400), 3)
	tests.AssertEqual(t, HttpCode2GrpcCode(401), 16)
	tests.AssertEqual(t, HttpCode2GrpcCode(404), 5)
	tests.AssertEqual(t, HttpCode2GrpcCode(403), 7)
	tests.AssertEqual(t, HttpCode2GrpcCode(499), 1)
	tests.AssertEqual(t, HttpCode2GrpcCode(504), 4)
	tests.AssertEqual(t, HttpCode2GrpcCode(409), 6)
	tests.AssertEqual(t, HttpCode2GrpcCode(501), 11)
	tests.AssertEqual(t, HttpCode2GrpcCode(503), 14)
	tests.AssertEqual(t, HttpCode2GrpcCode(566), 13)
}

func TestError(t *testing.T) {
	err := Error(4000001)
	tests.AssertEqual(t, err.Error(), "rpc error: code = InvalidArgument desc = ")
}

func TestErrorf(t *testing.T) {
	err := Errorf(4010000, "blah")
	tests.AssertEqual(t, err.Error(), "rpc error: code = Unauthenticated desc = blah")
}
