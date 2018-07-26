// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// NativeStatsItemsStats native stats items stats
// swagger:model nativeStatsItemsStats
type NativeStatsItemsStats struct {

	// act
	Act *int64 `json:"act,omitempty"`

	// addr
	Addr string `json:"addr,omitempty"`

	// agent code
	AgentCode *int64 `json:"agent_code,omitempty"`

	// agent desc
	AgentDesc string `json:"agent_desc,omitempty"`

	// agent duration
	AgentDuration *int64 `json:"agent_duration,omitempty"`

	// agent fall
	AgentFall *int64 `json:"agent_fall,omitempty"`

	// agent health
	AgentHealth *int64 `json:"agent_health,omitempty"`

	// agent rise
	AgentRise *int64 `json:"agent_rise,omitempty"`

	// agent status
	// Enum: [UNK INI SOCKERR L40K L4TOUT L4CON L7OK L7STS]
	AgentStatus string `json:"agent_status,omitempty"`

	// algo
	Algo string `json:"algo,omitempty"`

	// bck
	Bck *int64 `json:"bck,omitempty"`

	// bin
	Bin *int64 `json:"bin,omitempty"`

	// bout
	Bout *int64 `json:"bout,omitempty"`

	// check code
	CheckCode *int64 `json:"check_code,omitempty"`

	// check desc
	CheckDesc string `json:"check_desc,omitempty"`

	// check duration
	CheckDuration *int64 `json:"check_duration,omitempty"`

	// check fall
	CheckFall *int64 `json:"check_fall,omitempty"`

	// check health
	CheckHealth *int64 `json:"check_health,omitempty"`

	// check rise
	CheckRise *int64 `json:"check_rise,omitempty"`

	// check status
	// Enum: [UNK INI SOCKERR L40K L4TOUT L4CON L6OK L6TOUT L6RSP L7OK L7OKC L7TOUT L7RSP L7STS]
	CheckStatus string `json:"check_status,omitempty"`

	// chkdown
	Chkdown *int64 `json:"chkdown,omitempty"`

	// chkfail
	Chkfail *int64 `json:"chkfail,omitempty"`

	// cli abrt
	CliAbrt *int64 `json:"cli_abrt,omitempty"`

	// comp byp
	CompByp *int64 `json:"comp_byp,omitempty"`

	// comp in
	CompIn *int64 `json:"comp_in,omitempty"`

	// comp out
	CompOut *int64 `json:"comp_out,omitempty"`

	// comp rsp
	CompRsp *int64 `json:"comp_rsp,omitempty"`

	// conn rate
	ConnRate *int64 `json:"conn_rate,omitempty"`

	// conn rate max
	ConnRateMax *int64 `json:"conn_rate_max,omitempty"`

	// conn tot
	ConnTot *int64 `json:"conn_tot,omitempty"`

	// cookie
	Cookie string `json:"cookie,omitempty"`

	// ctime
	Ctime *int64 `json:"ctime,omitempty"`

	// dcon
	Dcon *int64 `json:"dcon,omitempty"`

	// downtime
	Downtime *int64 `json:"downtime,omitempty"`

	// dreq
	Dreq *int64 `json:"dreq,omitempty"`

	// dresp
	Dresp *int64 `json:"dresp,omitempty"`

	// dses
	Dses *int64 `json:"dses,omitempty"`

	// econ
	Econ *int64 `json:"econ,omitempty"`

	// ereq
	Ereq *int64 `json:"ereq,omitempty"`

	// eresp
	Eresp *int64 `json:"eresp,omitempty"`

	// hanafail
	Hanafail string `json:"hanafail,omitempty"`

	// hrsp 1xx
	Hrsp1xx *int64 `json:"hrsp_1xx,omitempty"`

	// hrsp 2xx
	Hrsp2xx *int64 `json:"hrsp_2xx,omitempty"`

	// hrsp 3xx
	Hrsp3xx *int64 `json:"hrsp_3xx,omitempty"`

	// hrsp 4xx
	Hrsp4xx *int64 `json:"hrsp_4xx,omitempty"`

	// hrsp 5xx
	Hrsp5xx *int64 `json:"hrsp_5xx,omitempty"`

	// hrsp other
	HrspOther *int64 `json:"hrsp_other,omitempty"`

	// iid
	Iid *int64 `json:"iid,omitempty"`

	// intercepted
	Intercepted *int64 `json:"intercepted,omitempty"`

	// lastchg
	Lastchg *int64 `json:"lastchg,omitempty"`

	// lastsess
	Lastsess *int64 `json:"lastsess,omitempty"`

	// lbtot
	Lbtot *int64 `json:"lbtot,omitempty"`

	// mode
	// Enum: [tcp http health unknown]
	Mode string `json:"mode,omitempty"`

	// pid
	Pid *int64 `json:"pid,omitempty"`

	// qcur
	Qcur *int64 `json:"qcur,omitempty"`

	// qlimit
	Qlimit *int64 `json:"qlimit,omitempty"`

	// qmax
	Qmax *int64 `json:"qmax,omitempty"`

	// qtime
	Qtime *int64 `json:"qtime,omitempty"`

	// rate
	Rate *int64 `json:"rate,omitempty"`

	// rate lim
	RateLim *int64 `json:"rate_lim,omitempty"`

	// rate max
	RateMax *int64 `json:"rate_max,omitempty"`

	// req rate
	ReqRate *int64 `json:"req_rate,omitempty"`

	// req rate max
	ReqRateMax *int64 `json:"req_rate_max,omitempty"`

	// req tot
	ReqTot *int64 `json:"req_tot,omitempty"`

	// rtime
	Rtime *int64 `json:"rtime,omitempty"`

	// scur
	Scur *int64 `json:"scur,omitempty"`

	// sid
	Sid *int64 `json:"sid,omitempty"`

	// slim
	Slim *int64 `json:"slim,omitempty"`

	// smax
	Smax *int64 `json:"smax,omitempty"`

	// srv abrt
	SrvAbrt *int64 `json:"srv_abrt,omitempty"`

	// status
	// Enum: [UP DOWN NOLB MAINT no check]
	Status string `json:"status,omitempty"`

	// stot
	Stot *int64 `json:"stot,omitempty"`

	// throttle
	Throttle *int64 `json:"throttle,omitempty"`

	// tracked
	Tracked *int64 `json:"tracked,omitempty"`

	// ttime
	Ttime *int64 `json:"ttime,omitempty"`

	// weight
	Weight *int64 `json:"weight,omitempty"`

	// wredis
	Wredis *int64 `json:"wredis,omitempty"`

	// wretr
	Wretr *int64 `json:"wretr,omitempty"`
}

// Validate validates this native stats items stats
func (m *NativeStatsItemsStats) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAgentStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCheckStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateMode(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var nativeStatsItemsStatsTypeAgentStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNK","INI","SOCKERR","L40K","L4TOUT","L4CON","L7OK","L7STS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatsItemsStatsTypeAgentStatusPropEnum = append(nativeStatsItemsStatsTypeAgentStatusPropEnum, v)
	}
}

const (

	// NativeStatsItemsStatsAgentStatusUNK captures enum value "UNK"
	NativeStatsItemsStatsAgentStatusUNK string = "UNK"

	// NativeStatsItemsStatsAgentStatusINI captures enum value "INI"
	NativeStatsItemsStatsAgentStatusINI string = "INI"

	// NativeStatsItemsStatsAgentStatusSOCKERR captures enum value "SOCKERR"
	NativeStatsItemsStatsAgentStatusSOCKERR string = "SOCKERR"

	// NativeStatsItemsStatsAgentStatusL40K captures enum value "L40K"
	NativeStatsItemsStatsAgentStatusL40K string = "L40K"

	// NativeStatsItemsStatsAgentStatusL4TOUT captures enum value "L4TOUT"
	NativeStatsItemsStatsAgentStatusL4TOUT string = "L4TOUT"

	// NativeStatsItemsStatsAgentStatusL4CON captures enum value "L4CON"
	NativeStatsItemsStatsAgentStatusL4CON string = "L4CON"

	// NativeStatsItemsStatsAgentStatusL7OK captures enum value "L7OK"
	NativeStatsItemsStatsAgentStatusL7OK string = "L7OK"

	// NativeStatsItemsStatsAgentStatusL7STS captures enum value "L7STS"
	NativeStatsItemsStatsAgentStatusL7STS string = "L7STS"
)

// prop value enum
func (m *NativeStatsItemsStats) validateAgentStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nativeStatsItemsStatsTypeAgentStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NativeStatsItemsStats) validateAgentStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.AgentStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateAgentStatusEnum("agent_status", "body", m.AgentStatus); err != nil {
		return err
	}

	return nil
}

var nativeStatsItemsStatsTypeCheckStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UNK","INI","SOCKERR","L40K","L4TOUT","L4CON","L6OK","L6TOUT","L6RSP","L7OK","L7OKC","L7TOUT","L7RSP","L7STS"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatsItemsStatsTypeCheckStatusPropEnum = append(nativeStatsItemsStatsTypeCheckStatusPropEnum, v)
	}
}

const (

	// NativeStatsItemsStatsCheckStatusUNK captures enum value "UNK"
	NativeStatsItemsStatsCheckStatusUNK string = "UNK"

	// NativeStatsItemsStatsCheckStatusINI captures enum value "INI"
	NativeStatsItemsStatsCheckStatusINI string = "INI"

	// NativeStatsItemsStatsCheckStatusSOCKERR captures enum value "SOCKERR"
	NativeStatsItemsStatsCheckStatusSOCKERR string = "SOCKERR"

	// NativeStatsItemsStatsCheckStatusL40K captures enum value "L40K"
	NativeStatsItemsStatsCheckStatusL40K string = "L40K"

	// NativeStatsItemsStatsCheckStatusL4TOUT captures enum value "L4TOUT"
	NativeStatsItemsStatsCheckStatusL4TOUT string = "L4TOUT"

	// NativeStatsItemsStatsCheckStatusL4CON captures enum value "L4CON"
	NativeStatsItemsStatsCheckStatusL4CON string = "L4CON"

	// NativeStatsItemsStatsCheckStatusL6OK captures enum value "L6OK"
	NativeStatsItemsStatsCheckStatusL6OK string = "L6OK"

	// NativeStatsItemsStatsCheckStatusL6TOUT captures enum value "L6TOUT"
	NativeStatsItemsStatsCheckStatusL6TOUT string = "L6TOUT"

	// NativeStatsItemsStatsCheckStatusL6RSP captures enum value "L6RSP"
	NativeStatsItemsStatsCheckStatusL6RSP string = "L6RSP"

	// NativeStatsItemsStatsCheckStatusL7OK captures enum value "L7OK"
	NativeStatsItemsStatsCheckStatusL7OK string = "L7OK"

	// NativeStatsItemsStatsCheckStatusL7OKC captures enum value "L7OKC"
	NativeStatsItemsStatsCheckStatusL7OKC string = "L7OKC"

	// NativeStatsItemsStatsCheckStatusL7TOUT captures enum value "L7TOUT"
	NativeStatsItemsStatsCheckStatusL7TOUT string = "L7TOUT"

	// NativeStatsItemsStatsCheckStatusL7RSP captures enum value "L7RSP"
	NativeStatsItemsStatsCheckStatusL7RSP string = "L7RSP"

	// NativeStatsItemsStatsCheckStatusL7STS captures enum value "L7STS"
	NativeStatsItemsStatsCheckStatusL7STS string = "L7STS"
)

// prop value enum
func (m *NativeStatsItemsStats) validateCheckStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nativeStatsItemsStatsTypeCheckStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NativeStatsItemsStats) validateCheckStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.CheckStatus) { // not required
		return nil
	}

	// value enum
	if err := m.validateCheckStatusEnum("check_status", "body", m.CheckStatus); err != nil {
		return err
	}

	return nil
}

var nativeStatsItemsStatsTypeModePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["tcp","http","health","unknown"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatsItemsStatsTypeModePropEnum = append(nativeStatsItemsStatsTypeModePropEnum, v)
	}
}

const (

	// NativeStatsItemsStatsModeTCP captures enum value "tcp"
	NativeStatsItemsStatsModeTCP string = "tcp"

	// NativeStatsItemsStatsModeHTTP captures enum value "http"
	NativeStatsItemsStatsModeHTTP string = "http"

	// NativeStatsItemsStatsModeHealth captures enum value "health"
	NativeStatsItemsStatsModeHealth string = "health"

	// NativeStatsItemsStatsModeUnknown captures enum value "unknown"
	NativeStatsItemsStatsModeUnknown string = "unknown"
)

// prop value enum
func (m *NativeStatsItemsStats) validateModeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nativeStatsItemsStatsTypeModePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NativeStatsItemsStats) validateMode(formats strfmt.Registry) error {

	if swag.IsZero(m.Mode) { // not required
		return nil
	}

	// value enum
	if err := m.validateModeEnum("mode", "body", m.Mode); err != nil {
		return err
	}

	return nil
}

var nativeStatsItemsStatsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["UP","DOWN","NOLB","MAINT","no check"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		nativeStatsItemsStatsTypeStatusPropEnum = append(nativeStatsItemsStatsTypeStatusPropEnum, v)
	}
}

const (

	// NativeStatsItemsStatsStatusUP captures enum value "UP"
	NativeStatsItemsStatsStatusUP string = "UP"

	// NativeStatsItemsStatsStatusDOWN captures enum value "DOWN"
	NativeStatsItemsStatsStatusDOWN string = "DOWN"

	// NativeStatsItemsStatsStatusNOLB captures enum value "NOLB"
	NativeStatsItemsStatsStatusNOLB string = "NOLB"

	// NativeStatsItemsStatsStatusMAINT captures enum value "MAINT"
	NativeStatsItemsStatsStatusMAINT string = "MAINT"

	// NativeStatsItemsStatsStatusNoCheck captures enum value "no check"
	NativeStatsItemsStatsStatusNoCheck string = "no check"
)

// prop value enum
func (m *NativeStatsItemsStats) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, nativeStatsItemsStatsTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *NativeStatsItemsStats) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *NativeStatsItemsStats) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *NativeStatsItemsStats) UnmarshalBinary(b []byte) error {
	var res NativeStatsItemsStats
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
