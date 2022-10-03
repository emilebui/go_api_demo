package rule

import (
	"github.com/google/uuid"
	"go_api/models"
	"go_api/proto_gen"
)

func AddRuleLogic(in *proto_gen.AddRule) error {
	newRule := models.Rule{
		ID:            uuid.NewString(),
		Description:   in.Description,
		Severity:      in.Severity,
		StringCompare: in.StringCompare,
	}
	res := models.DB.Create(&newRule)
	return res.Error
}

func DeleteRuleLogic(input *proto_gen.DeleteRuleReq) error {
	var rule models.Rule
	res := models.DB.Where("id = ?", input.Id).Delete(&rule)
	return res.Error
}

func EditRuleLogic(input *proto_gen.Rule) error {
	var rule models.Rule
	models.DB.First(&rule, "id = ?", input.Id)
	res := models.DB.Model(&rule).Updates(models.Rule{
		Description:   input.Description,
		Severity:      input.Severity,
		StringCompare: input.StringCompare,
	})
	return res.Error
}

func GetAllRulesLogic() (*proto_gen.GetAllRulesRes, error) {
	var rules []models.Rule
	res := models.DB.Find(&rules)
	if res.Error != nil {
		return nil, res.Error
	}
	var ruleList []*proto_gen.Rule
	for _, r := range rules {
		ruleList = append(ruleList, &proto_gen.Rule{
			Id:            r.ID,
			Description:   r.Description,
			Severity:      r.Severity,
			StringCompare: r.StringCompare,
		})
	}

	return &proto_gen.GetAllRulesRes{
		Rules: ruleList,
	}, nil
}
