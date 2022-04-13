package service

import (
	"context"
	v1 "kratos-realworld/api/realworld/v1"
)


func (s *RealWorldService) CreateArticles(ctx context.Context,req *v1.CreateArticlesRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}

func (s *RealWorldService) UpdateArticles(ctx context.Context, req *v1.UpdateArticlesRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}


func (s *RealWorldService) DeleteArticles(ctx context.Context, req *v1.DeleteArticlesRequest) (reply *v1.DeleteArticlesReply, err error) {
	return &v1.DeleteArticlesReply{}, nil
}

func (s *RealWorldService) AddComments(ctx context.Context,req *v1.AddCommentsRequest) (reply *v1.SingleCommentReply, err error) {
	return &v1.SingleCommentReply{}, nil
}



func (s *RealWorldService) GetComments(ctx context.Context, req *v1.GetCommentsRequest) (reply *v1.MultipleCommentsReply, err error) {
	return &v1.MultipleCommentsReply{}, nil
}


func (s *RealWorldService) DeleteComments(ctx context.Context, req *v1.DeleteCommentsRequest) (reply *v1.DeleteCommentsReply, err error) {
	return &v1.DeleteCommentsReply{}, nil
}


func (s *RealWorldService) FavouriteArticle(ctx context.Context, req *v1.FavouriteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}



func (s *RealWorldService) UnFavouriteArticle(ctx context.Context, req *v1.UnFavouriteArticleRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}



func (s *RealWorldService) ListArticles(ctx context.Context, req *v1.ListArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}


func (s *RealWorldService) FeedArticles(ctx context.Context, req *v1.FeedArticlesRequest) (reply *v1.MultipleArticlesReply, err error) {
	return &v1.MultipleArticlesReply{}, nil
}


func (s *RealWorldService) GetArticles(ctx context.Context, req *v1.GetArticlesRequest) (reply *v1.SingleArticleReply, err error) {
	return &v1.SingleArticleReply{}, nil
}




func (s *RealWorldService) GetTags(ctx context.Context, req *v1.GetTagsRequest) (reply *v1.TagListReply, err error) {
	return &v1.TagListReply{}, nil
}