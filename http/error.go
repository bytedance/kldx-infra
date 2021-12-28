package http

const (
	FaaSInfraSuccessCodeSuccess     = "0"
	FaaSInfraFailCodeInternalError  = "k_ec_000001"
	FaaSInfraFailCodeTokenExpire    = "k_ident_013000"
	FaaSInfraFailCodeIllegalToken   = "k_ident_013001"
	FaaSInfraFailCodeMissingToken   = "k_fs_ec_100001"
	FaaSInfraFailCodeRateLimitError = "k_fs_ec_000004"
)

func HasError(errCode string) bool {
	return errCode != FaaSInfraSuccessCodeSuccess
}

func IsSysError(errCode string) bool {
	return errCode == FaaSInfraFailCodeInternalError ||
		errCode == FaaSInfraFailCodeTokenExpire ||
		errCode == FaaSInfraFailCodeIllegalToken ||
		errCode == FaaSInfraFailCodeMissingToken ||
		errCode == FaaSInfraFailCodeRateLimitError
}
