/*
 * @Author: huangzhijie
 * @Description:
 * @File: params
 * @Version: 1.0.0
 * @Date: 2022/10/11 11:40
 */
package analyzer

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-state-types/abi"
	"github.com/filecoin-project/go-state-types/actors"
	"github.com/filecoin-project/go-state-types/builtin/v8/miner"
	"github.com/filecoin-project/go-state-types/network"
	"github.com/filecoin-project/lotus/build"
	"github.com/filecoin-project/lotus/chain/consensus/filcns"
	"github.com/filecoin-project/lotus/chain/stmgr"
	"github.com/filecoin-project/lotus/chain/vm"
	"github.com/filecoin-project/lotus/conformance/chaos"
	"github.com/ipfs/go-cid"
	cbg "github.com/whyrusleeping/cbor-gen"
	"golang.org/x/xerrors"
	"reflect"
)

var NetworkBundle = "calibrationnet"

var MethodsMiner = struct {
	Constructor              abi.MethodNum
	ControlAddresses         abi.MethodNum
	ChangeWorkerAddress      abi.MethodNum
	ChangePeerID             abi.MethodNum
	SubmitWindowedPoSt       abi.MethodNum
	PreCommitSector          abi.MethodNum
	ProveCommitSector        abi.MethodNum
	ExtendSectorExpiration   abi.MethodNum
	TerminateSectors         abi.MethodNum
	DeclareFaults            abi.MethodNum
	DeclareFaultsRecovered   abi.MethodNum
	OnDeferredCronEvent      abi.MethodNum
	CheckSectorProven        abi.MethodNum
	ApplyRewards             abi.MethodNum
	ReportConsensusFault     abi.MethodNum
	WithdrawBalance          abi.MethodNum
	ConfirmSectorProofsValid abi.MethodNum
	ChangeMultiaddrs         abi.MethodNum
	CompactPartitions        abi.MethodNum
	CompactSectorNumbers     abi.MethodNum
	ConfirmUpdateWorkerKey   abi.MethodNum
	RepayDebt                abi.MethodNum
	ChangeOwnerAddress       abi.MethodNum
	DisputeWindowedPoSt      abi.MethodNum
	PreCommitSectorBatch     abi.MethodNum
	ProveCommitAggregate     abi.MethodNum
	ProveReplicaUpdates      abi.MethodNum
}{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20, 21, 22, 23, 24, 25, 26, 27}

func GetMethod(me abi.MethodNum) interface{} {

	sType := reflect.TypeOf(MethodsMiner)
	for i := 0; i < sType.NumField(); i++ {
		fieldType := sType.Field(i)
		value, _ := sType.FieldByName(fieldType.Name)

		if int(value.Index[0]) == int(me-1) {
			return fieldType.Name
		}

	}
	return nil

}
func GetMethonHandler(method string) interface{} {
	// ProveCommitSector
	// SubmitWindowedPoSt
	var a miner.ProveCommitSectorParams
	//var a  miner.SubmitWindowedPoStParams{}

	fmt.Println(a)
	return a
	//miner.

}

type MethodMeta struct {
	Name string

	Params reflect.Type
	Ret    reflect.Type
}

var methodsRegistry = make(map[abi.MethodNum]reflect.Type)
var invoker *vm.ActorRegistry

func init() {
	av, _ := actors.VersionForNetwork(network.Version16)

	build.UseNetworkBundle(NetworkBundle)
	invoker = filcns.NewActorRegistry()

	invoker.Register(av, nil, chaos.Actor{})

}

func registerType(number int, elem interface{}) {
	methodsRegistry[abi.MethodNum(number)] = reflect.TypeOf(elem).Elem()
}

func GetStruct(code cid.Cid, number int) (interface{}, bool) {
	elem, ok := invoker.Methods[code][abi.MethodNum(number)]
	if !ok {
		return nil, false
	}
	return reflect.New(elem.Params.Elem()).Interface(), true
}

func JsonParams(code cid.Cid, method abi.MethodNum, params []byte) (string, error) {
	p, err := stmgr.GetParamType(filcns.NewActorRegistry(), code, method) // todo use api for correct actor registry
	if err != nil {
		return "", err
	}
	if err := p.UnmarshalCBOR(bytes.NewReader(params)); err != nil {
		return "", err
	}

	b, err := json.MarshalIndent(p, "", "  ")
	return string(b), err
}

func ParseParamsV2(cid cid.Cid, method abi.MethodNum, params []byte) (string, error) {

	pstr, err := JsonParams(cid, abi.MethodNum(method), params)
	if err != nil {
		return "", err
	}
	return pstr, nil

}

func ParseParamsV1(cid cid.Cid, method int, params []byte) (string, error) {
	if bytes.Equal(params, []byte(nil)) {
		return "", nil
	}
	me, ok := GetStruct(cid, method)
	if !ok {
		return "", xerrors.Errorf("ParseParams GetStruct error!")

	}
	meheadler := me.(cbg.CBORUnmarshaler)

	err := (meheadler).UnmarshalCBOR(bytes.NewReader(params))
	if err != nil {
		return "", xerrors.Errorf("ParseParams UnmarshalCBOR error!")
	}
	s, err := json.Marshal(meheadler)
	if err != nil {
		return "", xerrors.Errorf("ParseParams Marshal error!")
	}
	//fmt.Println(string(s))
	return string(s), nil
}
