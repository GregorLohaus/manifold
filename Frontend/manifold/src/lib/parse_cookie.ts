type cookie = {name:string,value:string,options:any}
export const parseCookie = (str:string) => {
  let cookie = str.split(';');
  let cookieObj:cookie = { name:"",value:"", options: {}}
  for (const [i,e] of cookie.entries()) {
    let valueParts = e.split("=")
    if (i==0) {
      cookieObj.name = valueParts[0]
      cookieObj.value = valueParts[1]
    } else if (valueParts.length > 1) {
      cookieObj.options[valueParts[0].trim()] = valueParts[1].trim()
    } else {
      cookieObj.options[valueParts[0].trim()] = true
    }
  }
  return cookieObj;
}
