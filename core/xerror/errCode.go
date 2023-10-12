package xerror

// 成功返回

const OK int = 200

/**(前3位代表业务,后三位代表具体功能)**/

// 全局错误码

const ServerCommonError int = 100001
const RequestParamError int = 100002
const TokenExpireError int = 100003
const TokenGenerateError int = 100004
const DbError int = 100005
const DbUpdateAffectedZeroError int = 100006

// 用户模块
