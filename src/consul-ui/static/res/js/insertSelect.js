layui.define(['jquery', 'form'], function (exports) {
    var modelName = 'insertSelect',
        $ = layui.jquery,
        form = layui.form;


    /**
     * 初始化select控件且定义事件处理
     * @param {any} opt
     */
    var InsertSelect =
    {
        render: function (opt) {
            var elem = $(opt.elem);//找到控件id
            var data = opt.data;//获取对应的数据
            var tempInput = "";//保存输入的值
         
            //追加可搜索属性            
            var attr = $(elem).attr('lay-search');
            if (attr == undefined) {
                elem.attr('lay-search', '');
            }
            elem.next('div').find('input').first().prop("autocomplete", "off");
            form.render('select');
          
           
            $(document).on('focusout keydown keyup', elem.next('div').find('input').first(), function (e) {
                if (e.type == 'keydown' && e.keyCode == 13) {
                    if (tempInput != "" && tempInput != undefined && data.indexOf(tempInput) < 0) {
                        if (elem.next('div').find('.layui-this')[0].innerText.indexOf(tempInput) < 0) {
                            elem.append('<option value="' + tempInput + '" selected>' + tempInput + '</option>');
                            form.render('select');
                            elem.next('div').find('input').first().prop("autocomplete", "off");  
                            data.push(tempInput);
                            tempInput = "";
                        }
                    }
                } else if (e.type == 'keyup') {
                    tempInput = elem.next('div').find('input').first().val();//找到select对应的input
                }              
                else if (e.type == "focusout") {
                    if (tempInput != "" && tempInput != undefined && data.indexOf(tempInput) < 0) {
                        if (elem.next('div').find('.layui-this')[0].innerText.indexOf(tempInput) < 0) {
                            elem.append('<option value="' + tempInput + '" selected>' + tempInput + '</option>');
                            form.render('select');
                            elem.next('div').find('input').first().prop("autocomplete", "off");
                            data.push(tempInput);
                            tempInput = "";
                        }
                    }
                } 

            });
        }
    };
    //外部接口
    exports(modelName, InsertSelect);
});