package rule

import (
	"fmt"
	"github.com/google/uuid"
	"go_api/logger"
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
	LogAndHandleError(res.Error, fmt.Sprintf("Created Rule - %s", newRule.ID))
	return res.Error
}

func DeleteRuleLogic(input *proto_gen.DeleteRuleReq) error {
	var rule models.Rule
	res := models.DB.Where("id = ?", input.Id).Delete(&rule)
	LogAndHandleError(res.Error, fmt.Sprintf("Deleted Rule - %s", input.Id))
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
	LogAndHandleError(res.Error, fmt.Sprintf("Edited Rule - %s", input.Id))
	return res.Error
}

func GetAllRulesLogic() (*proto_gen.GetAllRulesRes, error) {
	var rules []models.Rule
	res := models.DB.Find(&rules)
	LogAndHandleError(res.Error, fmt.Sprintf("Get ALl Rules"))
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

func LogAndHandleError(err error, logStr string) {
	if err != nil {
		logger.LogError(err, "rule")
	} else {
		logger.LogInfo(logStr, "rule")
	}
}
