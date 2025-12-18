# 爱心宠物领养平台 - UML模型图

## 目录

1. [功能结构图](#1-功能结构图)
2. [总体业务流程图](#2-总体业务流程图)
3. [功能模块业务流程图](#3-功能模块业务流程图)
4. [用例图](#4-用例图)
5. [类图](#5-类图)
6. [包图](#6-包图)
7. [ER图](#7-er图)
8. [状态图](#8-状态图)
9. [活动图](#9-活动图)
10. [时序图](#10-时序图)

---

## 1. 功能结构图

### 1.1 系统功能结构总图

```mermaid
graph TB
    System[爱心宠物领养平台]
    
    System --> UserModule[用户管理模块]
    System --> PetModule[宠物管理模块]
    System --> AdoptionModule[领养管理模块]
    System --> OrgModule[机构管理模块]
    System --> CommunityModule[社区互动模块]
    System --> DonationModule[捐赠管理模块]
    System --> AdminModule[后台管理模块]
    
    UserModule --> UserReg[用户注册]
    UserModule --> UserLogin[用户登录]
    UserModule --> UserProfile[个人信息管理]
    UserModule --> UserAuth[身份认证]
    
    PetModule --> PetPublish[宠物发布]
    PetModule --> PetBrowse[宠物浏览]
    PetModule --> PetSearch[宠物搜索]
    PetModule --> PetDetail[宠物详情]
    PetModule --> PetManage[我的宠物]
    
    AdoptionModule --> AdoptApply[提交申请]
    AdoptionModule --> AdoptReview[申请审核]
    AdoptionModule --> AdoptRecord[领养记录]
    AdoptionModule --> AdoptFollow[领养回访]
    
    OrgModule --> OrgRegister[机构注册]
    OrgModule --> OrgVerify[机构认证]
    OrgModule --> OrgManage[机构管理]
    OrgModule --> OrgDisplay[机构展示]
    
    CommunityModule --> PostPublish[发布动态]
    CommunityModule --> PostBrowse[浏览动态]
    CommunityModule --> PostComment[评论互动]
    CommunityModule --> PostLike[点赞收藏]
    
    DonationModule --> DonateCreate[发起捐赠]
    DonationModule --> DonateRecord[捐赠记录]
    DonationModule --> DonateWall[爱心墙]
    DonationModule --> DonateStats[捐赠统计]
    
    AdminModule --> AdminUser[用户管理]
    AdminModule --> AdminPet[宠物审核]
    AdminModule --> AdminAdoption[领养审核]
    AdminModule --> AdminOrg[机构审核]
    AdminModule --> AdminStats[数据统计]
    
    style System fill:#e1f5ff
    style UserModule fill:#fff4e1
    style PetModule fill:#e8f5e9
    style AdoptionModule fill:#f3e5f5
    style OrgModule fill:#fff3e0
    style CommunityModule fill:#fce4ec
    style DonationModule fill:#e0f2f1
    style AdminModule fill:#ffebee
```

### 1.2 用户管理模块功能结构

```mermaid
graph LR
    UserModule[用户管理模块]
    
    UserModule --> Register[用户注册]
    UserModule --> Login[用户登录]
    UserModule --> Profile[个人中心]
    UserModule --> Security[安全设置]
    
    Register --> RegForm[注册表单]
    Register --> RegValidate[信息验证]
    Register --> RegSubmit[提交注册]
    
    Login --> LoginForm[登录表单]
    Login --> LoginAuth[身份认证]
    Login --> LoginToken[生成Token]
    
    Profile --> ViewProfile[查看信息]
    Profile --> EditProfile[编辑信息]
    Profile --> UploadAvatar[上传头像]
    
    Security --> ChangePwd[修改密码]
    Security --> BindPhone[绑定手机]
    Security --> BindEmail[绑定邮箱]
```

### 1.3 宠物管理模块功能结构

```mermaid
graph LR
    PetModule[宠物管理模块]
    
    PetModule --> Browse[宠物浏览]
    PetModule --> Publish[宠物发布]
    PetModule --> Manage[宠物管理]
    PetModule --> Detail[宠物详情]
    
    Browse --> List[列表展示]
    Browse --> Filter[条件筛选]
    Browse --> Search[关键词搜索]
    Browse --> Recommend[推荐宠物]
    
    Publish --> FillInfo[填写信息]
    Publish --> UploadPhoto[上传照片]
    Publish --> Submit[提交发布]
    Publish --> Review[等待审核]
    
    Manage --> MyPets[我的宠物]
    Manage --> Edit[编辑信息]
    Manage --> Delete[删除宠物]
    Manage --> Offline[下架宠物]
    
    Detail --> ViewDetail[查看详情]
    Detail --> ViewPhotos[查看照片]
    Detail --> Contact[联系发布者]
    Detail --> ApplyAdopt[申请领养]
```

### 1.4 领养管理模块功能结构

```mermaid
graph LR
    AdoptionModule[领养管理模块]
    
    AdoptionModule --> Application[领养申请]
    AdoptionModule --> Review[申请审核]
    AdoptionModule --> Record[领养记录]
    AdoptionModule --> FollowUp[回访管理]
    
    Application --> FillForm[填写申请表]
    Application --> UploadDoc[上传证明]
    Application --> SubmitApp[提交申请]
    Application --> TrackStatus[跟踪状态]
    
    Review --> ViewApp[查看申请]
    Review --> Interview[安排面试]
    Review --> HomeVisit[安排家访]
    Review --> Approve[审核通过]
    Review --> Reject[审核拒绝]
    
    Record --> CreateRecord[创建记录]
    Record --> ViewRecord[查看记录]
    Record --> UpdateStatus[更新状态]
    Record --> Agreement[签署协议]
    
    FollowUp --> PlanFollow[制定计划]
    FollowUp --> ExecuteFollow[执行回访]
    FollowUp --> RecordFollow[记录回访]
    FollowUp --> Feedback[反馈处理]
```

### 1.5 机构管理模块功能结构

```mermaid
graph LR
    OrgModule[机构管理模块]
    
    OrgModule --> Register[机构注册]
    OrgModule --> Verify[机构认证]
    OrgModule --> Manage[机构管理]
    OrgModule --> Display[机构展示]
    
    Register --> FillOrgInfo[填写机构信息]
    Register --> UploadCredential[上传资质证明]
    Register --> SubmitReg[提交注册]
    Register --> WaitVerify[等待审核]
    
    Verify --> ReviewInfo[审核信息]
    Verify --> VerifyCredential[验证资质]
    Verify --> ApproveOrg[通过认证]
    Verify --> RejectOrg[拒绝认证]
    
    Manage --> ViewOrg[查看机构]
    Manage --> EditOrg[编辑信息]
    Manage --> ManagePets[管理宠物]
    Manage --> ManageAdoptions[管理领养]
    
    Display --> OrgList[机构列表]
    Display --> OrgDetail[机构详情]
    Display --> OrgPets[机构宠物]
    Display --> OrgStats[机构统计]
```

### 1.6 社区互动模块功能结构

```mermaid
graph LR
    CommunityModule[社区互动模块]
    
    CommunityModule --> Post[动态发布]
    CommunityModule --> Browse[动态浏览]
    CommunityModule --> Interact[互动功能]
    CommunityModule --> Manage[动态管理]
    
    Post --> CreatePost[创建动态]
    Post --> UploadImages[上传图片]
    Post --> SelectCategory[选择分类]
    Post --> PublishPost[发布动态]
    
    Browse --> PostList[动态列表]
    Browse --> PostDetail[动态详情]
    Browse --> SearchPost[搜索动态]
    Browse --> FilterPost[筛选动态]
    
    Interact --> Like[点赞]
    Interact --> Comment[评论]
    Interact --> Reply[回复]
    Interact --> Share[分享]
    
    Manage --> MyPosts[我的动态]
    Manage --> EditPost[编辑动态]
    Manage --> DeletePost[删除动态]
    Manage --> ViewStats[查看统计]
```

### 1.7 捐赠管理模块功能结构

```mermaid
graph LR
    DonationModule[捐赠管理模块]
    
    DonationModule --> Donate[发起捐赠]
    DonationModule --> Record[捐赠记录]
    DonationModule --> Display[捐赠展示]
    DonationModule --> Stats[捐赠统计]
    
    Donate --> SelectOrg[选择机构]
    Donate --> InputAmount[输入金额]
    Donate --> FillInfo[填写信息]
    Donate --> Payment[支付捐赠]
    
    Record --> MyDonations[我的捐赠]
    Record --> ViewDetail[查看详情]
    Record --> DownloadReceipt[下载凭证]
    Record --> TrackUsage[追踪使用]
    
    Display --> LoveWall[爱心墙]
    Display --> TopDonors[捐赠榜]
    Display --> RecentDonations[最新捐赠]
    Display --> OrgReceived[机构收款]
    
    Stats --> TotalAmount[总金额]
    Stats --> DonorCount[捐赠人数]
    Stats --> OrgStats[机构统计]
    Stats --> TrendChart[趋势图表]
```

### 1.8 后台管理模块功能结构

```mermaid
graph LR
    AdminModule[后台管理模块]
    
    AdminModule --> Dashboard[数据仪表盘]
    AdminModule --> UserMgmt[用户管理]
    AdminModule --> PetMgmt[宠物管理]
    AdminModule --> AdoptionMgmt[领养管理]
    AdminModule --> OrgMgmt[机构管理]
    AdminModule --> ContentMgmt[内容管理]
    AdminModule --> SystemMgmt[系统管理]
    
    Dashboard --> Overview[总览统计]
    Dashboard --> Charts[数据图表]
    Dashboard --> Reports[报表生成]
    
    UserMgmt --> UserList[用户列表]
    UserMgmt --> UserSearch[用户搜索]
    UserMgmt --> UserStatus[状态管理]
    UserMgmt --> UserRole[角色管理]
    
    PetMgmt --> PetReview[宠物审核]
    PetMgmt --> PetList[宠物列表]
    PetMgmt --> PetEdit[宠物编辑]
    PetMgmt --> PetStats[宠物统计]
    
    AdoptionMgmt --> AppReview[申请审核]
    AdoptionMgmt --> RecordMgmt[记录管理]
    AdoptionMgmt --> FollowMgmt[回访管理]
    AdoptionMgmt --> AdoptStats[领养统计]
    
    OrgMgmt --> OrgReview[机构审核]
    OrgMgmt --> OrgList[机构列表]
    OrgMgmt --> OrgEdit[机构编辑]
    OrgMgmt --> OrgStats[机构统计]
    
    ContentMgmt --> PostMgmt[动态管理]
    ContentMgmt --> CommentMgmt[评论管理]
    ContentMgmt --> DonationMgmt[捐赠管理]
    
    SystemMgmt --> ConfigMgmt[配置管理]
    SystemMgmt --> LogMgmt[日志管理]
    SystemMgmt --> BackupMgmt[备份管理]
```

## 2. 总体业务流程图

### 2.1 平台整体业务流程

```mermaid
flowchart TD
    Start([用户访问平台]) --> Register{是否注册}
    
    Register -->|否| RegProcess[注册流程]
    Register -->|是| Login[登录系统]
    
    RegProcess --> RegForm[填写注册信息]
    RegForm --> RegValidate{信息验证}
    RegValidate -->|失败| RegForm
    RegValidate -->|成功| RegSubmit[提交注册]
    RegSubmit --> Login
    
    Login --> LoginAuth{身份认证}
    LoginAuth -->|失败| Login
    LoginAuth -->|成功| SelectRole{选择角色}
    
    SelectRole -->|普通用户| UserFlow[用户业务流程]
    SelectRole -->|机构用户| OrgFlow[机构业务流程]
    SelectRole -->|管理员| AdminFlow[管理员业务流程]
    
    UserFlow --> UserAction{用户操作}
    UserAction -->|浏览宠物| BrowsePet[浏览宠物信息]
    UserAction -->|申请领养| ApplyAdopt[提交领养申请]
    UserAction -->|发布宠物| PublishPet[发布宠物信息]
    UserAction -->|社区互动| Community[参与社区互动]
    UserAction -->|爱心捐赠| Donate[发起捐赠]
    
    BrowsePet --> ViewDetail[查看宠物详情]
    ViewDetail --> ApplyAdopt
    
    ApplyAdopt --> FillApp[填写申请表]
    FillApp --> SubmitApp[提交申请]
    SubmitApp --> WaitReview[等待审核]
    WaitReview --> ReviewResult{审核结果}
    ReviewResult -->|通过| AdoptSuccess[领养成功]
    ReviewResult -->|拒绝| ApplyEnd[申请结束]
    ReviewResult -->|需面试| Interview[安排面试]
    Interview --> HomeVisit[安排家访]
    HomeVisit --> ReviewResult
    
    PublishPet --> FillPetInfo[填写宠物信息]
    FillPetInfo --> UploadPhoto[上传宠物照片]
    UploadPhoto --> SubmitPet[提交发布]
    SubmitPet --> PetReview[等待审核]
    PetReview --> PetResult{审核结果}
    PetResult -->|通过| PetOnline[宠物上线]
    PetResult -->|拒绝| PetEnd[发布结束]
    
    Community --> PostAction{社区操作}
    PostAction -->|发布动态| CreatePost[创建动态]
    PostAction -->|浏览动态| BrowsePost[浏览动态]
    PostAction -->|评论互动| CommentPost[评论动态]
    PostAction -->|点赞收藏| LikePost[点赞动态]
    
    Donate --> SelectOrg[选择捐赠机构]
    SelectOrg --> InputAmount[输入捐赠金额]
    InputAmount --> Payment[支付捐赠]
    Payment --> PayResult{支付结果}
    PayResult -->|成功| DonateSuccess[捐赠成功]
    PayResult -->|失败| Donate
    
    OrgFlow --> OrgAction{机构操作}
    OrgAction -->|管理宠物| ManagePets[管理宠物]
    OrgAction -->|审核申请| ReviewApps[审核领养申请]
    OrgAction -->|接收捐赠| ReceiveDonate[接收捐赠]
    OrgAction -->|发布动态| OrgPost[发布机构动态]
    
    AdminFlow --> AdminAction{管理操作}
    AdminAction -->|用户管理| ManageUsers[管理用户]
    AdminAction -->|宠物审核| ReviewPets[审核宠物]
    AdminAction -->|领养审核| ReviewAdoptions[审核领养]
    AdminAction -->|机构审核| ReviewOrgs[审核机构]
    AdminAction -->|数据统计| ViewStats[查看统计]
    
    AdoptSuccess --> End([流程结束])
    ApplyEnd --> End
    PetOnline --> End
    PetEnd --> End
    DonateSuccess --> End
    ManagePets --> End
    ReviewApps --> End
    ReceiveDonate --> End
    OrgPost --> End
    ManageUsers --> End
    ReviewPets --> End
    ReviewAdoptions --> End
    ReviewOrgs --> End
    ViewStats --> End
    CreatePost --> End
    BrowsePost --> End
    CommentPost --> End
    LikePost --> End
    
    style Start fill:#e1f5ff
    style End fill:#e1f5ff
    style UserFlow fill:#fff4e1
    style OrgFlow fill:#e8f5e9
    style AdminFlow fill:#ffebee
```



## 3. 功能模块业务流程图

### 3.1 用户注册登录流程

```mermaid
flowchart TD
    Start([开始]) --> ChooseAction{选择操作}
    
    ChooseAction -->|注册| RegStart[进入注册页面]
    ChooseAction -->|登录| LoginStart[进入登录页面]
    
    %% 注册流程
    RegStart --> InputRegInfo[输入注册信息]
    InputRegInfo --> RegValidate{验证信息}
    RegValidate -->|用户名已存在| RegError1[提示用户名已存在]
    RegValidate -->|手机号已注册| RegError2[提示手机号已注册]
    RegValidate -->|邮箱已注册| RegError3[提示邮箱已注册]
    RegValidate -->|密码强度不够| RegError4[提示密码强度不够]
    RegError1 --> InputRegInfo
    RegError2 --> InputRegInfo
    RegError3 --> InputRegInfo
    RegError4 --> InputRegInfo
    
    RegValidate -->|验证通过| EncryptPwd[加密密码]
    EncryptPwd --> SaveUser[保存用户信息]
    SaveUser --> RegSuccess[注册成功]
    RegSuccess --> LoginStart
    
    %% 登录流程
    LoginStart --> InputLoginInfo[输入登录信息]
    InputLoginInfo --> CheckUser{验证用户}
    CheckUser -->|用户不存在| LoginError1[提示用户不存在]
    CheckUser -->|密码错误| LoginError2[提示密码错误]
    CheckUser -->|账号被禁用| LoginError3[提示账号被禁用]
    LoginError1 --> InputLoginInfo
    LoginError2 --> InputLoginInfo
    LoginError3 --> End
    
    CheckUser -->|验证通过| GenerateToken[生成JWT Token]
    GenerateToken --> CacheUser[缓存用户信息]
    CacheUser --> UpdateLoginInfo[更新登录信息]
    UpdateLoginInfo --> LoginSuccess[登录成功]
    LoginSuccess --> RedirectHome[跳转到首页]
    RedirectHome --> End([结束])
    
    style Start fill:#e1f5ff
    style End fill:#e1f5ff
    style RegSuccess fill:#c8e6c9
    style LoginSuccess fill:#c8e6c9
    style RegError1 fill:#ffcdd2
    style RegError2 fill:#ffcdd2
    style RegError3 fill:#ffcdd2
    style RegError4 fill:#ffcdd2
    style LoginError1 fill:#ffcdd2
    style LoginError2 fill:#ffcdd2
    style LoginError3 fill:#ffcdd2
```

### 3.2 宠物发布审核流程

```mermaid
flowchart TD
    Start([开始]) --> CheckLogin{检查登录}
    CheckLogin -->|未登录| ToLogin[跳转登录]
    ToLogin --> End([结束])
    
    CheckLogin -->|已登录| EnterPublish[进入发布页面]
    EnterPublish --> FillBasicInfo[填写基本信息]
    FillBasicInfo --> UploadPhotos[上传宠物照片]
    UploadPhotos --> FillDetailInfo[填写详细信息]
    FillDetailInfo --> FillLocationInfo[填写位置信息]
    FillLocationInfo --> PreviewInfo[预览信息]
    
    PreviewInfo --> ConfirmSubmit{确认提交}
    ConfirmSubmit -->|取消| EnterPublish
    ConfirmSubmit -->|确认| ValidateInfo{验证信息}
    
    ValidateInfo -->|必填项缺失| ShowError1[提示必填项]
    ValidateInfo -->|照片未上传| ShowError2[提示上传照片]
    ValidateInfo -->|信息格式错误| ShowError3[提示格式错误]
    ShowError1 --> FillBasicInfo
    ShowError2 --> UploadPhotos
    ShowError3 --> FillDetailInfo
    
    ValidateInfo -->|验证通过| SavePet[保存宠物信息]
    SavePet --> SetStatusPending[设置状态为待审核]
    SetStatusPending --> NotifyAdmin[通知管理员审核]
    NotifyAdmin --> ShowSuccess[显示提交成功]
    ShowSuccess --> WaitReview[等待审核]
    
    WaitReview --> AdminReview{管理员审核}
    AdminReview -->|审核通过| SetStatusAvailable[设置状态为可领养]
    AdminReview -->|审核拒绝| SetStatusRejected[设置状态为已拒绝]
    AdminReview -->|需要修改| RequestModify[要求修改]
    
    SetStatusAvailable --> NotifyUserApprove[通知用户审核通过]
    NotifyUserApprove --> PetOnline[宠物上线展示]
    PetOnline --> End
    
    SetStatusRejected --> SaveRejectReason[保存拒绝原因]
    SaveRejectReason --> NotifyUserReject[通知用户审核拒绝]
    NotifyUserReject --> End
    
    RequestModify --> NotifyUserModify[通知用户修改]
    NotifyUserModify --> UserModify{用户修改}
    UserModify -->|修改并重新提交| FillBasicInfo
    UserModify -->|放弃修改| End
    
    style Start fill:#e1f5ff
    style End fill:#e1f5ff
    style PetOnline fill:#c8e6c9
    style ShowError1 fill:#ffcdd2
    style ShowError2 fill:#ffcdd2
    style ShowError3 fill:#ffcdd2
```

### 3.3 领养申请审核流程

```mermaid
flowchart TD
    Start([开始]) --> ViewPet[查看宠物详情]
    ViewPet --> CheckPetStatus{检查宠物状态}
    CheckPetStatus -->|已被领养| ShowAdopted[显示已被领养]
    CheckPetStatus -->|已下架| ShowOffline[显示已下架]
    ShowAdopted --> End([结束])
    ShowOffline --> End
    
    CheckPetStatus -->|可领养| CheckLogin{检查登录}
    CheckLogin -->|未登录| ToLogin[跳转登录]
    ToLogin --> End
    
    CheckLogin -->|已登录| ClickApply[点击申请领养]
    ClickApply --> EnterAppPage[进入申请页面]
    EnterAppPage --> FillPersonalInfo[填写个人信息]
    FillPersonalInfo --> FillHousingInfo[填写住房信息]
    FillHousingInfo --> FillFamilyInfo[填写家庭信息]
    FillFamilyInfo --> FillExperience[填写养宠经验]
    FillExperience --> FillReason[填写领养原因]
    FillReason --> UploadDocs[上传证明文件]
    UploadDocs --> PreviewApp[预览申请]
    
    PreviewApp --> ConfirmSubmit{确认提交}
    ConfirmSubmit -->|取消| EnterAppPage
    ConfirmSubmit -->|确认| ValidateApp{验证申请}
    
    ValidateApp -->|必填项缺失| ShowError1[提示必填项]
    ValidateApp -->|信息不完整| ShowError2[提示信息不完整]
    ValidateApp -->|文件未上传| ShowError3[提示上传文件]
    ShowError1 --> FillPersonalInfo
    ShowError2 --> FillFamilyInfo
    ShowError3 --> UploadDocs
    
    ValidateApp -->|验证通过| GenerateAppNo[生成申请编号]
    GenerateAppNo --> SaveApplication[保存申请信息]
    SaveApplication --> SetStatusPending[设置状态为待审核]
    SetStatusPending --> NotifyOrg[通知机构审核]
    NotifyOrg --> ShowAppSuccess[显示申请成功]
    ShowAppSuccess --> WaitReview[等待审核]
    
    WaitReview --> OrgReview{机构初审}
    OrgReview -->|初审通过| SetStatusReviewing[设置为审核中]
    OrgReview -->|初审拒绝| RejectApp1[拒绝申请]
    
    SetStatusReviewing --> ArrangeInterview[安排面试]
    ArrangeInterview --> SetStatusInterview[设置为待面试]
    SetStatusInterview --> ConductInterview[进行面试]
    ConductInterview --> InterviewResult{面试结果}
    
    InterviewResult -->|通过| ArrangeHomeVisit[安排家访]
    InterviewResult -->|不通过| RejectApp2[拒绝申请]
    
    ArrangeHomeVisit --> SetStatusHomeVisit[设置为待家访]
    SetStatusHomeVisit --> ConductHomeVisit[进行家访]
    ConductHomeVisit --> HomeVisitResult{家访结果}
    
    HomeVisitResult -->|通过| SetStatusApproved[设置为已通过]
    HomeVisitResult -->|不通过| RejectApp3[拒绝申请]
    
    SetStatusApproved --> NotifyUserApprove[通知用户审核通过]
    NotifyUserApprove --> CreateAdoption[创建领养记录]
    CreateAdoption --> UpdatePetStatus[更新宠物状态为已领养]
    UpdatePetStatus --> SignAgreement[签署领养协议]
    SignAgreement --> ArrangeHandover[安排交接]
    ArrangeHandover --> AdoptSuccess[领养成功]
    AdoptSuccess --> PlanFollowUp[制定回访计划]
    PlanFollowUp --> End
    
    RejectApp1 --> SaveRejectReason1[保存拒绝原因]
    RejectApp2 --> SaveRejectReason2[保存拒绝原因]
    RejectApp3 --> SaveRejectReason3[保存拒绝原因]
    SaveRejectReason1 --> NotifyUserReject[通知用户拒绝]
    SaveRejectReason2 --> NotifyUserReject
    SaveRejectReason3 --> NotifyUserReject
    NotifyUserReject --> End
    
    style Start fill:#e1f5ff
    style End fill:#e1f5ff
    style AdoptSuccess fill:#c8e6c9
    style ShowError1 fill:#ffcdd2
    style ShowError2 fill:#ffcdd2
    style ShowError3 fill:#ffcdd2
    style RejectApp1 fill:#ffcdd2
    style RejectApp2 fill:#ffcdd2
    style RejectApp3 fill:#ffcdd2
```

### 3.4 机构注册认证流程

```mermaid
flowchart TD
    Start([开始]) --> EnterRegPage[进入机构注册页面]
    EnterRegPage --> SelectType[选择机构类型]
    SelectType --> FillOrgInfo[填写机构信息]
    FillOrgInfo --> FillContactInfo[填写联系信息]
    FillContactInfo --> UploadLogo[上传机构Logo]
    UploadLogo --> UploadCredentials[上传资质证明]
    UploadCredentials --> FillDescription[填写机构描述]
    FillDescription --> PreviewOrg[预览机构信息]
    
    PreviewOrg --> ConfirmSubmit{确认提交}
    ConfirmSubmit -->|取消| EnterRegPage
    ConfirmSubmit -->|确认| ValidateOrg{验证信息}
    
    ValidateOrg -->|机构名称已存在| ShowError1[提示名称已存在]
    ValidateOrg -->|联系方式无效| ShowError2[提示联系方式无效]
    ValidateOrg -->|资质证明缺失| ShowError3[提示上传资质]
    ValidateOrg -->|信息不完整| ShowError4[提示信息不完整]
    ShowError1 --> FillOrgInfo
    ShowError2 --> FillContactInfo
    ShowError3 --> UploadCredentials
    ShowError4 --> FillOrgInfo
    
    ValidateOrg -->|验证通过| SaveOrg[保存机构信息]
    SaveOrg --> SetStatusPending[设置状态为待审核]
    SetStatusPending --> NotifyAdmin[通知管理员审核]
    NotifyAdmin --> ShowRegSuccess[显示注册成功]
    ShowRegSuccess --> WaitVerify[等待审核]
    
    WaitVerify --> AdminVerify{管理员审核}
    AdminVerify -->|审核通过| VerifyCredentials[验证资质真实性]
    AdminVerify -->|审核拒绝| RejectOrg[拒绝认证]
    AdminVerify -->|需要补充材料| RequestMore[要求补充材料]
    
    VerifyCredentials --> CheckResult{验证结果}
    CheckResult -->|验证通过| SetStatusApproved[设置状态为已认证]
    CheckResult -->|验证失败| RejectOrg
    
    SetStatusApproved --> GrantPermissions[授予机构权限]
    GrantPermissions --> NotifyOrgApprove[通知机构审核通过]
    NotifyOrgApprove --> OrgOnline[机构上线]
    OrgOnline --> EnableFeatures[启用机构功能]
    EnableFeatures --> End([结束])
    
    RejectOrg --> SaveRejectReason[保存拒绝原因]
    SaveRejectReason --> NotifyOrgReject[通知机构拒绝]
    NotifyOrgReject --> End
    
    RequestMore --> NotifyOrgMore[通知机构补充材料]
    NotifyOrgMore --> OrgSupply{机构补充}
    OrgSupply -->|补充材料| UploadCredentials
    OrgSupply -->|放弃| End
    
    style Start fill:#e1f5ff
    style End fill:#e1f5ff
    style OrgOnline fill:#c8e6c9
    style ShowError1 fill:#ffcdd2
    style ShowError2 fill:#ffcdd2
    style ShowError3 fill:#ffcdd2
    style ShowError4 fill:#ffcdd2
```

### 3.5 社区动态发布流程

```mermaid
flowchart TD
    Start([开始]) --> CheckLogin{检查登录}
    CheckLogin -->|未登录| ToLogin[跳转登录]
    ToLogin --> End([结束])
    
    CheckLogin -->|已登录| EnterPostPage[进入发布页面]
    EnterPostPage --> InputTitle[输入动态标题]
    InputTitle --> InputContent[输入动态内容]
    InputContent --> SelectCategory[选择动态分类]
    SelectCategory --> UploadImages[上传图片]
    UploadImages --> AddTags[添加标签]
    AddTags --> PreviewPost[预览动态]
    
    PreviewPost --> ConfirmPublish{确认发布}
    ConfirmPublish -->|取消| EnterPostPage
    ConfirmPublish -->|保存草稿| SaveDraft[保存为草稿]
    ConfirmPublish -->|立即发布| ValidatePost{验证内容}
    
    SaveDraft --> ShowDraftSuccess[显示保存成功]
    ShowDraftSuccess --> End
    
    ValidatePost -->|标题为空| ShowError1[提示输入标题]
    ValidatePost -->|内容为空| ShowError2[提示输入内容]
    ValidatePost -->|内容过长| ShowError3[提示内容过长]
    ValidatePost -->|包含敏感词| ShowError4[提示包含敏感词]
    ShowError1 --> InputTitle
    ShowError2 --> InputContent
    ShowError3 --> InputContent
    ShowError4 --> InputContent
    
    ValidatePost -->|验证通过| SavePost[保存动态]
    SavePost --> SetStatusPublished[设置状态为已发布]
    SetStatusPublished --> NotifyFollowers[通知关注者]
    NotifyFollowers --> ShowPublishSuccess[显示发布成功]
    ShowPublishSuccess --> PostOnline[动态上线]
    PostOnline --> End
    
    style Start fill:#e1f5ff
    style End fill:#e1f5ff
    style PostOnline fill:#c8e6c9
    style ShowError1 fill:#ffcdd2
    style ShowError2 fill:#ffcdd2
    style ShowError3 fill:#ffcdd2
    style ShowError4 fill:#ffcdd2
```

### 3.6 捐赠支付流程

```mermaid
flowchart TD
    Start([开始]) --> ViewDonate[进入捐赠页面]
    ViewDonate --> SelectOrg[选择捐赠机构]
    SelectOrg --> ViewOrgInfo[查看机构信息]
    ViewOrgInfo --> SelectAmount{选择金额}
    
    SelectAmount -->|预设金额| ClickPreset[点击预设金额]
    SelectAmount -->|自定义金额| InputCustom[输入自定义金额]
    
    ClickPreset --> ValidateAmount1{验证金额}
    InputCustom --> ValidateAmount2{验证金额}
    
    ValidateAmount1 -->|金额无效| ShowError1[提示金额无效]
    ValidateAmount2 -->|金额无效| ShowError2[提示金额无效]
    ValidateAmount2 -->|金额过小| ShowError3[提示金额过小]
    ValidateAmount2 -->|金额过大| ShowError4[提示金额过大]
    ShowError1 --> SelectAmount
    ShowError2 --> InputCustom
    ShowError3 --> InputCustom
    ShowError4 --> InputCustom
    
    ValidateAmount1 -->|验证通过| FillDonorInfo[填写捐赠人信息]
    ValidateAmount2 -->|验证通过| FillDonorInfo
    
    FillDonorInfo --> InputMessage[输入留言]
    InputMessage --> SelectAnonymous{选择是否匿名}
    SelectAnonymous -->|匿名| SetAnonymous[设置为匿名]
    SelectAnonymous -->|实名| SetRealName[设置为实名]
    
    SetAnonymous --> SelectPayment[选择支付方式]
    SetRealName --> SelectPayment
    
    SelectPayment --> ConfirmDonate{确认捐赠}
    ConfirmDonate -->|取消| End([结束])
    ConfirmDonate -->|确认| CreateDonation[创建捐赠记录]
    
    CreateDonation --> GenerateDonationNo[生成捐赠编号]
    GenerateDonationNo --> SetStatusPending[设置状态为待支付]
    SetStatusPending --> RedirectPayment[跳转支付页面]
    
    RedirectPayment --> ProcessPayment{处理支付}
    ProcessPayment -->|支付成功| PaymentSuccess[支付成功]
    ProcessPayment -->|支付失败| PaymentFailed[支付失败]
    ProcessPayment -->|支付取消| PaymentCancelled[支付取消]
    
    PaymentSuccess --> UpdateStatusPaid[更新状态为已支付]
    UpdateStatusPaid --> NotifyOrg[通知机构]
    NotifyOrg --> UpdateOrgBalance[更新机构余额]
    UpdateOrgBalance --> GenerateReceipt[生成捐赠凭证]
    GenerateReceipt --> ShowOnLoveWall[显示在爱心墙]
    ShowOnLoveWall --> SendThankYou[发送感谢信]
    SendThankYou --> DonateSuccess[捐赠成功]
    DonateSuccess --> End
    
    PaymentFailed --> ShowFailMsg[显示失败信息]
    ShowFailMsg --> RetryPayment{重试支付}
    RetryPayment -->|是| RedirectPayment
    RetryPayment -->|否| CancelDonation[取消捐赠]
    
    PaymentCancelled --> CancelDonation
    CancelDonation --> UpdateStatusCancelled[更新状态为已取消]
    UpdateStatusCancelled --> End
    
    style Start fill:#e1f5ff
    style End fill:#e1f5ff
    style DonateSuccess fill:#c8e6c9
    style ShowError1 fill:#ffcdd2
    style ShowError2 fill:#ffcdd2
    style ShowError3 fill:#ffcdd2
    style ShowError4 fill:#ffcdd2
    style PaymentFailed fill:#ffcdd2
```



## 4. 用例图 (Use Case Diagram)

### 4.1 系统总体用例图

```mermaid
graph TB
    %% 参与者
    User((普通用户))
    Org((机构用户))
    Admin((管理员))
    Guest((访客))
    
    %% 用例
    subgraph "爱心宠物领养平台"
        %% 访客用例
        UC1[浏览宠物]
        UC2[查看宠物详情]
        UC3[浏览社区动态]
        UC4[查看机构信息]
        UC5[用户注册]
        UC6[用户登录]
        
        %% 普通用户用例
        UC7[发布宠物]
        UC8[管理我的宠物]
        UC9[申请领养]
        UC10[查看我的申请]
        UC11[发布动态]
        UC12[评论互动]
        UC13[发起捐赠]
        UC14[查看捐赠记录]
        UC15[个人信息管理]
        
        %% 机构用户用例
        UC16[机构注册]
        UC17[管理机构信息]
        UC18[审核领养申请]
        UC19[管理机构宠物]
        UC20[接收捐赠]
        UC21[发布机构动态]
        
        %% 管理员用例
        UC22[用户管理]
        UC23[宠物审核]
        UC24[领养审核]
        UC25[机构审核]
        UC26[内容管理]
        UC27[数据统计]
        UC28[系统配置]
    end
    
    %% 关联关系
    Guest --> UC1
    Guest --> UC2
    Guest --> UC3
    Guest --> UC4
    Guest --> UC5
    Guest --> UC6
    
    User --> UC1
    User --> UC2
    User --> UC3
    User --> UC4
    User --> UC6
    User --> UC7
    User --> UC8
    User --> UC9
    User --> UC10
    User --> UC11
    User --> UC12
    User --> UC13
    User --> UC14
    User --> UC15
    
    Org --> UC1
    Org --> UC2
    Org --> UC3
    Org --> UC6
    Org --> UC16
    Org --> UC17
    Org --> UC18
    Org --> UC19
    Org --> UC20
    Org --> UC21
    
    Admin --> UC22
    Admin --> UC23
    Admin --> UC24
    Admin --> UC25
    Admin --> UC26
    Admin --> UC27
    Admin --> UC28
    
    %% 包含关系
    UC7 -.->|include| UC6
    UC9 -.->|include| UC6
    UC11 -.->|include| UC6
    UC13 -.->|include| UC6
    UC18 -.->|include| UC6
    UC22 -.->|include| UC6
    UC23 -.->|include| UC6
    UC24 -.->|include| UC6
    UC25 -.->|include| UC6
```

### 4.2 用户管理用例图

```mermaid
graph TB
    User((用户))
    System((系统))
    
    subgraph "用户管理"
        UC1[用户注册]
        UC2[用户登录]
        UC3[查看个人信息]
        UC4[修改个人信息]
        UC5[上传头像]
        UC6[修改密码]
        UC7[绑定手机]
        UC8[绑定邮箱]
        UC9[找回密码]
        UC10[退出登录]
    end
    
    User --> UC1
    User --> UC2
    User --> UC3
    User --> UC4
    User --> UC5
    User --> UC6
    User --> UC7
    User --> UC8
    User --> UC9
    User --> UC10
    
    UC1 -.->|include| UC11[验证信息]
    UC1 -.->|include| UC12[加密密码]
    UC2 -.->|include| UC11
    UC2 -.->|include| UC13[生成Token]
    UC4 -.->|include| UC11
    UC6 -.->|include| UC11
    UC6 -.->|include| UC12
    UC7 -.->|include| UC14[发送验证码]
    UC8 -.->|include| UC14
    UC9 -.->|include| UC14
    
    UC11 --> System
    UC12 --> System
    UC13 --> System
    UC14 --> System
```

### 4.3 宠物管理用例图

```mermaid
graph TB
    User((用户))
    Org((机构))
    Admin((管理员))
    
    subgraph "宠物管理"
        UC1[浏览宠物列表]
        UC2[搜索宠物]
        UC3[筛选宠物]
        UC4[查看宠物详情]
        UC5[收藏宠物]
        UC6[分享宠物]
        UC7[发布宠物]
        UC8[编辑宠物]
        UC9[删除宠物]
        UC10[下架宠物]
        UC11[查看我的宠物]
        UC12[审核宠物]
        UC13[查看宠物统计]
    end
    
    User --> UC1
    User --> UC2
    User --> UC3
    User --> UC4
    User --> UC5
    User --> UC6
    User --> UC7
    User --> UC8
    User --> UC9
    User --> UC10
    User --> UC11
    
    Org --> UC1
    Org --> UC2
    Org --> UC3
    Org --> UC4
    Org --> UC7
    Org --> UC8
    Org --> UC9
    Org --> UC10
    Org --> UC11
    
    Admin --> UC12
    Admin --> UC13
    
    UC7 -.->|include| UC14[上传照片]
    UC7 -.->|include| UC15[填写信息]
    UC8 -.->|include| UC14
    UC8 -.->|include| UC15
    UC12 -.->|extend| UC16[通过审核]
    UC12 -.->|extend| UC17[拒绝审核]
```

### 4.4 领养管理用例图

```mermaid
graph TB
    User((用户))
    Org((机构))
    Admin((管理员))
    
    subgraph "领养管理"
        UC1[提交领养申请]
        UC2[查看我的申请]
        UC3[修改申请]
        UC4[取消申请]
        UC5[查看申请详情]
        UC6[查看领养记录]
        UC7[审核申请]
        UC8[安排面试]
        UC9[安排家访]
        UC10[创建领养记录]
        UC11[签署协议]
        UC12[安排交接]
        UC13[制定回访计划]
        UC14[执行回访]
        UC15[查看领养统计]
    end
    
    User --> UC1
    User --> UC2
    User --> UC3
    User --> UC4
    User --> UC5
    User --> UC6
    
    Org --> UC7
    Org --> UC8
    Org --> UC9
    Org --> UC10
    Org --> UC11
    Org --> UC12
    Org --> UC13
    Org --> UC14
    
    Admin --> UC7
    Admin --> UC15
    
    UC1 -.->|include| UC16[填写申请表]
    UC1 -.->|include| UC17[上传证明]
    UC7 -.->|extend| UC18[通过申请]
    UC7 -.->|extend| UC19[拒绝申请]
    UC7 -.->|extend| UC8
    UC8 -.->|extend| UC9
    UC18 -.->|include| UC10
    UC10 -.->|include| UC11
    UC11 -.->|include| UC12
```

### 4.5 社区互动用例图

```mermaid
graph TB
    User((用户))
    Admin((管理员))
    
    subgraph "社区互动"
        UC1[浏览动态列表]
        UC2[搜索动态]
        UC3[查看动态详情]
        UC4[发布动态]
        UC5[编辑动态]
        UC6[删除动态]
        UC7[点赞动态]
        UC8[取消点赞]
        UC9[发表评论]
        UC10[删除评论]
        UC11[回复评论]
        UC12[点赞评论]
        UC13[分享动态]
        UC14[举报动态]
        UC15[管理动态]
        UC16[管理评论]
    end
    
    User --> UC1
    User --> UC2
    User --> UC3
    User --> UC4
    User --> UC5
    User --> UC6
    User --> UC7
    User --> UC8
    User --> UC9
    User --> UC10
    User --> UC11
    User --> UC12
    User --> UC13
    User --> UC14
    
    Admin --> UC15
    Admin --> UC16
    
    UC4 -.->|include| UC17[输入内容]
    UC4 -.->|include| UC18[上传图片]
    UC5 -.->|include| UC17
    UC5 -.->|include| UC18
    UC9 -.->|include| UC17
    UC11 -.->|include| UC17
```

## 5. 类图 (Class Diagram)

### 5.1 核心业务类图

```mermaid
classDiagram
    %% 用户相关
    class User {
        +int64 ID
        +string Username
        +string Password
        +string RealName
        +string Phone
        +string Email
        +string Avatar
        +Gender Gender
        +Date Birthday
        +string IDCard
        +string Address
        +UserRole Role
        +UserStatus Status
        +DateTime LastLoginAt
        +string LastLoginIP
        +DateTime CreatedAt
        +DateTime UpdatedAt
        +IsAdmin() bool
        +IsOrganization() bool
        +IsActive() bool
        +MaskSensitiveInfo() User
    }
    
    class UserRole {
        <<enumeration>>
        USER
        ORGANIZATION
        VOLUNTEER
        ADMIN
    }
    
    class UserStatus {
        <<enumeration>>
        DISABLED
        NORMAL
    }
    
    %% 宠物相关
    class Pet {
        +uint64 ID
        +string Name
        +PetType Type
        +string Breed
        +PetGender Gender
        +int Age
        +PetSize Size
        +string Color
        +float Weight
        +bool IsVaccinated
        +bool IsSterilized
        +string HealthStatus
        +string Description
        +string Character
        +string[] Photos
        +string CoverPhoto
        +Location Location
        +int64 UserID
        +PetStatus Status
        +int ViewCount
        +int FavoriteCount
        +DateTime CreatedAt
        +DateTime UpdatedAt
        +IncrementViewCount()
        +IsAvailable() bool
    }
    
    class PetType {
        <<enumeration>>
        DOG
        CAT
        RABBIT
        BIRD
        OTHER
    }
    
    class PetGender {
        <<enumeration>>
        MALE
        FEMALE
        UNKNOWN
    }
    
    class PetStatus {
        <<enumeration>>
        PENDING
        AVAILABLE
        ADOPTED
        OFFLINE
    }
    
    class PetSize {
        <<enumeration>>
        SMALL
        MEDIUM
        LARGE
    }
    
    class Location {
        +string Province
        +string City
        +string District
        +string Address
        +GetFullAddress() string
    }
    
    %% 领养相关
    class AdoptionApplication {
        +uint64 ID
        +string ApplicationNo
        +int64 UserID
        +uint64 PetID
        +uint64 OrganizationID
        +ApplicantInfo ApplicantInfo
        +HousingInfo HousingInfo
        +FamilyInfo FamilyInfo
        +PetExperience PetExperience
        +string AdoptionReason
        +string HowToCare
        +string EmergencyPlan
        +ApplicationStatus Status
        +ReviewInfo ReviewInfo
        +DateTime CreatedAt
        +DateTime UpdatedAt
        +CanModify() bool
        +CanCancel() bool
    }
    
    class ApplicantInfo {
        +string Name
        +string Phone
        +string IDCard
        +string Address
        +string IDCardImage
        +string HousingProof
    }
    
    class HousingInfo {
        +HousingType Type
        +int Area
        +bool HasYard
    }
    
    class FamilyInfo {
        +int Members
        +bool HasChildren
        +string ChildrenAge
        +bool FamilyAgree
    }
    
    class PetExperience {
        +bool HasExperience
        +string Experience
        +string CurrentPets
    }
    
    class ReviewInfo {
        +int64 ReviewerID
        +string Comment
        +DateTime InterviewTime
        +DateTime HomeVisitTime
        +DateTime ApprovedAt
        +DateTime RejectedAt
        +string RejectionReason
    }
    
    class ApplicationStatus {
        <<enumeration>>
        PENDING
        REVIEWING
        INTERVIEW
        HOME_VISIT
        APPROVED
        REJECTED
        CANCELLED
    }
    
    class HousingType {
        <<enumeration>>
        APARTMENT
        HOUSE
        VILLA
        OTHER
    }
    
    class Adoption {
        +uint64 ID
        +uint64 ApplicationID
        +int64 UserID
        +uint64 PetID
        +uint64 OrganizationID
        +Date AdoptionDate
        +string HandoverLocation
        +string AgreementURL
        +bool AgreementSigned
        +FollowUpPlan FollowUpPlan
        +AdoptionStatus Status
        +string Notes
        +DateTime CreatedAt
        +DateTime UpdatedAt
        +IsActive() bool
    }
    
    class FollowUpPlan {
        +FollowUpSchedule[] Schedules
        +AddSchedule(schedule FollowUpSchedule)
        +GetNextSchedule() FollowUpSchedule
    }
    
    class FollowUpSchedule {
        +Date ScheduledDate
        +string Type
        +string Status
        +string Result
        +DateTime CompletedAt
    }
    
    class AdoptionStatus {
        <<enumeration>>
        ACTIVE
        RETURNED
        DECEASED
    }
    
    %% 机构相关
    class Organization {
        +uint64 ID
        +string Name
        +string Type
        +string Logo
        +string Description
        +Location Location
        +ContactInfo ContactInfo
        +string[] CredentialUrls
        +OrganizationStatus Status
        +string RejectReason
        +uint64 CreatedBy
        +DateTime CreatedAt
        +DateTime UpdatedAt
        +IsApproved() bool
    }
    
    class ContactInfo {
        +string ContactName
        +string ContactPhone
        +string ContactEmail
        +string Phone
        +string Email
    }
    
    class OrganizationStatus {
        <<enumeration>>
        PENDING
        APPROVED
        REJECTED
    }
    
    %% 社区相关
    class Post {
        +uint64 ID
        +int64 UserID
        +string Title
        +string Content
        +string[] Images
        +string Category
        +int ViewCount
        +int LikeCount
        +int CommentCount
        +PostStatus Status
        +DateTime CreatedAt
        +DateTime UpdatedAt
        +IncrementViewCount()
        +IncrementLikeCount()
        +IncrementCommentCount()
    }
    
    class PostStatus {
        <<enumeration>>
        DRAFT
        PUBLISHED
        HIDDEN
    }
    
    class Comment {
        +uint64 ID
        +uint64 PostID
        +int64 UserID
        +uint64 ParentID
        +string Content
        +int LikeCount
        +DateTime CreatedAt
        +DateTime UpdatedAt
        +IsReply() bool
    }
    
    class Like {
        +uint64 ID
        +int64 UserID
        +string TargetType
        +uint64 TargetID
        +DateTime CreatedAt
    }
    
    %% 捐赠相关
    class Donation {
        +uint64 ID
        +string DonationNo
        +int64 UserID
        +uint64 OrganizationID
        +decimal Amount
        +string PaymentMethod
        +DonorInfo DonorInfo
        +string Message
        +bool IsAnonymous
        +DonationStatus Status
        +DateTime PaidAt
        +DateTime ConfirmedAt
        +DateTime CreatedAt
        +DateTime UpdatedAt
        +IsPaid() bool
        +IsConfirmed() bool
    }
    
    class DonorInfo {
        +string Name
        +string Phone
    }
    
    class DonationStatus {
        <<enumeration>>
        PENDING
        PAID
        CONFIRMED
        CANCELLED
    }
    
    %% 关联关系
    User "1" --> "0..*" Pet : publishes
    User "1" --> "0..*" AdoptionApplication : submits
    User "1" --> "0..*" Adoption : adopts
    User "1" --> "0..*" Organization : creates
    User "1" --> "0..*" Post : publishes
    User "1" --> "0..*" Comment : writes
    User "1" --> "0..*" Like : likes
    User "1" --> "0..*" Donation : donates
    
    Pet "1" --> "0..*" AdoptionApplication : receives
    Pet "1" --> "0..1" Adoption : adopted_as
    Pet "*" --> "1" User : belongs_to
    Pet --> Location : has
    Pet --> PetType : has
    Pet --> PetGender : has
    Pet --> PetStatus : has
    Pet --> PetSize : has
    
    AdoptionApplication "1" --> "0..1" Adoption : becomes
    AdoptionApplication --> ApplicantInfo : has
    AdoptionApplication --> HousingInfo : has
    AdoptionApplication --> FamilyInfo : has
    AdoptionApplication --> PetExperience : has
    AdoptionApplication --> ReviewInfo : has
    AdoptionApplication --> ApplicationStatus : has
    
    Adoption --> FollowUpPlan : has
    Adoption --> AdoptionStatus : has
    FollowUpPlan "1" --> "*" FollowUpSchedule : contains
    
    Organization --> Location : has
    Organization --> ContactInfo : has
    Organization --> OrganizationStatus : has
    Organization "1" --> "0..*" Pet : manages
    Organization "1" --> "0..*" AdoptionApplication : reviews
    Organization "1" --> "0..*" Adoption : facilitates
    Organization "1" --> "0..*" Donation : receives
    
    Post "1" --> "0..*" Comment : has
    Post "1" --> "0..*" Like : receives
    Post --> PostStatus : has
    
    Comment "1" --> "0..*" Like : receives
    Comment "0..1" --> "0..*" Comment : replies
    
    Donation --> DonorInfo : has
    Donation --> DonationStatus : has
    
    User --> UserRole : has
    User --> UserStatus : has
```



## 6. 包图 (Package Diagram)

### 6.1 后端系统包图

```mermaid
graph TB
    subgraph "应用层 Application"
        Main[main]
    end
    
    subgraph "配置层 Configuration"
        Config[config]
    end
    
    subgraph "表现层 Presentation"
        Router[router]
        Middleware[middleware]
        Controller[controller]
    end
    
    subgraph "业务层 Business"
        Service[service]
    end
    
    subgraph "持久层 Persistence"
        DAO[dao]
        Model[model]
    end
    
    subgraph "基础设施层 Infrastructure"
        Database[database]
        Cache[cache]
        Logger[logger]
        Utils[utils]
        Response[response]
    end
    
    subgraph "外部服务 External Services"
        MySQL[(MySQL)]
        Redis[(Redis)]
        OSS[OSS]
        Email[Email]
        SMS[SMS]
    end
    
    Main --> Config
    Main --> Router
    Main --> Database
    Main --> Cache
    Main --> Logger
    
    Router --> Middleware
    Router --> Controller
    
    Middleware --> Utils
    Middleware --> Logger
    
    Controller --> Service
    Controller --> Response
    Controller --> Logger
    
    Service --> DAO
    Service --> Cache
    Service --> Utils
    Service --> Logger
    
    DAO --> Model
    DAO --> Database
    DAO --> Logger
    
    Database --> MySQL
    Cache --> Redis
    Utils --> OSS
    Utils --> Email
    Utils --> SMS
    
    style Main fill:#e1f5ff
    style Config fill:#fff4e1
    style Router fill:#e8f5e9
    style Service fill:#f3e5f5
    style DAO fill:#fff3e0
    style Database fill:#ffebee
```

### 6.2 前端系统包图

```mermaid
graph TB
    subgraph "入口层 Entry"
        Main[main]
        App[App]
    end
    
    subgraph "路由层 Router"
        RouterConfig[router]
    end
    
    subgraph "视图层 View"
        Pages[pages]
        Components[components]
    end
    
    subgraph "状态层 State"
        Store[store]
    end
    
    subgraph "服务层 Service"
        API[api]
        AdminAPI[adminApi]
    end
    
    subgraph "工具层 Utility"
        Utils[utils]
    end
    
    subgraph "样式层 Style"
        CSS[styles]
        TailwindCSS[tailwindcss]
    end
    
    subgraph "后端服务 Backend"
        Backend[Backend API]
    end
    
    Main --> App
    Main --> RouterConfig
    Main --> Store
    Main --> CSS
    
    App --> RouterConfig
    App --> Components
    
    RouterConfig --> Pages
    
    Pages --> Components
    Pages --> Store
    Pages --> API
    Pages --> AdminAPI
    Pages --> Utils
    
    Components --> Store
    Components --> Utils
    Components --> TailwindCSS
    
    API --> Backend
    AdminAPI --> Backend
    API --> Utils
    AdminAPI --> Utils
    
    Store --> API
    
    style Main fill:#e1f5ff
    style App fill:#e1f5ff
    style RouterConfig fill:#fff4e1
    style Pages fill:#e8f5e9
    style Store fill:#f3e5f5
    style API fill:#fff3e0
    style Backend fill:#ffebee
```

### 6.3 模块依赖包图

```mermaid
graph LR
    subgraph "用户模块 User Module"
        UserController[UserController]
        UserService[UserService]
        UserDAO[UserDAO]
        UserModel[User Model]
    end
    
    subgraph "宠物模块 Pet Module"
        PetController[PetController]
        PetService[PetService]
        PetDAO[PetDAO]
        PetModel[Pet Model]
    end
    
    subgraph "领养模块 Adoption Module"
        AdoptionController[AdoptionController]
        AdoptionService[AdoptionService]
        AdoptionDAO[AdoptionDAO]
        AdoptionModel[Adoption Model]
    end
    
    subgraph "机构模块 Organization Module"
        OrgController[OrganizationController]
        OrgService[OrganizationService]
        OrgDAO[OrganizationDAO]
        OrgModel[Organization Model]
    end
    
    subgraph "社区模块 Community Module"
        CommunityController[CommunityController]
        CommunityService[CommunityService]
        PostDAO[PostDAO]
        CommentDAO[CommentDAO]
        LikeDAO[LikeDAO]
        PostModel[Post Model]
        CommentModel[Comment Model]
    end
    
    subgraph "捐赠模块 Donation Module"
        DonationController[DonationController]
        DonationService[DonationService]
        DonationDAO[DonationDAO]
        DonationModel[Donation Model]
    end
    
    subgraph "公共模块 Common Module"
        Database[Database]
        Cache[Cache]
        Logger[Logger]
        Utils[Utils]
        Response[Response]
    end
    
    UserController --> UserService
    UserService --> UserDAO
    UserDAO --> UserModel
    UserDAO --> Database
    UserService --> Cache
    
    PetController --> PetService
    PetService --> PetDAO
    PetDAO --> PetModel
    PetDAO --> Database
    PetService --> Cache
    
    AdoptionController --> AdoptionService
    AdoptionService --> AdoptionDAO
    AdoptionDAO --> AdoptionModel
    AdoptionDAO --> Database
    AdoptionService --> PetService
    AdoptionService --> UserService
    
    OrgController --> OrgService
    OrgService --> OrgDAO
    OrgDAO --> OrgModel
    OrgDAO --> Database
    
    CommunityController --> CommunityService
    CommunityService --> PostDAO
    CommunityService --> CommentDAO
    CommunityService --> LikeDAO
    PostDAO --> PostModel
    CommentDAO --> CommentModel
    PostDAO --> Database
    CommentDAO --> Database
    LikeDAO --> Database
    
    DonationController --> DonationService
    DonationService --> DonationDAO
    DonationDAO --> DonationModel
    DonationDAO --> Database
    DonationService --> OrgService
    
    UserController --> Response
    PetController --> Response
    AdoptionController --> Response
    OrgController --> Response
    CommunityController --> Response
    DonationController --> Response
    
    UserService --> Logger
    PetService --> Logger
    AdoptionService --> Logger
    OrgService --> Logger
    CommunityService --> Logger
    DonationService --> Logger
    
    UserService --> Utils
    PetService --> Utils
    AdoptionService --> Utils
```

## 7. ER图 (Entity Relationship Diagram)

### 7.1 完整数据库ER图

```mermaid
erDiagram
    USERS ||--o{ PETS : "publishes"
    USERS ||--o{ ADOPTION_APPLICATIONS : "submits"
    USERS ||--o{ ADOPTIONS : "adopts"
    USERS ||--o{ ORGANIZATIONS : "creates"
    USERS ||--o{ POSTS : "publishes"
    USERS ||--o{ COMMENTS : "writes"
    USERS ||--o{ LIKES : "likes"
    USERS ||--o{ DONATIONS : "donates"
    USERS ||--o{ FAVORITES : "favorites"
    
    PETS ||--o{ ADOPTION_APPLICATIONS : "receives"
    PETS ||--o| ADOPTIONS : "adopted_as"
    PETS }o--|| ORGANIZATIONS : "managed_by"
    PETS ||--o{ FAVORITES : "favorited"
    
    ORGANIZATIONS ||--o{ ADOPTION_APPLICATIONS : "reviews"
    ORGANIZATIONS ||--o{ ADOPTIONS : "facilitates"
    ORGANIZATIONS ||--o{ DONATIONS : "receives"
    ORGANIZATIONS ||--o{ PETS : "manages"
    
    ADOPTION_APPLICATIONS ||--o| ADOPTIONS : "becomes"
    
    POSTS ||--o{ COMMENTS : "has"
    POSTS ||--o{ LIKES : "receives"
    
    COMMENTS ||--o{ COMMENTS : "replies_to"
    COMMENTS ||--o{ LIKES : "receives"
    
    USERS {
        bigint id PK "主键ID"
        varchar username UK "用户名(唯一)"
        varchar password "密码(加密)"
        varchar real_name "真实姓名"
        varchar phone UK "手机号(唯一)"
        varchar email UK "邮箱(唯一)"
        varchar avatar "头像URL"
        tinyint gender "性别(0未知1男2女)"
        date birthday "生日"
        varchar id_card "身份证号"
        text address "地址"
        enum role "角色(user/organization/volunteer/admin)"
        tinyint status "状态(0禁用1正常)"
        timestamp last_login_at "最后登录时间"
        varchar last_login_ip "最后登录IP"
        timestamp created_at "创建时间"
        timestamp updated_at "更新时间"
        timestamp deleted_at "删除时间(软删除)"
    }
    
    PETS {
        bigint id PK "主键ID"
        varchar name "宠物名称"
        varchar type "宠物类型(dog/cat/rabbit/bird/other)"
        varchar breed "品种"
        varchar gender "性别(male/female/unknown)"
        int age "年龄(月)"
        varchar size "体型(small/medium/large)"
        varchar color "毛色"
        decimal weight "体重(kg)"
        boolean is_vaccinated "是否接种疫苗"
        boolean is_sterilized "是否绝育"
        varchar health_status "健康状况"
        text description "详细描述"
        text character "性格特点"
        text photos "照片URL(JSON数组)"
        varchar cover_photo "封面图"
        varchar province "省份"
        varchar city "城市"
        varchar district "区县"
        varchar address "详细地址"
        bigint user_id FK "发布用户ID"
        bigint organization_id FK "所属机构ID"
        tinyint status "状态(0待审核1可领养2已领养3已下架)"
        int view_count "浏览次数"
        int favorite_count "收藏次数"
        bigint adopted_by FK "领养人ID"
        timestamp adopted_at "领养时间"
        bigint reviewed_by FK "审核人ID"
        timestamp reviewed_at "审核时间"
        varchar reject_reason "拒绝原因"
        timestamp created_at "创建时间"
        timestamp updated_at "更新时间"
        timestamp deleted_at "删除时间"
    }
    
    ADOPTION_APPLICATIONS {
        bigint id PK "主键ID"
        varchar application_no UK "申请编号(唯一)"
        bigint user_id FK "申请人ID"
        bigint pet_id FK "宠物ID"
        bigint organization_id FK "机构ID"
        varchar applicant_name "申请人姓名"
        varchar applicant_phone "申请人电话"
        varchar applicant_id_card "身份证号"
        text applicant_address "地址"
        varchar housing_type "住房类型(apartment/house/villa/other)"
        int housing_area "住房面积(平方米)"
        boolean has_yard "是否有院子"
        int family_members "家庭成员数"
        boolean has_children "是否有孩子"
        varchar children_age "孩子年龄"
        boolean family_agree "家人是否同意"
        boolean has_pet_experience "是否有养宠经验"
        text pet_experience "养宠经历"
        text current_pets "现有宠物"
        text adoption_reason "领养原因"
        text how_to_care "如何照顾"
        text emergency_plan "应急计划"
        varchar id_card_image "身份证照片URL"
        varchar housing_proof "住房证明URL"
        text additional_files "其他附件(JSON)"
        tinyint status "状态(0待审核1审核中2待面试3待家访4已通过5已拒绝6已取消)"
        bigint reviewer_id FK "审核人ID"
        text review_comment "审核意见"
        timestamp interview_time "面试时间"
        timestamp home_visit_time "家访时间"
        timestamp approved_at "批准时间"
        timestamp rejected_at "拒绝时间"
        text rejection_reason "拒绝原因"
        timestamp created_at "创建时间"
        timestamp updated_at "更新时间"
        timestamp deleted_at "删除时间"
    }
    
    ADOPTIONS {
        bigint id PK "主键ID"
        bigint application_id FK UK "申请ID(唯一)"
        bigint user_id FK "领养人ID"
        bigint pet_id FK "宠物ID"
        bigint organization_id FK "机构ID"
        date adoption_date "领养日期"
        varchar handover_location "交接地点"
        varchar agreement_url "协议文件URL"
        boolean agreement_signed "协议是否签署"
        text follow_up_plan "回访计划(JSON)"
        varchar status "状态(active/returned/deceased)"
        text notes "备注"
        timestamp created_at "创建时间"
        timestamp updated_at "更新时间"
        timestamp deleted_at "删除时间"
    }
    
    ORGANIZATIONS {
        bigint id PK "主键ID"
        varchar name "机构名称"
        varchar type "机构类型"
        varchar logo "机构Logo"
        text description "机构描述"
        varchar province "省份"
        varchar city "城市"
        varchar address "机构地址"
        varchar contact_name "联系人姓名"
        varchar contact_phone "联系人电话"
        varchar contact_email "联系人邮箱"
        varchar phone "联系电话"
        varchar email "联系邮箱"
        json credential_urls "资历证明图片(JSON数组)"
        tinyint status "状态(0待审核1已通过2已拒绝)"
        varchar reject_reason "拒绝原因"
        bigint created_by FK "创建人ID"
        bigint updated_by FK "更新人ID"
        timestamp created_at "创建时间"
        timestamp updated_at "更新时间"
        timestamp deleted_at "删除时间"
    }
    
    POSTS {
        bigint id PK "主键ID"
        bigint user_id FK "用户ID"
        varchar title "动态标题"
        text content "动态内容"
        text images "图片URL(JSON数组)"
        varchar category "分类"
        int view_count "浏览次数"
        int like_count "点赞数"
        int comment_count "评论数"
        tinyint status "状态(0草稿1已发布2已隐藏)"
        timestamp created_at "创建时间"
        timestamp updated_at "更新时间"
        timestamp deleted_at "删除时间"
    }
    
    COMMENTS {
        bigint id PK "主键ID"
        bigint post_id FK "动态ID"
        bigint user_id FK "用户ID"
        bigint parent_id FK "父评论ID(回复)"
        text content "评论内容"
        int like_count "点赞数"
        timestamp created_at "创建时间"
        timestamp updated_at "更新时间"
        timestamp deleted_at "删除时间"
    }
    
    LIKES {
        bigint id PK "主键ID"
        bigint user_id FK "用户ID"
        varchar target_type "目标类型(post/comment)"
        bigint target_id "目标ID"
        timestamp created_at "创建时间"
    }
    
    DONATIONS {
        bigint id PK "主键ID"
        varchar donation_no UK "捐赠编号(唯一)"
        bigint user_id FK "用户ID"
        bigint organization_id FK "机构ID"
        decimal amount "捐赠金额"
        varchar payment_method "支付方式"
        varchar donor_name "捐赠人姓名"
        varchar donor_phone "捐赠人电话"
        text message "留言"
        boolean is_anonymous "是否匿名"
        varchar status "状态(pending/paid/confirmed/cancelled)"
        timestamp paid_at "支付时间"
        timestamp confirmed_at "确认时间"
        timestamp created_at "创建时间"
        timestamp updated_at "更新时间"
        timestamp deleted_at "删除时间"
    }
    
    FAVORITES {
        bigint id PK "主键ID"
        bigint user_id FK "用户ID"
        bigint pet_id FK "宠物ID"
        timestamp created_at "创建时间"
    }
```

### 7.2 核心实体关系图

```mermaid
erDiagram
    USER ||--o{ PET : "1:N"
    USER ||--o{ APPLICATION : "1:N"
    USER ||--o{ ADOPTION : "1:N"
    PET ||--o{ APPLICATION : "1:N"
    PET ||--o| ADOPTION : "1:0..1"
    APPLICATION ||--o| ADOPTION : "1:0..1"
    ORGANIZATION ||--o{ PET : "1:N"
    ORGANIZATION ||--o{ APPLICATION : "1:N"
    ORGANIZATION ||--o{ ADOPTION : "1:N"
    
    USER {
        id PK
        username
        password
        role
        status
    }
    
    PET {
        id PK
        name
        type
        user_id FK
        organization_id FK
        status
    }
    
    APPLICATION {
        id PK
        user_id FK
        pet_id FK
        organization_id FK
        status
    }
    
    ADOPTION {
        id PK
        application_id FK
        user_id FK
        pet_id FK
        organization_id FK
        status
    }
    
    ORGANIZATION {
        id PK
        name
        status
    }
```

