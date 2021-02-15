package main

import (
	"fmt"
	"strings"

	"github.com/skandragon/dysonsphere/types"
)

// VeinData holds information about a specific vein node,
// which has a position, a list of miners, etc.
type VeinData struct {
	ID         int32          `json:"id"`
	Type       types.VeinType `json:"type"`
	ModelIndex int16          `json:"model_index,omitempty"`
	GroupIndex int16          `json:"group_index,omitempty"`
	Amount     int32          `json:"amount"`
	ProductID  int32          `json:"product_id,omitempty"`
	Position   *Vector3       `json:"-"`
	MinerCount int32          `json:"miner_count"`
	MinerID0   int32          `json:"miner_id_0,omitempty"`
	MinerID1   int32          `json:"miner_id_1,omitempty"`
	MinerID2   int32          `json:"miner_id_2,omitempty"`
	MinerID3   int32          `json:"miner_id_3,omitempty"`
}

func (vd *VeinData) isJunk() bool {
	return vd.ID == 0 && vd.Amount == 0
}

func compressVeinDataSlice(l []*VeinData) []*VeinData {
	ll := make([]*VeinData, 0)
	for _, vd := range l {
		if vd != nil && !vd.isJunk() {
			ll = append(ll, vd)
		}
	}
	return ll
}

func (vd *VeinData) String() string {
	r := []string{
		fmt.Sprintf("ID=%d", vd.ID),
		fmt.Sprintf("Type=%s", vd.Type),
		fmt.Sprintf("ModelIndex=%d", vd.ModelIndex),
		fmt.Sprintf("GroupIndex=%d", vd.GroupIndex),
		fmt.Sprintf("Amount=%d", vd.Amount),
		fmt.Sprintf("ProductID=%d", vd.ProductID),
		fmt.Sprintf("Position=%s", vd.Position),
		fmt.Sprintf("MinerCount=%d", vd.MinerCount),
		fmt.Sprintf("MinerID0=%d", vd.MinerID0),
		fmt.Sprintf("MinerID1=%d", vd.MinerID1),
		fmt.Sprintf("MinerID2=%d", vd.MinerID2),
		fmt.Sprintf("MinerID3=%d", vd.MinerID3),
	}

	return "VeinData{" + strings.Join(r, ",") + "}"
}

func parseVeinData(b *Buffer) *VeinData {
	const VERSION = 0
	checkVersByte(b, VERSION, "VeinData")

	return &VeinData{
		ID:         b.GetInt32le(),
		Type:       types.VeinType(b.GetInt16le()),
		ModelIndex: b.GetInt16le(),
		GroupIndex: b.GetInt16le(),
		Amount:     b.GetInt32le(),
		ProductID:  b.GetInt32le(),
		Position: &Vector3{
			X: b.GetFloat32(),
			Y: b.GetFloat32(),
			Z: b.GetFloat32(),
		},
		MinerCount: b.GetInt32le(),
		MinerID0:   b.GetInt32le(),
		MinerID1:   b.GetInt32le(),
		MinerID2:   b.GetInt32le(),
		MinerID3:   b.GetInt32le(),
	}
}
