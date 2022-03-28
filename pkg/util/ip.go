package util

//func GetClientIp(c *context.Context) string {
//	ip := c.ClientIP()
//	if ip == "::1" {
//		ip = "127.0.0.1"
//	}
//	return ip
//}
//func IpStringToInt(ipstring string) int {
//	if net.ParseIP(ipstring) == nil {
//		return 0
//	}
//	ipSegs := strings.Split(ipstring, ".")
//	var ipInt int = 0
//	var pos uint = 24
//	for _, ipSeg := range ipSegs {
//		tempInt, _ := strconv.Atoi(ipSeg)
//		tempInt = tempInt << pos
//		ipInt = ipInt | tempInt
//		pos -= 8
//	}
//	return ipInt
//}
