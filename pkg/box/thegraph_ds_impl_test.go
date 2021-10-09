package box

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestThegraphResp_dto(t *testing.T) {
	type fields struct {
		Data struct {
			Summoners []struct {
				ID    string `json:"id"`
				Owner string `json:"owner"`
			} `json:"summoners"`
		}
	}
	tests := []struct {
		name    string
		fields  fields
		want    []uint64
		wantErr bool
	}{
		{
			name: "test hex string to uint64",
			fields: fields{
				Data: struct {
					Summoners []struct {
						ID    string `json:"id"`
						Owner string `json:"owner"`
					} `json:"summoners"`
				}{
					Summoners: []struct {
						ID    string `json:"id"`
						Owner string `json:"owner"`
					}{{
						ID:    "0x1f76b9",
						Owner: "0x6d81145629f154dbf07fdab18d2892818626fecf",
					}},
				},
			},
			want:    []uint64{2062009},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := &ThegraphResp{
				Data: tt.fields.Data,
			}
			got, err := r.dto()
			if (err != nil) != tt.wantErr {
				t.Errorf("dto() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("dto() got = %v, want %v", got, tt.want)
			}
		})
	}
}

type testThegraphDao struct {
	idsCheckLength int
}

func (t *testThegraphDao) SaveOrUpdate(address string, tokenIDs []uint64) error {
	if t.idsCheckLength != len(tokenIDs) {
		return fmt.Errorf("want get ids length: %d, actual: %d", t.idsCheckLength, len(tokenIDs))
	}
	return nil
}

func Test_thegraphDataSynchronizer_Sync(t *testing.T) {
	tds := newThegraphDataSynchronizer(&testThegraphDao{309})
	assert.NoError(t, tds.Sync("0x6d81145629f154dbf07fdab18d2892818626fecf"))
}
