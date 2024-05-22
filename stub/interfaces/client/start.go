package interfaces

import (
	"encoding/json"
	"stub/domain"
	"stub/domain/handler"
)

// 起動シナリオのテスト
type start struct {
	mqtt   handler.MqttRepository
	logger handler.LoggerRepository
}

func NewStart(mqtt handler.MqttRepository,
	logger handler.LoggerRepository) Start {
	return &start{
		mqtt:   mqtt,
		logger: logger}
}

func (s *start) Senario1() {
	//tex/unifunc/money/notice_status_cashを受信する
	s.mqtt.Subscribe("/tex/unifunc/money/notice_status_cash", s.RecvNoticeStatusCash)
}

// 通知のチェック
func (c *start) RecvNoticeStatusCash(message string) {
	var statusCash domain.StatusCash

	err := json.Unmarshal([]byte(message), &statusCash)
	if err != nil {
		c.logger.Error("json unmarshal error: %v", err)
		return
	}
	status := 1
	var testStatusCash = domain.StatusCash{}
	switch status {
	case 1:
		/*テストデータ
		[INFO ] 2024/04/23 11:14:40.874798 [Send](MQTT)/tex/unifunc/money/notice_status_cash,
		  {"cashControlId":"","statusReady":true,"statusMode":1,"statusLine":true,"statusError":true,
		  "statusCover":false,"statusAction":0,"statusInsert":false,"statusExit":false,"statusRjbox":false,
		  "billStatusTbl":{"statusUnitSet":false,"statusInCassette":false,"statusOutCassette":false},
		  "coinStatusTbl":{"statusUnitSet":false,"statusInCassette":false,"statusOutCassette":false},
		  "billResidueInfoTbl":[],"coinResidueInfoTbl":[],"deviceStatusInfoTbl":["","","","","",""],
		  "warningInfoTbl":[0,0,0,0,0,0,0,0,0,0]}*/
		testStatusCash = domain.StatusCash{
			CashControlId: "",
			StatusReady:   true,
			StatusMode:    1,
			StatusLine:    true,
			StatusError:   true,
			StatusCover:   false,
			StatusAction:  0,
			StatusInsert:  false,
			StatusExit:    false,
			StatusRjbox:   false,
			BillStatusTbl: domain.TexmyBillStatusTbl{
				StatusUnitSet:     false,
				StatusInCassette:  false,
				StatusOutCassette: false,
				StatusAmountCount: 0,
			},
			CoinStatusTbl: domain.CoinStatusTbl{
				StatusUnitSet:     false,
				StatusInCassette:  false,
				StatusOutCassette: false,
				StatusAmountCount: 0,
			},
			BillResidueInfoTbl:  []domain.BillResidueInfo{},
			CoinResidueInfoTbl:  []domain.CoinResidueInfo{},
			DeviceStatusInfoTbl: []string{"", "", "", "", "", ""},
			WarningInfoTbl:      []int{0, 0, 0, 0, 0, 0, 0, 0, 0, 0},
		}

	}

	// 比較 //TODO: isEqualStatusCash関数の実装
	if !isEqualStatusCash(testStatusCash, statusCash) {
		c.logger.Error("Senalio1_RecvNoticeStatus is not equal: %v", statusCash)
	}
}
