package domain

// 現金入出金機制御:回収開始要求
type RequestCollectStart struct {
	RequestInfo RequestInfo       `json:"requestInfo"`
	CollectMode int               `json:"collectMode"` //回収種別
	CountTbl    [CASH_TYPE_UI]int `json:"countTbl"`    //枚数情報
	Amount      int               `json:"amount"`      //金額情報
}

type ResultCollectStart struct {
	RequestInfo   RequestInfo `json:"requestInfo"`
	Result        bool        `json:"result"`                //処理結果
	ErrorCode     string      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail   string      `json:"errorDetail,omitempty"` //エラー詳細
	CashControlId string      `json:"cashControlId"`         //入出金機制御管理番号
}

// 現金入出金機制御:回収停止要求
type RequestCollectStop struct {
	RequestInfo   RequestInfo `json:"requestInfo"`
	CashControlId string      `json:"cashControlId"`
}

type ResultCollectStop struct {
	RequestInfo RequestInfo `json:"requestInfo"`
	Result      bool        `json:"result"`                //処理結果
	ErrorCode   string      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail string      `json:"errorDetail,omitempty"` //エラー詳細
}

// 現金入出金機制御:入金終了要求情報
type RequestInEnd struct {
	RequestInfo   RequestInfo `json:"requestInfo"`
	CashControlId string      `json:"cashControlId"` //入出金機制御管理番号
	TargetDevice  int         `json:"targetDevice"`  //対象デバイス
	StatusMode    int         `json:"statusMode"`    //動作モード
}

type ResultInEnd struct {
	RequestInfo   RequestInfo `json:"requestInfo"`
	Result        bool        `json:"result"`                //処理結果
	ErrorCode     string      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail   string      `json:"errorDetail,omitempty"` //エラー詳細
	CashControlId string      `json:"cashControlId"`         //入出金機制御管理番号
}

// 現金入出金機制御:入金開始要求情報
type RequestInStart struct {
	RequestInfo   RequestInfo `json:"requestInfo"`
	ModeOperation int         `json:"modeOperation"` //運用モード
	CountClear    bool        `json:"countClear"`    //入金枚数クリア
	TargetDevice  int         `json:"targetDevice"`  //対象デバイス
}

type ResultInStart struct {
	RequestInfo   RequestInfo `json:"requestInfo"`
	Result        bool        `json:"result"`                //処理結果
	ErrorCode     string      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail   string      `json:"errorDetail,omitempty"` //エラー詳細
	CashControlId string      `json:"cashControlId"`         //入出金機制御管理番号
}

// 現金入出金制御:有高ステータス通知
type AmountStatus struct {
	CoinStatusCode int                         `json:"coinStatusCode"`        //硬貨結果通知コード
	BillStatusCode int                         `json:"billStatusCode"`        //紙幣結果通知コード
	Amount         int                         `json:"amount"`                //金額
	CountTbl       [CASH_TYPE_SHITEI]int       `json:"countTbl"`              //通常金種別枚数
	ExCountTbl     [EXTRA_CASH_TYPE_SHITEI]int `json:"exCountTbl"`            //拡張金種別枚数
	DepositTbl     [CASH_TYPE_SHITEI]int       `json:"depositTbl"`            //入金可能枚数
	ErrorCode      string                      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail    string                      `json:"errorDetail,omitempty"` //エラー詳細
}

// 現金入出金制御:回収ステータス通知
type CollectStatus struct {
	CashControlId  string                      `json:"cashControlId"`         //入出金機制御管理番号
	CoinStatusCode int                         `json:"coinStatusCode"`        //硬貨結果通知コード
	BillStatusCode int                         `json:"billStatusCode"`        //紙幣結果通知コード
	Amount         int                         `json:"amount"`                //金額
	CountTbl       [CASH_TYPE_SHITEI]int       `json:"countTbl"`              //通常金種別枚数
	ExCountTbl     [EXTRA_CASH_TYPE_SHITEI]int `json:"exCountTbl"`            //拡張金種別枚数
	ErrorCode      string                      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail    string                      `json:"errorDetail,omitempty"` //エラー詳細
}

// 現金入出金制御:入金ステータス通知
type InStatus struct {
	CashControlId  string                      `json:"cashControlId"`         //入出金機制御管理番号
	CoinStatusCode int                         `json:"coinStatusCode"`        //硬貨結果通知コード
	BillStatusCode int                         `json:"billStatusCode"`        //紙幣結果通知コード
	Amount         int                         `json:"amount"`                //金額
	CountTbl       [CASH_TYPE_SHITEI]int       `json:"countTbl"`              //通常金種別枚数
	ExCountTbl     [EXTRA_CASH_TYPE_SHITEI]int `json:"exCountTbl"`            //拡張金種別枚数
	ErrorCode      string                      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail    string                      `json:"errorDetail,omitempty"` //エラー詳細
}

// 現金入出金制御:出金ステータス通知
type OutStatus struct {
	CashControlId  string                      `json:"cashControlId"`         //入出金機制御管理番号
	CoinStatusCode int                         `json:"coinStatusCode"`        //硬貨結果通知コード
	BillStatusCode int                         `json:"billStatusCode"`        //紙幣結果通知コード
	Amount         int                         `json:"amount"`                //金額
	CountTbl       [CASH_TYPE_SHITEI]int       `json:"countTbl"`              //通常金種別枚数
	ExCountTbl     [EXTRA_CASH_TYPE_SHITEI]int `json:"exCountTbl"`            //拡張金種別枚数
	ErrorCode      string                      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail    string                      `json:"errorDetail,omitempty"` //エラー詳細
}

// 現金入出金制御：入出金機ステータス通知
type NoticeStatus struct {
	CoinStatusCode           int              `json:"coinStatusCode"`      //硬貨結果通知コード
	BillStatusCode           int              `json:"billStatusCode"`      //紙幣結果通知コード
	CoinNoticeStatusTbl      NoticeStatusTbl  `json:"coinStatusTbl"`       //硬貨ステータス情報
	NoticeCoinResidueInfoTbl []ResidueInfoTbl `json:"coinResidueInfoTbl"`  //硬貨残留情報
	BillNoticeStatusTbl      NoticeStatusTbl  `json:"billStatusTbl"`       //紙幣ステータス情報
	BillNoticeResidueInfoTbl []ResidueInfoTbl `json:"billResidueInfoTbl"`  //紙幣残留情報
	DeviceStatusInfoTbl      []string         `json:"deviceStatusInfoTbl"` //デバイス詳細情報
	WarningInfoTbl           []int            `json:"warningInfoTbl"`      //警告情報
}

// ステータス情報
type NoticeStatusTbl struct {
	StatusError       bool   `json:"statusError"`           //処理結果
	ErrorCode         string `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail       string `json:"errorDetail,omitempty"` //エラー詳細
	StatusCover       bool   `json:"statusCover"`           //トビラ状態
	StatusUnitSet     bool   `json:"statusUnitSet"`         //ユニットセット状態
	StatusInCassette  bool   `json:"statusInCassette"`      //補充カセット状態
	StatusOutCassette bool   `json:"statusOutCassette"`     //回収カセット状態
	StatusInsert      bool   `json:"statusInsert"`          //入金口状態
	StatusExit        bool   `json:"statusExit"`            //出金口状態
	StatusRjbox       bool   `json:"statusRjbox"`           //リジェクトBOX状態
}

// 現金入出金機制御:入出金機ステータス取得要求
type RequestStatus struct {
	RequestInfo RequestInfo `json:"requestInfo"`
}

type ResultStatus struct {
	RequestInfo         RequestInfo      `json:"requestInfo"`
	Result              bool             `json:"result"`                //処理結果
	ErrorCode           string           `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail         string           `json:"errorDetail,omitempty"` //エラー詳細
	CoinStatusCode      int              `json:"coinStatusCode"`        //硬貨結果通知コード
	BillStatusCode      int              `json:"billStatusCode"`        //紙幣結果通知コード
	CoinStatusTbl       StatusTbl        `json:"coinStatusTbl"`         //硬貨ステータス情報
	CoinResidueInfoTbl  []ResidueInfoTbl `json:"coinResidueInfoTbl"`    //硬貨残留情報
	BillStatusTbl       StatusTbl        `json:"billStatusTbl"`         //紙幣ステータス情報
	BillResidueInfoTbl  []ResidueInfoTbl `json:"billResidueInfoTbl"`    //紙幣残留情報
	DeviceStatusInfoTbl []string         `json:"deviceStatusInfoTbl"`   //デバイス詳細情報
	WarningInfoTbl      []int            `json:"warningInfoTbl"`        //警告情報
}

// ステータス情報
type StatusTbl struct {
	StatusCover       bool `json:"statusCover"`       //トビラ状態
	StatusUnitSet     bool `json:"statusUnitSet"`     //ユニットセット状態
	StatusInCassette  bool `json:"statusInCassette"`  //補充カセット状態
	StatusOutCassette bool `json:"statusOutCassette"` //回収カセット状態
	StatusInsert      bool `json:"statusInsert"`      //入金口状態
	StatusExit        bool `json:"statusExit"`        //出金口状態
	StatusRjbox       bool `json:"statusRjbox"`       //リジェクトBOX状態
}

// 残留情報
type ResidueInfoTbl struct {
	Title  string `json:"title"`  //管理名称
	Status bool   `json:"status"` //状態
}

// 現金入出金機制御:出金開始要求
type RequestOutStart struct {
	RequestInfo        RequestInfo `json:"requestInfo"`
	StatusOutRejectBox bool        `json:"statusOutRejectBox"` //出金種別
	OutMode            int         `json:"outMode"`            //出金種別
	OutStatusCashInfoTbl
}

type ResultOutStart struct {
	RequestInfo   RequestInfo `json:"requestInfo"`
	Result        bool        `json:"result"`                //処理結果
	ErrorCode     string      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail   string      `json:"errorDetail,omitempty"` //エラー詳細
	CashControlId string      `json:"cashControlId"`         //入出金機制御管理番号
}

type OutStatusCashInfoTbl struct {
	Amount   int               `json:"amount"`   //金種
	CountTbl [CASH_TYPE_UI]int `json:"countTbl"` //金種別枚数
}

// 現金入出金機制御:出金停止要求
type RequestOutStop struct {
	RequestInfo   RequestInfo `json:"requestInfo"`
	CashControlId string      `json:"cashControlId"`
}

type ResultOutStop struct {
	RequestInfo RequestInfo `json:"requestInfo"`
	Result      bool        `json:"result"`                //処理結果
	ErrorCode   string      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail string      `json:"errorDetail,omitempty"` //エラー詳細
}

// 現金入出金機制御:有高ステータス取得要求
type RequestAmountStatus struct {
	RequestInfo RequestInfo `json:"requestInfo"`
}

type ResultAmountStatus struct {
	RequestInfo    RequestInfo                 `json:"requestInfo"`
	Result         bool                        `json:"result"`                //処理結果
	CoinStatusCode int                         `json:"coinStatusCode"`        //硬貨結果通知コード
	BillStatusCode int                         `json:"billStatusCode"`        //紙幣結果通知コード
	Amount         int                         `json:"amount"`                //金額
	CountTbl       [CASH_TYPE_SHITEI]int       `json:"countTbl"`              //通常金種別枚数
	ExCountTbl     [EXTRA_CASH_TYPE_SHITEI]int `json:"exCountTbl"`            //拡張金種別枚数
	DepositTbl     [CASH_TYPE_SHITEI]int       `json:"depositTbl"`            //入金可能枚数
	ErrorCode      string                      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail    string                      `json:"errorDetail,omitempty"` //エラー詳細
}

// 現金入出金機制御:回収ステータス取得要求
type RequestCollectStatus struct {
	RequestInfo RequestInfo `json:"requestInfo"`
}

type ResultCollectStatus struct {
	RequestInfo      RequestInfo         `json:"requestInfo"`
	Result           bool                `json:"result"`                //処理結果
	ErrorCode        string              `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail      string              `json:"errorDetail,omitempty"` //エラー詳細
	CoinStatusAction bool                `json:"coinStatusAction"`      //硬貨動作状況
	CoinStatusCode   int                 `json:"coinStatusCode"`        //硬貨結果通知コード
	BillStatusAction bool                `json:"billStatusAction"`      //紙幣動作状況
	BillStatusCode   int                 `json:"billStatusCode"`        //紙幣結果通知コード
	CollectCountKin  int                 `json:"collectCountKin"`       //回収金額
	CashTbl          []CollectStatusCash `json:"cashTbl"`               //回収枚数
}

type CollectStatusCash struct {
	CashType  string `json:"cashType"`  //金種
	CashCount int    `json:"cashCount"` //枚数
}

// 現金入出金機制御:入金ステータス取得要求
type RequestRequestInStatus struct {
	RequestInfo RequestInfo `json:"requestInfo"`
}

type ResultRequestInStatus struct {
	RequestInfo RequestInfo `json:"requestInfo"`
	Result      bool        `json:"result"`                //処理結果
	ErrorCode   string      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail string      `json:"errorDetail,omitempty"` //エラー詳細
}

// 現金入出金機制御:出金ステータス取得要求
type RequestOutStatus struct {
	RequestInfo RequestInfo `json:"requestInfo"`
}

type ResultOutStatus struct {
	RequestInfo      RequestInfo     `json:"requestInfo"`
	Result           bool            `json:"result"`                //処理結果
	ErrorCode        string          `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail      string          `json:"errorDetail,omitempty"` //エラー詳細
	CoinStatusAction bool            `json:"coinStatusAction"`      //硬貨動作状況
	CoinStatusCode   int             `json:"coinStatusCode"`        //硬貨結果通知コード
	BillStatusAction bool            `json:"billStatusAction"`      //紙幣動作状況
	BillStatusCode   int             `json:"billStatusCode"`        //紙幣結果通知コード
	OutCountKin      int             `json:"outCountKin"`           //出金金額
	CashTbl          []OutStatusCash `json:"cashTbl"`               //出金枚数
}

type OutStatusCash struct {
	CashType  string `json:"cashType"`  //金種
	CashCount int    `json:"cashCount"` //枚数
}

// 現金入出金機制御:有高枚数変更要求
type RequestCashctlSetAmount struct {
	RequestInfo   RequestInfo                 `json:"requestInfo"`
	OperationMode int                         `json:"operationMode"` //操作モード
	Amount        int                         `json:"amount"`        //金額
	CountTbl      [CASH_TYPE_SHITEI]int       `json:"countTbl"`      //通常金種別枚数
	ExCountTbl    [EXTRA_CASH_TYPE_SHITEI]int `json:"exCountTbl"`    //拡張金種別枚数
}

type ResultCashctlSetAmount struct {
	RequestInfo   RequestInfo `json:"requestInfo"`
	CashControlId string      `json:"cashControlId"`
	Result        bool        `json:"result"`                //処理結果
	ErrorCode     string      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail   string      `json:"errorDetail,omitempty"` //エラー詳細
}

type RequestScrutinyStart struct {
	RequestInfo  RequestInfo `json:"requestInfo"`
	TargetDevice int         `json:"targetDevice"`
}

type ResultScrutinyStart struct {
	RequestInfo   RequestInfo `json:"requestInfo"`
	Result        bool        `json:"result"`                //処理結果
	ErrorCode     string      `json:"errorCode,omitempty"`   //エラーコード
	ErrorDetail   string      `json:"errorDetail,omitempty"` //エラー詳細
	CashControlId string      `json:"cashControlId"`
}
