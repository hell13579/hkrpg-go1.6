package Game

import (
	"strconv"

	"github.com/Eichs/hkrpg-go/gdconf"
	"github.com/Eichs/hkrpg-go/protocol/cmd"
	"github.com/Eichs/hkrpg-go/protocol/proto"
)

func (g *Game) ExpUpEquipmentCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.ExpUpEquipmentCsReq, payloadMsg)
	req := msg.(*proto.ExpUpEquipmentCsReq)
	if req.EquipmentUniqueId == 0 {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}

	var equipmentList []uint32 // 需要删除的equipmentList
	var pileItem []*Material   // 需要删除的升级材料
	var delScoin uint32        // 扣除的信用点
	var addExp uint32          // 增加的经验

	// 从背包获取需要升级的光锥
	dbEquipment := g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId]
	if dbEquipment == nil {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}
	// 获取需要升级光锥的配置信息
	equConf := gdconf.GetEquipmentConfigById(strconv.Itoa(int(dbEquipment.Tid)))
	if equConf == nil {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.send(cmd.ExpUpEquipmentScRsp, rsp)
		return
	}

	// 遍历用来升级的材料
	for _, pileList := range req.ItemCostList.ItemList {
		// 如果没有则退出
		if pileList.GetPileItem() == nil {
			break
		}
		pile := new(Material)
		pile.Tid = pileList.GetPileItem().ItemId
		pile.Num = pileList.GetPileItem().ItemNum

		pileItem = append(pileItem, pile)
		// 获取材料配置
		pileconf := gdconf.GetEquipmentConfigById(strconv.Itoa(int(pileList.GetPileItem().ItemId)))
		if pileconf == nil {
			rsp := &proto.ExpUpEquipmentScRsp{}
			g.send(cmd.ExpUpEquipmentScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += pileconf.CoinCost * pileList.GetPileItem().ItemNum
		// 获取能添加多少经验
		addExp += pileconf.ExpProvide * pileList.GetPileItem().ItemNum
	}

	// 遍历用来升级的光锥
	for _, equipment := range req.ItemCostList.ItemList {
		// 如果没有则退出
		if equipment.GetEquipmentUniqueId() == 0 {
			break
		}
		equipmentList = append(equipmentList, equipment.GetEquipmentUniqueId())
		// 获取光锥配置
		equipmentconfig := gdconf.GetEquipmentConfigById(strconv.Itoa(int(g.Player.DbItem.EquipmentMap[equipment.GetEquipmentUniqueId()].Tid)))
		if equipmentconfig == nil {
			rsp := &proto.ExpUpEquipmentScRsp{}
			g.send(cmd.ExpUpEquipmentScRsp, rsp)
			return
		}
		// 获取要扣多少信用点
		delScoin += equipmentconfig.CoinCost
		// 获取能添加多少经验
		addExp += equipmentconfig.ExpProvide
	}

	// 计算添加后有多少经验
	exp := addExp + dbEquipment.Exp

	// 获取能升级到的等级和升级后经验
	level, exp := gdconf.GetEquipmentExpByLevel(equConf.ExpType, exp, dbEquipment.Level, dbEquipment.Promotion)
	if level == 0 && exp == 0 {
		rsp := &proto.ExpUpEquipmentScRsp{}
		g.send(cmd.ExpUpEquipmentScRsp, rsp)
	}

	// 扣除本次升级需要的信用点
	g.Player.DbItem.MaterialMap[2].Num -= delScoin
	// 更新需要升级的光锥状态
	g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].Level = level
	g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].Exp = exp

	// 删除用来升级的材料
	if pileItem != nil {
		g.DelMaterialPlayerSyncScNotify(pileItem)
	}
	if len(equipmentList) != 0 {
		// 删除用来升级的光锥
		g.DelEquipmentPlayerSyncScNotify(equipmentList)
	}
	// 通知角色还有多少信用点
	g.PlayerPlayerSyncScNotify()
	// 通知升级后光锥消息
	g.EquipmentPlayerSyncScNotify(dbEquipment.Tid, req.EquipmentUniqueId)
	rsp := &proto.ExpUpEquipmentScRsp{}
	g.send(cmd.ExpUpEquipmentScRsp, rsp)
}

func (g *Game) PlayerPlayerSyncScNotify() {
	notify := &proto.PlayerSyncScNotify{
		BasicInfo: &proto.PlayerBasicInfo{
			Nickname:   g.Player.NickName,
			Level:      g.Player.Level,
			Exp:        g.Player.Exp,
			Stamina:    g.Player.Stamina,
			Mcoin:      0,
			Hcoin:      g.Player.DbItem.MaterialMap[1].Num,
			Scoin:      g.Player.DbItem.MaterialMap[2].Num,
			WorldLevel: g.Player.WorldLevel,
		},
	}

	g.send(cmd.PlayerSyncScNotify, notify)
}

func (g *Game) DelEquipmentPlayerSyncScNotify(equipmentList []uint32) {
	for _, equipment := range equipmentList {
		delete(g.Player.DbItem.EquipmentMap, equipment)
	}

	notify := &proto.PlayerSyncScNotify{DelEquipmentList: equipmentList}
	g.send(cmd.PlayerSyncScNotify, notify)
}

func (g *Game) DelMaterialPlayerSyncScNotify(pileItem []*Material) {
	notify := &proto.PlayerSyncScNotify{MaterialList: make([]*proto.Material, 0)}

	for _, item := range pileItem {
		g.Player.DbItem.MaterialMap[item.Tid].Num -= item.Num
		material := &proto.Material{
			Tid: item.Tid,
			Num: g.Player.DbItem.MaterialMap[item.Tid].Num,
		}
		notify.MaterialList = append(notify.MaterialList, material)
	}
	g.send(cmd.PlayerSyncScNotify, notify)
}

func (g *Game) RankUpEquipmentCsReq(payloadMsg []byte) {
	msg := g.decodePayloadToProto(cmd.RankUpEquipmentCsReq, payloadMsg)
	req := msg.(*proto.RankUpEquipmentCsReq)

	var equipmentList []uint32 // 需要删除的equipmentList

	// 从背包获取需要叠影的光锥
	dbEquipment := g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId]
	if dbEquipment == nil {
		rsp := new(proto.GetChallengeScRsp)
		g.send(cmd.RankUpEquipmentScRsp, rsp)
		return
	}
	// 遍历用来叠影的光锥
	for _, equipment := range req.ItemCostList.ItemList {
		// 如果没有则退出
		if equipment.GetEquipmentUniqueId() == 0 {
			break
		}
		if g.Player.DbItem.EquipmentMap[equipment.GetEquipmentUniqueId()].Tid != dbEquipment.Tid {
			rsp := new(proto.GetChallengeScRsp)
			g.send(cmd.RankUpEquipmentScRsp, rsp)
			return
		}
		equipmentList = append(equipmentList, equipment.GetEquipmentUniqueId())
		g.Player.DbItem.EquipmentMap[req.EquipmentUniqueId].Rank++
	}

	if len(equipmentList) != 0 {
		// 删除用来叠影的光锥
		g.DelEquipmentPlayerSyncScNotify(equipmentList)
	}
	// 通知叠影后光锥消息
	g.EquipmentPlayerSyncScNotify(dbEquipment.Tid, req.EquipmentUniqueId)

	rsp := new(proto.GetChallengeScRsp)
	g.send(cmd.RankUpEquipmentScRsp, rsp)
}
