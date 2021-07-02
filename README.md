# eve-api-service
a esi service, provide type price search for other users


design

redis数据结构
1. 索引  
   a. 中文索引  TYPE_INDEX:[type name] typeId  
   b. 英文索引  TYPE_INDEX:[type name] typeId 
   
2. 热点数据记录
TYPE_HOT:TYPE_ID count
   
3. Order数据
   TypeOrder:TypeId:Order score
   
4. TypeId数据

TYPE_ID list
   