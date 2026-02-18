export const transformURL = (originUrl:string):string =>{
        const path = originUrl.split('\\')
        const dir = path[1]
        const filename = path[2]
        const baseUrl = "http://localhost:8080/static/"
        return baseUrl + dir + "/" + filename
}