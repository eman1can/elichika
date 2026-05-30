package config

import (
	"encoding/json"
	"log"
	"reflect"

	"elichika/internal/enum"
	"elichika/internal/utils"
)

type RuntimeConfig struct {
	ServerAddress            *string `json:"server_address" of_label:"Server's address"`
	ServerCdn                *string `json:"server_cdn" of_label:"CDN Server to request files from. Defaults to internal w/ through cache."`
	StaticProxyCdn           *string `json:"static_proxy_cdn" of_label:"Upstream CDN to proxy-cache missing static files from (leave empty to disable)"`
	AdminPassword            *string `json:"admin_password" of_label:"Admin password" of_type:"password""`
	TapBondGain              *int32  `json:"tap_bond_gain" of_label:"Partner tap bond reward" of_attrs:"min=\"0\" max=\"20000000\"`
	AutoJudgeType            *int32  `json:"auto_judge_type" of_type:"select" of_options:"None\n1\nMiss\n10\nBad\n12\nGood\n14\nGreat\n20\nPerfect\n30" of_label:"Autoplay judge type"`
	Tutorial                 *bool   `json:"tutorial" of_label:"Enable tutorial"`                                                          // whether to turn on tutorial when starting a new account
	LoginBonusSecond         *int32  `json:"login_bonus_second" of_type:"time" of_label:"Login bonus reset time"`                          // the second from mid-night till login bonus
	TimeZone                 *string `json:"timezone" of_label:"Timezone (from tz database)"`                                              // https://en.wikipedia.org/wiki/List_of_tz_database_time_zones
	DefaultContentAmount     *int32  `json:"default_content_amount" of_label:"Default item count" of_attrs:"min=\"0\" max=\"1000000000\""` // the amount of items to give an user if they don't have that item
	MissionMultiplier        *int32  `json:"mission_multiplier" of_label:"Mission progress multiplier" of_attrs:"min=\"0\" max=\"10000\""` // multiply the progress of missions. Only work for do "x" of things, not for "get x different thing or reach x level"
	ResourceConfigType       *string `json:"resource_config_type" of_type:"select" of_options:"Original\noriginal\nComfortable\ncomfortable\nFree\nfree" of_label:"Resource config"`
	EventAutoSchedulerPeriod *string `json:"event_auto_scheduler_period" of_type:"select" of_options:"once per day\n1_day\nonce per week\n7_days" of_label:"Event frequency"`
	MaintenanceUrl           *string `json:"maintenance_url" of_label:"Maintenance Url"`
}

func defaultConfigs() *RuntimeConfig {
	configs := RuntimeConfig{
		ServerAddress:            new(string),
		ServerCdn:                new(string),
		StaticProxyCdn:           new(string),
		AdminPassword:            new(string),
		TapBondGain:              new(int32),
		AutoJudgeType:            new(int32),
		Tutorial:                 new(bool),
		LoginBonusSecond:         new(int32),
		TimeZone:                 new(string),
		DefaultContentAmount:     new(int32),
		MissionMultiplier:        new(int32),
		ResourceConfigType:       new(string),
		EventAutoSchedulerPeriod: new(string),
		MaintenanceUrl:           new(string),
	}
	// TODO: Move TapBondGain/Multipliers other "soft" settings to table w/ WebUI interaction
	*configs.ServerAddress = "0.0.0.0:8080"
	*configs.ServerCdn = DefaultServerCdn
	*configs.StaticProxyCdn = DefaultProxyCdn
	*configs.AdminPassword = ""
	*configs.TapBondGain = 20
	*configs.AutoJudgeType = enum.JudgeTypeGreat
	*configs.Tutorial = true
	*configs.LoginBonusSecond = enum.HourSecondCount * 4
	*configs.TimeZone = "Asia/Tokyo"
	*configs.DefaultContentAmount = 10000000
	*configs.MissionMultiplier = 1
	*configs.ResourceConfigType = "comfortable"
	*configs.EventAutoSchedulerPeriod = "7_days"
	*configs.MaintenanceUrl = "http://127.0.0.1:8080/webui/admin/"
	return &configs
}

func Load(p string) *RuntimeConfig {
	if !utils.PathExists(p) {
		c := defaultConfigs()
		c.Save(p)
		return c
	}

	c := RuntimeConfig{}
	err := json.Unmarshal([]byte(utils.ReadAllText(p)), &c)
	if err != nil {
		log.Panic("config file is wrong, change/delete it and try again")
	}
	d := defaultConfigs()
	for i := 0; i < reflect.TypeOf(c).NumField(); i++ {
		f := reflect.ValueOf(&c).Elem().Field(i)
		if f.IsNil() {
			f.Set(reflect.ValueOf(d).Elem().Field(i))
		}
	}
	return &c
}

func (c *RuntimeConfig) Save(p string) error {
	data, err := json.Marshal(c)
	utils.CheckErr(err)
	utils.WriteAllText(p, string(data))
	return nil
}

var confs = []*RuntimeConfig{}

func UpdateConfig(newConfig *RuntimeConfig) {
	confs = append(confs, Conf)
	newConfig.Save("./data/config.json") // this has lock so the file can't be corrupted
	// this should be safe because we overwrite the pointer, not the object
	// if someone had an old version of confs then they would just have an access to the old config until they discard it
	// this also assume the pointer size is less than machine word, which it should be
	Conf = newConfig
}

func (r RuntimeConfig) ResourceConfig() *ResourceConfig {
	return resourceConfigs[*r.ResourceConfigType]
}
