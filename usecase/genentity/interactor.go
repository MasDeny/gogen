package genentity

import (
  "context"
  "fmt"
  "github.com/mirzaakhena/gogen/domain/entity"
  "github.com/mirzaakhena/gogen/domain/service"
)

//go:generate mockery --name Outport -output mocks/

type genEntityInteractor struct {
  outport Outport
}

// NewUsecase is constructor for create default implementation of usecase GenEntity
func NewUsecase(outputPort Outport) Inport {
  return &genEntityInteractor{
    outport: outputPort,
  }
}

// Execute the usecase GenEntity
func (r *genEntityInteractor) Execute(ctx context.Context, req InportRequest) (*InportResponse, error) {

  res := &InportResponse{}

  // buat object entity
  obj, err := entity.NewObjEntity(req.EntityName)
  if err != nil {
    return nil, err
  }

  _, err = r.outport.CreateFolderIfNotExist(ctx, "domain/entity/")
  if err != nil {
    return nil, err
  }

  exist, err := obj.IsEntityExist()
  if err != nil {
    return nil, err
  }

  if exist {
    return res, nil
  }

  err = service.CreateEverythingExactly("default/", "domain/entity", map[string]string{
    "entityname": obj.EntityName.SnakeCase(),
  }, obj.GetData())
  if err != nil {
    return nil, err
  }

  // reformat entity.go
  err = r.outport.Reformat(ctx, fmt.Sprintf("domain/entity/%s.go", obj.EntityName.SnakeCase()), nil)
  if err != nil {
    return nil, err
  }

  return res, nil
}
