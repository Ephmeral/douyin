// Code generated by Fastpb v0.0.2. DO NOT EDIT.

package comment

import (
	fmt "fmt"
	fastpb "github.com/cloudwego/fastpb"
)

var (
	_ = fmt.Errorf
	_ = fastpb.Skip
)

func (x *BaseResp) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_BaseResp[number], err)
}

func (x *BaseResp) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.StatusCode, offset, err = fastpb.ReadInt32(buf, _type)
	return offset, err
}

func (x *BaseResp) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.StatusMessage, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *BaseResp) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.ServiceTime, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 7:
		offset, err = x.fastReadField7(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 8:
		offset, err = x.fastReadField8(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 9:
		offset, err = x.fastReadField9(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 10:
		offset, err = x.fastReadField10(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 11:
		offset, err = x.fastReadField11(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_User[number], err)
}

func (x *User) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.Name, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.FollowCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.FollowerCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.IsFollow, offset, err = fastpb.ReadBool(buf, _type)
	return offset, err
}

func (x *User) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.Avatar, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField7(buf []byte, _type int8) (offset int, err error) {
	x.BackgroundImage, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField8(buf []byte, _type int8) (offset int, err error) {
	x.Signature, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *User) fastReadField9(buf []byte, _type int8) (offset int, err error) {
	x.TotalFavorited, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField10(buf []byte, _type int8) (offset int, err error) {
	x.WorkCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *User) fastReadField11(buf []byte, _type int8) (offset int, err error) {
	x.FavoriteCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Comment) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 4:
		offset, err = x.fastReadField4(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 5:
		offset, err = x.fastReadField5(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 6:
		offset, err = x.fastReadField6(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_Comment[number], err)
}

func (x *Comment) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Id, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v User
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.User = &v
	return offset, nil
}

func (x *Comment) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.Content, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField4(buf []byte, _type int8) (offset int, err error) {
	x.CreateDate, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField5(buf []byte, _type int8) (offset int, err error) {
	x.LikeCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *Comment) fastReadField6(buf []byte, _type int8) (offset int, err error) {
	x.TeaseCount, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateCommentRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateCommentRequest[number], err)
}

func (x *CreateCommentRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateCommentRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CreateCommentRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.CommentText, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CreateCommentResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CreateCommentResponse[number], err)
}

func (x *CreateCommentResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *CreateCommentResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Comment = &v
	return offset, nil
}

func (x *DeleteCommentRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 3:
		offset, err = x.fastReadField3(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteCommentRequest[number], err)
}

func (x *DeleteCommentRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *DeleteCommentRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteCommentRequest) fastReadField3(buf []byte, _type int8) (offset int, err error) {
	x.CommentId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *DeleteCommentResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_DeleteCommentResponse[number], err)
}

func (x *DeleteCommentResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *DeleteCommentResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.Comment = &v
	return offset, nil
}

func (x *CommentListRequest) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CommentListRequest[number], err)
}

func (x *CommentListRequest) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	x.Token, offset, err = fastpb.ReadString(buf, _type)
	return offset, err
}

func (x *CommentListRequest) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	x.VideoId, offset, err = fastpb.ReadInt64(buf, _type)
	return offset, err
}

func (x *CommentListResponse) FastRead(buf []byte, _type int8, number int32) (offset int, err error) {
	switch number {
	case 1:
		offset, err = x.fastReadField1(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	case 2:
		offset, err = x.fastReadField2(buf, _type)
		if err != nil {
			goto ReadFieldError
		}
	default:
		offset, err = fastpb.Skip(buf, _type, number)
		if err != nil {
			goto SkipFieldError
		}
	}
	return offset, nil
SkipFieldError:
	return offset, fmt.Errorf("%T cannot parse invalid wire-format data, error: %s", x, err)
ReadFieldError:
	return offset, fmt.Errorf("%T read field %d '%s' error: %s", x, number, fieldIDToName_CommentListResponse[number], err)
}

func (x *CommentListResponse) fastReadField1(buf []byte, _type int8) (offset int, err error) {
	var v BaseResp
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.BaseResp = &v
	return offset, nil
}

func (x *CommentListResponse) fastReadField2(buf []byte, _type int8) (offset int, err error) {
	var v Comment
	offset, err = fastpb.ReadMessage(buf, _type, &v)
	if err != nil {
		return offset, err
	}
	x.CommentList = append(x.CommentList, &v)
	return offset, nil
}

func (x *BaseResp) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *BaseResp) fastWriteField1(buf []byte) (offset int) {
	if x.StatusCode == 0 {
		return offset
	}
	offset += fastpb.WriteInt32(buf[offset:], 1, x.StatusCode)
	return offset
}

func (x *BaseResp) fastWriteField2(buf []byte) (offset int) {
	if x.StatusMessage == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.StatusMessage)
	return offset
}

func (x *BaseResp) fastWriteField3(buf []byte) (offset int) {
	if x.ServiceTime == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.ServiceTime)
	return offset
}

func (x *User) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	offset += x.fastWriteField7(buf[offset:])
	offset += x.fastWriteField8(buf[offset:])
	offset += x.fastWriteField9(buf[offset:])
	offset += x.fastWriteField10(buf[offset:])
	offset += x.fastWriteField11(buf[offset:])
	return offset
}

func (x *User) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *User) fastWriteField2(buf []byte) (offset int) {
	if x.Name == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 2, x.Name)
	return offset
}

func (x *User) fastWriteField3(buf []byte) (offset int) {
	if x.FollowCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.FollowCount)
	return offset
}

func (x *User) fastWriteField4(buf []byte) (offset int) {
	if x.FollowerCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 4, x.FollowerCount)
	return offset
}

func (x *User) fastWriteField5(buf []byte) (offset int) {
	if !x.IsFollow {
		return offset
	}
	offset += fastpb.WriteBool(buf[offset:], 5, x.IsFollow)
	return offset
}

func (x *User) fastWriteField6(buf []byte) (offset int) {
	if x.Avatar == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 6, x.Avatar)
	return offset
}

func (x *User) fastWriteField7(buf []byte) (offset int) {
	if x.BackgroundImage == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 7, x.BackgroundImage)
	return offset
}

func (x *User) fastWriteField8(buf []byte) (offset int) {
	if x.Signature == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 8, x.Signature)
	return offset
}

func (x *User) fastWriteField9(buf []byte) (offset int) {
	if x.TotalFavorited == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 9, x.TotalFavorited)
	return offset
}

func (x *User) fastWriteField10(buf []byte) (offset int) {
	if x.WorkCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 10, x.WorkCount)
	return offset
}

func (x *User) fastWriteField11(buf []byte) (offset int) {
	if x.FavoriteCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 11, x.FavoriteCount)
	return offset
}

func (x *Comment) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	offset += x.fastWriteField4(buf[offset:])
	offset += x.fastWriteField5(buf[offset:])
	offset += x.fastWriteField6(buf[offset:])
	return offset
}

func (x *Comment) fastWriteField1(buf []byte) (offset int) {
	if x.Id == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 1, x.Id)
	return offset
}

func (x *Comment) fastWriteField2(buf []byte) (offset int) {
	if x.User == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.User)
	return offset
}

func (x *Comment) fastWriteField3(buf []byte) (offset int) {
	if x.Content == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.Content)
	return offset
}

func (x *Comment) fastWriteField4(buf []byte) (offset int) {
	if x.CreateDate == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 4, x.CreateDate)
	return offset
}

func (x *Comment) fastWriteField5(buf []byte) (offset int) {
	if x.LikeCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 5, x.LikeCount)
	return offset
}

func (x *Comment) fastWriteField6(buf []byte) (offset int) {
	if x.TeaseCount == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 6, x.TeaseCount)
	return offset
}

func (x *CreateCommentRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *CreateCommentRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.Token)
	return offset
}

func (x *CreateCommentRequest) fastWriteField2(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.VideoId)
	return offset
}

func (x *CreateCommentRequest) fastWriteField3(buf []byte) (offset int) {
	if x.CommentText == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 3, x.CommentText)
	return offset
}

func (x *CreateCommentResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CreateCommentResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *CreateCommentResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Comment == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.Comment)
	return offset
}

func (x *DeleteCommentRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	offset += x.fastWriteField3(buf[offset:])
	return offset
}

func (x *DeleteCommentRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.Token)
	return offset
}

func (x *DeleteCommentRequest) fastWriteField2(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.VideoId)
	return offset
}

func (x *DeleteCommentRequest) fastWriteField3(buf []byte) (offset int) {
	if x.CommentId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 3, x.CommentId)
	return offset
}

func (x *DeleteCommentResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *DeleteCommentResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *DeleteCommentResponse) fastWriteField2(buf []byte) (offset int) {
	if x.Comment == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 2, x.Comment)
	return offset
}

func (x *CommentListRequest) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CommentListRequest) fastWriteField1(buf []byte) (offset int) {
	if x.Token == "" {
		return offset
	}
	offset += fastpb.WriteString(buf[offset:], 1, x.Token)
	return offset
}

func (x *CommentListRequest) fastWriteField2(buf []byte) (offset int) {
	if x.VideoId == 0 {
		return offset
	}
	offset += fastpb.WriteInt64(buf[offset:], 2, x.VideoId)
	return offset
}

func (x *CommentListResponse) FastWrite(buf []byte) (offset int) {
	if x == nil {
		return offset
	}
	offset += x.fastWriteField1(buf[offset:])
	offset += x.fastWriteField2(buf[offset:])
	return offset
}

func (x *CommentListResponse) fastWriteField1(buf []byte) (offset int) {
	if x.BaseResp == nil {
		return offset
	}
	offset += fastpb.WriteMessage(buf[offset:], 1, x.BaseResp)
	return offset
}

func (x *CommentListResponse) fastWriteField2(buf []byte) (offset int) {
	if x.CommentList == nil {
		return offset
	}
	for i := range x.CommentList {
		offset += fastpb.WriteMessage(buf[offset:], 2, x.CommentList[i])
	}
	return offset
}

func (x *BaseResp) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *BaseResp) sizeField1() (n int) {
	if x.StatusCode == 0 {
		return n
	}
	n += fastpb.SizeInt32(1, x.StatusCode)
	return n
}

func (x *BaseResp) sizeField2() (n int) {
	if x.StatusMessage == "" {
		return n
	}
	n += fastpb.SizeString(2, x.StatusMessage)
	return n
}

func (x *BaseResp) sizeField3() (n int) {
	if x.ServiceTime == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.ServiceTime)
	return n
}

func (x *User) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	n += x.sizeField7()
	n += x.sizeField8()
	n += x.sizeField9()
	n += x.sizeField10()
	n += x.sizeField11()
	return n
}

func (x *User) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *User) sizeField2() (n int) {
	if x.Name == "" {
		return n
	}
	n += fastpb.SizeString(2, x.Name)
	return n
}

func (x *User) sizeField3() (n int) {
	if x.FollowCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.FollowCount)
	return n
}

func (x *User) sizeField4() (n int) {
	if x.FollowerCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(4, x.FollowerCount)
	return n
}

func (x *User) sizeField5() (n int) {
	if !x.IsFollow {
		return n
	}
	n += fastpb.SizeBool(5, x.IsFollow)
	return n
}

func (x *User) sizeField6() (n int) {
	if x.Avatar == "" {
		return n
	}
	n += fastpb.SizeString(6, x.Avatar)
	return n
}

func (x *User) sizeField7() (n int) {
	if x.BackgroundImage == "" {
		return n
	}
	n += fastpb.SizeString(7, x.BackgroundImage)
	return n
}

func (x *User) sizeField8() (n int) {
	if x.Signature == "" {
		return n
	}
	n += fastpb.SizeString(8, x.Signature)
	return n
}

func (x *User) sizeField9() (n int) {
	if x.TotalFavorited == 0 {
		return n
	}
	n += fastpb.SizeInt64(9, x.TotalFavorited)
	return n
}

func (x *User) sizeField10() (n int) {
	if x.WorkCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(10, x.WorkCount)
	return n
}

func (x *User) sizeField11() (n int) {
	if x.FavoriteCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(11, x.FavoriteCount)
	return n
}

func (x *Comment) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	n += x.sizeField4()
	n += x.sizeField5()
	n += x.sizeField6()
	return n
}

func (x *Comment) sizeField1() (n int) {
	if x.Id == 0 {
		return n
	}
	n += fastpb.SizeInt64(1, x.Id)
	return n
}

func (x *Comment) sizeField2() (n int) {
	if x.User == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.User)
	return n
}

func (x *Comment) sizeField3() (n int) {
	if x.Content == "" {
		return n
	}
	n += fastpb.SizeString(3, x.Content)
	return n
}

func (x *Comment) sizeField4() (n int) {
	if x.CreateDate == "" {
		return n
	}
	n += fastpb.SizeString(4, x.CreateDate)
	return n
}

func (x *Comment) sizeField5() (n int) {
	if x.LikeCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(5, x.LikeCount)
	return n
}

func (x *Comment) sizeField6() (n int) {
	if x.TeaseCount == 0 {
		return n
	}
	n += fastpb.SizeInt64(6, x.TeaseCount)
	return n
}

func (x *CreateCommentRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *CreateCommentRequest) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.Token)
	return n
}

func (x *CreateCommentRequest) sizeField2() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.VideoId)
	return n
}

func (x *CreateCommentRequest) sizeField3() (n int) {
	if x.CommentText == "" {
		return n
	}
	n += fastpb.SizeString(3, x.CommentText)
	return n
}

func (x *CreateCommentResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CreateCommentResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *CreateCommentResponse) sizeField2() (n int) {
	if x.Comment == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.Comment)
	return n
}

func (x *DeleteCommentRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	n += x.sizeField3()
	return n
}

func (x *DeleteCommentRequest) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.Token)
	return n
}

func (x *DeleteCommentRequest) sizeField2() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.VideoId)
	return n
}

func (x *DeleteCommentRequest) sizeField3() (n int) {
	if x.CommentId == 0 {
		return n
	}
	n += fastpb.SizeInt64(3, x.CommentId)
	return n
}

func (x *DeleteCommentResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *DeleteCommentResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *DeleteCommentResponse) sizeField2() (n int) {
	if x.Comment == nil {
		return n
	}
	n += fastpb.SizeMessage(2, x.Comment)
	return n
}

func (x *CommentListRequest) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CommentListRequest) sizeField1() (n int) {
	if x.Token == "" {
		return n
	}
	n += fastpb.SizeString(1, x.Token)
	return n
}

func (x *CommentListRequest) sizeField2() (n int) {
	if x.VideoId == 0 {
		return n
	}
	n += fastpb.SizeInt64(2, x.VideoId)
	return n
}

func (x *CommentListResponse) Size() (n int) {
	if x == nil {
		return n
	}
	n += x.sizeField1()
	n += x.sizeField2()
	return n
}

func (x *CommentListResponse) sizeField1() (n int) {
	if x.BaseResp == nil {
		return n
	}
	n += fastpb.SizeMessage(1, x.BaseResp)
	return n
}

func (x *CommentListResponse) sizeField2() (n int) {
	if x.CommentList == nil {
		return n
	}
	for i := range x.CommentList {
		n += fastpb.SizeMessage(2, x.CommentList[i])
	}
	return n
}

var fieldIDToName_BaseResp = map[int32]string{
	1: "StatusCode",
	2: "StatusMessage",
	3: "ServiceTime",
}

var fieldIDToName_User = map[int32]string{
	1:  "Id",
	2:  "Name",
	3:  "FollowCount",
	4:  "FollowerCount",
	5:  "IsFollow",
	6:  "Avatar",
	7:  "BackgroundImage",
	8:  "Signature",
	9:  "TotalFavorited",
	10: "WorkCount",
	11: "FavoriteCount",
}

var fieldIDToName_Comment = map[int32]string{
	1: "Id",
	2: "User",
	3: "Content",
	4: "CreateDate",
	5: "LikeCount",
	6: "TeaseCount",
}

var fieldIDToName_CreateCommentRequest = map[int32]string{
	1: "Token",
	2: "VideoId",
	3: "CommentText",
}

var fieldIDToName_CreateCommentResponse = map[int32]string{
	1: "BaseResp",
	2: "Comment",
}

var fieldIDToName_DeleteCommentRequest = map[int32]string{
	1: "Token",
	2: "VideoId",
	3: "CommentId",
}

var fieldIDToName_DeleteCommentResponse = map[int32]string{
	1: "BaseResp",
	2: "Comment",
}

var fieldIDToName_CommentListRequest = map[int32]string{
	1: "Token",
	2: "VideoId",
}

var fieldIDToName_CommentListResponse = map[int32]string{
	1: "BaseResp",
	2: "CommentList",
}
