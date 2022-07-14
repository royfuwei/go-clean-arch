# 專案目錄及架構

我們開發的架構參考自[go-clean-arch](https://github.com/bxcodec/go-clean-arch), 目的是為了讓我們的 code 能夠解耦合、更容易的進行單元測試及抽換元件。

[詳細文章請閱讀這邊](https://medium.com/hackernoon/golang-clean-archithecture-efd6d7c43047)

目錄結構

```
config 啟動參數配置。
doc markdown說明文件。
docs swagger api文件。
infrastructures 框架跟driver...等。
    |- mongodb
    |- ginrest
    |- tools 
core 功能模組。
    |- delivery 開放給外面存取的界面(restful、grpc...)。
    |- repository DB抽象化，存儲資料的地方(mongo、mysql...)。
    |- usecase 操控各個Repository並回應給Delivery層。
    |- service(optional) 存放一些共用的功能，開放給別的功能模組使用。
domain 功能開發前可以先到這一層來定義介面
    |- mocks 測試用的假物件
    |- errcode 錯誤代碼
    |- errors.go response error、use case error...
    |- modules 所有開發的模組放在modules 底下
        |- users 模組
            |- UsersDelivery
            |- UsersRepository
            |- UsersUsecase
```

Reference:

- [[Android 十全大補] Clean Architecture](https://ithelp.ithome.com.tw/articles/10224386)

## 實際開發步驟
1. 需求確定後，