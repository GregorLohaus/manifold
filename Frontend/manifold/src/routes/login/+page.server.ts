import { redirect, type Actions, type RequestEvent, type ServerLoadEvent } from "@sveltejs/kit"
import { loginSchema } from "./schema";
import { zod } from "sveltekit-superforms/adapters";
import { parseCookie } from "$lib/parse_cookie";
import { fail, superValidate } from "sveltekit-superforms";
import { DefaultApiPath } from "$lib/default_api_url";
export const load = async(e:ServerLoadEvent) => {
  let form:any = null;
  let formString = e.url.searchParams.get("form")?.toString()
  if (formString !== undefined) {
    try {
      form = JSON.parse(atob(formString))
    } catch (e) {
      form = null
    }
  }
  if (form !== null) {
    return {form: form};
  }
  return {
    form: await superValidate(zod(loginSchema))
  }
}

export const actions: Actions = {
  default: async(event:RequestEvent) => {
    console.log("action")
    const form = await superValidate(event,zod(loginSchema))
    if (!form.valid) {
      console.log("form not valid")
      return fail(400,{form})
    }
    console.log("form valid")
    let body = ""
    try {
      body = JSON.stringify(form.data)
    } catch (e) {
      console.log(e)
      return fail(400,{form})
    }
    let res = await fetch(DefaultApiPath("login"), {
      method: "POST",
      body : body,
      headers: {
        "Content-Type" : "application/json"
      }
    })
    console.log(res)
    if (res.ok) {
      console.log("ok")
      let cookie = parseCookie(res.headers.getSetCookie()[0])
      event.cookies.set(cookie.name,cookie.value,cookie.options)
      return redirect(300,"/app")
    }
    let responseBody:any = {}
    try {
      let txt = await res.text()
      responseBody = JSON.parse(txt)
      console.log(responseBody)
    } catch (e) {
      return fail(400,{form})
    }
    if (responseBody.error_text !== null) {
      form.message = responseBody.error_text
      return fail(400,{form})
    }
    switch(responseBody.error_code) {
      case 102: 
        form.message = `User not found.`
        return fail(400,{form})
      case 103: 
        form.message = `Wrong Password.`
        return fail(400,{form})
      case 104: 
        form.message = `No registrationkey found for user.`
        return fail(400,{form})
      case 105: 
        form.message = `Wrong registrationkey.`
        return fail(400,{form})
      default:
        form.message = `Unkown error: ${responseBody.error_code}`
        return fail(400,{form})
    }
    // console.log(res)
  }
}
