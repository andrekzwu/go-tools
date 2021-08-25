package acm

import (
	"errors"
	"fmt"
	"github.com/nacos-group/nacos-sdk-go/clients"
	"github.com/nacos-group/nacos-sdk-go/clients/config_client"
	"github.com/nacos-group/nacos-sdk-go/common/constant"
	"github.com/nacos-group/nacos-sdk-go/vo"
	"sync"
)

var (
	acm *ACMManager
)

func init() {
	acm = New()
}

func New(entry ...*ACMEntry) *ACMManager {
	if entry == nil || len(entry) == 0 {
		return &ACMManager{}
	}
	if err := checkACMEntry(entry[0]); err != nil {
		panic(err)
	}
	return &ACMManager{entry: entry[0]}
}

type ACMEntry struct {
	Endpoint    string
	AccessKey   string
	SecretKey   string
	NamespaceId string // default namespace id
}

type ACMManager struct {
	entry        *ACMEntry
	acmClientMap sync.Map // client
	lock         sync.RWMutex
}

func (mng *ACMManager) String() string {
	mng.lock.RLock()
	defer mng.lock.RUnlock()
	return fmt.Sprintf("acm endpoint: %s", mng.entry.Endpoint)
}

type ACMParam struct {
	dataId      string
	group       string
	namespaceId string
}

func NewACMParam(dataId, group string) *ACMParam {
	return &ACMParam{
		dataId: dataId,
		group:  group,
	}
}

func (param *ACMParam) SetNamespaceId(namespaceId string) *ACMParam {
	param.namespaceId = namespaceId
	return param
}

// set entry
func RegisterACM(entry *ACMEntry) {
	acm.RegisterACM(entry)
}

func (mng *ACMManager) RegisterACM(entry *ACMEntry) {
	if err := checkACMEntry(entry); err != nil {
		panic(err)
	}
	if entry != nil {
		mng.lock.Lock()
		acm.entry = entry
		mng.lock.Unlock()
	}
}

func GetString(param *ACMParam) string {
	return acm.GetString(param)
}

func (mng *ACMManager) GetString(param *ACMParam) string {
	content, err := mng.getClient(param.namespaceId).instance().GetConfig(vo.ConfigParam{DataId: param.dataId, Group: param.group})
	if err != nil {
		panic(err)
	}
	return content
}

func (mng *ACMManager) getClient(namespaceId ...string) *acmClient {
	_namespaceId := mng.getNamespaceId(namespaceId...)
	client, ok := mng.acmClientMap.Load(_namespaceId)
	if !ok {
		return mng.setClient(_namespaceId)
	}
	return client.(*acmClient)
}

func (mng *ACMManager) getNamespaceId(namespaceId ...string) string {
	mng.lock.RLock()
	_namespaceId := mng.entry.NamespaceId
	mng.lock.RUnlock()
	if namespaceId != nil && len(namespaceId) > 0 && namespaceId[0] != "" {
		_namespaceId = namespaceId[0]
	}
	return _namespaceId
}

func (mng *ACMManager) setClient(namespaceId string) *acmClient {
	mng.lock.RLock()
	client := &acmClient{entry: &ACMEntry{
		Endpoint:    mng.entry.Endpoint,
		AccessKey:   mng.entry.AccessKey,
		SecretKey:   mng.entry.SecretKey,
		NamespaceId: namespaceId,
	}}
	mng.lock.RUnlock()
	// save
	mng.acmClientMap.Store(namespaceId, client)
	return client
}

type acmClient struct {
	entry  *ACMEntry
	once   sync.Once
	client config_client.IConfigClient
}

func (client *acmClient) instance() config_client.IConfigClient {
	client.once.Do(func() {
		clientConfig := constant.ClientConfig{
			Endpoint:       client.entry.Endpoint + ":8080",
			NamespaceId:    client.entry.NamespaceId,
			AccessKey:      client.entry.AccessKey,
			SecretKey:      client.entry.SecretKey,
			TimeoutMs:      5 * 1000,
			ListenInterval: 30 * 1000,
		}

		// Initialize client.
		cc, err := clients.CreateConfigClient(map[string]interface{}{
			"clientConfig": clientConfig,
		})

		if err != nil {
			panic(err)
		}

		client.client = cc
	})
	return client.client
}

func checkACMEntry(entry *ACMEntry) error {
	if entry == nil {
		return errors.New("acm config error;entry is nil")
	}
	if entry.Endpoint == "" {
		return errors.New("acm config error;invalid endpoint")
	}
	if entry.AccessKey == "" {
		return errors.New("acm config error;invalid access key")
	}
	if entry.SecretKey == "" {
		return errors.New("acm config error;invalid secret key")
	}
	return nil
}
