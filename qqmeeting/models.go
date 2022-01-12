package qqmeeting

// API LIST
var (
	RequestDescriptorMeetingCreate                = MeetingRequestDescriptor{"/meetings", "POST", "Create"}
	RequestDescriptorMeetingQueryByID             = MeetingRequestDescriptor{"/meetings/%s", "GET", "ID"}
	RequestDescriptorMeetingQueryByCode           = MeetingRequestDescriptor{"/meetings", "GET", "Code"}
	RequestDescriptorMeetingCancelByID            = MeetingRequestDescriptor{"/meetings/%s/cancel", "POST", "CANCEL"}
	RequestDescriptorMeetingUpdateByID            = MeetingRequestDescriptor{"/meetings/%s", "PUT", "UPDATE"}
	RequestDescriptorMeetingQueryParticipantsByID = MeetingRequestDescriptor{"/meetings/%s/participants", "GET", "QUERY"}
	RequestDescriptorMeetingQueryUserMeetingList  = MeetingRequestDescriptor{"/meetings", "GET", "QUERY_MEETING_LIST"}

	RequestDescriptorUserCreate      = MeetingRequestDescriptor{"/users", "POST", "Create"}
	RequestDescriptorUserDetailQuery = MeetingRequestDescriptor{"/users/%s", "GET", "QUERY"}
	RequestDescriptorUserUpdate      = MeetingRequestDescriptor{"/users/%s", "PUT", "UPDATE"}
	RequestDescriptorUserDelete      = MeetingRequestDescriptor{"/users/%s", "DELETE", "DELETE"}
	RequestDescriptorUserList        = MeetingRequestDescriptor{"/users/list", "GET", "QUERY"}

	RequestDescriptorLiveReplaysQueryByID   = MeetingRequestDescriptor{"/meetings/%s/live_play/replays", "GET", "QUERY"}
	RequestDescriptorLiveReplaysQueryByCode = MeetingRequestDescriptor{"/meetings/live_play/replays", "GET", "QUERY"}

	RequestDescriptorMeetLiveConfigUpdateByID   = MeetingRequestDescriptor{"/meetings/%s/live_play/config", "PUT", "UPDATE"}
	RequestDescriptorMeetLiveConfigUpdateByCode = MeetingRequestDescriptor{"/meetings/live_play/config", "PUT", "UPDATE"}
	RequestDescriptorMeetLiveDelete             = MeetingRequestDescriptor{"/meetings/%s/live_play/%s/replays", "DELETE", "DELETE"}
)

// UserObj  用户对象 https://cloud.tencent.com/document/product/1095/42417
type UserObj struct {
	UserID      string `json:"userid"`              // 调用方用于标示用户的唯一 ID
	IsAnonymous bool   `json:"is_anonymous"`        // 匿名入会
	NickName    string `json:"nick_name,omitempty"` // 用户匿名字符串
}

// Settings 会议设置
type Settings struct {
	MuteEnableJoin             bool   `json:"mute_enable_join"`                        // 入会时静音
	AllowUnmuteSelf            bool   `json:"allow_unmute_self"`                       // 允许参会者取消静音
	MuteAll                    bool   `json:"mute_all,omitempty"`                      // 全体静音
	HostVideo                  bool   `json:"host_video,omitempty"`                    // 入会时主持人视频是否开启，暂时不支持。
	ParticipantVideo           bool   `json:"participant_video,omitempty"`             // 入会时参会者视频是否开启，暂时不支持。
	EnableRecord               bool   `json:"enable_record,omitempty"`                 // 开启录播，暂时不支持。
	PlayIVROnLeave             bool   `json:"play_ivr_on_leave,omitempty"`             // 参会者离开时播放提示音。
	PlayIVROnJoin              bool   `json:"play_ivr_on_join,omitempty"`              // 有新的入会者加入时播放提示音。
	LiveUrl                    bool   `json:"live_url,omitempty"`                      // 开启直播, 暂时不支持。
	AutoInWaitingRoom          bool   `json:"auto_in_waiting,omitempty"`               // 参会者离开时播放提示音，暂时不支持，可在客户端设置。
	AllowInBeforeHost          bool   `json:"allow_in_before_host"`                    // 是否允许成员在主持人进会前加入会议，默认值为 true。 true：允许 false：不允许
	AllowScreenSharedWatermark bool   `json:"allow_screen_shared_watermark,omitempty"` // 是否开启屏幕共享水印，默认值为 false。true： 开启false：不开启
	OnlyEnterpriseUserAllowed  bool   `json:"only_enterprise_user_allowed"`            // 是否仅企业内部成员可入会，默认值为 false。 true：仅企业内部用户可入会 false：所有人可入会
	ParticipantJoinAutoRecord  bool   `json:"participant_join_auto_record,omitempty"`  // 当有嘉宾入会时立即开启云录制，默认值为 false 关闭，关闭时，主持人入会自动开启云录制；当设置为开启时，则有嘉宾入会自动开启云录制。 说明： 该参数必须 auto_record_type 设置为“cloud”时才生效，该参数依赖企业账户设置，当企业强制锁定后，该参数必须与企业配置保持一致。 仅客户端2.7及以上版本生效。
	EnableHostPauseAutoRecord  bool   `json:"enable_host_pause_auto_record,omitempty"` // 允许主持人暂停或者停止云录制，默认值为 true 开启，开启时，主持人允许暂停和停止云录制；当设置为关闭时，则主持人不允许暂停和关闭云录制。 说明： 该参数必须 auto_record_type 设置为“cloud”时才生效，该参数依赖企业账户设置，当企业强制锁定后，该参数必须与企业配置保持一致。 仅客户端2.7及以上版本生效。
	WaterMarkType              int    `json:"watermark_type,omitempty"`                // 水印样式，默认为单排。当屏幕共享水印参数为开启时，此参数才生效。 0：单排 1：多排
	AutoRecordType             string `json:"auto_record_type,omitempty"`              // 自动会议录制类型。 none：禁用，表示不开启自动会议录制。 local：本地录制，表示主持人入会后自动开启本地录制。 cloud：云录制，表示主持人入会后自动开启云录制。 说明： 该参数依赖企业账户设置，当企业强制锁定后，该参数必须与企业配置保持一致。 仅客户端2.7及以上版本可生效

}

// LiveConfig 直播配置
type LiveConfig struct {
	LiveSubject        string `json:"live_subject,omitempty"`         // 直播主题，开启之后。
	LiveSummary        string `json:"live_summary,omitempty"`         // 直播简介。
	EnableLivePassword bool   `json:"enable_live_password,omitempty"` // 是否开启直播密码，默认值为 false。
	LivePassword       string `json:"live_password,omitempty"`        // 直播密码。当设置开启直播密码时，该参数必填。
	EnableLiveIm       bool   `json:"enable_live_im,omitempty"`       // 允许观众讨论，默认值为 false。
	EnableLiveReplay   bool   `json:"enable_live_replay,omitempty"`   // 开启直播回看，默认值为 false。
}

// RecurringRule 周期性会议
type RecurringRule struct {
	RecurringType int `json:"recurring_type,omitempty"` // 重复类型，默认值为0。 0：每天 1：每个工作日 2：每周 3：每两周 4：每月
	UntilType     int `json:"until_type,omitempty"`     // 结束重复类型，默认值为0。 0：按日期结束重复 1：按次数结束重复
	UntilDate     int `json:"until_date,omitempty"`     // 结束日期时间戳，默认值为当前日期 + 7天。 最大支持预定50场子会议。
	UntilCount    int `json:"until_count,omitempty"`    // 限定会议次数（1-50次），默认值为7次。
}

// UserInfo 用户信息
type UserInfo struct {
	Email    string `json:"email" binding:"required"`    // 邮箱地址
	Phone    string `json:"phone" binding:"required"`    // 手机号码
	Username string `json:"username" binding:"required"` // 用户昵称
	UserID   string `json:"userid" binding:"required"`   // 调用方用于表示用户的唯一ID
}

// UserDetail 用户详情
type UserDetail struct {
	Username   string `json:"username"`    // 用户昵称
	UpdateTime string `json:"update_time"` // 更新时间
	Status     string `json:"status"`      // 用户状态，1：正常；2：删除
	Email      string `json:"email"`       // 邮箱地址
	Phone      string `json:"phone"`       // 手机号码
	UserID     string `json:"userid"`      // 调用方用于标示用户的唯一 ID
	AreaCode   string `json:"area"`        // 地区编码（国内默认86）
	AvatarUrl  string `json:"avatar_url"`  // 头像
}

// MeetingInfo 会议信息
type MeetingInfo struct {
	MeetingID    string     `json:"meeting_id"`             // 会议的唯一标示
	MeetingCode  string     `json:"meeting_code"`           // 会议 App 的呼入号码
	Subject      string     `json:"subject"`                // 会议主题
	Hosts        []*UserObj `json:"hosts,omitempty"`        // 会议主持人的用户 ID，如果没有指定，主持人被设定为参数 userid 的用户，即 API 调用者。
	Participants []*UserObj `json:"participants,omitempty"` // 会议邀请的参会者，可为空
	StartTime    string     `json:"start_time"`             // 会议开始时间戳（单位秒）。
	EndTime      string     `json:"end_time"`               // 会议结束时间戳（单位秒）。
	Password     *string    `json:"password,omitempty"`     // 会议密码，可不填
	JoinUrl      string     `json:"join_url"`               // 加入会议　URL（点击链接直接加入会议）
	Settings     *Settings  `json:"settings,omitempty"`     // 会议设置
}

// MeetingQueryInfo 带状态的会议信息
type MeetingQueryInfo struct {
	MeetingInfo
	Status string `json:"status"` // 当前会议状态
}

// SimplifiedMeetingInfo 简要会议信息
type SimplifiedMeetingInfo struct {
	MeetingID   string `json:"meeting_id"`
	MeetingCode string `json:"meeting_code"`
}

// MeetingParticipants 参会人员信息
type MeetingParticipants struct {
	UserID                string `json:"userid"`    // 参会者用户 ID。
	Base64EncodedUsername string `json:"user_name"` // 入会用户名（base64）
	SHA256HashedPhone     string `json:"phone"`     // 参会者手机号 hash 值 SHA256（手机号/secretid）。
	JoinTime              string `json:"join_time"` // 参会者加入会议时间戳（单位秒）。
	LeftTime              string `json:"left_time"` // 参会者离开会议时间戳（单位秒）。
}

// MeetingQueryUserListMeetingInfo 查询用户的会议列表中会议对象
type MeetingQueryUserListMeetingInfo struct {
	MeetingID       string     `json:"meeting_id"`        // 会议的唯一标示
	MeetingCode     string     `json:"meeting_code"`      // 会议 App 的呼入号码
	Subject         string     `json:"subject"`           // 会议主题
	Hosts           []*UserObj `json:"hosts,omitempty"`   // 会议主持人的用户 ID，如果没有指定，主持人被设定为参数 userid 的用户，即 API 调用者。
	StartTime       string     `json:"start_time"`        // 会议开始时间戳（单位秒）。
	EndTime         string     `json:"end_time"`          // 会议结束时间戳（单位秒）。
	JoinMeetingRole string     `json:"join_meeting_role"` // 查询者在会议中的角色
}

// LiveReplaysList 直播回放会议对象数组
type LiveReplaysList struct {
	MeetingId      string             `json:"meeting_id"`                 // 会议 ID。
	MeetingCode    string             `json:"meeting_code"`               // 会议 CODE。
	Subject        string             `json:"subject"`                    // 会议主题。
	LiveReplayList []*LiveReplaysInfo `json:"live_replay_list,omitempty"` // 直播回看数组（会议创建人才有权限查询）。
}

// LiveReplaysInfo 直播回放对象
type LiveReplaysInfo struct {
	LiveRoomId  string `json:"live_room_id"` // 直播房间号。
	LiveSubject string `json:"live_subject"` // 直播主题。
	VideoUrl    string `json:"video_url"`    // 直播回放链接。
}
