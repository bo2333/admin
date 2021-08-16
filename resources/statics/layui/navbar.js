/*
// 引入
<script src="/resources/statics/layui/navbar.js"></script>
// 使用
    layui.use(['element', 'navbar'], function () {
      let element = layui.element,
              $ = layui.jquery,
              navbar = layui.navbar();

      navbar.set({
        elem: $("#nav-list-html"),　　// 存在组件的容器
        target: "rIframe",           // 输出到的目标
        data: navs,　　               // 提供给组件的数据列表，请严格按照规定格式提供。
      });

      //渲染navbar
      navbar.render();

      //监听点击事件
      navbar.on('click(side)', function (data) {
        //let childHtml = data.field.href;    // 获取当前点击的路径
        //$("#rIframe").attr("src",childHtml);// 将路径赋值给iframe展示
        windowChange($(window).width())
      });
    })
* */
layui.define(['element'], function (exports) {
    "use strict";
    let $ = layui.jquery,element = layui.element;
    let Navbar = function () {
		// 默认配置
        this.config = {
            elem: undefined, //容器
            target: undefined, // 输出目标
            data: undefined, //数据源
            skinParent: undefined, // 主导航图标与文字的色彩
            skinBg:undefined, // 主导航背景色
        };
        this.v = '1.00';
    };

    //渲染
    Navbar.prototype.render = function () {
        let _that = this;
        let _config = _that.config;
        if (_config.data === undefined) {
            throw new Error('Navbar error: 请为Navbar配置数据源.');
        }
        if (typeof (_config.elem) !== 'object') {
            throw new Error('Navbar error: elem参数需要 object ');
        }

        let $container = _config.elem;
        if (typeof (_config.data) === 'object') {
            let html = getHtml(_config.data, _config.target, _config.skinParent, _config.skinBg);
            $container.html(html);
            element.init();
            _that.config.elem = $container;
        }
        return _that;
    };

	/**
	 * 配置Navbar
	 * @param {Object} options
	 */
    Navbar.prototype.set = function (options) {
        let that = this;
        that.config.data = undefined;
        $.extend(true, that.config, options);
        return that;
    };

    Navbar.prototype.changeBg = function (bg) {
        let that = this;
        let config = that.config;
        config.skinBg = bg
        let $container = config.elem;
        let html = getHtml(config.data, config.target, config.skinParent, config.skinBg);
        $container.html(html);
        element.init();
        config.elem = $container;
        that.render()
        return that;
    };

	/**
	 * 绑定事件
	 * @param {String} events
	 * @param {Function} callback
	 */

    Navbar.prototype.on = function (events, callback) {

        let that = this;
        let _con = that.config.elem;
        if (typeof (events) !== 'string') {
            throw new Error('Navbar error:事件名配置出错，请参考API文档.');
        }
        let lIndex = events.indexOf('(');
        let eventName = events.substr(0, lIndex);
        let filter = events.substring(lIndex + 1, events.indexOf(')'));
        if (eventName === 'click') {
            if (_con.attr('lay-filter') !== undefined) {
                _con.children('ul').find('li').each(function () {
                    let $this = $(this);
                    if ($this.find('dl').length > 0) {
                        let $dd = $this.find('dd').each(function () {
                            $(this).on('click', function () {
                                let $a = $(this).children('a');
                                let href = $a.data('url');
                                let icon = $a.children('i:first').data('icon');
                                let title = $a.children('cite').text();
                                let data = {
                                    elem: $a,
                                    field: {
                                        href: href,
                                        icon: icon,
                                        title: title
                                    }
                                }
                                callback(data);
                            });
                        });
                    } else {
                        $this.on('click', function () {
                            let $a = $this.children('a');
                            let href = $a.data('url');
                            let icon = $a.children('i:first').data('icon');
                            let title = $a.children('cite').text();
                            let data = {
                                elem: $a,
                                field: {
                                    href: href,
                                    icon: icon,
                                    title: title
                                }
                            }
                            callback(data);
                        });
                    }
                });
            }
        }
    };


    /**
     * 父节点
     * @param data
     * @param target
     * @param skinParent
     * @returns {string}
     */
    function parentNode(data , target, skinParent) {
        let html = '';

        if (data.href !== undefined &&  data.href !== '') {
            html += `<a href="${data.href}"`
            if (target !== undefined) {
                html += ` target="${target}"`
            }
            html += '>'
        } else {
            html += '<a href="javascript:;">';
        }
        if (data.icon !== undefined && data.icon !== '') {
            html += `<i class="${data.icon} ${skinParent}"></i>`;
        }
        html += `<span class="nav-label ${skinParent}">${data.title}</span>`
        html += '</a>';

        return html
    }

    /**
     * 子节点
     * @param chdData
     * @param target
     * @returns {string}
     */
    function childrenNode(chdData, target) {
        let html = ''
        if (chdData !== undefined && chdData !== null && chdData.length > 0) {
            html = '<dl class="layui-nav-child">'
            for (let j = 0; j < chdData.length; ++j) {
                if (chdData[j].href === '') {
                    html += `<dd><a><hr class="${chdData[j].title}"></a></dd>`;
                    continue
                }

                html += `<dd title="${chdData[j].title}">`;
                // html += `<a href="javascript:;" data-url="${chdData[j].href}">`;

                html += `<a href="${chdData[j].href}"`;
                if (target !== undefined) {
                    html += ` target="${target}"`
                }
                html += '>'

                if (chdData[j].icon !== undefined && chdData[j].icon !== '') {
                    html += `<i class="${chdData[j].icon}" style="color: rgb(167, 177, 194);"></i>`;
                }
                html += `<span>${chdData[j].title}</span>`;
                html += '</a>';
                html += '</dd>';
            }
            html += '</dl>';
        }
        return html
    }


    /**
     * 获取html字符串
     * @param data
     * @param target
     * @param skinParent
     * @param skinBg
     * @returns {string}
     */
    function getHtml(data,target, skinParent, skinBg) {
        let ulHtml = `<ul class="layui-nav layui-input-inline ${skinBg}">`;
        for (let i = 0; i < data.length; i++) {
            ulHtml += '<li class="layui-nav-item">';
            ulHtml += parentNode(data[i], target, skinParent)
            ulHtml += childrenNode(data[i].children, target)
            ulHtml += '</li>';
        }
        ulHtml += '</ul>';
        return ulHtml;
    }

    let navbar = new Navbar();

    exports('navbar', function (options) {
        return navbar.set(options);
    });
});