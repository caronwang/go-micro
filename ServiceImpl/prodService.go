package ServiceImpl

import (
	"context"
	Service "go-micro-grpc/Services"
	"strconv"
)
//测试方法
func newProd(id int32,pname string) *Service.ProdModel{
	return &Service.ProdModel{ProID: id,ProdName: pname}
}


type ProdService struct {

}

func (*ProdService) GetProdList(ctx context.Context,in *Service.ProdRequest,res *Service.ProdListResponse) error {
    models := make([]*Service.ProdModel,0)
    var i int32
    for i=0;i<in.Size;i++{
    	models = append(models,newProd(100+i+1,"prod_"+strconv.Itoa(100+int(i))))
	}
	res.Data=models
	return nil
}