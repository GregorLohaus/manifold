import { fail, superValidate } from "sveltekit-superforms"
import { registerSchema } from "./schema"
import { zod } from "sveltekit-superforms/adapters"
import { redirect, type Actions } from "@sveltejs/kit"
import { DefaultApiPath } from "$lib/default_api_url"
import { verifySchema } from "../verify/schema"

export const load = async() => {
  return {
    form: await superValidate(zod(registerSchema))
  }
}

export const actions: Actions = {
  default: async(event:any) => {
    const form = await superValidate(event,zod(registerSchema));
    if (!form.valid) {
      return fail(400, {form});
    }
    if (form.data.password !== form.data.re_password){
      form.message = "Passwords dont match."
      return fail(400,{form})
    }
    let body = ""
    try {
      body = JSON.stringify(form.data)
    } catch (e) {
      return fail(400, {form})
    }
    let res = await fetch(DefaultApiPath("register"), {
      method: "POST",
      body: body,
      headers: {
        "Content-Type" : "application/json"
      }
    })
    let responseBody:any = {}
    if (res.ok) {
      let validateform = await superValidate(zod(verifySchema))
      validateform.data.email = form.data.email
      let param = btoa(JSON.stringify(validateform))
      let urlparam = new URLSearchParams()
      urlparam.set("form",param)
      redirect(300,`/verify?${urlparam.toString()}`)
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
      case 101:
        let validateform = await superValidate(zod(verifySchema))
        validateform.data.email = form.data.email
        validateform.message = "User already exists."
        let param = btoa(JSON.stringify(validateform))
        let urlparam = new URLSearchParams()
        urlparam.set("form",param)
        redirect(300,`/verify?${urlparam.toString()}`)
      default:
        form.message = `Unkown error: ${responseBody.error_code}`
        return fail(400,{form})
    }
  }
}
