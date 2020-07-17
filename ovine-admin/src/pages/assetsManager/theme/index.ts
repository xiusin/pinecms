export const schema =
  {
    "title": "网站主题",
    "remark": "选择想要的主题以后，需要刷新缓存才能生效",
    "body": {
      "type": "crud",
      "api": "GET assets-manager/theme",
      "mode": "cards",
      // "quickSaveApi": "https://houtai.baidu.com/api/sample/bulkUpdate",
      // "quickSaveItemApi": "https://houtai.baidu.com/api/sample/$id",
      "draggable": true,
      "card": {
        "header": {
          "title": "$name",
          "subTitle": "作者： $author",
          "avatar": "<%= data.avatar || \"https://internal-amis-res.cdn.bcebos.com/images/2019-12/1575350573496/4873dbfaf6a5.png\" %>",
          "avatarClassName": "pull-top thumb b-3x m-r"
        },
        "actions": [
          {
            "type": "button",
            "label": "设为主题",
            "actionType": "ajax",
            "level": "<%= data.is_default ? success : error %>",  //todo 如何通过条件判断按钮颜色
            "confirmText": "确定要修改主题吗， 您可能要刷新缓存后才能生效?",
            "api": "POST assets-manager/set-theme?theme=$dir"
          }
        ],
        "body": [
          {
            "name": "description",
            "label": "描述"
          }
        ]
      }
    }
  }
