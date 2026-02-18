import type { UploadFile } from "ant-design-vue"

//所有类别
export type Category = '前端开发'|'后端开发'|'移动端开发'|'全部'|''
// 后端返回的包装结构
export interface ApiResponse<T> {
    code: number
    msg: string
    data: T
}
//PPT卡片的结构
export interface PPTcard{
    id:number
    title:string
    cover_url:string
    file_url:string
    category:Category
    created_at:Date
    updated_at:Date
}
//搜索内容的结构
export interface SearchContent{
    category:Category
    title:string
}
//上传ppt的结构
export interface UploadPPT{
    title:string
    category:Category
    coverImgList:Array<UploadFile>
    filelist:Array<UploadFile>
}