package gdconf

import (
	"fmt"
	"os"

	"github.com/Eichs/hkrpg-go/pkg/logger"
	"github.com/hjson/hjson-go/v4"
)

type ItemConfigEquipment struct {
	ID                  uint32              `json:"ID"`
	ItemMainType        string              `json:"ItemMainType"`
	ItemSubType         string              `json:"ItemSubType"`
	InventoryDisplayTag uint32              `json:"InventoryDisplayTag"`
	Rarity              string              `json:"Rarity"`
	IsVisible           bool                `json:"isVisible"`
	PileLimit           uint32              `json:"PileLimit"`
	IsSellable          bool                `json:"IsSellable"`
	ReturnItemIDList    []*ReturnItemIDList `json:"ReturnItemIDList"`
	SellType            string              `json:"SellType"`
}

type ReturnItemIDList struct {
	ItemID  uint32 `json:"ItemID"`
	ItemNum uint32 `json:"ItemNum"`
}

func (g *GameDataConfig) loadItemConfigEquipment() {
	g.ItemConfigEquipmentMap = make(map[string]*ItemConfigEquipment)
	playerElementsFilePath := g.excelPrefix + "ItemConfigEquipment.json"
	playerElementsFile, err := os.ReadFile(playerElementsFilePath)
	if err != nil {
		info := fmt.Sprintf("open file error: %v", err)
		panic(info)
	}

	err = hjson.Unmarshal(playerElementsFile, &g.ItemConfigEquipmentMap)
	if err != nil {
		info := fmt.Sprintf("parse file error: %v", err)
		panic(info)
	}
	logger.Info("load %v ItemConfigEquipment", len(g.ItemConfigEquipmentMap))
}

func GetItemConfigEquipmentById(ID string) *ItemConfigEquipment {
	return CONF.ItemConfigEquipmentMap[ID]
}

func GetItemConfigEquipmentMap() map[string]*ItemConfigEquipment {
	return CONF.ItemConfigEquipmentMap
}
