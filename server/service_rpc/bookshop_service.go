package service_rpc

import (
	"context"
	"grpc_dev/server/pb/bookshop"
)

type BookshopServer struct {
	bookshop.UnimplementedInventoryServer
}

func (s *BookshopServer) GetBookList(c context.Context, in *bookshop.GetBookListRequest) (*bookshop.GetBookListResponse, error) {
	return &bookshop.GetBookListResponse{
		Books: getSampleBooks(),
	}, nil
}

func getSampleBooks() []*bookshop.Book {
	return []*bookshop.Book{
		{Title: "NAGN", Author: "Tran", PageCount: 30},
		{Title: "NAGN 1", Author: "Ngoc", PageCount: 32},
		{Title: "NAGN 2", Author: "Hoang", PageCount: 34},
	}
}
