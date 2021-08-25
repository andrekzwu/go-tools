package context

// InPackage std in param
type InPackage struct {
	ClientInfo interface{}     `json:"ClientInfo"`
	Data       interface{}     `json:"data"`
}

// OutPackage std out param
type OutPackage struct {
	Status int32       `json:"status"`
	Msg    string      `json:"msg"`
	Data   interface{} `json:"data"`
}

// InPackage new and method
// NewInPackage create a InPackage
func NewInPackage(clientInfo, data interface{}) *InPackage {
	return &InPackage{
		ClientInfo: clientInfo,
		Data:       data,
	}
}

func NewSimpleInPackage(data interface{}) *InPackage {
	return &InPackage{
		Data: data,
	}
}


func (in *InPackage) GetData() interface{} {
	if in != nil && in.Data != nil {
		return in.Data
	}
	return nil
}

func (in *InPackage) GetClientInfo() interface{}{
	if in != nil && in.ClientInfo != nil {
		return in.ClientInfo
	}
	return nil
}

// OutPackage new and method
// NewOutPackage create a OutPackage
func NewOutPackage(data interface{}) *OutPackage {
	out := &OutPackage{
		Status: 0,
		Msg:    "success",
		Data:   data}
	return out
}

// GetStatus get the OutPackage.Status
func (out *OutPackage) GetStatus() int32 {
	if out != nil {
		return out.Status
	}
	return -1
}

// GetData get the OutPackage.Data
func (out *OutPackage) GetData() interface{} {
	if out != nil && out.Data != nil {
		return out.Data
	}
	return nil
}

// SetStatus set the OutPackage.Status
func (out *OutPackage) SetStatus(status int32) {
	if out != nil {
		out.Status = status
	}
	return
}

// SetMsg set the OutPackage.Msg
func (out *OutPackage) SetMsg(msg string) {
	if out != nil {
		out.Msg = msg
	}
	return
}

// SetData set the OutPackage.Data
func (out *OutPackage) SetData(data interface{}) {
	if out != nil && data != nil {
		out.Data = data
	}
	return
}

// SetStatusMsg set the staus and msg
func (out *OutPackage) SetStatusMsg(outStd *OutPackage) {
	if out != nil && outStd != nil {
		out.Status, out.Msg = outStd.Status, outStd.Msg
	}
	return
}
