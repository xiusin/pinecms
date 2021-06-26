import cloneDeep from "lodash/cloneDeep";

/**
 * 去除字符串首尾空格
 * @param {string} s
 */
export const trim = s => {
  if (s) {
    return s.replace(/(^\s*)|(\s*$)/g, "");
  }
  return "";
};

/**
 * 获取指定数量的空格
 * @param {int} indent
 */
export const getIndent = indent => {
  let string = "";
  for (let i = 0; i < indent; i++) {
    string += "\xa0";
  }
  return string;
};

/**
 * 根据参数生成json字符串
 * @param {array}} params
 * @param {int} indent
 * @param {*} notes
 */
export const renderParamsCode = (params, indent = 0, notes) => {
  const indentContent = getIndent(indent);
  const valueIndentContent = getIndent(indent + 2);

  let code = indentContent + "{\n";
  if (indent > 0) {
    code = "";
  }

  if (params && params.length) {
    params.forEach(item => {
      let fieldValue = item.default;
      if (!fieldValue) {
        switch (item.type) {
          case "int":
            fieldValue = 0;
            break;
          case "boolean":
            fieldValue = false;
            break;
          case "date":
            fieldValue = getNowTime();
            break;

          case "datetime":
            fieldValue = getNowTime("yyyy-MM-dd HH:mm:ss");
            break;
          case "time":
            fieldValue = getNowTime("HH:mm:ss");
            break;

          default:
            // eslint-disable-next-line no-undef
            fieldValue = config.USE_TYPE_DEFAULT_VALUE ? item.type : "";
            break;
        }
      }
      if (item.type === "array" && (fieldValue === "array" || !fieldValue)) {
        fieldValue = "[]";
      }

      let value = ["int", "float", "boolean", "array"].includes(item.type)
        ? fieldValue
        : `"${trim(fieldValue)}"`;
      let type = "string";
      let noteText = "";
      if (notes) {
        noteText = "//" + item.desc + "";
      }

      if (item.type == "object" && item.params && item.params.length) {
        let arrayCode = "{    " + noteText + "\n";
        arrayCode += renderParamsCode(item.params, indent + 2, notes);
        arrayCode += valueIndentContent + "},\n";
        value = arrayCode;

        type = "object";
      } else if (item.type == "array" && item.params && item.params.length) {
        let childrenTypeCodeStart = valueIndentContent + getIndent(2);
        let childrenTypeCodeEnd = valueIndentContent + getIndent(2);
        // 根据子节点类型渲染不同的结构
        if (item.childrenType == "array") {
          childrenTypeCodeStart = childrenTypeCodeStart + "[\n";
          childrenTypeCodeEnd = childrenTypeCodeEnd + "]\n";
          item.params = item.params.map(p => {
            p.name = null;
            return p;
          });
        } else if (["string", "int"].includes(item.childrenType)) {
          childrenTypeCodeStart = "";
          childrenTypeCodeEnd = "";
          item.params = item.params.map(p => {
            p.name = null;
            return p;
          });
        } else {
          childrenTypeCodeStart = childrenTypeCodeStart + "{\n";
          childrenTypeCodeEnd = childrenTypeCodeEnd + "}\n";
        }
        let arrayCode = "[    " + noteText + "\n";
        arrayCode += childrenTypeCodeStart;
        arrayCode += renderParamsCode(item.params, indent + 4, notes);
        arrayCode += childrenTypeCodeEnd;
        arrayCode += valueIndentContent + "],\n";
        value = arrayCode;
        type = "array";
      } else if (item.type == "tree" && item.params && item.params.length) {
        let arrayCode = "[    " + noteText + "\n";
        arrayCode += valueIndentContent + getIndent(2) + "{\n";
        arrayCode += renderParamsCode(item.params, indent + 4);
        arrayCode += valueIndentContent + getIndent(2) + "}\n";
        arrayCode += valueIndentContent + "],\n";
        value = arrayCode;
        type = "tree";
      }
      let desc = "";
      if (!(type === "array" || type == "object" || type === "tree")) {
        if (notes) {
          desc = `,  // ${item.desc}\n`;
        } else {
          desc = `,\n`;
        }
      }

      code += `${valueIndentContent}${
        item.name ? item.name + ": " : ""
      }${value}${desc}`;
    });
  }
  if (indent == 0) {
    code += indentContent + "}\n";
  }
  return code;
};

/**
 * 设置当前Url
 * @param {string} url
 */
export const setCurrentUrl = url => {
  window.history.replaceState(
    {
      path: url
    },
    "",
    url
  );
};

export const changeUrlArg = (url, arg, val) => {
  var pattern = arg + "=([^&]*)";
  var replaceText = arg + "=" + val;
  return url.match(pattern)
    ? url.replace(eval("/(" + arg + "=)([^&]*)/gi"), replaceText)
    : url.match("[?]")
    ? url + "&" + replaceText
    : url + "?" + replaceText;
};

export const deleteUrlArg = (url, name) => {
  var urlArr = url.split("?");
  if (urlArr.length > 1 && urlArr[1].indexOf(name) > -1) {
    var query = urlArr[1];
    var obj = {};
    var arr = query.split("&");
    for (var i = 0; i < arr.length; i++) {
      arr[i] = arr[i].split("=");
      obj[arr[i][0]] = arr[i][1];
    }
    delete obj[name];
    var urlte =
      urlArr[0] +
      "?" +
      JSON.stringify(obj)
        .replace(/[\\"\\{\\}]/g, "")
        .replace(/\\:/g, "=")
        .replace(/\\,/g, "&");
    return urlte;
  } else {
    return url;
  }
};

/**
 * 获取当前url参数
 */
export const getUrlQuery = () => {
  const query = window.location.search.substring(1);
  const vars = query.split("&");
  let values = {};
  for (let i = 0; i < vars.length; i++) {
    const pair = vars[i].split("=");
    values[pair[0]] = pair[1];
  }
  return values;
};

export const treeTransArray = (tree, key) => {
  return [].concat(
    ...tree.map(item => {
      if (item[key] && item[key].length) {
        const currentItem = cloneDeep(item);
        delete currentItem[key];
        return [].concat(currentItem, ...treeTransArray(item[key], key));
      } else {
        return item;
      }
    })
  );
};

export const getTreeMaxlevel = (treeData, childrenField = "children") => {
  // let level = 0;
  // let v = this;
  let maxLevel = 0;
  function loop(data, level) {
    data.forEach(item => {
      item.level = level;
      if (level > maxLevel) {
        maxLevel = level;
      }
      if (childrenField in item) {
        if (item[childrenField].length > 0) {
          loop(item[childrenField], level + 1);
        }
      }
    });
  }
  loop(treeData, 1);
  return maxLevel;
};

export const getTreeFirstNode = (tree, childrenField = "children") => {
  var temp = [];
  var forFn = function(arr) {
    if (arr && arr.length > 0) {
      temp.push(arr[0]);
      if (arr[0][childrenField]) {
        forFn(arr[0][childrenField]);
      }
    }
  };
  forFn(tree);
  return temp;
};

/**
 * 将文本内的特殊标记替换成html
 * @param {string} text
 * @returns
 */
export const textToHtml = text => {
  return text ? text.replace(/ /g, "&nbsp;").replace(/\r\n/g, "<br>") : "";
};

export const getNowTime = (fmt = "yyyy-MM-dd") => {
  const date = new Date();
  var o = {
    "M+": date.getMonth() + 1, //月份
    "d+": date.getDate(), //日
    "H+": date.getHours(), //小时
    "h+": date.getHours(), //小时
    "m+": date.getMinutes(), //分
    "s+": date.getSeconds(), //秒
    "q+": Math.floor((date.getMonth() + 3) / 3), //季度
    S: date.getMilliseconds() //毫秒
  };
  if (/(y+)/.test(fmt)) {
    fmt = fmt.replace(
      RegExp.$1,
      (date.getFullYear() + "").substr(4 - RegExp.$1.length)
    );
  }
  for (var k in o) {
    if (new RegExp("(" + k + ")").test(fmt)) {
      fmt = fmt.replace(
        RegExp.$1,
        RegExp.$1.length == 1 ? o[k] : ("00" + o[k]).substr(("" + o[k]).length)
      );
    }
  }
  return fmt;
};

export const getCurrentAppConfig = (appKey, apps) => {
  if (appKey && appKey.indexOf("_") > -1) {
    let list = apps;
    let currentApp = "";
    const arr = appKey.split("_");
    for (let i = 0; i < arr.length; i++) {
      const item = arr[i];
      const find = list.find(p => p.folder === item);
      currentApp = find;
      if (find && find.items && find.items.length) {
        list = find.items;
      }
    }
    return currentApp;
  } else if (appKey) {
    const find = apps.find(p => p.folder === appKey);
    return find;
  }
  return "";
};
