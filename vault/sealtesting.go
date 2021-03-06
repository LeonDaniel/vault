package vault

import (
	"bytes"
	"fmt"
)

type TestSeal struct {
	defseal        *DefaultSeal
	barrierKeys    [][]byte
	recoveryKey    []byte
	recoveryConfig *SealConfig
}

func (d *TestSeal) checkCore() error {
	if d.defseal.core == nil {
		return fmt.Errorf("seal does not have a core set")
	}
	return nil
}

func (d *TestSeal) SetCore(core *Core) {
	d.defseal = &DefaultSeal{}
	d.defseal.core = core
}

func (d *TestSeal) Init() error {
	d.barrierKeys = [][]byte{}
	return d.defseal.Init()
}

func (d *TestSeal) Finalize() error {
	return d.defseal.Finalize()
}

func (d *TestSeal) BarrierType() string {
	return "shamir"
}

func (d *TestSeal) StoredKeysSupported() bool {
	return true
}

func (d *TestSeal) RecoveryKeySupported() bool {
	return true
}

func (d *TestSeal) SetStoredKeys(keys [][]byte) error {
	d.barrierKeys = keys
	return nil
}

func (d *TestSeal) GetStoredKeys() ([][]byte, error) {
	return d.barrierKeys, nil
}

func (d *TestSeal) BarrierConfig() (*SealConfig, error) {
	return d.defseal.BarrierConfig()
}

func (d *TestSeal) SetBarrierConfig(config *SealConfig) error {
	return d.defseal.SetBarrierConfig(config)
}

func (d *TestSeal) RecoveryType() string {
	return "shamir"
}

func (d *TestSeal) RecoveryConfig() (*SealConfig, error) {
	return d.recoveryConfig, nil
}

func (d *TestSeal) SetRecoveryConfig(config *SealConfig) error {
	d.recoveryConfig = config
	return nil
}

func (d *TestSeal) VerifyRecoveryKey(key []byte) error {
	if bytes.Equal(d.recoveryKey, key) {
		return nil
	}
	return fmt.Errorf("not equivalent")
}

func (d *TestSeal) SetRecoveryKey(key []byte) error {
	newbuf := bytes.NewBuffer(nil)
	newbuf.Write(key)
	d.recoveryKey = newbuf.Bytes()
	return nil
}
