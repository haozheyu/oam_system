/**
 * api公共响应错误码
@Getter
@AllArgsConstructor
@Accessors(chain = true)
public enum BaseResultCode implements IResultCode {

    SUCCESS("200", "操作成功"),

    FAILURE("400", "业务异常"),

    UN_AUTHORIZED("401", "请求未授权"),

    REQ_FORBIDDEN("403", "请求被拒绝"),

    NOT_FOUND("404", "请求未找到"),

    METHOD_NOT_SUPPORTED("405", "不支持当前请求方法"),

    MEDIA_TYPE_NOT_SUPPORTED("415", "不支持当前媒体类型"),

    INTERNAL_SERVER_ERROR("500", "服务器异常"),

    REQ_NOT_SUPPORT("40004", "暂不支持当前请求操作"),

    PARAM_MISS("40010", "缺少必要的请求参数"),

    PARAM_TYPE_ERROR("40011", "请求参数类型错误"),

    PARAM_BIND_ERROR("40012", "请求参数绑定错误"),

    PARAM_VALID_ERROR("40013", "参数校验失败"),

    DATA_EXCEPTION("40014", "数据异常"),

    HANDLE_OBJECT_NOT_EXIST("50000", "操作对象不存在"),

    UNAUTHORIZED_HEADER_IS_EMPTY("00401", "无权访问,请求头【Authorization】为空"),

    INVALID_TOKEN("00402", "无效token"),

    FORBIDDEN("00403", "无权访问，请联系管理员!"),

    GATEWAY_ERROR("00404", "网关异常"),

    GATEWAY_CONNECT_TIME_OUT("00405", "网关超时"),

    BAD_GATEWAY("00406", "错误网关"),

    GATEWAY_NOT_FOUND_SERVICE("00407", "服务未找到"),

    SYSTEM_ERROR("00408", "系统异常"),

    SYSTEM_BUSY("00409", "系统繁忙,请稍候再试"),

    LOGIN_FAIL("00410", "登陆失败"),

    LOGIN_SUCCESS("00411", "登陆成功"),

    LOGOUT_SUCCESS("00412", "退出成功"),

    REFRESH_TOKEN_EXPIRE("00413", "刷新令牌过期"),

    GET_TOKEN_KEY_ERROR("00414", "认证服务器获取Token异常"),

    GEN_PUBLIC_KEY_ERROR("00415", "生成公钥异常"),

    ACCOUNT_NOT_EXIST("00416", "该账号不存在，请联系管理员!"),

    ACCOUNT_DISABLED("00417", "该账号已被禁用，请联系管理员!"),

    ACCOUNT_LOCKED("00418", "该账号已被锁定，请联系管理员!"),

    ACCOUNT_EXPIRED("00419", "该账号已过期，请联系管理员!"),

    CREDENTIALS_EXPIRED("00420", "该账户的登录凭证已过期，请重新登录!"),

    BAD_CREDENTIALS("00421", "该账户的登录凭证是无效的"),

    USERNAME_PASSWORD_ERROR("00422", "账号或密码错误!"),

    AUTHENTICATION_FAIL("00423", "认证失败"),

    UNTRUSTED_CREDENTIALS("00424", "该账户的登录凭证是不可信的"),

    INVALID_CLIENT("00405", "无效的客户端"),

    REQUIRE_ACCESS_TOKEN("00406", "缺少必需的access token"),

    DIGEST_NONCE_EXPIRED("00407", "摘要nonce已过期"),

    UNSUPPORTED_GRANT_TYPE("00408", "不支持的授权类型"),

    MESSAGE_FAILED_TO_SEND("00409","短信发送失败"),

    NO_GOODS("00501", "无任何商品信息"),

    GOODS_OFF_THE_SHELF("00502", "商品已下架");

    code编码
    final String code;
    中文信息描述
    final String msg;
}
 */

export const SUCCESS = 200

export const UN_AUTHORIZED = 401

export const LOADING_STATUS = {
  DEFAULT: 'DEFAULT',
  NET_ERROR: 'NET_ERROR',
  LOADING: 'LOADING',
  FINISH: 'FINISH'
}
