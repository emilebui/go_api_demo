package rule

import (
	"go_api/models"
	"go_api/proto_gen"
	"gorm.io/gorm/utils/tests"
	"testing"
)

func TestAddRuleLogic(t *testing.T) {
	models.DB = models.InitTestDB()

	err := AddRuleLogic(&proto_gen.AddRule{
		Description:   "blah",
		Severity:      "blah",
		StringCompare: "blah",
	})

	if err != nil {
		t.Fatalf("fail in logic")
	}
}

func TestEditRuleLogic(t *testing.T) {
	models.DB = models.InitTestDB()

	err := EditRuleLogic(&proto_gen.Rule{
		Id:            "test",
		Description:   "blah",
		Severity:      "blah",
		StringCompare: "blah",
	})

	if err != nil {
		t.Fatalf("fail in logic")
	}

	err = EditRuleLogic(&proto_gen.Rule{
		Id:            "test2",
		Description:   "blah",
		Severity:      "blah",
		StringCompare: "blah",
	})

	if err == nil {
		t.Fatalf("fail in logic")
	}
}

func TestDeleteRuleLogic(t *testing.T) {
	models.DB = models.InitTestDB()

	err := DeleteRuleLogic(&proto_gen.DeleteRuleReq{
		Id: "test",
	})

	if err != nil {
		t.Fatalf("fail in logic")
	}
}

func TestGetAllRulesLogic(t *testing.T) {
	models.DB = models.InitTestDB()

	res, err := GetAllRulesLogic()

	if err != nil {
		t.Fatalf("fail in logic")
	}

	tests.AssertEqual(t, len(res.Rules), 1)
}
