// Code generated by protoc-gen-entgrpc. DO NOT EDIT.
package entpb

import (
	context "context"
	base64 "encoding/base64"
	entproto "entgo.io/contrib/entproto"
	sqlgraph "entgo.io/ent/dialect/sql/sqlgraph"
	fmt "fmt"
	ent "github.com/rotemtam/ent-grpc-example/ent"
	category "github.com/rotemtam/ent-grpc-example/ent/category"
	user "github.com/rotemtam/ent-grpc-example/ent/user"
	codes "google.golang.org/grpc/codes"
	status "google.golang.org/grpc/status"
	emptypb "google.golang.org/protobuf/types/known/emptypb"
	strconv "strconv"
)

// CategoryService implements CategoryServiceServer
type CategoryService struct {
	client *ent.Client
	UnimplementedCategoryServiceServer
}

// NewCategoryService returns a new CategoryService
func NewCategoryService(client *ent.Client) *CategoryService {
	return &CategoryService{
		client: client,
	}
}

// toProtoCategory transforms the ent type to the pb type
func toProtoCategory(e *ent.Category) (*Category, error) {
	v := &Category{}
	id := int64(e.ID)
	v.Id = id
	name := e.Name
	v.Name = name
	if edg := e.Edges.Admin; edg != nil {
		id := int64(edg.ID)
		v.Admin = &User{
			Id: id,
		}
	}
	return v, nil
}

// toProtoCategoryList transforms a list of ent type to a list of pb type
func toProtoCategoryList(e []*ent.Category) ([]*Category, error) {
	var pbList []*Category
	for _, entEntity := range e {
		pbEntity, err := toProtoCategory(entEntity)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		pbList = append(pbList, pbEntity)
	}
	return pbList, nil
}

// Create implements CategoryServiceServer.Create
func (svc *CategoryService) Create(ctx context.Context, req *CreateCategoryRequest) (*Category, error) {
	category := req.GetCategory()
	m, err := svc.createBuilder(category)
	if err != nil {
		return nil, err
	}
	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoCategory(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Get implements CategoryServiceServer.Get
func (svc *CategoryService) Get(ctx context.Context, req *GetCategoryRequest) (*Category, error) {
	var (
		err error
		get *ent.Category
	)
	id := int(req.GetId())
	switch req.GetView() {
	case GetCategoryRequest_VIEW_UNSPECIFIED, GetCategoryRequest_BASIC:
		get, err = svc.client.Category.Get(ctx, id)
	case GetCategoryRequest_WITH_EDGE_IDS:
		get, err = svc.client.Category.Query().
			Where(category.ID(id)).
			WithAdmin(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			Only(ctx)
	default:
		return nil, status.Error(codes.InvalidArgument, "invalid argument: unknown view")
	}
	switch {
	case err == nil:
		return toProtoCategory(get)
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Update implements CategoryServiceServer.Update
func (svc *CategoryService) Update(ctx context.Context, req *UpdateCategoryRequest) (*Category, error) {
	category := req.GetCategory()
	categoryID := int(category.GetId())
	m := svc.client.Category.UpdateOneID(categoryID)
	categoryName := category.GetName()
	m.SetName(categoryName)
	if category.GetAdmin() != nil {
		categoryAdmin := int(category.GetAdmin().GetId())
		m.SetAdminID(categoryAdmin)
	}

	res, err := m.Save(ctx)
	switch {
	case err == nil:
		proto, err := toProtoCategory(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return proto, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// Delete implements CategoryServiceServer.Delete
func (svc *CategoryService) Delete(ctx context.Context, req *DeleteCategoryRequest) (*emptypb.Empty, error) {
	var err error
	id := int(req.GetId())
	err = svc.client.Category.DeleteOneID(id).Exec(ctx)
	switch {
	case err == nil:
		return &emptypb.Empty{}, nil
	case ent.IsNotFound(err):
		return nil, status.Errorf(codes.NotFound, "not found: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// List implements CategoryServiceServer.List
func (svc *CategoryService) List(ctx context.Context, req *ListCategoryRequest) (*ListCategoryResponse, error) {
	var (
		err      error
		entList  []*ent.Category
		pageSize int
	)
	pageSize = int(req.GetPageSize())
	switch {
	case pageSize < 0:
		return nil, status.Errorf(codes.InvalidArgument, "page size cannot be less than zero")
	case pageSize == 0 || pageSize > entproto.MaxPageSize:
		pageSize = entproto.MaxPageSize
	}
	listQuery := svc.client.Category.Query().
		Order(ent.Desc(category.FieldID)).
		Limit(pageSize + 1)
	if req.GetPageToken() != "" {
		bytes, err := base64.StdEncoding.DecodeString(req.PageToken)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		token, err := strconv.ParseInt(string(bytes), 10, 32)
		if err != nil {
			return nil, status.Errorf(codes.InvalidArgument, "page token is invalid")
		}
		pageToken := int(token)
		listQuery = listQuery.
			Where(category.IDLTE(pageToken))
	}
	switch req.GetView() {
	case ListCategoryRequest_VIEW_UNSPECIFIED, ListCategoryRequest_BASIC:
		entList, err = listQuery.All(ctx)
	case ListCategoryRequest_WITH_EDGE_IDS:
		entList, err = listQuery.
			WithAdmin(func(query *ent.UserQuery) {
				query.Select(user.FieldID)
			}).
			All(ctx)
	}
	switch {
	case err == nil:
		var nextPageToken string
		if len(entList) == pageSize+1 {
			nextPageToken = base64.StdEncoding.EncodeToString(
				[]byte(fmt.Sprintf("%v", entList[len(entList)-1].ID)))
			entList = entList[:len(entList)-1]
		}
		protoList, err := toProtoCategoryList(entList)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &ListCategoryResponse{
			CategoryList:  protoList,
			NextPageToken: nextPageToken,
		}, nil
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

// BatchCreate implements CategoryServiceServer.BatchCreate
func (svc *CategoryService) BatchCreate(ctx context.Context, req *BatchCreateCategoriesRequest) (*BatchCreateCategoriesResponse, error) {
	requests := req.GetRequests()
	if len(requests) > entproto.MaxBatchCreateSize {
		return nil, status.Errorf(codes.InvalidArgument, "batch size cannot be greater than %d", entproto.MaxBatchCreateSize)
	}
	bulk := make([]*ent.CategoryCreate, len(requests))
	for i, req := range requests {
		category := req.GetCategory()
		var err error
		bulk[i], err = svc.createBuilder(category)
		if err != nil {
			return nil, err
		}
	}
	res, err := svc.client.Category.CreateBulk(bulk...).Save(ctx)
	switch {
	case err == nil:
		protoList, err := toProtoCategoryList(res)
		if err != nil {
			return nil, status.Errorf(codes.Internal, "internal error: %s", err)
		}
		return &BatchCreateCategoriesResponse{
			Categories: protoList,
		}, nil
	case sqlgraph.IsUniqueConstraintError(err):
		return nil, status.Errorf(codes.AlreadyExists, "already exists: %s", err)
	case ent.IsConstraintError(err):
		return nil, status.Errorf(codes.InvalidArgument, "invalid argument: %s", err)
	default:
		return nil, status.Errorf(codes.Internal, "internal error: %s", err)
	}

}

func (svc *CategoryService) createBuilder(category *Category) (*ent.CategoryCreate, error) {
	m := svc.client.Category.Create()
	categoryName := category.GetName()
	m.SetName(categoryName)
	if category.GetAdmin() != nil {
		categoryAdmin := int(category.GetAdmin().GetId())
		m.SetAdminID(categoryAdmin)
	}
	return m, nil
}
