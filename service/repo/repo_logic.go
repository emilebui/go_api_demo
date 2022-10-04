package repo

import (
	"fmt"
	"go_api/logger"
	"go_api/models"
	"go_api/proto_gen"
)

func HelloWorldLogic() string {
	return "Hello World"
}

func CreateRepoLogic(input *proto_gen.CreateRepoReq) error {
	newRepo := models.Repo{
		Name: input.Name,
		Url:  input.Url,
	}
	res := models.DB.Create(&newRepo)
	LogAndHandleError(res.Error, fmt.Sprintf("Created Repo - %s", input.Name))
	return res.Error
}

func DeleteRepoLogic(input *proto_gen.DeleteRepoReq) error {
	var repo models.Repo
	res := models.DB.Where("name = ?", input.Name).Delete(&repo)
	LogAndHandleError(res.Error, fmt.Sprintf("Deleted Repo - %s", input.Name))
	return res.Error
}

func GetRepoLogic(input *proto_gen.GetRepoReq) (proto_gen.GetRepoResp, error) {
	var repo models.Repo
	res := models.DB.First(&repo, "name = ?", input.Name)
	LogAndHandleError(res.Error, fmt.Sprintf("Retrived Repo - %s", input.Name))
	return proto_gen.GetRepoResp{
		Name: repo.Name,
		Url:  repo.Url,
	}, res.Error
}

func UpdateRepoLogic(input *proto_gen.UpdateRepoReq) error {
	var repo models.Repo
	models.DB.First(&repo, "name = ?", input.Name)
	res := models.DB.Model(&repo).Updates(models.Repo{
		Name: input.Name,
		Url:  input.Url,
	})
	LogAndHandleError(res.Error, fmt.Sprintf("Updated Repo - %s", input.Name))
	return res.Error
}

func GetAllRepoLogic() (*proto_gen.GetAllRepoRes, error) {
	var repos []models.Repo
	res := models.DB.Find(&repos)
	LogAndHandleError(res.Error, fmt.Sprintf("Get All Repo"))
	if res.Error != nil {
		return nil, res.Error
	}
	var repoList []*proto_gen.GetRepoResp
	for _, r := range repos {
		repoList = append(repoList, &proto_gen.GetRepoResp{
			Name: r.Name,
			Url:  r.Url,
		})
	}

	return &proto_gen.GetAllRepoRes{
		Repositories: repoList,
	}, nil

}

func LogAndHandleError(err error, logStr string) {
	if err != nil {
		logger.LogError(err, "repo")
	} else {
		logger.LogInfo(logStr, "repo")
	}
}
