package main

import (
	"fmt"
	"strings"
)

type EntityData struct {
	Version       int32           `json:"version"`
	ID            int32           `json:"id"`
	ProtoID       int16           `json:"proto_id,omitempty"`
	ModelIndex    int16           `json:"model_index,omitempty"`
	Position      *PositionSingle `json:"position,omitempty"`
	Rotation      *RotationSingle `json:"rotation,omitempty"`
	BeltID        int32           `json:"belt_id,omitempty"`
	SplitterID    int32           `json:"splitter_id,omitempty"`
	StorageID     int32           `json:"storage_id,omitempty"`
	TankID        int32           `json:"tank_id,omitempty"`
	MinerID       int32           `json:"miner_id,omitempty"`
	InserterID    int32           `json:"inserter_id,omitempty"`
	AssemblerID   int32           `json:"assembler_id,omitempty"`
	FractionateID int32           `json:"fractionate_id,omitempty"`
	EjectorID     int32           `json:"ejector_id,omitempty"`
	SiloID        int32           `json:"silo_id,omitempty"`
	LabID         int32           `json:"lab_id,omitempty"`
	StationID     int32           `json:"station_id,omitempty"`
	PowerNodeID   int32           `json:"power_node_id,omitempty"`
	PowerGenID    int32           `json:"power_gen_id,omitempty"`
	PowerConID    int32           `json:"power_con_id,omitempty"`
	PowerAccID    int32           `json:"power_acc_id,omitempty"`
	PowerExcID    int32           `json:"power_exc_id,omitempty"`
	MonsterID     int32           `json:"monster_id,omitempty"`
}

func (e *EntityData) String() string {
	items := []string{
		fmt.Sprintf("ID=%d", e.ID),
		fmt.Sprintf("ProtoID=%d", e.ProtoID),
		fmt.Sprintf("ModelIndex=%d", e.ModelIndex),
		e.Position.String(),
		e.Rotation.String(),
	}
	if e.BeltID > 0 {
		items = append(items, fmt.Sprintf("BeltID=%d", e.BeltID))
	}
	if e.SplitterID > 0 {
		items = append(items, fmt.Sprintf("SplitterID=%d", e.SplitterID))
	}
	if e.StorageID > 0 {
		items = append(items, fmt.Sprintf("StorageID=%d", e.StorageID))
	}
	if e.TankID > 0 {
		items = append(items, fmt.Sprintf("TankID=%d", e.TankID))
	}
	if e.MinerID > 0 {
		items = append(items, fmt.Sprintf("MinerID=%d", e.MinerID))
	}
	if e.InserterID > 0 {
		items = append(items, fmt.Sprintf("InserterID=%d", e.InserterID))
	}
	if e.AssemblerID > 0 {
		items = append(items, fmt.Sprintf("AssemblerID=%d", e.AssemblerID))
	}
	if e.FractionateID > 0 {
		items = append(items, fmt.Sprintf("FractionateID=%d", e.FractionateID))
	}
	if e.EjectorID > 0 {
		items = append(items, fmt.Sprintf("EjectorID=%d", e.EjectorID))
	}
	if e.SiloID > 0 {
		items = append(items, fmt.Sprintf("SiloID=%d", e.SiloID))
	}
	if e.LabID > 0 {
		items = append(items, fmt.Sprintf("LabID=%d", e.LabID))
	}
	if e.StationID > 0 {
		items = append(items, fmt.Sprintf("StationID=%d", e.StationID))
	}
	if e.PowerNodeID > 0 {
		items = append(items, fmt.Sprintf("PowerNodeID=%d", e.PowerNodeID))
	}
	if e.PowerGenID > 0 {
		items = append(items, fmt.Sprintf("PowerGenID=%d", e.PowerGenID))
	}
	if e.PowerConID > 0 {
		items = append(items, fmt.Sprintf("PowerConID=%d", e.PowerConID))
	}
	if e.PowerAccID > 0 {
		items = append(items, fmt.Sprintf("PowerAccID=%d", e.PowerAccID))
	}
	if e.PowerExcID > 0 {
		items = append(items, fmt.Sprintf("PowerExcID=%d", e.PowerExcID))
	}
	if e.MonsterID > 0 {
		items = append(items, fmt.Sprintf("MonsterID=%d", e.MonsterID))
	}

	return "EntityData{" + strings.Join(items, ",") + "}"
}

func parseEntityData(b *Buffer) *EntityData {
	const VERSION = 0
	checkVersByte(b, VERSION, "EntityData")

	e := &EntityData{
		Version:    VERSION,
		ID:         b.GetInt32le(),
		ProtoID:    b.GetInt16le(),
		ModelIndex: b.GetInt16le(),
		Position: &PositionSingle{
			X: b.GetFloat32(),
			Y: b.GetFloat32(),
			Z: b.GetFloat32(),
		},
		Rotation: &RotationSingle{
			X: b.GetFloat32(),
			Y: b.GetFloat32(),
			Z: b.GetFloat32(),
			W: b.GetFloat32(),
		},
		BeltID:        b.GetInt32le(),
		SplitterID:    b.GetInt32le(),
		StorageID:     b.GetInt32le(),
		TankID:        b.GetInt32le(),
		MinerID:       b.GetInt32le(),
		InserterID:    b.GetInt32le(),
		AssemblerID:   b.GetInt32le(),
		FractionateID: b.GetInt32le(),
		EjectorID:     b.GetInt32le(),
		SiloID:        b.GetInt32le(),
		LabID:         b.GetInt32le(),
		StationID:     b.GetInt32le(),
		PowerNodeID:   b.GetInt32le(),
		PowerGenID:    b.GetInt32le(),
		PowerConID:    b.GetInt32le(),
		PowerAccID:    b.GetInt32le(),
		PowerExcID:    b.GetInt32le(),
		MonsterID:     b.GetInt32le(),
	}

	//fmt.Printf("Entity: %v\n", e)
	return e
}
