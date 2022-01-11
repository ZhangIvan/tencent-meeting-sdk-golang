//
//	直播管理 API
//
package qqmeeting

/*
	获取直播会看地址-模块
*/

// LiveReplaysQueryByIDRequest
// GET 通过会议 ID 查询直播回看 Request
type LiveReplaysQueryByIDRequest struct {
	MeetingID  string `param:"meeting_id"` // 有效的会议 ID。
	UserID     string `query:"userid"`     // 调用 API 的用户 ID。
	InstanceID int    `query:"instanceid"` // 用户的终端设备类型：
}

// LiveReplaysQueryByIDReply
// GET 通过会议 ID 查询直播回看 Reply
type LiveReplaysQueryByIDReply struct {
	MeetingNumber   int                `json:"meeting_number"`
	MeetingInfoList []*LiveReplaysList `json:"meeting_info_list"`
}

func (req LiveReplaysQueryByIDRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorLiveReplaysQueryByID
}

func (req LiveReplaysQueryByIDRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

// LiveReplaysQueryByCodeRequest
// GET 通过会议 Code 查询直播回看信息。 Request
type LiveReplaysQueryByCodeRequest struct {
	MeetingCode string `param:"meeting_code"` // 有效的会议 Code。
	UserID      string `query:"userid"`       // 调用 API 的用户 ID。
	InstanceID  int    `query:"instanceid"`   // 用户的终端设备类型：
}

func (req LiveReplaysQueryByCodeRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorLiveReplaysQueryByCode
}

func (req LiveReplaysQueryByCodeRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

/*
	修改直播配置-模块
*/
// MeetLiveConfigUpdateByIDRequest
// PUT 通过会议 ID 修改直播配置信息。 PUT
type MeetLiveConfigUpdateByIDRequest struct {
	MeetingID  string      `param:"meeting_id"` // 有效的会议 ID。
	UserID     string      `json:"userid"`      // 调用 API 的用户 ID。
	InstanceID int         `json:"instanceid"`  // 用户的终端设备类型：
	LiveConfig *LiveConfig `json:"live_config"` // 直播配置。
}

func (req MeetLiveConfigUpdateByIDRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetLiveConfigUpdateByID
}

func (req MeetLiveConfigUpdateByIDRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

// MeetLiveConfigUpdateByCodeRequest
// PUT 通过会议 CODE 修改直播配置信息。 PUT
type MeetLiveConfigUpdateByCodeRequest struct {
	MeetingCode string      `param:"meeting_code"`                 // 会议号码。
	UserID      string      `json:"userid"`                        // 调用 API 的用户 ID。
	InstanceID  int         `json:"instanceid" query:"instanceid"` // 用户的终端设备类型：
	LiveConfig  *LiveConfig `json:"live_config"`                   // 直播配置。
}

func (req MeetLiveConfigUpdateByCodeRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetLiveConfigUpdateByCode
}

func (req MeetLiveConfigUpdateByCodeRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}

// LiveReplayFileDeleteRequest
// DELETE 删除直播回看文件 DELETE
// 根据会议 ID 和直播房间 ID 删除该直播房间 ID 下的直播回看文件。
type LiveReplayFileDeleteRequest struct {
	MeetingID  string `path:"meeting_id"`   // 有效的会议 ID。
	LiveRoomID string `path:"live_room_id"` // 直播房间 ID。
	UserID     string `query:"userid"`      // 调用 API 的用户 ID。
	InstanceID int    `query:"instanceid"`  // 用户的终端设备类型：
}

func (req LiveReplayFileDeleteRequest) getDescriptor() *MeetingRequestDescriptor {
	return &RequestDescriptorMeetLiveDelete
}

func (req LiveReplayFileDeleteRequest) fillPlaceholder(args ...interface{}) string {
	return req.getDescriptor().Url
}
