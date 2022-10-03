package repo

import (
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
	return res.Error
}

func DeleteRepoLogic(input *proto_gen.DeleteRepoReq) error {
	var repo models.Repo
	res := models.DB.Where("name = ?", input.Name).Delete(&repo)
	return res.Error
}

func GetRepoLogic(input *proto_gen.GetRepoReq) (proto_gen.GetRepoResp, error) {
	var repo models.Repo
	res := models.DB.First(&repo, "name = ?", input.Name)
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
	return res.Error
}

func GetAllRepoLogic() (*proto_gen.GetAllRepoRes, error) {
	var repos []models.Repo
	res := models.DB.Find(&repos)
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
