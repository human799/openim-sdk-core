package msgtest

import (
	"fmt"

	"github.com/openimsdk/openim-sdk-core/v3/sdk_struct"
	"github.com/openimsdk/protocol/constant"
)

// config here

// system
var (
	TESTIP        = "wg8686.com"
	APIADDR       = fmt.Sprintf("http://%v/api", TESTIP)
	WSADDR        = fmt.Sprintf("ws://%v/msg_gateway", TESTIP)
	SECRET        = "openIM123"
	MANAGERUSERID = "openIMAdmin"

	PLATFORMID = constant.WindowsPlatformID
	LogLevel   = uint32(6)
)

func GetConfig() *sdk_struct.IMConfig {
	var cf sdk_struct.IMConfig
	cf.ApiAddr = APIADDR
	cf.PlatformID = int32(PLATFORMID)
	cf.WsAddr = WSADDR
	cf.DataDir = "./"
	cf.LogLevel = LogLevel
	cf.IsExternalExtensions = true
	cf.IsLogStandardOutput = true
	cf.LogFilePath = ""
	return &cf

}
