package server

import (
	"github.com/JobNing/user-rpc/internal/model"
	"github.com/JobNing/user-rpc/pb/type"
)

func modelToPb(in *model.Type) *_type.TypeInfo {
	return &_type.TypeInfo{
		ID:   in.ID,
		Name: in.Name,
		Icon: in.Icon,
		LV:   in.Lv,
		PID:  in.Pid,
	}
}

func GetByLV1ID(id int64) (info *_type.TypeInfo, err error) {
	mod := new(model.Type)

	//获取一级分类信息
	lv1Info, err := mod.GetTypeByID(id)
	if err != nil {
		return nil, err
	}

	//转换一级分类信息
	info = modelToPb(lv1Info)

	//通过一级分类id获取二级分类
	lv2Infos, err := mod.GetTypeListByPID(id)
	if err != nil {
		return nil, err
	}

	//定义二级分类id的切片
	var lv2IDs []int64
	//通过for循环将二级分类id全部
	for _, val := range lv2Infos {
		lv2IDs = append(lv2IDs, val.ID)
	}

	//根据二级分类的切片，查询所有的三级分类
	lv3Infos, err := mod.GetTypeListByPIDs(lv2IDs)
	if err != nil {
		return nil, err
	}

	//将所有的三级分类放到map里，key是这个三级分类的PID,val是一个切片type.TypeInfo
	lv3Map := make(map[int64][]*_type.TypeInfo)
	for _, val := range lv3Infos {
		lv3Map[val.Pid] = append(lv3Map[val.Pid], modelToPb(val))
	}

	//拼接结构
	for _, val := range lv2Infos {
		//转换二级分类的结构
		tInfo := modelToPb(val)
		//通过二级分类的id去map里面取所有对应的三级分类然后赋值
		tInfo.Infos = lv3Map[val.ID]
		//将二级分类拼到一级分类的infos二级分类切片下
		info.Infos = append(info.Infos, tInfo)
	}

	return info, nil
}
