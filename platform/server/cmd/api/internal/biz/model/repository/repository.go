// Code generated by thriftgo (0.2.12). DO NOT EDIT.

package repository

import (
	"context"
	"fmt"

	"github.com/cloudwego/cwgo/platform/server/shared/kitex_gen/model"
)

type AddRepositoryReq struct {
	RepositoryType int32  `thrift:"repository_type,1" form:"repository_type,required" json:"repository_type,required"`
	RepositoryURL  string `thrift:"repository_url,2" form:"repository_url,required" json:"repository_url,required" vd:"len($)>0"`
	Branch         string `thrift:"branch,3" form:"branch,required" json:"branch,required"`
	StoreType      int32  `thrift:"store_type,4" form:"store_type,required" json:"store_type,required"`
}

func NewAddRepositoryReq() *AddRepositoryReq {
	return &AddRepositoryReq{}
}

func (p *AddRepositoryReq) GetRepositoryType() (v int32) {
	return p.RepositoryType
}

func (p *AddRepositoryReq) GetRepositoryURL() (v string) {
	return p.RepositoryURL
}

func (p *AddRepositoryReq) GetBranch() (v string) {
	return p.Branch
}

func (p *AddRepositoryReq) GetStoreType() (v int32) {
	return p.StoreType
}

func (p *AddRepositoryReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddRepositoryReq(%+v)", *p)
}

type AddRepositoryRes struct {
	Code int32  `thrift:"code,1" form:"code" json:"code" query:"code"`
	Msg  string `thrift:"msg,2" form:"msg" json:"msg" query:"msg"`
}

func NewAddRepositoryRes() *AddRepositoryRes {
	return &AddRepositoryRes{}
}

func (p *AddRepositoryRes) GetCode() (v int32) {
	return p.Code
}

func (p *AddRepositoryRes) GetMsg() (v string) {
	return p.Msg
}

func (p *AddRepositoryRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("AddRepositoryRes(%+v)", *p)
}

type DeleteRepositoriesReq struct {
	Ids []int64 `thrift:"ids,1" form:"ids,required" json:"ids,required" vd:"len($)>0"`
}

func NewDeleteRepositoriesReq() *DeleteRepositoriesReq {
	return &DeleteRepositoriesReq{}
}

func (p *DeleteRepositoriesReq) GetIds() (v []int64) {
	return p.Ids
}

func (p *DeleteRepositoriesReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteRepositoriesReq(%+v)", *p)
}

type DeleteRepositoriesRes struct {
	Code int32  `thrift:"code,1" form:"code" json:"code" query:"code"`
	Msg  string `thrift:"msg,2" form:"msg" json:"msg" query:"msg"`
}

func NewDeleteRepositoriesRes() *DeleteRepositoriesRes {
	return &DeleteRepositoriesRes{}
}

func (p *DeleteRepositoriesRes) GetCode() (v int32) {
	return p.Code
}

func (p *DeleteRepositoriesRes) GetMsg() (v string) {
	return p.Msg
}

func (p *DeleteRepositoriesRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("DeleteRepositoriesRes(%+v)", *p)
}

type UpdateRepositoryReq struct {
	ID     int64  `thrift:"id,1" form:"id" json:"id"`
	Branch string `thrift:"branch,2" form:"branch" json:"branch"`
	Status int32  `thrift:"status,3" form:"status" json:"status"`
}

func NewUpdateRepositoryReq() *UpdateRepositoryReq {
	return &UpdateRepositoryReq{}
}

func (p *UpdateRepositoryReq) GetID() (v int64) {
	return p.ID
}

func (p *UpdateRepositoryReq) GetBranch() (v string) {
	return p.Branch
}

func (p *UpdateRepositoryReq) GetStatus() (v int32) {
	return p.Status
}

func (p *UpdateRepositoryReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateRepositoryReq(%+v)", *p)
}

type UpdateRepositoryRes struct {
	Code int32  `thrift:"code,1" form:"code" json:"code" query:"code"`
	Msg  string `thrift:"msg,2" form:"msg" json:"msg" query:"msg"`
}

func NewUpdateRepositoryRes() *UpdateRepositoryRes {
	return &UpdateRepositoryRes{}
}

func (p *UpdateRepositoryRes) GetCode() (v int32) {
	return p.Code
}

func (p *UpdateRepositoryRes) GetMsg() (v string) {
	return p.Msg
}

func (p *UpdateRepositoryRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("UpdateRepositoryRes(%+v)", *p)
}

type GetRepositoriesReq struct {
	Page             int32  `thrift:"page,1" json:"page" query:"page" vd:"$>=0"`
	Limit            int32  `thrift:"limit,2" json:"limit" query:"limit" vd:"$>=0"`
	Order            int32  `thrift:"order,3" json:"order" query:"order" vd:"$>=0"`
	OrderBy          string `thrift:"order_by,4" json:"order_by" query:"order_by"`
	RepositoryType   int32  `thrift:"repository_type,5" json:"repository_type" query:"repository_type"`
	StoreType        int32  `thrift:"store_type,6" json:"store_type" query:"store_type"`
	RepositoryDomain string `thrift:"repository_domain,7" json:"repository_domain" query:"repository_domain"`
	RepositoryOwner  string `thrift:"repository_owner,8" json:"repository_owner" query:"repository_owner"`
	RepositoryName   string `thrift:"repository_name,9" json:"repository_name" query:"repository_name"`
}

func NewGetRepositoriesReq() *GetRepositoriesReq {
	return &GetRepositoriesReq{}
}

func (p *GetRepositoriesReq) GetPage() (v int32) {
	return p.Page
}

func (p *GetRepositoriesReq) GetLimit() (v int32) {
	return p.Limit
}

func (p *GetRepositoriesReq) GetOrder() (v int32) {
	return p.Order
}

func (p *GetRepositoriesReq) GetOrderBy() (v string) {
	return p.OrderBy
}

func (p *GetRepositoriesReq) GetRepositoryType() (v int32) {
	return p.RepositoryType
}

func (p *GetRepositoriesReq) GetStoreType() (v int32) {
	return p.StoreType
}

func (p *GetRepositoriesReq) GetRepositoryDomain() (v string) {
	return p.RepositoryDomain
}

func (p *GetRepositoriesReq) GetRepositoryOwner() (v string) {
	return p.RepositoryOwner
}

func (p *GetRepositoriesReq) GetRepositoryName() (v string) {
	return p.RepositoryName
}

func (p *GetRepositoriesReq) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetRepositoriesReq(%+v)", *p)
}

type GetRepositoriesRes struct {
	Code int32                   `thrift:"code,1" form:"code" json:"code" query:"code"`
	Msg  string                  `thrift:"msg,2" form:"msg" json:"msg" query:"msg"`
	Data *GetRepositoriesResData `thrift:"data,3" form:"data" json:"data" query:"data"`
}

func NewGetRepositoriesRes() *GetRepositoriesRes {
	return &GetRepositoriesRes{}
}

func (p *GetRepositoriesRes) GetCode() (v int32) {
	return p.Code
}

func (p *GetRepositoriesRes) GetMsg() (v string) {
	return p.Msg
}

var GetRepositoriesRes_Data_DEFAULT *GetRepositoriesResData

func (p *GetRepositoriesRes) GetData() (v *GetRepositoriesResData) {
	if !p.IsSetData() {
		return GetRepositoriesRes_Data_DEFAULT
	}
	return p.Data
}

func (p *GetRepositoriesRes) IsSetData() bool {
	return p.Data != nil
}

func (p *GetRepositoriesRes) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetRepositoriesRes(%+v)", *p)
}

type GetRepositoriesResData struct {
	Repositories []*model.Repository `thrift:"repositories,1" form:"repositories" json:"repositories" query:"repositories"`
	Total        int32               `thrift:"total,2" form:"total" json:"total" query:"total"`
}

func NewGetRepositoriesResData() *GetRepositoriesResData {
	return &GetRepositoriesResData{}
}

func (p *GetRepositoriesResData) GetRepositories() (v []*model.Repository) {
	return p.Repositories
}

func (p *GetRepositoriesResData) GetTotal() (v int32) {
	return p.Total
}

func (p *GetRepositoriesResData) String() string {
	if p == nil {
		return "<nil>"
	}
	return fmt.Sprintf("GetRepositoriesResData(%+v)", *p)
}

type RepositoryService interface {
	AddRepository(ctx context.Context, req *AddRepositoryReq) (r *AddRepositoryRes, err error)

	DeleteRepository(ctx context.Context, req *DeleteRepositoriesReq) (r *DeleteRepositoriesRes, err error)

	UpdateRepository(ctx context.Context, req *UpdateRepositoryReq) (r *UpdateRepositoryRes, err error)

	GetRepositories(ctx context.Context, req *GetRepositoriesReq) (r *GetRepositoriesRes, err error)
}
