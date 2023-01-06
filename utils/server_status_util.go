package utils

import (
	"../Adoter_Asset"
	"github.com/garyburd/redigo/redis"
	"github.com/golang/protobuf/proto"
)

func getThresholdKey(zoneId string) string {
	return zoneId + ":serverstatus"
}

func GetServerStatus(redisConn redis.Conn, serverId string) Adoter_Asset.ServerStatus {
	serverStatus := Adoter_Asset.ServerStatus{}

	val, _ := redisConn.Do("GET", getThresholdKey(serverId))
	value, ok := val.([]byte)
	if ok {
		proto.Unmarshal(value, &serverStatus)
	} else {
		serverStatus.Liuchang = 30000
		serverStatus.Huobao = 50000
	}

	return serverStatus
}

func SetServerStatus(redisConn redis.Conn, zoneId string, serverStatus Adoter_Asset.ServerStatus) {
	redisConn.Do("SELECT", 0)

	bytes, _ := proto.Marshal(&serverStatus)
	redisConn.Do("SET", getThresholdKey(zoneId), bytes)
}
