### GET

`/recommend` 推荐列表

```json
[
  {
    "category": "Python 项目",
    "categoryID": 1,
    "contents": [
      {
        "category": "Python 项目",
        "categoryID": 1,
        "description":
          "廖老师的 Python 入门教程中的实践项目的代码，[教程在线阅读](http://www.liaoxuefeng.com/wiki/001374738125095c955c1e6d8bb493182103fac9270762a000/001397616003925a3d157284cd24bc0952d6c4a7c9d8c55000)\n\n",
        "projectURL": "https://github.com/michaelliao/awesome-python-webapp",
        "title": "awesome-python-webapp"
      }
    ]
  },
  {
    "category": "JavaScript 项目",
    "categoryID": 2,
    "contents": [
      {
        "category": "JavaScript 项目",
        "categoryID": 2,
        "description":
          "阿里开源的一套企业级的 UI 设计语言和 React 实现。[中文文档](https://ant.design/docs/react/introduce-cn)，样式偏向于后端，展示效果十分漂亮\n\n",
        "projectURL": "https://github.com/ant-design/ant-design",
        "title": "ant-design"
      }
    ]
  },
  {
    "category": "开源书籍",
    "categoryID": 4,
    "contents": [
      {
        "category": "开源书籍",
        "categoryID": 4,
        "description":
          "《Python Cookbook 3rd 中文版》，[在线阅读](http://python3-cookbook.readthedocs.org/zh_CN/latest/)\n\n",
        "projectURL": "https://github.com/yidao620c/python3-cookbook",
        "title": "python3-cookbook"
      }
    ]
  }
]
```

---

### GET

`/recommend/:categoryID` 根据项目category推荐列表

`TODO: category enumerics`

- @categoryID Enum{}

```json
{
  "category": "Python 项目",
  "categoryID": 1,
  "contents": [
    {
      "category": "Python 项目",
      "categoryID": 1,
      "description":
        "廖老师的 Python 入门教程中的实践项目的代码，[教程在线阅读](http://www.liaoxuefeng.com/wiki/001374738125095c955c1e6d8bb493182103fac9270762a000/001397616003925a3d157284cd24bc0952d6c4a7c9d8c55000)\n\n",
      "id": 1,
      "projectURL": "https://github.com/michaelliao/awesome-python-webapp",
      "title": "awesome-python-webapp"
    }
  ]
}
```
