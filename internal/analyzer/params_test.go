package analyzer

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"github.com/filecoin-project/go-state-types/builtin/v8/miner"
	"reflect"
	"testing"
)

func Test_GetMethod(t *testing.T) {
	params := "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"

	bin, err := base64.StdEncoding.DecodeString(params)
	if err != nil {
		panic(err)
		return
	}
	me := GetMethod(7)
	fmt.Println(me)
	me, ok := GetStruct(7)
	if !ok {
		panic(ok)

	}
	fmt.Println(me)
	if t, ok := me.(miner.ProveCommitSectorParams); ok {
		t.UnmarshalCBOR(bytes.NewReader(bin))
		fmt.Println(t)
		s, _ := json.Marshal(t)
		fmt.Println(string(s))
	} else {
		fmt.Println(reflect.TypeOf(t))
	}

}

type Test struct {
	Name string
	Sex  int
}

func Test_GetMethodHandler(t *testing.T) {
	//structName := "Test"

	s, ok := GetStruct(1)
	if !ok {
		panic(ok)

	}

	fmt.Println(s, reflect.TypeOf(s))

	if t, ok := s.(Test); ok {
		t.Name = "i am test"
		fmt.Println(t, reflect.TypeOf(t))
	}
	fmt.Println(t, reflect.TypeOf(t))

}

func Test_GetMethodInit(t *testing.T) {
	sType := reflect.TypeOf(MethodsMiner)
	for i := 0; i < sType.NumField(); i++ {
		fieldType := sType.Field(i)
		value, _ := sType.FieldByName(fieldType.Name)
		fmt.Printf("registerType(%d, (*miner.%s)(nil))\n", int(value.Index[0]), fieldType.Name+"Params")

	}

}

func Test_ParseParams(t *testing.T) {
	params := "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"

	bin, err := base64.StdEncoding.DecodeString(params)
	if err != nil {
		panic(err)
		return
	}
	as, err := ParseParams(7, bin)
	if err != nil {
		t.Fatal(err)
	}
	t.Logf("-------: %+v", as)
}

func Test_builtinsActor(t *testing.T) {

	t.Logf("%+v", builtinsActor.Methods)
}
