package database

// func init() {
//         //open a db connection
//         if databaseConn == nil {
//                 d, err := gorm.Open("mysql", config.GetConfig().Db)
//                 d.LogMode(false)
//                 if err != nil {
//                         log.Println(err)
//                         panic(err)
//                 }
//                 // skip save associations of gorm -> manual save by code
//                 databaseConn = d.Set("gorm:save_associations", false)
//                 databaseConn.DB().SetMaxOpenConns(20)
//                 databaseConn.DB().SetMaxIdleConns(10)
//         }
// }
//
// func DB() *gorm.DB {
//         return databaseConn
// }
