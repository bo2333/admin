let loadId

function dateFormat(value) {
    return value ? new Date(value*1000).format("yyyy-MM-dd hh:mm:ss") : "";
}

function handleData(data, tips) {
    if (data.code === 0) {
        if (tips.length > 0) {
            layer.msg(tips + "：成功!", {icon: 6, shade: 0.3, time: 1000});
        }
        return true;
    }
    layer.msg(tips + "（失败）:" + data.msg, {icon: 5, shade: 0.3, time: 2000});
    console.log(data.msg)
    return false;
}

function loadingTips(state) {
    if (state) {
        loadId = layer.load(0, {
            shade: false,
            shadeClose: true, //开启遮罩关闭
        });
    } else {
        layer.close(loadId)
    }

}
