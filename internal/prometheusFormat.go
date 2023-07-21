/**
 * @Author Jinxiao Zhang · Az · JokerDevops
 * @Description //TODO
 * @Date 20:35 2023/4/15
 **/

package internal

type Labels struct {
	Alertname    string `json:"alertname"`
	Instance     string `json:"instance"`
	Level        string `json:"level"` //2019年11月20日 16:03:10更改告警级别定义位置,适配prometheus alertmanager rule
	Severity     string `json:"severity"`
	BusinessType string `json:"businessType"`
	DomainId     int64  `json:"domainId"`
	DomainName   string `json:"domainName"`
	SendType     string `json:"sendType"`
	Job          string `json:"job"`
	Hostgroup    string `json:"hostgroup,omitempty"`
	Hostname     string `json:"hostname,omitempty"`
}
type Annotations struct {
	Description string `json:"description"`
	Summary     string `json:"summary"`
	Mobile      string `json:"mobile"`  //2019年2月25日 19:09:23 增加手机号支持
	Ddurl       string `json:"ddurl"`   //2019年3月12日 20:33:38 增加多个钉钉告警支持
	Wxurl       string `json:"wxurl"`   //2019年3月12日 20:33:38 增加多个钉钉告警支持
	Fsurl       string `json:"fsurl"`   //2020年4月25日 17:33:38 增加多个飞书告警支持
	Email       string `json:"email"`   //2020年7月4日 10:15:20 增加多个email告警支持
	Groupid     string `json:"groupid"` //2021年2月2日 17:28:23 增加多个如流告警支持
	AtSomeOne   string `json:"at"`      //2021年6月23日 14:02:21 增加@某人支持
	Rr          string `json:"rr"`      //2021年9月14日 14:48:08 增加随机轮询参数支持
}
type Alerts struct {
	Status       string      `json:"status"`
	Labels       Labels      `json:"labels"`
	Annotations  Annotations `json:"annotations"`
	StartsAt     string      `json:"startsAt"`
	EndsAt       string      `json:"endsAt"`
	GeneratorUrl string      `json:"generatorURL"` //prometheus 告警返回地址
}
type Prometheus struct {
	Status      string
	Alerts      []Alerts
	Externalurl string `json:"externalURL"` //alertmanage 返回地址
}

// 按照 Alert.Level 从大到小排序
type AlerMessages []Alerts

func (a AlerMessages) Len() int { // 重写 Len() 方法
	return len(a)
}
func (a AlerMessages) Swap(i, j int) { // 重写 Swap() 方法
	a[i], a[j] = a[j], a[i]
}
func (a AlerMessages) Less(i, j int) bool { // 重写 Less() 方法， 从大到小排序
	return a[j].Labels.Level < a[i].Labels.Level
}
