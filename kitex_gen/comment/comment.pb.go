// Code generated by protoc-gen-go. DO NOT EDIT.
// versions:
// 	protoc-gen-go v1.28.1
// 	protoc        v3.6.1
// source: idl/comment.proto

package comment

import (
	context "context"
	protoreflect "google.golang.org/protobuf/reflect/protoreflect"
	protoimpl "google.golang.org/protobuf/runtime/protoimpl"
	reflect "reflect"
	sync "sync"
)

const (
	// Verify that this generated code is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(20 - protoimpl.MinVersion)
	// Verify that runtime/protoimpl is sufficiently up-to-date.
	_ = protoimpl.EnforceVersion(protoimpl.MaxVersion - 20)
)

type BaseResp struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	StatusCode    int32  `protobuf:"varint,1,opt,name=status_code,json=statusCode,proto3" json:"status_code,omitempty"`         //状态码
	StatusMessage string `protobuf:"bytes,2,opt,name=status_message,json=statusMessage,proto3" json:"status_message,omitempty"` //状态描述
	ServiceTime   int64  `protobuf:"varint,3,opt,name=service_time,json=serviceTime,proto3" json:"service_time,omitempty"`      //服务时间
}

func (x *BaseResp) Reset() {
	*x = BaseResp{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[0]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *BaseResp) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*BaseResp) ProtoMessage() {}

func (x *BaseResp) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[0]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use BaseResp.ProtoReflect.Descriptor instead.
func (*BaseResp) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{0}
}

func (x *BaseResp) GetStatusCode() int32 {
	if x != nil {
		return x.StatusCode
	}
	return 0
}

func (x *BaseResp) GetStatusMessage() string {
	if x != nil {
		return x.StatusMessage
	}
	return ""
}

func (x *BaseResp) GetServiceTime() int64 {
	if x != nil {
		return x.ServiceTime
	}
	return 0
}

type User struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id              int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                                 //用户id
	Name            string `protobuf:"bytes,2,opt,name=name,proto3" json:"name,omitempty"`                                              //用户名称
	FollowCount     int64  `protobuf:"varint,3,opt,name=follow_count,json=followCount,proto3" json:"follow_count,omitempty"`            //关注总数
	FollowerCount   int64  `protobuf:"varint,4,opt,name=follower_count,json=followerCount,proto3" json:"follower_count,omitempty"`      //粉丝总数
	IsFollow        bool   `protobuf:"varint,5,opt,name=is_follow,json=isFollow,proto3" json:"is_follow,omitempty"`                     //true-已关注,false-未关注
	Avatar          string `protobuf:"bytes,6,opt,name=avatar,proto3" json:"avatar,omitempty"`                                          // 用户头像
	BackgroundImage string `protobuf:"bytes,7,opt,name=background_image,json=backgroundImage,proto3" json:"background_image,omitempty"` // 用户个人页顶部大图
	Signature       string `protobuf:"bytes,8,opt,name=signature,proto3" json:"signature,omitempty"`                                    // 个人简介
	TotalFavorited  int64  `protobuf:"varint,9,opt,name=total_favorited,json=totalFavorited,proto3" json:"total_favorited,omitempty"`   // 获赞数量
	WorkCount       int64  `protobuf:"varint,10,opt,name=work_count,json=workCount,proto3" json:"work_count,omitempty"`                 // 作品数量
	FavoriteCount   int64  `protobuf:"varint,11,opt,name=favorite_count,json=favoriteCount,proto3" json:"favorite_count,omitempty"`     // 点赞数量
}

func (x *User) Reset() {
	*x = User{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[1]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *User) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*User) ProtoMessage() {}

func (x *User) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[1]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use User.ProtoReflect.Descriptor instead.
func (*User) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{1}
}

func (x *User) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *User) GetName() string {
	if x != nil {
		return x.Name
	}
	return ""
}

func (x *User) GetFollowCount() int64 {
	if x != nil {
		return x.FollowCount
	}
	return 0
}

func (x *User) GetFollowerCount() int64 {
	if x != nil {
		return x.FollowerCount
	}
	return 0
}

func (x *User) GetIsFollow() bool {
	if x != nil {
		return x.IsFollow
	}
	return false
}

func (x *User) GetAvatar() string {
	if x != nil {
		return x.Avatar
	}
	return ""
}

func (x *User) GetBackgroundImage() string {
	if x != nil {
		return x.BackgroundImage
	}
	return ""
}

func (x *User) GetSignature() string {
	if x != nil {
		return x.Signature
	}
	return ""
}

func (x *User) GetTotalFavorited() int64 {
	if x != nil {
		return x.TotalFavorited
	}
	return 0
}

func (x *User) GetWorkCount() int64 {
	if x != nil {
		return x.WorkCount
	}
	return 0
}

func (x *User) GetFavoriteCount() int64 {
	if x != nil {
		return x.FavoriteCount
	}
	return 0
}

type Comment struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Id         int64  `protobuf:"varint,1,opt,name=id,proto3" json:"id,omitempty"`                                  //视频评论id
	User       *User  `protobuf:"bytes,2,opt,name=user,proto3" json:"user,omitempty"`                               //评论用户信息
	Content    string `protobuf:"bytes,3,opt,name=content,proto3" json:"content,omitempty"`                         //评论内容
	CreateDate string `protobuf:"bytes,4,opt,name=create_date,json=createDate,proto3" json:"create_date,omitempty"` //评论发布日期，格式mm-dd
}

func (x *Comment) Reset() {
	*x = Comment{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[2]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *Comment) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*Comment) ProtoMessage() {}

func (x *Comment) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[2]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use Comment.ProtoReflect.Descriptor instead.
func (*Comment) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{2}
}

func (x *Comment) GetId() int64 {
	if x != nil {
		return x.Id
	}
	return 0
}

func (x *Comment) GetUser() *User {
	if x != nil {
		return x.User
	}
	return nil
}

func (x *Comment) GetContent() string {
	if x != nil {
		return x.Content
	}
	return ""
}

func (x *Comment) GetCreateDate() string {
	if x != nil {
		return x.CreateDate
	}
	return ""
}

type CreateCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token       string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                                //用户鉴权token
	VideoId     int64  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`            //视频id
	CommentText string `protobuf:"bytes,3,opt,name=comment_text,json=commentText,proto3" json:"comment_text,omitempty"` ////用户填写的评论内容
}

func (x *CreateCommentRequest) Reset() {
	*x = CreateCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[3]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentRequest) ProtoMessage() {}

func (x *CreateCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[3]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentRequest.ProtoReflect.Descriptor instead.
func (*CreateCommentRequest) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{3}
}

func (x *CreateCommentRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CreateCommentRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *CreateCommentRequest) GetCommentText() string {
	if x != nil {
		return x.CommentText
	}
	return ""
}

type CreateCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	Comment  *Comment  `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"` //评论成功返回评论内容，不需要重新拉取整个列表
}

func (x *CreateCommentResponse) Reset() {
	*x = CreateCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[4]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CreateCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CreateCommentResponse) ProtoMessage() {}

func (x *CreateCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[4]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CreateCommentResponse.ProtoReflect.Descriptor instead.
func (*CreateCommentResponse) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{4}
}

func (x *CreateCommentResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *CreateCommentResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type DeleteCommentRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token     string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                           //用户鉴权token
	VideoId   int64  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"`       //视频id
	CommentId int64  `protobuf:"varint,3,opt,name=comment_id,json=commentId,proto3" json:"comment_id,omitempty"` //要删除的评论id
}

func (x *DeleteCommentRequest) Reset() {
	*x = DeleteCommentRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[5]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentRequest) ProtoMessage() {}

func (x *DeleteCommentRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[5]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentRequest.ProtoReflect.Descriptor instead.
func (*DeleteCommentRequest) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{5}
}

func (x *DeleteCommentRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *DeleteCommentRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

func (x *DeleteCommentRequest) GetCommentId() int64 {
	if x != nil {
		return x.CommentId
	}
	return 0
}

type DeleteCommentResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp *BaseResp `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	Comment  *Comment  `protobuf:"bytes,2,opt,name=comment,proto3" json:"comment,omitempty"` //评论成功返回评论内容，不需要重新拉取整个列表
}

func (x *DeleteCommentResponse) Reset() {
	*x = DeleteCommentResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[6]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *DeleteCommentResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*DeleteCommentResponse) ProtoMessage() {}

func (x *DeleteCommentResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[6]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use DeleteCommentResponse.ProtoReflect.Descriptor instead.
func (*DeleteCommentResponse) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{6}
}

func (x *DeleteCommentResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *DeleteCommentResponse) GetComment() *Comment {
	if x != nil {
		return x.Comment
	}
	return nil
}

type CommentListRequest struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	Token   string `protobuf:"bytes,1,opt,name=token,proto3" json:"token,omitempty"`                     //用户鉴权token
	VideoId int64  `protobuf:"varint,2,opt,name=video_id,json=videoId,proto3" json:"video_id,omitempty"` //视频id
}

func (x *CommentListRequest) Reset() {
	*x = CommentListRequest{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[7]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListRequest) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListRequest) ProtoMessage() {}

func (x *CommentListRequest) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[7]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListRequest.ProtoReflect.Descriptor instead.
func (*CommentListRequest) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{7}
}

func (x *CommentListRequest) GetToken() string {
	if x != nil {
		return x.Token
	}
	return ""
}

func (x *CommentListRequest) GetVideoId() int64 {
	if x != nil {
		return x.VideoId
	}
	return 0
}

type CommentListResponse struct {
	state         protoimpl.MessageState
	sizeCache     protoimpl.SizeCache
	unknownFields protoimpl.UnknownFields

	BaseResp    *BaseResp  `protobuf:"bytes,1,opt,name=base_resp,json=baseResp,proto3" json:"base_resp,omitempty"`
	CommentList []*Comment `protobuf:"bytes,2,rep,name=comment_list,json=commentList,proto3" json:"comment_list,omitempty"` //评论列表
}

func (x *CommentListResponse) Reset() {
	*x = CommentListResponse{}
	if protoimpl.UnsafeEnabled {
		mi := &file_idl_comment_proto_msgTypes[8]
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		ms.StoreMessageInfo(mi)
	}
}

func (x *CommentListResponse) String() string {
	return protoimpl.X.MessageStringOf(x)
}

func (*CommentListResponse) ProtoMessage() {}

func (x *CommentListResponse) ProtoReflect() protoreflect.Message {
	mi := &file_idl_comment_proto_msgTypes[8]
	if protoimpl.UnsafeEnabled && x != nil {
		ms := protoimpl.X.MessageStateOf(protoimpl.Pointer(x))
		if ms.LoadMessageInfo() == nil {
			ms.StoreMessageInfo(mi)
		}
		return ms
	}
	return mi.MessageOf(x)
}

// Deprecated: Use CommentListResponse.ProtoReflect.Descriptor instead.
func (*CommentListResponse) Descriptor() ([]byte, []int) {
	return file_idl_comment_proto_rawDescGZIP(), []int{8}
}

func (x *CommentListResponse) GetBaseResp() *BaseResp {
	if x != nil {
		return x.BaseResp
	}
	return nil
}

func (x *CommentListResponse) GetCommentList() []*Comment {
	if x != nil {
		return x.CommentList
	}
	return nil
}

var File_idl_comment_proto protoreflect.FileDescriptor

var file_idl_comment_proto_rawDesc = []byte{
	0x0a, 0x11, 0x69, 0x64, 0x6c, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x12, 0x07, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x75, 0x0a, 0x08,
	0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x1f, 0x0a, 0x0b, 0x73, 0x74, 0x61, 0x74,
	0x75, 0x73, 0x5f, 0x63, 0x6f, 0x64, 0x65, 0x18, 0x01, 0x20, 0x01, 0x28, 0x05, 0x52, 0x0a, 0x73,
	0x74, 0x61, 0x74, 0x75, 0x73, 0x43, 0x6f, 0x64, 0x65, 0x12, 0x25, 0x0a, 0x0e, 0x73, 0x74, 0x61,
	0x74, 0x75, 0x73, 0x5f, 0x6d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0d, 0x73, 0x74, 0x61, 0x74, 0x75, 0x73, 0x4d, 0x65, 0x73, 0x73, 0x61, 0x67, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x5f, 0x74, 0x69, 0x6d, 0x65,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x73, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65, 0x54,
	0x69, 0x6d, 0x65, 0x22, 0xe1, 0x02, 0x0a, 0x04, 0x55, 0x73, 0x65, 0x72, 0x12, 0x0e, 0x0a, 0x02,
	0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02, 0x69, 0x64, 0x12, 0x12, 0x0a, 0x04,
	0x6e, 0x61, 0x6d, 0x65, 0x18, 0x02, 0x20, 0x01, 0x28, 0x09, 0x52, 0x04, 0x6e, 0x61, 0x6d, 0x65,
	0x12, 0x21, 0x0a, 0x0c, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0b, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x43, 0x6f,
	0x75, 0x6e, 0x74, 0x12, 0x25, 0x0a, 0x0e, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x65, 0x72, 0x5f,
	0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18, 0x04, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x6f, 0x6c,
	0x6c, 0x6f, 0x77, 0x65, 0x72, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x12, 0x1b, 0x0a, 0x09, 0x69, 0x73,
	0x5f, 0x66, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x18, 0x05, 0x20, 0x01, 0x28, 0x08, 0x52, 0x08, 0x69,
	0x73, 0x46, 0x6f, 0x6c, 0x6c, 0x6f, 0x77, 0x12, 0x16, 0x0a, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61,
	0x72, 0x18, 0x06, 0x20, 0x01, 0x28, 0x09, 0x52, 0x06, 0x61, 0x76, 0x61, 0x74, 0x61, 0x72, 0x12,
	0x29, 0x0a, 0x10, 0x62, 0x61, 0x63, 0x6b, 0x67, 0x72, 0x6f, 0x75, 0x6e, 0x64, 0x5f, 0x69, 0x6d,
	0x61, 0x67, 0x65, 0x18, 0x07, 0x20, 0x01, 0x28, 0x09, 0x52, 0x0f, 0x62, 0x61, 0x63, 0x6b, 0x67,
	0x72, 0x6f, 0x75, 0x6e, 0x64, 0x49, 0x6d, 0x61, 0x67, 0x65, 0x12, 0x1c, 0x0a, 0x09, 0x73, 0x69,
	0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x18, 0x08, 0x20, 0x01, 0x28, 0x09, 0x52, 0x09, 0x73,
	0x69, 0x67, 0x6e, 0x61, 0x74, 0x75, 0x72, 0x65, 0x12, 0x27, 0x0a, 0x0f, 0x74, 0x6f, 0x74, 0x61,
	0x6c, 0x5f, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x64, 0x18, 0x09, 0x20, 0x01, 0x28,
	0x03, 0x52, 0x0e, 0x74, 0x6f, 0x74, 0x61, 0x6c, 0x46, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65,
	0x64, 0x12, 0x1d, 0x0a, 0x0a, 0x77, 0x6f, 0x72, 0x6b, 0x5f, 0x63, 0x6f, 0x75, 0x6e, 0x74, 0x18,
	0x0a, 0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x77, 0x6f, 0x72, 0x6b, 0x43, 0x6f, 0x75, 0x6e, 0x74,
	0x12, 0x25, 0x0a, 0x0e, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x75,
	0x6e, 0x74, 0x18, 0x0b, 0x20, 0x01, 0x28, 0x03, 0x52, 0x0d, 0x66, 0x61, 0x76, 0x6f, 0x72, 0x69,
	0x74, 0x65, 0x43, 0x6f, 0x75, 0x6e, 0x74, 0x22, 0x77, 0x0a, 0x07, 0x43, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x12, 0x0e, 0x0a, 0x02, 0x69, 0x64, 0x18, 0x01, 0x20, 0x01, 0x28, 0x03, 0x52, 0x02,
	0x69, 0x64, 0x12, 0x21, 0x0a, 0x04, 0x75, 0x73, 0x65, 0x72, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b,
	0x32, 0x0d, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x55, 0x73, 0x65, 0x72, 0x52,
	0x04, 0x75, 0x73, 0x65, 0x72, 0x12, 0x18, 0x0a, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74,
	0x18, 0x03, 0x20, 0x01, 0x28, 0x09, 0x52, 0x07, 0x63, 0x6f, 0x6e, 0x74, 0x65, 0x6e, 0x74, 0x12,
	0x1f, 0x0a, 0x0b, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x64, 0x61, 0x74, 0x65, 0x18, 0x04,
	0x20, 0x01, 0x28, 0x09, 0x52, 0x0a, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x44, 0x61, 0x74, 0x65,
	0x22, 0x6c, 0x0a, 0x16, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14, 0x0a, 0x05, 0x74, 0x6f,
	0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e,
	0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64, 0x18, 0x02, 0x20, 0x01,
	0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x12, 0x21, 0x0a, 0x0c, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x74, 0x65, 0x78, 0x74, 0x18, 0x03, 0x20, 0x01, 0x28,
	0x09, 0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x54, 0x65, 0x78, 0x74, 0x22, 0x75,
	0x0a, 0x17, 0x63, 0x72, 0x65, 0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x62, 0x61, 0x73,
	0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52,
	0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x68, 0x0a, 0x16, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12,
	0x14, 0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05,
	0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69,
	0x64, 0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64,
	0x12, 0x1d, 0x0a, 0x0a, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x69, 0x64, 0x18, 0x03,
	0x20, 0x01, 0x28, 0x03, 0x52, 0x09, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x49, 0x64, 0x22,
	0x75, 0x0a, 0x17, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x62, 0x61,
	0x73, 0x65, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e,
	0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70,
	0x52, 0x08, 0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x2a, 0x0a, 0x07, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x18, 0x02, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x10, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x52, 0x07, 0x63,
	0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x22, 0x47, 0x0a, 0x14, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x12, 0x14,
	0x0a, 0x05, 0x74, 0x6f, 0x6b, 0x65, 0x6e, 0x18, 0x01, 0x20, 0x01, 0x28, 0x09, 0x52, 0x05, 0x74,
	0x6f, 0x6b, 0x65, 0x6e, 0x12, 0x19, 0x0a, 0x08, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x5f, 0x69, 0x64,
	0x18, 0x02, 0x20, 0x01, 0x28, 0x03, 0x52, 0x07, 0x76, 0x69, 0x64, 0x65, 0x6f, 0x49, 0x64, 0x22,
	0x7c, 0x0a, 0x15, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x5f,
	0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x12, 0x2e, 0x0a, 0x09, 0x62, 0x61, 0x73, 0x65,
	0x5f, 0x72, 0x65, 0x73, 0x70, 0x18, 0x01, 0x20, 0x01, 0x28, 0x0b, 0x32, 0x11, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x42, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x52, 0x08,
	0x62, 0x61, 0x73, 0x65, 0x52, 0x65, 0x73, 0x70, 0x12, 0x33, 0x0a, 0x0c, 0x63, 0x6f, 0x6d, 0x6d,
	0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73, 0x74, 0x18, 0x02, 0x20, 0x03, 0x28, 0x0b, 0x32, 0x10,
	0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x52, 0x0b, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x32, 0x8c, 0x02,
	0x0a, 0x0e, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x53, 0x65, 0x72, 0x76, 0x69, 0x63, 0x65,
	0x12, 0x54, 0x0a, 0x0d, 0x43, 0x72, 0x65, 0x61, 0x74, 0x65, 0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x65, 0x61,
	0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65,
	0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x72, 0x65,
	0x61, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x72, 0x65, 0x73, 0x70,
	0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x54, 0x0a, 0x0d, 0x44, 0x65, 0x6c, 0x65, 0x74, 0x65,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x12, 0x1f, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74,
	0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x20, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65,
	0x6e, 0x74, 0x2e, 0x64, 0x65, 0x6c, 0x65, 0x74, 0x65, 0x5f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x12, 0x4e, 0x0a, 0x0b,
	0x43, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x4c, 0x69, 0x73, 0x74, 0x12, 0x1d, 0x2e, 0x63, 0x6f,
	0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69,
	0x73, 0x74, 0x5f, 0x72, 0x65, 0x71, 0x75, 0x65, 0x73, 0x74, 0x1a, 0x1e, 0x2e, 0x63, 0x6f, 0x6d,
	0x6d, 0x65, 0x6e, 0x74, 0x2e, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x5f, 0x6c, 0x69, 0x73,
	0x74, 0x5f, 0x72, 0x65, 0x73, 0x70, 0x6f, 0x6e, 0x73, 0x65, 0x22, 0x00, 0x42, 0x2e, 0x5a, 0x2c,
	0x67, 0x69, 0x74, 0x68, 0x75, 0x62, 0x2e, 0x63, 0x6f, 0x6d, 0x2f, 0x65, 0x70, 0x68, 0x6d, 0x65,
	0x72, 0x61, 0x6c, 0x2f, 0x64, 0x6f, 0x75, 0x79, 0x69, 0x6e, 0x2f, 0x6b, 0x69, 0x74, 0x65, 0x78,
	0x5f, 0x67, 0x65, 0x6e, 0x2f, 0x63, 0x6f, 0x6d, 0x6d, 0x65, 0x6e, 0x74, 0x62, 0x06, 0x70, 0x72,
	0x6f, 0x74, 0x6f, 0x33,
}

var (
	file_idl_comment_proto_rawDescOnce sync.Once
	file_idl_comment_proto_rawDescData = file_idl_comment_proto_rawDesc
)

func file_idl_comment_proto_rawDescGZIP() []byte {
	file_idl_comment_proto_rawDescOnce.Do(func() {
		file_idl_comment_proto_rawDescData = protoimpl.X.CompressGZIP(file_idl_comment_proto_rawDescData)
	})
	return file_idl_comment_proto_rawDescData
}

var file_idl_comment_proto_msgTypes = make([]protoimpl.MessageInfo, 9)
var file_idl_comment_proto_goTypes = []interface{}{
	(*BaseResp)(nil),              // 0: comment.BaseResp
	(*User)(nil),                  // 1: comment.User
	(*Comment)(nil),               // 2: comment.Comment
	(*CreateCommentRequest)(nil),  // 3: comment.create_comment_request
	(*CreateCommentResponse)(nil), // 4: comment.create_comment_response
	(*DeleteCommentRequest)(nil),  // 5: comment.delete_comment_request
	(*DeleteCommentResponse)(nil), // 6: comment.delete_comment_response
	(*CommentListRequest)(nil),    // 7: comment.comment_list_request
	(*CommentListResponse)(nil),   // 8: comment.comment_list_response
}
var file_idl_comment_proto_depIdxs = []int32{
	1,  // 0: comment.Comment.user:type_name -> comment.User
	0,  // 1: comment.create_comment_response.base_resp:type_name -> comment.BaseResp
	2,  // 2: comment.create_comment_response.comment:type_name -> comment.Comment
	0,  // 3: comment.delete_comment_response.base_resp:type_name -> comment.BaseResp
	2,  // 4: comment.delete_comment_response.comment:type_name -> comment.Comment
	0,  // 5: comment.comment_list_response.base_resp:type_name -> comment.BaseResp
	2,  // 6: comment.comment_list_response.comment_list:type_name -> comment.Comment
	3,  // 7: comment.CommentService.CreateComment:input_type -> comment.create_comment_request
	5,  // 8: comment.CommentService.DeleteComment:input_type -> comment.delete_comment_request
	7,  // 9: comment.CommentService.CommentList:input_type -> comment.comment_list_request
	4,  // 10: comment.CommentService.CreateComment:output_type -> comment.create_comment_response
	6,  // 11: comment.CommentService.DeleteComment:output_type -> comment.delete_comment_response
	8,  // 12: comment.CommentService.CommentList:output_type -> comment.comment_list_response
	10, // [10:13] is the sub-list for method output_type
	7,  // [7:10] is the sub-list for method input_type
	7,  // [7:7] is the sub-list for extension type_name
	7,  // [7:7] is the sub-list for extension extendee
	0,  // [0:7] is the sub-list for field type_name
}

func init() { file_idl_comment_proto_init() }
func file_idl_comment_proto_init() {
	if File_idl_comment_proto != nil {
		return
	}
	if !protoimpl.UnsafeEnabled {
		file_idl_comment_proto_msgTypes[0].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*BaseResp); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[1].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*User); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[2].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*Comment); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[3].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[4].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CreateCommentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[5].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[6].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*DeleteCommentResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[7].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListRequest); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
		file_idl_comment_proto_msgTypes[8].Exporter = func(v interface{}, i int) interface{} {
			switch v := v.(*CommentListResponse); i {
			case 0:
				return &v.state
			case 1:
				return &v.sizeCache
			case 2:
				return &v.unknownFields
			default:
				return nil
			}
		}
	}
	type x struct{}
	out := protoimpl.TypeBuilder{
		File: protoimpl.DescBuilder{
			GoPackagePath: reflect.TypeOf(x{}).PkgPath(),
			RawDescriptor: file_idl_comment_proto_rawDesc,
			NumEnums:      0,
			NumMessages:   9,
			NumExtensions: 0,
			NumServices:   1,
		},
		GoTypes:           file_idl_comment_proto_goTypes,
		DependencyIndexes: file_idl_comment_proto_depIdxs,
		MessageInfos:      file_idl_comment_proto_msgTypes,
	}.Build()
	File_idl_comment_proto = out.File
	file_idl_comment_proto_rawDesc = nil
	file_idl_comment_proto_goTypes = nil
	file_idl_comment_proto_depIdxs = nil
}

var _ context.Context

// Code generated by Kitex v0.4.4. DO NOT EDIT.

type CommentService interface {
	CreateComment(ctx context.Context, req *CreateCommentRequest) (res *CreateCommentResponse, err error)
	DeleteComment(ctx context.Context, req *DeleteCommentRequest) (res *DeleteCommentResponse, err error)
	CommentList(ctx context.Context, req *CommentListRequest) (res *CommentListResponse, err error)
}
