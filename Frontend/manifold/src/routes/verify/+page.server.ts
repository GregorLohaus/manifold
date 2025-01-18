import { verifySchema } from "./schema"
import { fail, superValidate } from "sveltekit-superforms"
import { zod } from "sveltekit-superforms/adapters"
import { redirect, type Actions, type ServerLoadEvent } from "@sveltejs/kit" 
import { DefaultApiPath } from "$lib/default_api_url"
import { loginSchema } from "../login/schema"
export const load = async( e:ServerLoadEvent) => {
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
    form: await superValidate(zod(verifySchema))
  }
}

export const actions: Actions = {
  default: async(event:any) => {
    const form = await superValidate(event,zod(verifySchema))
    if (!form.valid) {
      return fail(400,{form})
    }
    let body = ""
    try {
      body = JSON.stringify(form.data)
    } catch (e) {
      return fail(400,{form})
    }
    let res = await fetch(DefaultApiPath("verify"), {
      method: "POST",
      body : body,
      headers: {
        "Content-Type" : "application/json"
      }
    })
    let responseBody:any = {}
    if (res.ok) {
      let loginform = await superValidate(zod(loginSchema))
      loginform.data.email = form.data.email
      let param = btoa(JSON.stringify(loginform))
      let urlparam = new URLSearchParams()
      urlparam.set("form",param)
      redirect(300,`/login?${urlparam.toString()}`)
    }
    try {
      let txt = await res.text()
      responseBody = JSON.parse(txt)
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
      case 106: 
        let loginform = await superValidate(zod(loginSchema))
        loginform.data.email = form.data.email
        loginform.message = "User already exists."
        let param = btoa(JSON.stringify(loginform))
        let urlparam = new URLSearchParams()
        urlparam.set("form",param)
        redirect(300,`/login?${urlparam.toString()}`)
      default:
        form.message = `Unkown error: ${responseBody.error_code}`
        return fail(400,{form})
    }
  }
}
