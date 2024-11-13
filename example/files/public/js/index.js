// 树节点数据
var treeObj = null
var setting = {
  data: {
    simpleData: {
      enable: true
    }
  },
  check: {
    enable: true
  }
}
$.post('getTables', res => {
  Object.keys(res.data).forEach(key => {
    if (key !== 'list') {
      $('#' + key).val(res.data[key])
    }
  })
  $.fn.zTree.init($('#tree'), setting, res.data.list)
  treeObj = $.fn.zTree.getZTreeObj('tree')
  treeObj.expandAll(true)
})

// 表单数据
$('.generate-btn').click(function () {
  var checkedData = treeObj.getCheckedNodes().filter(item => {
    return !item.isParent;
  }).map(item => {
    return {
      id: item.id,
      pId: item.pId
    }
  })
  var sendData = {
    tableList: JSON.stringify(checkedData),
    path: $('#path').val(),
    ignoreTablePrefix: $('#ignoreTablePrefix').val(),
    ignoreTableSuffix: $('#ignoreTableSuffix').val(),
    ignoreFieldPrefix: $('#ignoreFieldPrefix').val(),
    ignoreFieldSuffix: $('#ignoreFieldSuffix').val()
  }
  $('.generate-type').find('input[type="checkbox"]').each(function () {
    sendData[$(this).attr('id')] = $(this).prop('checked')
  })
  $.ajax({
    type: 'post',
    url: 'generate',
    data: sendData,
    success: res => {
      if (res.code === 0) {
        window.alert('模板生成成功')
      }
    }
  })
})